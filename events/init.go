package events

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

func (b *VKAIUserBot) Init() {
	_, err := b.Rdb.Ping(context.Background()).Result()

	if err != nil {
		slog.Error(
			fmt.Sprintf("Redis: %s", err.Error()),
		)
		os.Exit(1)
	}

	b.Vk.Updater.Messages.Add(b.NewMessage)

	// b.CacheFriends()
	//
	go b.EventsFromUnread()
	// go b.SendFriendRequests()
	go b.Vk.Start()
}
