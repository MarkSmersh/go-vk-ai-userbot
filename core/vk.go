package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/MarkSmersh/go-vk-ai-userbot/types/vk/general"
	"github.com/MarkSmersh/go-vk-ai-userbot/types/vk/methods"
)

type VK struct {
	Token   string
	Version string
}

type Response struct {
	Error    general.Error `json:"error"`
	Response any           `json:"response"`
}

func (vk *VK) Request(method string, params any) ([]byte, error) {
	paramsValues := url.Values{}

	if params != nil {
		var paramsMap map[string]any

		tmp, _ := json.Marshal(params)

		d := json.NewDecoder(strings.NewReader(string(tmp[:])))

		d.UseNumber()

		d.Decode(&paramsMap)

		for k, v := range paramsMap {
			paramsValues.Add(k, fmt.Sprintf("%v", v))
		}
	}

	url := fmt.Sprintf(
		"https://api.vk.com/method/%s?%s&access_token=%s&v=%s",
		method,
		paramsValues.Encode(),
		vk.Token,
		vk.Version,
	)

	println(url)

	res, err := http.Get(url)

	if err != nil {
		return []byte{}, err
	}

	v := Response{}

	d := json.NewDecoder(res.Body)
	d.UseNumber()
	d.Decode(&v)

	if v.Error.ErrorCode != 0 {
		e := errors.New(
			fmt.Sprintf("VK API Error. Code: %d. Message: %s", v.Error.ErrorCode, v.Error.ErrorMsg),
		)

		println(e.Error())

		return nil, e
	}

	result, _ := json.Marshal(v.Response)

	return result, nil
}

func (vk *VK) UsersGet(params methods.UsersGet) []general.User {
	res, _ := vk.Request("users.get", params)
	v := []general.User{}
	json.Unmarshal(res, &v)
	return v
}
