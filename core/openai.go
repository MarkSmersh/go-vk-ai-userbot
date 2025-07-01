package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MarkSmersh/go-vk-ai-userbot/types/openai"
)

type OpenAI struct {
	Token string
}

func (o *OpenAI) Request(b RequestBuilder) openai.Response {
	body, _ := json.Marshal(b.Req)

	println(string(body))

	r := bytes.NewReader(body)

	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/responses", r)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", o.Token))

	res, err := http.DefaultClient.Do(req)

	resBody, _ := io.ReadAll(res.Body)

	println(string(resBody))

	if err != nil {
		println(err)
	}

	v := openai.Response{}

	json.Unmarshal(resBody, &v)

	return v
}

type RequestBuilder struct {
	Req openai.ModelRequest
}

func CreateRequestBuilder(model string) RequestBuilder {
	b := RequestBuilder{}
	b.Req.Model = model
	return b
}

func (b *RequestBuilder) AddInput(content string, role string) {
	b.Req.Input = append(b.Req.Input, openai.Input{Role: role, Content: content})
}
