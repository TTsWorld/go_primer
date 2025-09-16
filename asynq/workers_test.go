package asynq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/hibiken/asynq"
)

// workers.go
func TestWorkers(t *testing.T) {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{Concurrency: 10},
	)

	// 使用asynq.HandlerFunc适配器来处理函数
	if err := srv.Run(asynq.HandlerFunc(handler)); err != nil {
		log.Fatal(err)
	}
}

func handler(ctx context.Context, t *asynq.Task) error {
	switch t.Type() {
	case "email:welcome":
		var p EmailTaskPayload
		if err := json.Unmarshal(t.Payload(), &p); err != nil {
			return err
		}
		log.Printf(" [*] 给用户 %d 发送欢迎邮件", p.UserID)

	case "email:reminder":
		var p EmailTaskPayload
		if err := json.Unmarshal(t.Payload(), &p); err != nil {
			return err
		}
		log.Printf(" [*] 给用户 %d 发送提醒邮件", p.UserID)

	default:
		return fmt.Errorf("意外的任务类型：%s", t.Type())
	}
	return nil
}
