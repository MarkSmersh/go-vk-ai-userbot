package events

import (
	"github.com/MarkSmersh/go-vk-ai-userbot/types/vk/events"
	"github.com/MarkSmersh/go-vk-ai-userbot/types/vk/methods"
)

func (b *VKAIUserBot) EventsFromUnread() {
	chats := b.Vk.MessagesGetConversations(methods.MessagesGetConversations{
		Filter: "unread",
	})

	for _, i := range chats.Items {
		if i.Conversation.Peer.Type == "user" {
			b.Vk.Updater.Messages.Invoke(
				events.NewMessage{
					MessageId: i.LastMessage.ID,
					Flags:     0, // there is no flags
					PeerId:    i.LastMessage.PeerID,
					Timestamp: i.LastMessage.Date,
					Text:      i.LastMessage.Text,
					RandomId:  i.LastMessage.RandomID,
					// Attachments: i.LastMessage.Attachments,
				},
			)
		}
	}
}
