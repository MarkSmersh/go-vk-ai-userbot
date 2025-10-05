package events

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

func (b *VKAIUserBot) Init() {
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

	b.CacheFriendsAndRequests()

	go b.EventsFromNewFriends()
	// go b.EventsFromUnread()
	go b.SendFriendRequests()
	go b.Vk.Start()
}
