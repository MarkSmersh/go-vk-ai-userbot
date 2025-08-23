package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/MarkSmersh/go-vk-ai-userbot/types/deepseek"
)

type Deepseek struct {
	Token string
}

func (o *Deepseek) Request(b DeepseekRequestBuilder) deepseek.Response {
	body, _ := json.Marshal(b.Req)

	r := bytes.NewReader(body)

	req, _ := http.NewRequest("POST", "https://api.deepseek.com/chat/completions", r)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", o.Token))
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

type DeepseekRequestBuilder struct {
	Req deepseek.Request
}

func CreateRequestDeepseekBuilder(model string) DeepseekRequestBuilder {
	b := DeepseekRequestBuilder{}
	b.Req.Model = model
	return b
}

func (b *DeepseekRequestBuilder) AddMessage(content string, role string, name string) {
	b.Req.Messages = append(b.Req.Messages, deepseek.Message{Role: role, Content: content, Name: name})
}
