package asynq

import "github.com/hibiken/asynq"

var redisConnOpt = asynq.RedisClientOpt{
	Addr: "localhost:6379",
	// 如果不需要密码，可省略
	Password: "",
	// 为asynq使用一个专用的数据库编号。
	DB: 0,
}
