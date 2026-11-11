package handlers

import (
	"finbot/database/models"
	"fmt"
	"strconv"
	"strings"

	tele "gopkg.in/telebot.v4"
)

var (
	btnRefPrev = menu.Data("◀️", "ref_prev")
	btnRefNext = menu.Data("▶️", "ref_next")
)

func (h *Handler) Referral(c tele.Context) error {
	initRefMenu()
	return c.Send("Выберите пункт в меню", &RefMarkUp)
}

func (h *Handler) GetRefUrl(c tele.Context) error {
	BotName := h.bot.Me.Username
	tgID := c.Sender().ID
	refUrl := fmt.Sprintf("https://t.me/%s?start=ref_%d", BotName, tgID)

	c.Send("Ваша реферальная ссылка:\n" + refUrl)

	return c.Respond()
}

func (h *Handler) RefList(c tele.Context) error {
	referrals, err := h.referrals.GetRefList(c.Sender().ID)
	if err != nil {
		return err
	}

	BotName := h.bot.Me.Username
	tgID := c.Sender().ID
	refUrl := fmt.Sprintf("https://t.me/%s?start=ref_%d", BotName, tgID)
	if len(referrals) == 0 {
		return c.Send("У вас пока нет рефералов. Ваша реферальная ссылка:\n" + refUrl)
	}
	text, menu := BuildRefPage(referrals, 1)

	c.Send(text, menu)
	return c.Respond()
}

func (h *Handler) RefPrev(c tele.Context) error {
	page, _ := strconv.Atoi(c.Callback().Data)
	referrals, err := h.referrals.GetRefList(c.Sender().ID)
	if err != nil {
		return err
	}
	text, menu := BuildRefPage(referrals, page)
	if err := c.Edit(text, menu); err != nil {
		return err
	}
	return c.Respond()
}

func (h *Handler) RefNext(c tele.Context) error {
	page, _ := strconv.Atoi(c.Callback().Data)
	referrals, err := h.referrals.GetRefList(c.Sender().ID)
	if err != nil {
		return err
	}
	text, menu := BuildRefPage(referrals, page)
	if err := c.Edit(text, menu); err != nil {
		return err
	}
	return c.Respond()
}

func BuildRefPage(referrals []models.User, page int) (string, *tele.ReplyMarkup) {
	if page < 1 {
		page = 1
	}
	const pageSize = 4
	pages := (len(referrals) + 3) / pageSize
	start := (page - 1) * pageSize
	end := min(start+pageSize, len(referrals))
	pageItems := referrals[start:end]

	menu := tele.ReplyMarkup{}
	var btns []tele.Btn
	if page > 1 {
		btns = append(btns, menu.Data("◀️", "ref_prev", strconv.Itoa(page-1)))
	}

	if page < pages {
		btns = append(btns, menu.Data("▶️", "ref_next", strconv.Itoa(page+1)))
	}

	menu.Inline(menu.Row(btns...))

	var str strings.Builder

	str.WriteString("Ваши рефералы:\n\n")

	for _, refs := range pageItems {
		str.WriteString(fmt.Sprintf("%s | %s\n",
			refs.Username,
			refs.CreatedAt.Format("02.01.2006 15:04")))
	}

	str.WriteString(fmt.Sprintf("\nСтр. %d из %d", page, pages))

	text := str.String()

	return text, &menu
}
