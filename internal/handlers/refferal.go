package handlers

import tele "gopkg.in/telebot.v4"

func Refferal(bot *tele.Bot) {
	bot.Handle(&btnRefUrl, func(c tele.Context) error {
		return c.Send("t.me/finfatherrrbot")
	})
}
