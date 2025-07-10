package events

func (b *VKAIUserBot) Init() {
	b.Vk.Updater.Messages.Add(b.NewMessage)

	b.CacheFriends()
	b.EventsFromUnread()

	go b.SendFriendRequests()
	go b.Vk.Start()
}
