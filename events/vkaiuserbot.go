package events

import (
	"github.com/MarkSmersh/go-vk-ai-userbot/core"
	"github.com/MarkSmersh/go-vk-ai-userbot/core/llm"
	"github.com/redis/go-redis/v9"
)

type VKAIUserBot struct {
	Vk  core.VK
	LLM llm.LLMModel
	Rdb *redis.Client
	// Groups you want to use to get friends from
	TargetGroups []int
	Config       VKAIUserBotConfig

	friends        []int
	friendRequests []int
	typing         core.State[int, bool]
}

func NewVKAIUserBot(
	vk core.VK,
	llm llm.LLMModel,
	rdb *redis.Client,
	targetGroups []int,
	config VKAIUserBotConfig,
) VKAIUserBot {
	bot := VKAIUserBot{
		Vk:             vk,
		LLM:            llm,
		Rdb:            rdb,
		TargetGroups:   targetGroups,
		Config:         config,
		friends:        []int{},
		friendRequests: []int{},
		typing:         core.State[int, bool]{},
	}

	bot.ApplyDefaultConfigValues()

	return bot
}

type VKAIUserBotConfig struct {
	// Count of messages from the chat that will be represented for the context of LLM model.
	MessagesInHistory int
	// Value that will be represented as highest value in the range (1...n) of seconds to wait until bot will read user's message (good way to make interection and economize tokens).
	SecondsBeforeRead int
	// Value that will be represented as highest value in the range (1...n) of seconds to wait until bot will start writing. That range of time will be used for gathering more context if there are more than one message from the user.
	SecondsBeforeWrite int
	// Used as a value in the next formula: tokenInTheLLMAnswer / n. Result of this will be used as the time of wait before bot will send an answer.
	SymbolsPerSecond int
	// That link will be used to share with users
	Link string
	// Time to wait in seconds before next friend request
	RequestWait int
	// From 1 to 2. 1 is default
	LLMTemparature float64
	// That phrase will be used by LMM to prevent the bot from proceeding conversation with a user
	SafePhrase string
	// Time to check are there any new applied friend requests (minutes)
	NewFriendsCheck int
}
