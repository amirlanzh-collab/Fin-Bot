package handlers

import (
	"finbot/service"
	"log"

	tele "gopkg.in/telebot.v4"
)

func Ballance(bot *tele.Bot, accountService *service.AccountService) {
	bot.Handle(&btnBallance, func(c tele.Context) error {
		balance, err := accountService.GetBalance(c.Sender().ID)
		if err != nil {
			log.Println("balane:", err)
			return c.Send("Не удалось получить ваш баланс, нажмите /start")
		}

		return c.Send("Ваш баланс:" + balance.StringFixed(2))
	})
}
