package handlers

import tele "gopkg.in/telebot.v3"

var (
	menu           = tele.ReplyMarkup{ResizeKeyboard: true}
	btnBallance    = menu.Text("Баланс")
	btnHistory     = menu.Text("История")
	btnTransaction = menu.Text("Перевод")
	btnHelp        = menu.Text("Помощь")
)

func SendMenu(bot *tele.Bot) {
	menu.Reply(
		menu.Row(btnBallance, btnTransaction),
		menu.Row(btnHistory, btnHelp),
	)
	bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Привет! я твой бот для финансовых транзакий", &menu)
	})
}
