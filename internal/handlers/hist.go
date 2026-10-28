package handlers

import (
	"finbot/service"
	"fmt"
	"strconv"
	"strings"

	tele "gopkg.in/telebot.v3"
)

func History(bot *tele.Bot, operationService *service.OperationService) {
	bot.Handle(&btnHistory, func(c tele.Context) error {
		history, err := operationService.GetHistory(c.Sender().ID)
		if err != nil {
			return c.Send("Не удалось получить вашу историю")
		}
		const pageSize = 3
		page := 1

		pages := (len(history) + pageSize - 1) / 3
		start := (page - 1) * pageSize
		end := min(start+pageSize, len(history))

		pageItems := history[start:end]

		var str strings.Builder

		str.WriteString("Ваша история транзакций\n\n")

		for _, op := range pageItems {
			str.WriteString(fmt.Sprintf("Сумма: %s | Дата: %s\n", op.Amount.StringFixed(2), op.CreatedAt.Format("02.01.2006 15:04")))
		}
		str.WriteString(fmt.Sprintf("\nСтр. %d из %d", page, pages))
		text := str.String()

		markup := tele.ReplyMarkup{}
		var btns []tele.Btn
		if pages > 1 {
			btns = append(btns, menu.Data("◀️", "hist_page", strconv.Itoa(page-1)))
		}

		if page < pages {
			btns = append(btns, menu.Data("▶️", "hist_page", strconv.Itoa(page+1)))
		}

		markup.Inline(markup.Row(btns...))

		return c.Send(text, &markup)
	})
}
