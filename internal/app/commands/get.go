package commands

import (
	"errors"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMsg *tgbotapi.Message) {

	answer, err := c.handleParse(inputMsg)
	if err != nil {
		c.Respond(inputMsg, "Ошибка \n"+err.Error())
		return

	}
	c.Respond(inputMsg, answer)
}

func (c *Commander) handleParse(inputMsg *tgbotapi.Message) (string, error) {
	args := inputMsg.CommandArguments()
	arg, err := strconv.Atoi(args)
	if err != nil {
		return "", errors.New("после гет должно идти одно число")
	}
	product, err := c.productServ.Get(arg)
	if err != nil {
		return "", errors.New("неверный индекс")
	}
	return product.Title, nil
}
