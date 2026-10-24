package main

import (
	"finbot/internal/config"
	"finbot/internal/database.go"
	"finbot/internal/handlers"
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	cfg := config.LoadEnv()

	if cfg.Token == "" {
		log.Fatal("❌ Ошибка: Токен бота пустой!")
	}
	pref := tele.Settings{
		Token:  cfg.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	database.Connect(cfg)
	database.AutoMigrate()

	handlers.SendMenu(bot)
	handlers.History(bot)
	handlers.Ballance(bot)
	handlers.Transaction(bot)
	handlers.Help(bot)

	log.Println("🚀 Сервер успешно запущен!")
	bot.Start()
}
