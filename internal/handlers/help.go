package handlers

import tele "gopkg.in/telebot.v4"

func (h *Handler) Help(c tele.Context) error {
	return c.Send("Этот бот умеет переводить суммы, показывать историю переводов и баланс")
}
