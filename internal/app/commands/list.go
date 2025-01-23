package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMsg *tgbotapi.Message) {
	products := c.productServ.List()
	output := "Here is all products \n\n"
	for _, p := range products {
		output += p.Title + "\n"
	}
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, output)
	msg.ReplyToMessageID = inputMsg.MessageID

	c.bot.Send(msg)
}
