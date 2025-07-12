package main

import (
	"log"
	"log/slog"
	"sync"

	"github.com/MarkSmersh/go-vk-ai-userbot/core"
	"github.com/MarkSmersh/go-vk-ai-userbot/events"
	"github.com/MarkSmersh/go-vk-ai-userbot/utils"
)

var ENV = utils.GetEnv()

var vk = core.VK{
	Token:   ENV["ACCESS_TOKEN"],
	Version: "5.199",
}

var oai = core.OpenAI{
	Token: ENV["OPENAI_TOKEN"],
}

var typingState = core.State[int, bool]{}

var inviteState = core.State[int, int]{}

var config = events.VKAIUserBotConfig{
	MessagesInHistory:  12,
	SecondsBeforeRead:  10,
	SecondsBeforeWrite: 20,
	SymbolsPerSecond:   300,
	Link:               "https://t.me/+4B_CFaEhQtM1MDMy",
	RequestWait:        60,
	LLMTemparature:     1.0,
}

var bot = events.VKAIUserBot{
	Vk:          vk,
	OAi:         oai,
	TypingState: typingState,
	InviteState: inviteState,
	Config:      config,
	// TargetGroups: []int{216632235, 204981431, 205912478}, // it should strike
	TargetGroups: []int{205912478}, // it should strike
	FriendsAdded: []int{},
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	slog.SetLogLoggerLevel(slog.LevelDebug)

	bot.Init()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
