package config

const OnStartMessage = "Ask me something."
const MessagePrefix = "ChatGPT"

type GPTConfig struct {
	Model          string
	Temperature    float32
	MaxTokens      int
	CompletionsUrl string
}

var DefaultGPTConfig = &GPTConfig{
	Model:          "text-babbage-001",
	Temperature:    0.9,
	MaxTokens:      256,
	CompletionsUrl: "https://api.openai.com/v1/completions",
}
