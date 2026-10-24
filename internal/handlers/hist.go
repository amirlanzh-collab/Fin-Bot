package handlers

import tele "gopkg.in/telebot.v3"

func History(bot *tele.Bot) {
	bot.Handle(&btnHistory, func(c tele.Context) error {
		return c.Send("пока в разработке")
	})
}
