package handlers

import (
	"finbot/service"
	"fmt"
	"strings"

	tele "gopkg.in/telebot.v3"
)

func History(bot *tele.Bot, operationService *service.OperationService) {
	bot.Handle(&btnHistory, func(c tele.Context) error {
		history, err := operationService.GetHistory(c.Sender().ID)
		if err != nil {
			return c.Send("Не удалось получить вашу историю")
		}
		var sb strings.Builder
		sb.WriteString("Вот ваша история:\n\n")

		for _, tx := range history {
			fmt.Fprintf(&sb, "Сумма: %s | Время: %s\n", tx.Amount, tx.CreatedAt)
		}

		return c.Send(sb.String())
	})
}
