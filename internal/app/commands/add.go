package commands

import (
	"errors"

	"github.com/YevgeniyMakkaveev/gobot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) AddNew(inputMsg *tgbotapi.Message) {

	res, err := c.handleAdd(inputMsg)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Ошибка \n"+err.Error())
		c.Respond(inputMsg, msg)
		return

	}
	var answer = "Добавление успешно"
	if !res {
		answer = "Произошла ошибка"
	}
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, answer)
	c.Respond(inputMsg, msg)
}

func (c *Commander) handleAdd(inputMsg *tgbotapi.Message) (bool, error) {
	args := inputMsg.CommandArguments()

	if args == "" {
		return false, errors.New("ошибка, пустая строка")
	}
	var addEl = product.Product{Title: args}
	res, err := c.productServ.AddElement(addEl)
	if err != nil {
		return false, errors.New("ошибка во время добавления")
	}
	return res, nil
}
