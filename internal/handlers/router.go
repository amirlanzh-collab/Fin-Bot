package handlers

import (
	tele "gopkg.in/telebot.v4"
)

func (h *Handler) Router(bot *tele.Bot) {
	bot.Handle("/start", h.SendMenu)
	bot.Handle(&btnBalance, h.Balance)
	bot.Handle(&btnTransaction, h.Transaction)
	bot.Handle(&btnCustomSum, h.CustomSum)
	bot.Handle(&btnRef, h.Referral)
	bot.Handle(&btnRefUrl, h.GetRefUrl)
	bot.Handle(&btnRefList, h.RefList)
	bot.Handle(&btnRefPrev, h.RefPrev)
	bot.Handle(&btnRefNext, h.RefNext)
	bot.Handle(&btnHelp, h.Help)
	bot.Handle(&btnHistory, h.History)
	bot.Handle(&btnPrev, h.PrevBtn)
	bot.Handle(&btnNext, h.NextBtn)
}
