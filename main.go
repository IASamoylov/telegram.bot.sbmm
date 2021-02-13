package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"ia.samoylov/telegram.bot.sbmm/config"
	"ia.samoylov/telegram.bot.sbmm/internal/storage"
	"ia.samoylov/telegram.bot.sbmm/internal/updatehandler"
)

func main() {
	var config = config.Export()

	bot, err := tgbotapi.NewBotAPI(config.Telegram.Token)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = config.Debug
	storage := storage.NewUserStorage()

	updateHandlerContex := updatehandler.HandlerContext{
		Bot:         bot,
		UserStorage: storage,
		AppConfig:   config,
	}

	log.Printf("Bot %s was started", bot.Self.UserName)

	uConfug := tgbotapi.NewUpdate(0)
	uConfug.Timeout = 60
	updates, err := bot.GetUpdatesChan(uConfug)

	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		switch true {
		case update.Message != nil && update.Message.IsCommand():
			updateHandlerContex.CommandHandle(&update)
		case update.CallbackQuery != nil:
			updateHandlerContex.CallbackQueryHandle(&update)
		}
	}
}
