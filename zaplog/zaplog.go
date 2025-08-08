package log

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

type Options struct {
	LogFileDir     string        //文件保存地方
	AppName        string        //日志文件前缀
	ErrorFileName  string        // error 级别日志文件名
	NormalFileName string        // 非 error 级别日志文件名
	Level          zapcore.Level //日志等级
	MaxSize        int           //日志文件小大（M）
	MaxBackups     int           // 最多存在多少个切片文件
	MaxAge         int           //保存的最大天数
	Development    bool          //是否是开发模式
	zap.Config
}

type ModOptions func(options *Options)

var (
	l               *Logger
	sp              = string(filepath.Separator)
	errWS, normalWS zapcore.WriteSyncer       //  文件IO输出
	debugConsoleWS  = zapcore.Lock(os.Stdout) // 控制台标准输出
	errorConsoleWS  = zapcore.Lock(os.Stderr)
)

type Logger struct {
	*zap.Logger
	sync.RWMutex
	Opts        *Options `json:"opts"`
	zapConfig   zap.Config
	initialized bool
}

func NewLogger(mod ...ModOptions) *zap.Logger {
	l = &Logger{}
	l.Lock()
	defer l.Unlock()
	if l.initialized {
		l.Info("[NewLogger] logger initEd")
		return nil
	}
	l.Opts = &Options{
		LogFileDir:     "",
		AppName:        "app",
		ErrorFileName:  "error.log",
		NormalFileName: "normal.log",
		Level:          zapcore.DebugLevel,
		MaxSize:        100,
		MaxBackups:     60,
		MaxAge:         30,
	}

	if l.Opts.LogFileDir == "" {
		l.Opts.LogFileDir, _ = filepath.Abs(filepath.Dir(filepath.Join(".")))
		l.Opts.LogFileDir += sp + "log" + sp
	}
	if l.Opts.Development {
		l.zapConfig = zap.NewDevelopmentConfig()
		l.zapConfig.EncoderConfig.EncodeTime = timeEncoder
	} else {
		l.zapConfig = zap.NewProductionConfig()
		l.zapConfig.EncoderConfig.EncodeTime = timeUnixNano
	}
	if l.Opts.OutputPaths == nil || len(l.Opts.OutputPaths) == 0 {
		l.zapConfig.OutputPaths = []string{"stdout"}
	}
	if l.Opts.ErrorOutputPaths == nil || len(l.Opts.ErrorOutputPaths) == 0 {
		l.zapConfig.OutputPaths = []string{"stderr"}
	}
	for _, fn := range mod {
		fn(l.Opts)
	}

	l.zapConfig.Level.SetLevel(l.Opts.Level)
	l.init()
	l.initialized = true
	return l.Logger
}

func (l *Logger) init() {
	l.setSyncs()
	var err error
	l.Logger, err = l.zapConfig.Build(l.cores(), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	defer l.Logger.Sync()
}

func (l *Logger) setSyncs() {
	f := func(fN string) zapcore.WriteSyncer {
		return zapcore.AddSync(&lumberjack.Logger{
			Filename:   l.Opts.LogFileDir + sp + l.Opts.AppName + "-" + fN,
			MaxSize:    l.Opts.MaxSize,
			MaxBackups: l.Opts.MaxBackups,
			MaxAge:     l.Opts.MaxAge,
			Compress:   false,
			LocalTime:  true,
		})
	}
	errWS = f(l.Opts.ErrorFileName)
	normalWS = f(l.Opts.NormalFileName)
	return
}

func SetMaxSize(MaxSize int) ModOptions {
	return func(option *Options) {
		option.MaxSize = MaxSize
	}
}
func SetMaxBackups(MaxBackups int) ModOptions {
	return func(option *Options) {
		option.MaxBackups = MaxBackups
	}
}
func SetMaxAge(MaxAge int) ModOptions {
	return func(option *Options) {
		option.MaxAge = MaxAge
	}
}

func SetLogFileDir(LogFileDir string) ModOptions {
	return func(option *Options) {
		option.LogFileDir = LogFileDir
	}
}

func SetAppName(AppName string) ModOptions {
	return func(option *Options) {
		option.AppName = AppName
	}
}

func SetLevel(Level zapcore.Level) ModOptions {
	return func(option *Options) {
		option.Level = Level
	}
}
func SetErrorFileName(ErrorFileName string) ModOptions {
	return func(option *Options) {
		option.ErrorFileName = ErrorFileName
	}
}
func SetNormalFileName(NormalFileName string) ModOptions {
	return func(option *Options) {
		option.NormalFileName = NormalFileName
	}
}
func SetDevelopment(Development bool) ModOptions {
	return func(option *Options) {
		option.Development = Development
	}
}

func (l *Logger) cores() zap.Option {
	fileEncoder := zapcore.NewJSONEncoder(l.zapConfig.EncoderConfig)
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = timeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	errPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.ErrorLevel && zapcore.ErrorLevel-l.zapConfig.Level.Level() > -1
	})
	normalPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return zapcore.DebugLevel-l.zapConfig.Level.Level() > -1
	})
	cores := []zapcore.Core{
		zapcore.NewCore(fileEncoder, errWS, errPriority),
		zapcore.NewCore(fileEncoder, normalWS, normalPriority),
	}
	if l.Opts.Development {
		cores = append(cores, []zapcore.Core{
			zapcore.NewCore(consoleEncoder, errorConsoleWS, errPriority),
			zapcore.NewCore(consoleEncoder, debugConsoleWS, normalPriority),
		}...)
	}
	return zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return zapcore.NewTee(cores...)
	})
}
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func timeUnixNano(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(t.UnixNano() / 1e6)
}

// InitLog dev=true 模式，日志会被输出到console
func InitLog(appname, level, fileDir string, devModule bool) {
	logLevel := zap.DebugLevel
	if "debug" == level {
		logLevel = zap.DebugLevel
	}

	if "info" == level {
		logLevel = zap.InfoLevel
	}

	if "warn" == level {
		logLevel = zap.WarnLevel
	}

	if "error" == level {
		logLevel = zap.ErrorLevel
	}
	fmt.Println("Init ZapLog, logLevel:", strings.ToUpper(logLevel.String()))

	logger = NewLogger(
		SetAppName(appname),
		SetDevelopment(devModule),
		SetErrorFileName("error.log"),
		SetNormalFileName("normal.log"),
		SetMaxAge(30),
		SetMaxBackups(30),
		SetMaxSize(1024),
		SetLevel(zap.DebugLevel),
	)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func DebugF(format string, args ...interface{}) {
	logger.Debug("", zap.Any("data", fmt.Sprintf(format, args...)))
}

func InfoF(format string, args ...interface{}) {
	logger.Info("", zap.Any("data", fmt.Sprintf(format, args...)))
}

func WarnF(format string, args ...interface{}) {
	logger.Warn("", zap.Any("data", fmt.Sprintf(format, args...)))
}

func ErrorF(format string, args ...interface{}) {
	logger.Error("", zap.Any("data", fmt.Sprintf(format, args...)))
}

func FatalF(format string, args ...interface{}) {
	logger.Fatal("", zap.Any("data", fmt.Sprintf(format, args...)))
}
