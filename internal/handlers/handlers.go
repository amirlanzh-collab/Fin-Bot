package handlers

import (
	tele "gopkg.in/telebot.v3"
)

func MainMenu(bot *tele.Bot) {
	menu := tele.ReplyMarkup{ResizeKeyboard: true}

	btnBallance := menu.Text("Баланс")
	btnHistory := menu.Text("История")
	btnTransaction := menu.Text("Перевод")
	btnHelp := menu.Text("Помощь")

	menu.Reply(
		menu.Row(btnBallance, btnTransaction),
		menu.Row(btnHistory, btnHelp),
	)

	TransMarkUp := tele.ReplyMarkup{}

	btn100 := TransMarkUp.Data("100", "trans_amt", "100")
	btn300 := TransMarkUp.Data("300", "trans_amt", "300")
	btn500 := TransMarkUp.Data("500", "trans_amt", "500")
	btn1000 := TransMarkUp.Data("1000", "trans_amt", "1000")
	btnCustomSum := TransMarkUp.Data("Введите свою сумму", "trans_CustSum")

	TransMarkUp.Inline(
		TransMarkUp.Row(btn100, btn300, btn500, btn1000),
		TransMarkUp.Row(btnCustomSum),
	)

	bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Привет! я твой бот для финансовых транзакий", &menu)
	})
	bot.Handle(&btnBallance, func(c tele.Context) error {
		return c.Send("пока в разработке")
	})

	bot.Handle(&btnTransaction, func(c tele.Context) error {

		return c.Send("Выберете сумму перевода или введите свою", &TransMarkUp)
	})

	bot.Handle(&btnCustomSum, func(c tele.Context) error {
		selector := tele.ReplyMarkup{}
		btnCancel := selector.Data("Отмена", "trans_cancel")

		selector.Inline(
			selector.Row(btnCancel),
		)
		return c.Send("Введите свою сумму", &selector)
	})

	bot.Handle(&btnHistory, func(c tele.Context) error {
		return c.Send("пока в разработке")
	})

	bot.Handle(&btnHelp, func(c tele.Context) error {
		return c.Send("Этот бот умеет переводить суммы, показывать историю переводов и баланс")
	})
}
