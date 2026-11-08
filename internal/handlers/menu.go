package handlers

import (
	"log"
	"strconv"
	"strings"

	tele "gopkg.in/telebot.v4"
)

func (h *Handler) SendMenu(c tele.Context) error {
	initMenu()
	payload := c.Message().Payload
	var referrerTgID int64
	if rest, ok := strings.CutPrefix(payload, "ref_"); ok {
		id, err := strconv.ParseInt(rest, 10, 64)
		if err == nil {
			referrerTgID = id
		}
	}

	log.Println("referrer:", referrerTgID)

	_, err := h.users.RegisterUser(c.Sender().ID)
	if err != nil {
		log.Fatal(err)
		return c.Send("Ошибка регистрации, попробуйте позже")
	}
	return c.Send("Привет! Я твой бот для финансовых операций!", &menu)
}
