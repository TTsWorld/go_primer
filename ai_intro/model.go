package main

import (
	"fmt"
	"os"

	"github.com/tmc/langchaingo/llms/openai"
)

func NewAIClient() (*openai.LLM, error) {
	config := GetModelConfig(ModelDeepSeekV3)
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		os.Exit(-1)
	}
	// 创建OpenAI兼容的LLM实例，DeepSeek使用OpenAI兼容的API
	llm, err := openai.New(
		openai.WithToken(apiKey),
		openai.WithBaseURL("https://model.in.zhihu.com/v1"),
		openai.WithModel(config.ModelName),
	)
	if err != nil {
		return nil, fmt.Errorf("创建DeepSeek客户端失败: %w", err)
	}

	return llm, nil
}
