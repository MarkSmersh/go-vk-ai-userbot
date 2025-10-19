package main

import (
	"log"
	"log/slog"
	"os"
	"sync"

	"github.com/MarkSmersh/go-vk-ai-userbot/core"
	"github.com/MarkSmersh/go-vk-ai-userbot/core/llm"
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

	var llmModel llm.LLMModel

	if token := os.Getenv("OPENAI_TOKEN"); len(token) > 0 {
		llmModel = llm.NewOpenAI(token)
	}

	if token := os.Getenv("DEEPSEEK_TOKEN"); len(token) > 0 {
		llmModel = llm.NewDeepseek(token)
	}

	if llmModel == nil {
		slog.Error("There is no any provided token for llm. Use an OPENAI_TOKEN or DEEPSEEK_TOKEN enviroment variable.")
		os.Exit(1)
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
		NewFriendsCheck:    utils.GetEnvInt("NEW_FRIENDS_CHECK"),
	}

	var rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       utils.GetEnvInt("REDIS_DB"),
	})

	var bot = events.NewVKAIUserBot(
		vk,
		llmModel,
		rdb,
		utils.GetEnvArray("TARGET_GROUPS"),
		config,
	)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	slog.SetLogLoggerLevel(slog.LevelDebug)

	bot.Init()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
