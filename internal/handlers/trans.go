package handlers

import tele "gopkg.in/telebot.v4"

func (h *Handler) Transaction(c tele.Context) error {
	initTransMenu()
	return c.Send("Выберете сумму перевода или введите свою", &TransMarkUp)
}

func (h *Handler) CustomSum(c tele.Context) error {
	selector := tele.ReplyMarkup{}
	btnCancel := selector.Data("Отмена", "trans_cancel")

	selector.Inline(
		selector.Row(btnCancel),
	)
	return c.Send("Введите свою сумму", &selector)
}
