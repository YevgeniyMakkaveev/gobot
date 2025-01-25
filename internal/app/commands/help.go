package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "/list-list \n /get-get product \n /delete-delete product \n /add-add product")
	c.Respond(inputMsg, msg)
}
