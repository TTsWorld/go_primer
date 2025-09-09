package main

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestDeepSeekIntroGenerator(t *testing.T) {
	ctx := context.Background()
	generator, err := NewIntroGenerator()
	if err != nil {
		log.Fatal("创建生成器失败:", err)
	}

	req, err := getNovel(ctx)
	if err != nil {
		log.Fatal("获取小说失败:", err)
	}

	t0 := time.Now().UnixMilli()
	response, err := generator.GenerateIntro(ctx, *req)
	if err != nil {
		log.Fatal("生成导语失败:", err)
	}
	t1 := time.Now().UnixMilli()
	fmt.Printf("cost: %dms\n", t1-t0)

	fmt.Printf("默认导语:\n%s\n\n", response.DefaultIntro)
	fmt.Printf("备选导语:\n%s\n", response.AlternativeIntro)
}
