package events

import (
	"github.com/MarkSmersh/go-vk-ai-userbot/core"
	"github.com/redis/go-redis/v9"
)

type VKAIUserBot struct {
	Vk   core.VK
	OAi  core.OpenAI
	Dpsk core.Deepseek
	Rdb  *redis.Client
	// Groups you want to use to get friends from
	TargetGroups []int
	Typing       core.State[int, bool]
	FriendsAdded []int
	Config       VKAIUserBotConfig
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
	// From 0 to 2. 1 is default
	LLMTemparature float64
	// That phrase will be used by LMM to prevent the bot from proceeding conversation with a user
	SafePhrase string
}
