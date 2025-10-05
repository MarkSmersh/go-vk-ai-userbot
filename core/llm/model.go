package llm

import "log/slog"

// import (
// 	"github.com/MarkSmersh/go-vk-ai-userbot/types/deepseek"
// 	"github.com/MarkSmersh/go-vk-ai-userbot/types/openai"
// )

type LLMModel interface {
	// AskGPT(RequestBuilder) openai.Response
	// AskDeepseek(RequestBuilder) deepseek.Response
	// Compiles GPT or Deepseek response into a standardized struct. Data may miss depending upon a specific model.
	// Ask(RequestBuilder) Response
	Builder() RequestBuilder
}

type RequestBuilder interface {
	// content, role, name
	Ask() Response
	AddInput(string, string, string)
	SetModel(string)
	AddInstruction(string)
	SetTemperature(float64)
}

type Response struct {
	ID     string
	Output []Output
}

func (r Response) IsEmpty() bool {
	if len(r.Output) <= 0 {
		return true
	}

	if len(r.Output[0].Content) <= 0 {
		return true
	}

	if len(r.Output[0].Content[0].Text) <= 0 {
		return true
	}

	return false
}

func (r Response) Text() string {
	if r.IsEmpty() {
		slog.Warn("Response from LLM is empty. Use a IsEmpty() method before getting text from the response.")
		return ""
	}

	return r.Output[0].Content[0].Text
}

type Output struct {
	ID           string
	FinishReason string
	Content      []OutputBlock
}

type OutputBlock struct {
	Text string
}
