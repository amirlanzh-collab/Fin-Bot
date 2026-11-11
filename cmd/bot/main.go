package main

import (
	"finbot/database"
	"finbot/database/models"
	"finbot/internal/config"
	"finbot/internal/handlers"
	"log"
	"time"

	"github.com/shopspring/decimal"
	tele "gopkg.in/telebot.v4"
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

	err = database.DB.AutoMigrate(
		&models.User{},
		&models.Account{},
		&models.Operation{},
		&models.Referral{},
	)

	if err != nil {
		log.Fatalf("Ошибка автомиграции")
	}

	log.Println("✅ Автомиграция прошла успешно!")

	welcomeBonus, err := decimal.NewFromString(cfg.WelcomeBonus)
	referralBonus, err := decimal.NewFromString(cfg.RefBonus)

	accountService := models.NewAccounrService(database.DB)
	userService := models.NewUserService(database.DB, welcomeBonus, referralBonus)
	operationService := models.NewOperationService(database.DB)
	referralSevice := models.NewReferralService(database.DB)

	h := handlers.NewHandler(bot, userService, accountService, operationService, referralSevice)

	h.Router(bot)

	log.Println("🚀 Сервер успешно запущен!")
	bot.Start()
}
