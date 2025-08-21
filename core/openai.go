package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/MarkSmersh/go-vk-ai-userbot/types/openai"
)

type OpenAI struct {
	Token string
}

func (o *OpenAI) Request(b OpenAIRequestBuilder) openai.Response {
	body, _ := json.Marshal(b.Req)

	r := bytes.NewReader(body)

	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/responses", r)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", o.Token))
	req.Header.Add("Connection", "keep-alive")

	res, err := http.DefaultClient.Do(req)

	resBody, _ := io.ReadAll(res.Body)

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

type OpenAIRequestBuilder struct {
	Req openai.Request
}

func CreateOpenAIBuilder(model string) OpenAIRequestBuilder {
	b := OpenAIRequestBuilder{}
	b.Req.Model = model
	return b
}

func (b *OpenAIRequestBuilder) AddInput(content string, role string) {
	b.Req.Input = append(b.Req.Input, openai.Input{Role: role, Content: content})
}
