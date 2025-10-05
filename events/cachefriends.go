package events

import (
	"github.com/MarkSmersh/go-vk-ai-userbot/types/vk/methods"
)

func (b *VKAIUserBot) CacheFriendsAndRequests() {
	friends := b.Vk.FriendsGet(methods.FriendsGet{}).Items
	requests := b.Vk.FriendsGetRequests(methods.FriendsGetRequests{
		Out:   1,
		Count: 1000,
	}).Items

	b.AddFriends(friends...)
	b.AddFriendRequests(requests...)
}
