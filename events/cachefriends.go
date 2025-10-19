package events

import (
	"time"

	"github.com/MarkSmersh/go-vk-ai-userbot/types/vk/methods"
)

func (b *VKAIUserBot) CacheFriendsAndRequests() {
	friends := b.Vk.FriendsGet(methods.FriendsGet{}).Items
	b.AddFriends(friends...)

	for {
		requests := b.Vk.FriendsGetRequests(methods.FriendsGetRequests{
			Out:   1,
			Count: 1000,
		}).Items

		b.AddFriendRequests(requests...)

		if len(requests) < 1000 {
			break
		}

		time.Sleep(time.Second * 10)
	}

}
