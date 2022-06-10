package log

import (
	"errors"
	"testing"

	"go.uber.org/zap"
)

//zaplogger  库测试
var err = errors.New("123")
var appname, level, fileDir, devModule = "zaplogger", "debug", "./logger", true

// init loggerger
func init() {
	InitLog(appname, level, fileDir, devModule)
}

func TestZapLog(t *testing.T) {
	logger.Debug("http start success， port " + "8080")
	logger.Info("http start success， port " + "8080")
	logger.Warn("http start success， port "+"8080", zap.String("err", err.Error()))
	logger.Error("http start success， port " + "8080")
	//logger.Fatal("http start success， port " + "8080" + err.Error())

}

func TestLoggerRotate(t *testing.T) {
	logger.Debug("http start success， port " + "8080")
	logger.Info("http start success， port " + "8080")
	logger.Warn("http start success， port " + "8080")
	logger.Error("http start success， port " + "8080")
	logger.Fatal("http start success， port " + "8080" + err.Error())

}

func TestLog(t *testing.T) {
	Debug("hhhhhhhhh")
	Info("hhhhhhhhh")
	Warn("hhhhhhhhh")
	Error("hhhhhhhhh")
	Fatal("hhhhhhhhh")

}
func TestDebugf(t *testing.T) {
	DebugF("it's a test log , %s, %d", "hhh", 10)
	InfoF("it's a test log , %s, %d", "hhh", 10)
	WarnF("it's a test log , %s, %d", "hhh", 10)
	ErrorF("it's a test log , %s, %d", "hhh", 10)
	FatalF("it's a test log , %s, %d", "hhh", 10)
}
