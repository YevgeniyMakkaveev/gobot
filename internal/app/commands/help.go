package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMsg *tgbotapi.Message) {
	c.Respond(inputMsg, "/list-list \n /get-get product")
}
