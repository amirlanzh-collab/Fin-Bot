package handlers

import (
	tele "gopkg.in/telebot.v4"
)

func (h *Handler) Router(bot *tele.Bot) {
	bot.Handle("/start", h.SendMenu)
	bot.Handle(&btnBalance, h.Balance)
	bot.Handle(&btnTransaction, h.Transaction)
	bot.Handle(&btnCustomSum, h.CustomSum)
	bot.Handle(&btnRefUrl, h.Referral)
	bot.Handle(&btnHelp, h.Help)
	bot.Handle(&btnHistory, h.History)
	bot.Handle(&btnPrev, h.PrevBtn)
	bot.Handle(&btnNext, h.NextBtn)
}
