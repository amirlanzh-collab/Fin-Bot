package handlers

import tele "gopkg.in/telebot.v3"

func Ballance(bot *tele.Bot) {
	bot.Handle(&btnBallance, func(c tele.Context) error {
		return c.Send("пока в разработке")
	})
}
