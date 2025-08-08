package lumberjack

import (
	"log"
	"testing"
	"time"

	"github.com/natefinch/lumberjack"
)

func Test01(t *testing.T) {
	logger := &lumberjack.Logger{
		Filename:   "myapp.log",
		MaxSize:    1, // 兆字节
		MaxBackups: 3,
		MaxAge:     28, // 天数
	}
	log.SetOutput(logger)
	for i := 0; i < 1000000; i++ {
		log.Println("hello", i)
	}
	defer logger.Close()
	time.Sleep(time.Second * 10)

}
