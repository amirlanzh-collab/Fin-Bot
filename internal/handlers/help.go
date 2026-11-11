package handlers

import tele "gopkg.in/telebot.v4"

func (h *Handler) Help(c tele.Context) error {
	return c.Send("Напишите сюда: @granzhev69")
}
