package langchain

import (
	"context"
	"fmt"
	"go_primer/utils"
	"log"
	"os"
	"testing"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func TestDeepSeek(t *testing.T) {
	llm, err := openai.New(
		openai.WithBaseURL(os.Getenv("OPENAI_BASE_URL")),
		openai.WithToken(os.Getenv("OPENAI_API_KEY")),
		openai.WithModel(DeepSeekV3),
	)
	if err != nil {
		log.Fatal("err", err)
	}
	fmt.Println(os.Getenv("OPENAI_API_KEY"))
	fmt.Println(os.Getenv("OPENAI_BASE_URL"))
	fmt.Println(utils.MarshalString(llm))
	ctx := context.Background()
	prompt := "responee use chinese language?"
	content := " 你是什么模型"

	resp, err := llm.GenerateContent(ctx, []llms.MessageContent{
		{
			Role: llms.ChatMessageTypeSystem,
			Parts: []llms.ContentPart{
				llms.TextContent{Text: prompt},
			},
		},
		{
			Role: llms.ChatMessageTypeHuman,
			Parts: []llms.ContentPart{
				llms.TextContent{Text: content},
			},
		},
	})

	if err != nil {
		log.Println("err2-->", err)
		log.Println("err2-->", err)
	}

	fmt.Println(resp.Choices[0].Content)
	fmt.Println(1111111111111111111)
	fmt.Println(utils.RenderMarkdown(resp.Choices[0].Content))

}
