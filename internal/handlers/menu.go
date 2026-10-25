package handlers

import (
	"finbot/service"

	tele "gopkg.in/telebot.v3"
)

var (
	menu           = tele.ReplyMarkup{ResizeKeyboard: true}
	btnBallance    = menu.Text("Баланс")
	btnHistory     = menu.Text("История")
	btnTransaction = menu.Text("Перевод")
	btnHelp        = menu.Text("Помощь")
	btnRefUrl      = menu.Text("Реферальная система")
)

func SendMenu(bot *tele.Bot, userService *service.UserService) {
	menu.Reply(
		menu.Row(btnBallance, btnTransaction),
		menu.Row(btnHistory, btnHelp),
		menu.Row(btnRefUrl),
	)
	bot.Handle("/start", func(c tele.Context) error {
		_, err := userService.RegisterUser(c.Sender().ID)
		if err != nil {
			return c.Send("произошла ошибка, попробуй позже")
		}
		return c.Send("Привет! я твой бот для финансовых транзакий", &menu)
	})
}
