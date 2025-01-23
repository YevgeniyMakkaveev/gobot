package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Default(inputMsg *tgbotapi.Message) {
	c.Respond(inputMsg, "You wrote "+inputMsg.Text)

}
