package events

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

func (b *VKAIUserBot) Init() {
	// check config and apply default if needed
	if b.Config.MessagesInHistory == 0 {
		b.Config.MessagesInHistory = 10
	}

	if b.Config.SecondsBeforeRead == 0 {
		b.Config.SecondsBeforeRead = 20
	}

	if b.Config.SecondsBeforeWrite == 0 {
		b.Config.SecondsBeforeWrite = 5
	}

	if b.Config.SymbolsPerSecond == 0 {
		b.Config.SymbolsPerSecond = 10
	}

	if b.Config.Link == "" {
		slog.Error("Link value in VKAIUserBot config is necessary.")
		os.Exit(1)
	}

	if b.Config.LLMTemparature == 0 {
		b.Config.LLMTemparature = 0
	}

	if b.Config.RequestWait == 0 {
		b.Config.RequestWait = 60
	}

	if b.Config.SafePhrase == "" {
		b.Config.SafePhrase = "damn"
	}

	// redis availability checkout

	_, err := b.Rdb.Ping(context.Background()).Result()

	if err != nil {
		slog.Error(
			fmt.Sprintf("Redis: %s", err.Error()),
		)
		os.Exit(1)
	}

	// init

	b.Vk.Updater.Messages.Add(b.NewMessage)

	b.CacheFriends()

	go b.EventsFromUnread()
	go b.SendFriendRequests()
	go b.Vk.Start()
}
