package events

func (b *VKAIUserBot) Init() {
	b.Vk.Updater.Messages.Add(b.NewMessage)

	go b.Vk.Start()

	// messages :=
}
