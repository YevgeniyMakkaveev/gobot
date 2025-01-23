package commands

import (
	"log"

	"github.com/YevgeniyMakkaveev/gobot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot         *tgbotapi.BotAPI
	productServ *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productServ *product.Service) *Commander {
	return &Commander{
		bot:         bot,
		productServ: productServ,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {

	defer func() {
		if panicV := recover(); panicV != nil {
			log.Printf("recover from panic %v", panicV)
		}
	}()
	if update.Message != nil { // If we got a message

		switch update.Message.Command() {
		case "help":
			c.Help(update.Message)
		case "list":
			c.List(update.Message)
		case "get":
			c.Get(update.Message)
		default:
			c.Default(update.Message)
		}

	}
}
func (c *Commander) Respond(inputMsg *tgbotapi.Message, text string) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, text)
	msg.ReplyToMessageID = inputMsg.MessageID

	c.bot.Send(msg)
}
