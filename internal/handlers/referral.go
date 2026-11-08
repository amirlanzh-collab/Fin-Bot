package handlers

import (
	"fmt"

	tele "gopkg.in/telebot.v4"
)

func (h *Handler) Referral(c tele.Context) error {
	BotName := h.bot.Me.Username
	tgID := c.Sender().ID
	refUrl := fmt.Sprintf("https://t.me/%s?start=ref_%d", BotName, tgID)

	return c.Send("Ваша реферальная ссылка:\n" + refUrl)
}
