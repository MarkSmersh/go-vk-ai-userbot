package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"maps"
	"net/http"
	"net/url"
	"strconv"
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

	slog.Info("Vk Userbot is started")

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

	slog.Debug(url)

	res, err := http.Get(url)

	if err != nil {
		slog.Error(err.Error())
	}

	v := events.Event{}

	d := json.NewDecoder(res.Body)

	d.UseNumber()

	d.Decode(&v)

	// jsonBytes, _ := json.Marshal(v)
	//
	// slog.Debug(string(jsonBytes[:]))

	for _, u := range v.Updates {
		code, _ := u[0].(json.Number).Int64()

		switch code {

		case 4:
			{
				attachmentsMap := map[int]events.Attachment{}

				// vk's' idea to put attachments as "attach${i}_product_id" is disgusting

				for k, v := range u[7].(map[string]any) {
					// "attach1": "62793",
					// "attach1_product_id": "1308",
					// "attach1_type": "sticker",
					//   "attachments": "[{\"type\":\"sticker\",\"sticker\":{\"images\":[{\"height\":64,\"url\":\"https://vk.com/sticker/1-62793-64\",\"width\":64},{\"height\":128,\"url\":\"https://vk.com/sticker/1-62793-128\",\"width\":128},{\"height\":256,\"url\":\"https://vk.com/sticker/1-62793-256\",\"width\":256},{\"height\":352,\"url\":\"https://vk.com/sticker/1-62793-352\",\"width\":352},{\"height\":512,\"url\":\"https://vk.com/sticker/1-62793-512\",\"width\":512}],\"images_with_background\":[{\"height\":64,\"url\":\"https://vk.com/sticker/1-62793-64b\",\"width\":64},{\"height\":128,\"url\":\"https://vk.com/sticker/1-62793-128b\",\"width\":128},{\"height\":256,\"url\":\"https://vk.com/sticker/1-62793-256b\",\"width\":256},{\"height\":352,\"url\":\"https://vk.com/sticker/1-62793-352b\",\"width\":352},{\"height\":512,\"url\":\"https://vk.com/sticker/1-62793-512b\",\"width\":512}],\"product_id\":1308,\"sticker_id\":62793}}]",
					// "attachments_count": "1"

					if strings.Contains(k, "attachments") {
						continue
					}

					// this is why replies won't work
					if !strings.Contains(k, "attach") {
						continue
					}

					// ['1', 'product', 'id']

					tags := strings.Split(strings.Split(k, "attach")[1], "_")

					if len(tags) <= 0 {
						continue
					}

					id, _ := strconv.Atoi(tags[0])

					attachment := attachmentsMap[id]

					if len(tags) == 1 {
						attachment.ID, _ = strconv.Atoi(v.(string))
					} else if len(tags) > 1 {
						typo := tags[1]

						if typo == "product" {
							attachment.ProductId, _ = strconv.Atoi(v.(string))
						}

						if typo == "type" {
							attachment.Type = v.(string)
						}
					}

					attachmentsMap[id] = attachment
				}

				attachments := []events.Attachment{}

				for v := range maps.Values(attachmentsMap) {
					attachments = append(attachments, v)
				}

				vk.Updater.Messages.Invoke(
					events.NewMessage{
						MessageId: jsonNumToInt(u[1]),
						Flags:     jsonNumToInt(u[2]),
						// MinorId:   jsonNumToInt(u[3]),
						PeerId:      jsonNumToInt(u[3]),
						Timestamp:   jsonNumToInt(u[4]),
						Text:        u[5].(string),
						Attachments: attachments,
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

	slog.Debug(url)

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

		slog.Error(e.Error())

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

type GetStickersKeywordsRes struct {
	Count      int          `json:"count"`
	Dictionary []Dictionary `json:"dictionary"`
}

type Dictionary struct {
	Words            []string `json:"words"`             // Keywords or Emoji
	UserStickers     []int    `json:"user_stickers"`     // IDs of owned stickers
	PromotedStickers []int    `json:"promoted_stickers"` // IDs of promoted stickers
}

func (vk *VK) StoreGetStickersKeywords(params methods.StoreGetStickersKeywords) GetStickersKeywordsRes {
	res, _ := vk.Request("store.getStickersKeywords", params)
	v := GetStickersKeywordsRes{}
	json.Unmarshal(res, &v)
	return v
}
