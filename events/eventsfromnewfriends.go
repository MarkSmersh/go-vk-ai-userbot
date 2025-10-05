package events

import (
	"slices"
	"time"

	"github.com/MarkSmersh/go-vk-ai-userbot/types/vk/events"
	"github.com/MarkSmersh/go-vk-ai-userbot/types/vk/methods"
)

func (b *VKAIUserBot) EventsFromNewFriends() {
	for {
		time.Sleep(1 * time.Minute)

		newFriends := b.Vk.FriendsGet(methods.FriendsGet{}).Items

		for _, newFriend := range newFriends {
			if slices.Contains(b.friends, newFriend) {
				continue
			}

			progress, _ := b.GetProgress(newFriend)

			if progress >= 0 {
				continue
			}

			b.Vk.Updater.Messages.Invoke(
				events.NewMessage{
					MessageId: 0,
					Flags:     0, // there is no flags
					PeerId:    newFriend,
					Timestamp: 0,
					Text:      "",
					RandomId:  0,
					// Attachments: i.LastMessage.Attachments,
				},
			)
		}

		b.AddFriends(newFriends...)
		b.RemoveFriendRequests(newFriends...)
	}
}
