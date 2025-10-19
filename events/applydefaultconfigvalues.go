package events

import (
	"log/slog"
	"os"
)

func (b *VKAIUserBot) ApplyDefaultConfigValues() {
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
		slog.Error("The Link value in VKAIUserBot config is necessary.")
		os.Exit(1)
	}

	if b.Config.LLMTemparature == 0 {
		b.Config.LLMTemparature = 1
	}

	if b.Config.RequestWait == 0 {
		b.Config.RequestWait = 60
	}

	if b.Config.SafePhrase == "" {
		b.Config.SafePhrase = "damn"
	}

	if b.Config.NewFriendsCheck == 0 {
		b.Config.NewFriendsCheck = 20
	}
}
