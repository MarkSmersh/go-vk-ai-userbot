package main

import (
	"log"
	"log/slog"
	"strconv"
	"sync"

	"github.com/MarkSmersh/go-vk-ai-userbot/core"
	"github.com/MarkSmersh/go-vk-ai-userbot/events"
	"github.com/MarkSmersh/go-vk-ai-userbot/utils"
	"github.com/redis/go-redis/v9"
)

var ENV = utils.GetEnv(".env")

var vk = core.VK{
	Token:   ENV["ACCESS_TOKEN"],
	Version: "5.199",
}

var dpsk = core.Deepseek{
	Token: ENV["DEEPSEEK_TOKEN"],
}

var config = events.VKAIUserBotConfig{
	MessagesInHistory:  14,
	SecondsBeforeRead:  19,
	SecondsBeforeWrite: 0,
	SymbolsPerSecond:   10,
	Link:               "https://t.me/+-7k-lwTGNZw2YWIy",
	RequestWait:        69,
	LLMTemparature:     1.5,
	SafePhrase:         "damn",
}

var redisDb, _ = strconv.Atoi(ENV["REDIS_DB"])

var rdb = redis.NewClient(&redis.Options{
	Addr:     ENV["REDIS_HOST"],
	Password: ENV["REDIS_PASSWORD"],
	DB:       redisDb,
})

var bot = events.VKAIUserBot{
	Vk:     vk,
	Dpsk:   dpsk,
	Config: config,
	Rdb:    rdb,
	// TargetGroups: []int{216632235, 204981431, 205912478}, // it should strike
	TargetGroups: []int{221578984, 141103259}, // it should strike
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
