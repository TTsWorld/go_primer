package main

// model
const (
	ModelDeepSeekV3 = "deepseek-v3-250324-baidubce"
	ModelGpt4oMini  = "gpt-4o-mini-2024-07-18"
)

const (
	URL = "https://api.deepseek.com/v1"
)

type ModelConfig struct {
	ModelName   string
	Temperature float64
	Tag         string
}

var ModelConfigMap = map[string]ModelConfig{
	ModelDeepSeekV3: {
		ModelName:   ModelDeepSeekV3,
		Temperature: 0.2,
		Tag:         "intro",
	},
	ModelGpt4oMini: {
		ModelName:   ModelGpt4oMini,
		Temperature: 0.2,
		Tag:         "intro",
	},
}

func GetModelConfig(modelName string) ModelConfig {
	return ModelConfigMap[modelName]
}
