package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"slices"
	"strings"

	"github.com/MarkSmersh/go-vk-ai-userbot/consts"
	"github.com/MarkSmersh/go-vk-ai-userbot/types/deepseek"
)

type Deepseek struct {
	token string
}

func NewDeepseek(token string) Deepseek {
	return Deepseek{
		token: token,
	}
}

func (d Deepseek) Request(request deepseek.Request) deepseek.Response {
	body, _ := json.Marshal(request)

	r := bytes.NewReader(body)

	req, _ := http.NewRequest("POST", "https://api.deepseek.com/chat/completions", r)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", d.token))
	req.Header.Add("Connection", "keep-alive")

	slog.Debug(req.URL.String())

	res, err := http.DefaultClient.Do(req)

	resBody, _ := io.ReadAll(res.Body)

	slog.Debug(string(resBody))

	var v deepseek.Response

	defer res.Body.Close()

	if err != nil {
		slog.Error(err.Error())
	}

	if res.StatusCode >= 400 {
		slog.Error(
			fmt.Sprintf("Error from Deepseek. Status: %s. Message: %s", res.Status, string(resBody)),
		)

		return v
	}

	json.Unmarshal(resBody, &v)

	return v
}

func (d Deepseek) Builder() RequestBuilder {
	var builder RequestBuilder

	builder = &DeepseekRequestBuilder{
		req:          deepseek.Request{},
		instructions: []string{},
		dpsk:         d,
	}

	return builder
}

type DeepseekRequestBuilder struct {
	req          deepseek.Request
	instructions []string
	dpsk         Deepseek
}

func NewRequestDeepseekBuilder(dpsk Deepseek) DeepseekRequestBuilder {
	b := DeepseekRequestBuilder{
		dpsk:         dpsk,
		req:          deepseek.Request{},
		instructions: []string{},
	}
	return b
}

// shittiest code that ever exists in the Universe
func (b DeepseekRequestBuilder) Ask() Response {
	res := b.dpsk.Request(b.req)

	result := Response{
		ID:     res.ID,
		Output: []Output{},
	}

	for _, c := range res.Choices {
		result.Output = append(result.Output, Output{
			ID: "",
			Content: []OutputBlock{
				{
					Text: c.Message.Content,
				},
			},
		})
	}

	return result
}

func (b *DeepseekRequestBuilder) AddInput(content string, role string, name string) {
	b.req.Messages = append(b.req.Messages, deepseek.Message{Role: role, Content: content, Name: name})
}

func (b *DeepseekRequestBuilder) SetModel(model string) {
	b.req.Model = model
}

func (b *DeepseekRequestBuilder) AddInstruction(instruction string) {
	b.instructions = append(b.instructions, instruction)

	if len(b.req.Messages) > 0 && b.req.Messages[0].Role == consts.SYSTEM {
		b.req.Messages[0].Content = strings.Join(b.instructions, "\n\n")
	} else {
		b.req.Messages = slices.Insert(
			b.req.Messages,
			0,
			deepseek.Message{
				Role:    "system",
				Content: strings.Join(b.instructions, "\n\n"),
			},
		)
	}
}

func (b *DeepseekRequestBuilder) SetTemperature(temperature float64) {
	b.req.Temperature = temperature
}
