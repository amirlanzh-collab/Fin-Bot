package handlers

import (
	"finbot/internal/database.go"
	"finbot/service"
	"fmt"
	"strconv"
	"strings"

	tele "gopkg.in/telebot.v3"
)

var (
	btnPrev = menu.Data("◀️", "hist_prev")
	btnNext = menu.Data("▶️", "hist_next")
)

func History(bot *tele.Bot, operationService *service.OperationService) {
	bot.Handle(&btnHistory, func(c tele.Context) error {
		history, err := operationService.GetHistory(c.Sender().ID)
		if err != nil {
			return c.Send("Не удалось получить вашу историю")
		}
		if len(history) == 0 {
			c.Send("Ваша история пуста")
		}
		text, markup := BuildHistPage(history, 1)
		if err := c.Send(text, markup); err != nil {
			return err
		}
		return c.Respond()
	})

	bot.Handle(&btnPrev, func(c tele.Context) error {
		page, _ := strconv.Atoi(c.Callback().Data)
		history, err := operationService.GetHistory(c.Sender().ID)
		if err != nil {
			return c.Send("Не удалось получить вашу историю")
		}
		text, markup := BuildHistPage(history, page)
		if err := c.Edit(text, markup); err != nil {
			return err
		}
		return c.Respond()
	})

	bot.Handle(&btnNext, func(c tele.Context) error {
		page, _ := strconv.Atoi(c.Callback().Data)
		history, err := operationService.GetHistory(c.Sender().ID)
		if err != nil {
			return c.Send("Не удалось получить вашу историю")
		}
		text, markup := BuildHistPage(history, page)
		if err := c.Edit(text, markup); err != nil {
			return err
		}
		return c.Respond()
	})
}

func BuildHistPage(history []database.Operation, page int) (string, *tele.ReplyMarkup) {
	if page < 1 {
		page = 1
	}

	const pageSize = 3
	pages := (len(history) + pageSize - 1) / 3
	start := (page - 1) * pageSize
	end := min(start+pageSize, len(history))
	pageItems := history[start:end]

	markup := tele.ReplyMarkup{}
	var btns []tele.Btn
	if page > 1 {
		btns = append(btns, menu.Data("◀️", "hist_prev", strconv.Itoa(page-1)))
	}

	if page < pages {
		btns = append(btns, menu.Data("▶️", "hist_next", strconv.Itoa(page+1)))
	}

	markup.Inline(markup.Row(btns...))

	var str strings.Builder

	str.WriteString("Ваша история транзакций\n\n")

	for _, op := range pageItems {
		str.WriteString(fmt.Sprintf("Сумма: %s | Дата: %s\n", op.Amount.StringFixed(2), op.CreatedAt.Format("02.01.2006 15:04")))
	}
	str.WriteString(fmt.Sprintf("\nСтр. %d из %d", page, pages))
	text := str.String()

	return text, &markup
}
