package fileratate

import (
	"log"
	"testing"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func Test01(t *testing.T) {
	rl, _ := rotatelogs.New("./access_log.%Y-%m-%d_%H-%M",
		rotatelogs.WithLinkName("./access_log.log"),
		rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationTime(time.Minute),
	)

	log.SetOutput(rl)

	for i := 0; i < 1000000; i++ {
		log.Printf("Hello, World! %d", i)
		time.Sleep(time.Second * 1)
	}

}
