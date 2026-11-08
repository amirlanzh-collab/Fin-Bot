package handlers

import (
	"log"

	tele "gopkg.in/telebot.v4"
)

func (h *Handler) Balance(c tele.Context) error {
	balance, err := h.accounts.GetBalance(c.Sender().ID)
	if err != nil {
		log.Println("Ошибка получения баланса", err)
		return c.Send("Не удалось получить ваш баланс")
	}
	return c.Send("Ваш баланс:" + balance.StringFixed(2))
}
