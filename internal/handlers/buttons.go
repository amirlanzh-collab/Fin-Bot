package handlers

import (
	tele "gopkg.in/telebot.v4"
)

var (
	menu           = tele.ReplyMarkup{ResizeKeyboard: true}
	btnBalance     = menu.Text("Баланс")
	btnHistory     = menu.Text("История")
	btnTransaction = menu.Text("Перевод")
	btnHelp        = menu.Text("Помощь")
	btnRefUrl      = menu.Text("Реферальная система")
)

func initMenu() {
	menu.Reply(
		menu.Row(btnBalance, btnTransaction),
		menu.Row(btnHistory, btnHelp),
		menu.Row(btnRefUrl),
	)
}

var (
	TransMarkUp  = tele.ReplyMarkup{}
	btn100       = TransMarkUp.Data("100", "trans_amt", "100")
	btn300       = TransMarkUp.Data("300", "trans_amt", "300")
	btn500       = TransMarkUp.Data("500", "trans_amt", "500")
	btn1000      = TransMarkUp.Data("1000", "trans_amt", "1000")
	btnCustomSum = TransMarkUp.Data("Введите свою сумму", "trans_CustSum")
)

func initTransMenu() {
	TransMarkUp.Inline(
		TransMarkUp.Row(btn100, btn300, btn500, btn1000),
		TransMarkUp.Row(btnCustomSum),
	)
}
