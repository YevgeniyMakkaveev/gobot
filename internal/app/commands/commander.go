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
	if update.CallbackQuery != nil {
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
		c.bot.Send(msg)
		return
	}
	if update.Message != nil { // If we got a message

		switch update.Message.Command() {
		case "help":
			c.Help(update.Message)
		case "list":
			c.List(update.Message)
		case "get":
			c.Get(update.Message)
		case "delete":
			c.Delete(update.Message)
		case "add":
			c.AddNew(update.Message)
		default:
			c.Default(update.Message)
		}

	}
}
func (c *Commander) Respond(inputMsg *tgbotapi.Message, msg tgbotapi.MessageConfig) {

	msg.ReplyToMessageID = inputMsg.MessageID

	c.bot.Send(msg)
}
