package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/MarkSmersh/go-vk-ai-userbot/consts"
	"github.com/MarkSmersh/go-vk-ai-userbot/types/openai"
)

type OpenAI struct {
	token string
}

func NewOpenAI(token string) OpenAI {
	return OpenAI{
		token: token,
	}
}

func (o *OpenAI) Request(request openai.Request) openai.Response {
	body, _ := json.Marshal(request)

	r := bytes.NewReader(body)

	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/responses", r)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", o.token))
	req.Header.Add("Connection", "keep-alive")

	slog.Debug(req.URL.String())

	res, err := http.DefaultClient.Do(req)

	resBody, _ := io.ReadAll(res.Body)

	slog.Debug(string(resBody))

	var v openai.Response

	defer res.Body.Close()

	if err != nil {
		slog.Error(err.Error())
	}

	if res.StatusCode >= 400 {
		slog.Error(
			fmt.Sprintf("Error from OpenAI. Status: %s. Message: %s", res.Status, string(resBody)),
		)

		return v
	}

	json.Unmarshal(resBody, &v)

	return v
}

func (o OpenAI) Builder() RequestBuilder {
	var builder RequestBuilder
	openaiBuilder := NewOpenAIBuilder(o)
	builder = &openaiBuilder
	return builder
}

type OpenAIRequestBuilder struct {
	req          openai.Request
	instructions []string
	openai       OpenAI
}

func NewOpenAIBuilder(openai OpenAI) OpenAIRequestBuilder {
	b := OpenAIRequestBuilder{
		openai: openai,
	}
	return b
}

func (b OpenAIRequestBuilder) Ask() Response {
	b.req.Instructions = strings.Join(b.instructions, "\n\n")

	res := b.openai.Request(b.req)

	result := Response{
		ID:     res.ID,
		Output: []Output{},
	}

	for _, output := range res.Output {
		o := Output{
			ID:           output.ID,
			FinishReason: output.Status,
			Content:      []OutputBlock{},
		}

		for _, block := range output.Content {
			o.Content = append(o.Content, OutputBlock{
				Text: block.Text,
			})
		}

		result.Output = append(result.Output, o)
	}

	return result
}

func (b *OpenAIRequestBuilder) AddInput(content string, role string, name string) {
	text := ""

	if name != "" {
		text = content
	} else {
		text = fmt.Sprintf("%s :\n%s", name, content)
	}

	b.req.Input = append(
		b.req.Input,
		openai.Input{
			Role:    role,
			Content: text,
		},
	)
}

func (b *OpenAIRequestBuilder) SetModel(model string) {
	b.req.Model = model
}

func (b *OpenAIRequestBuilder) SetTemperature(temperature float64) {
	if b.req.Model == consts.GPT_4_1_NANO || b.req.Model == consts.GPT_5_MINI {
		slog.Warn(
			fmt.Sprintf("Temperature cannot be set for model %s.", b.req.Model),
		)
		return
	}

	b.req.Temperature = temperature
}

func (b *OpenAIRequestBuilder) AddInstruction(instruction string) {
	b.instructions = append(b.instructions, instruction)
}
