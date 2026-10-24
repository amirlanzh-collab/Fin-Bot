package handlers

import tele "gopkg.in/telebot.v3"

func Help(bot *tele.Bot) {
	bot.Handle(&btnHelp, func(c tele.Context) error {
		return c.Send("Этот бот умеет переводить суммы, показывать историю переводов и баланс")
	})
}
