package events

import (
	"fmt"
	"log/slog"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/MarkSmersh/go-vk-ai-userbot/consts"
	"github.com/MarkSmersh/go-vk-ai-userbot/types/vk/events"
	"github.com/MarkSmersh/go-vk-ai-userbot/types/vk/methods"
)

func (b *VKAIUserBot) NewMessage(e events.NewMessage) {
	isTyping := b.IsTyping(e.PeerId)
	userProgress, _ := b.GetProgress(e.PeerId)

	if e.Flags&2 != 0 ||
		isTyping ||
		e.PeerId >= 2000000000 ||
		e.PeerId <= 0 ||
		userProgress == consts.INVITE_OVER {
		return
	}

	b.SetTyping(e.PeerId, true)

	ampl := math.Round(rand.ExpFloat64() * float64(b.Config.SecondsBeforeRead))

	time.Sleep(time.Second * time.Duration(ampl))

	b.Vk.MessagesMarkAsRead(methods.MessagesMarkAsRead{
		PeerID: strconv.Itoa(e.PeerId),
	})

	time.Sleep(time.Second * time.Duration(b.Config.SecondsBeforeWrite))

	self := b.Vk.UsersGet(methods.UsersGet{})[0]

	builder := b.LLM.Builder()

	builder.SetModel(consts.DEEPSEEK_CHAT)

	builder.AddInstruction(fmt.Sprintf(SYSTEM_PROMPT_HEADER, self.FirstName, self.LastName))

	builder.AddInstruction(SYSTEM_PROMPT_BASIS)

	builder.AddInstruction(SYSTEM_PROMPT_LORE)

	// builder.AddInstruction(SYSTEM_PROMPT_INSTRUCTIONS)

	builder.AddInstruction(fmt.Sprintf(SYSTEM_PROMPT_LINK, b.Config.Link))

	builder.AddInstruction(fmt.Sprintf(SYSTEM_PROMPT_SAFE_PHRASE, b.Config.SafePhrase))

	if userProgress == consts.INVITE_SENT {
		builder.AddInstruction(SYSTEM_PROMPT_SENT_INVITE)
	}

	builder.AddInstruction(fmt.Sprintf(SYSTEM_PROMPT_TIME, time.Now()))

	messages := b.Vk.MessagesGetHistory(methods.MessagesGetHistory{
		Count:  b.Config.MessagesInHistory,
		PeerID: e.PeerId,
	}).Items

	target := b.Vk.UsersGet(methods.UsersGet{
		UserIDs: strconv.Itoa(e.PeerId),
	})[0]

	for i := len(messages) - 1; i >= 0; i-- {
		m := messages[i]

		text := m.Text

		if len(text) <= 0 {
			if len(m.Attachments) > 0 {
				for _, v := range m.Attachments {
					if v.Type == "sticker" {
						keyWords := b.Vk.StoreGetStickersKeywords(methods.StoreGetStickersKeywords{
							StickersIDs: v.Sticker.StickerId,
						})

						if len(keyWords.Dictionary) > 0 {
							text = keyWords.Dictionary[0].Words[0]
						}
					}
				}
			}
		}

		if len(text) <= 0 {
			continue
		}

		if m.FromID == self.ID {
			builder.AddInput(
				text,
				consts.ASSISTANT,
				"",
			)
		} else {
			builder.AddInput(
				text,
				consts.USER,
				fmt.Sprintf("%s %s", target.FirstName, target.LastName),
			)
		}
	}

	builder.SetTemperature(1.4)

	text := ""

	for i := range 3 {
		res := builder.Ask()

		if res.IsEmpty() {
			slog.Warn(
				fmt.Sprintf("Try #%d (max 3) to get response from LLM was unsuccessful. Next try in 5 minutes.", i+1),
			)

			time.Sleep(5 * time.Minute)

			continue
		}

		text = res.Text()

		break
	}

	if len(text) <= 0 {
		slog.Error("Cannot get LLM's response after 3 tries. Looks like:\n- https://api-docs.deepseek.com/faq#why-are-empty-lines-continuously-returned-when-calling-the-api\n- You need to top up balance\n- API key is expired or absent\n- Input prompt is diabolical\n- LLM's servers are down\n- LLM's are not accesseble in your country\n- AI Gen bubble blew up")

		b.SetTyping(e.PeerId, false)

		return
	}

	if userProgress == consts.INVITE_SENT || strings.Contains(text, b.Config.SafePhrase) {
		b.SetProgress(e.PeerId, consts.INVITE_OVER)
	}

	if strings.Contains(text, b.Config.Link) {
		b.SetProgress(e.PeerId, consts.INVITE_SENT)
	}

	stopActivity := false

	go b.SetActivity(&stopActivity, e.PeerId)

	for _, row := range strings.Split(text, "\n\n") { // idc
		time.Sleep(time.Second * time.Duration(len(row)/b.Config.SymbolsPerSecond))

		b.Vk.MessagesSend(methods.MessagesSend{
			UserID:  e.PeerId,
			Message: row,
		})
	}

	stopActivity = true
	b.SetTyping(e.PeerId, false)
}
