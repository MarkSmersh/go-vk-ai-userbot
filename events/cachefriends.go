package events

import "github.com/MarkSmersh/go-vk-ai-userbot/types/vk/methods"

func (b *VKAIUserBot) CacheFriends() {
	friends := b.Vk.FriendsGet(methods.FriendsGet{}).Items
	requests := b.Vk.FriendsGetRequests(methods.FriendsGetRequests{}).Items

	b.FriendsAdded = append(b.FriendsAdded, friends...)
	b.FriendsAdded = append(b.FriendsAdded, requests...)
}
