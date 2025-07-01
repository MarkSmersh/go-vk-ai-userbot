package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/MarkSmersh/go-vk-ai-userbot/types/vk/events"
	"github.com/MarkSmersh/go-vk-ai-userbot/types/vk/general"
	"github.com/MarkSmersh/go-vk-ai-userbot/types/vk/methods"
)

type VKUpdater struct {
	Messages Caller[events.NewMessage]
}

type VK struct {
	Token          string
	Version        string
	Updater        VKUpdater
	LongpollServer general.LongPollServer
}

type Response struct {
	Error    general.Error `json:"error"`
	Response any           `json:"response"`
}

func (vk *VK) Start() {
	ls := vk.MessageGetLongPollServer(methods.MessagesGetLongPollServer{
		NeedPts:   0,
		LPVersion: 3,
	})

	vk.LongpollServer = ls

	vk.longpoll(0)
}

func (vk *VK) longpoll(ts int) {
	url := fmt.Sprintf(
		"https://%s?act=a_check&key=%s&ts=%d&wait=%d&mode=%d&version=%d",
		vk.LongpollServer.Server,
		vk.LongpollServer.Key,
		ts,
		90,
		2,
		3,
	)

	res, err := http.Get(url)

	if err != nil {
		println(url)
		println(err)
	}

	v := events.Event{}

	// resBody, _ := io.ReadAll(res.Body)
	// json.Unmarshal(resBody, &v)

	d := json.NewDecoder(res.Body)

	d.UseNumber()

	d.Decode(&v)

	for _, u := range v.Updates {
		code, _ := u[0].(json.Number).Int64()

		switch code {

		case 4:
			{
				vk.Updater.Messages.Invoke(
					events.NewMessage{
						MessageId: jsonNumToInt(u[1]),
						Flags:     jsonNumToInt(u[2]),
						// MinorId:   jsonNumToInt(u[3]),
						PeerId:    jsonNumToInt(u[3]),
						Timestamp: jsonNumToInt(u[4]),
						Text:      u[5].(string),
						// Attachments: u[7],(interface{}),
						// RandomId: jsonNumToInt(u[7]),
					},
				)
			}
		}
	}

	vk.longpoll(v.Ts)
}

func jsonNumToInt(jsonNumber any) int {
	i, _ := jsonNumber.(json.Number).Int64()
	return int(i)
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
		println(url)

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

func (vk *VK) MessageGetLongPollServer(params methods.MessagesGetLongPollServer) general.LongPollServer {
	res, _ := vk.Request("messages.getLongPollServer", params)
	v := general.LongPollServer{}
	json.Unmarshal(res, &v)
	return v
}

func (vk *VK) MessagesSend(params methods.MessagesSend) int {
	res, _ := vk.Request("messages.send", params)
	v := 0
	json.Unmarshal(res, &v)
	return v
}

func (vk *VK) MessagesSetActivity(params methods.MessagesSetActivity) int {
	res, _ := vk.Request("messages.setActivity", params)
	v := 0
	json.Unmarshal(res, &v)
	return v
}

func (vk *VK) MessagesMarkAsRead(params methods.MessagesMarkAsRead) int {
	res, _ := vk.Request("messages.markAsRead", params)
	v := 0
	json.Unmarshal(res, &v)
	return v
}

type MessagesGetHistoryRes struct {
	Count int               `json:"count"`
	Items []general.Message `json:"items,omitempty"`
}

func (vk *VK) MessagesGetHistory(params methods.MessagesGetHistory) MessagesGetHistoryRes {
	res, _ := vk.Request("messages.getHistory", params)
	v := MessagesGetHistoryRes{}
	json.Unmarshal(res, &v)
	return v
}
