package handlers

import (
	"finbot/database/models"

	tele "gopkg.in/telebot.v4"
)

type Handler struct {
	bot        *tele.Bot
	users      *models.UserService
	accounts   *models.AccountService
	operations *models.OperationService
}

func NewHandler(bot *tele.Bot, users *models.UserService, accounts *models.AccountService, operations *models.OperationService) *Handler {
	return &Handler{
		bot:        bot,
		users:      users,
		accounts:   accounts,
		operations: operations,
	}
}
