package main

import (
	"log"
	"log/slog"
	"os"
	"sync"

	"github.com/MarkSmersh/go-vk-ai-userbot/core"
	"github.com/MarkSmersh/go-vk-ai-userbot/events"
	"github.com/MarkSmersh/go-vk-ai-userbot/utils"
	"github.com/redis/go-redis/v9"
)

func main() {
	utils.LoadDotEnv(".env")

	var vk = core.VK{
		Token:   os.Getenv("ACCESS_TOKEN"),
		Version: "5.199",
	}

	// var dpsk = core.Deepseek{
	// 	Token: os.Getenv("DEEPSEEK_TOKEN"),
	// }

	var openai = core.OpenAI{
		Token: os.Getenv("OPENAI_TOKEN"),
	}

	var config = events.VKAIUserBotConfig{
		MessagesInHistory:  utils.GetEnvInt("MESSAGES_IN_HISTORY"),
		SecondsBeforeRead:  utils.GetEnvInt("SECONDS_BEFORE_READ"),
		SecondsBeforeWrite: utils.GetEnvInt("SECONDS_BEFORE_WRITE"),
		SymbolsPerSecond:   utils.GetEnvInt("SYMBOLS_PER_SECOND"),
		Link:               os.Getenv("LINK"),
		RequestWait:        utils.GetEnvInt("REQUEST_WAIT"),
		LLMTemparature:     utils.GetEnvFloat("LLM_TEMPERATURE"),
		SafePhrase:         os.Getenv("SAFE_PHRASE"),
	}

	var rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       utils.GetEnvInt("REDIS_DB"),
	})

	var typing = core.State[int, bool]{}

	var bot = events.VKAIUserBot{
		Vk:           vk,
		OAi:          openai,
		Config:       config,
		Rdb:          rdb,
		Typing:       typing,
		TargetGroups: []int{200651056, 117082191, 220858832, 220399071, 221529283}, // it should strike
		FriendsAdded: []int{},
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	slog.SetLogLoggerLevel(slog.LevelDebug)

	bot.Init()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
