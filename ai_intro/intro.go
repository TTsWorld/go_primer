package main

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/tmc/langchaingo/llms"
)

// 生成导语
func (g *DeepSeekIntroGenerator) GenerateIntro(ctx context.Context, req NovelRequest) (*IntroResponse, error) {
	// 构建完整的prompt

	// 调用LLM
	completion, err := g.llm.GenerateContent(ctx, []llms.MessageContent{
		{
			Role: llms.ChatMessageTypeSystem,
			Parts: []llms.ContentPart{
				llms.TextContent{Text: AIIntroPrompt},
			},
		},
		{
			Role: llms.ChatMessageTypeHuman,
			Parts: []llms.ContentPart{
				llms.TextContent{Text: fmt.Sprintf(AIResourcePrompt, req.Question, req.Content)},
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("调用DeepSeek API失败: %w", err)
	}

	// 解析响应
	response, err := g.parseResponse(completion.Choices[0].Content)
	if err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	return response, nil
}

// func (g *DeepSeekIntroGenerator) buildPrompt(content, question string) string {
// 	return fmt.Sprintf(AIIntroPrompt, content, question)
// }

// 解析AI响应，提取默认导语和备选导语
func (g *DeepSeekIntroGenerator) parseResponse(response string) (*IntroResponse, error) {
	// 新的响应格式应该是：
	// {
	//     "intros": ["默认导语", "备选导语1"]
	// }

	// 尝试提取JSON对象
	jsonPattern := regexp.MustCompile(`\{[^{}]*"intros"\s*:\s*\[[^\]]*\][^{}]*\}`)
	jsonMatch := jsonPattern.FindString(response)

	if jsonMatch == "" {
		return nil, fmt.Errorf("未找到有效的JSON格式响应")
	}

	// 解析JSON
	var aiResponse AIResponse
	err := json.Unmarshal([]byte(jsonMatch), &aiResponse)
	if err != nil {
		return nil, fmt.Errorf("解析JSON响应失败: %w", err)
	}

	// 检查数组长度
	if len(aiResponse.Intros) < 2 {
		return nil, fmt.Errorf("导语数组长度不足，期望至少2个元素，实际收到%d个", len(aiResponse.Intros))
	}

	// 提取默认导语和备选导语
	defaultIntro := aiResponse.Intros[0]
	alternativeIntro := aiResponse.Intros[1]

	return &IntroResponse{
		DefaultIntro:     defaultIntro,
		AlternativeIntro: alternativeIntro,
	}, nil
}

// 输入结构体
type NovelRequest struct {
	ResourceId int    `json:"id"`
	AnswerId   int    `json:"type"`
	Content    string `json:"content"`
	Question   string `json:"question"`
}

// 输出结构体
type IntroResponse struct {
	DefaultIntro     string `json:"default_intro"`     // 默认导语
	AlternativeIntro string `json:"alternative_intro"` // 备选导语
}

// AI导语生成接口
type IntroGenerator interface {
	GenerateIntro(ctx context.Context, req NovelRequest) (*IntroResponse, error)
}

// DeepSeek导语生成器实现
type DeepSeekIntroGenerator struct {
	llm    llms.Model
	prompt string
}

// AI返回的JSON结构体
type AIResponse struct {
	Intros []string `json:"intros"`
}

func NewIntroGenerator() (*DeepSeekIntroGenerator, error) {
	llm, err := NewAIClient()
	if err != nil {
		return nil, fmt.Errorf("创建DeepSeek客户端失败: %w", err)
	}
	return &DeepSeekIntroGenerator{
		llm:    llm,
		prompt: AIIntroPrompt,
	}, nil
}
