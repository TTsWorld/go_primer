package langchain

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"go_primer/utils"
	"log"
	"os"
	"testing"
)

func TestLangchain(t *testing.T) {
	llm, err := openai.New(
		openai.WithBaseURL(os.Getenv("OPENAI_PUB_URL")),
		openai.WithToken(os.Getenv("OPENAI_API_KEY")),
		openai.WithModel(ChatModelGPT4o2024_08_06),
	)
	if err != nil {
		log.Fatal("err", err)
	}
	fmt.Println(utils.MarshalString(llm))
	ctx := context.Background()

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
		{
			Role: llms.ChatMessageTypeHuman,
			Parts: []llms.ContentPart{
				llms.TextContent{Text: title},
			},
		},
	})

	if err != nil {
		log.Error(ctx, "err2-->", err)
	}

	if resp.Choices == nil || len(resp.Choices) == 0 {
		return errors.New("gpt 请求无返回")
	}

	err1 := json.Unmarshal([]byte(utils.ClearJson(resp.Choices[0].Content)), &res)
	if err1 != nil {
		log.Error(ctx, "gpt结果解析出错-->", utils.ClearJson(resp.Choices[0].Content))
		return err
	}

	if validator1.Struct(res) != nil {
		log.Error(ctx, "validate error", json_print.JsonPrint(utils.ClearJson(resp.Choices[0].Content)))
		return errors.New("validate error")
	}

	return nil

}
