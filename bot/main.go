package main

import (
	"log"
	"tgBot/internal/commander"
	"tgBot/internal/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	log.Println("start main")
	bot, err := tgbotapi.NewBotAPI("6031661197:AAHS2x5nay290yIIggdEFCsASRS4TwIGNs8")
	if err != nil {
		log.Panic(err)
	}

	cmd, err := commander.Init(bot)
	if err != nil {
		log.Panic(err)
	}
	handlers.AddHandlers(cmd)

	cmd.Run()
}
