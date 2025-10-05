package events

import (
	"time"

	"github.com/MarkSmersh/go-vk-ai-userbot/types/vk/methods"
)

func (b VKAIUserBot) SetActivity(stop *bool, targetPeerId int) {
	for {
		if *stop {
			break
		}

		b.Vk.MessagesSetActivity(methods.MessagesSetActivity{
			Type:   "typing",
			PeerID: targetPeerId,
		})

		time.Sleep(time.Second * 5)
	}
}
