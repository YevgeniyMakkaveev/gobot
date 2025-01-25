package commands

import (
	"errors"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Delete(inputMsg *tgbotapi.Message) {

	res, err := c.handleDelete(inputMsg)
	if err != nil {
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Ошибка \n"+err.Error())
		c.Respond(inputMsg, msg)
		return

	}
	var answer = "Удаление успешно"
	if !res {
		answer = "Произошла ошибка"
	}
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, answer)
	c.Respond(inputMsg, msg)
}

func (c *Commander) handleDelete(inputMsg *tgbotapi.Message) (bool, error) {
	args := inputMsg.CommandArguments()
	arg, err := strconv.Atoi(args)
	if err != nil {
		return false, errors.New("после гет должно идти одно число")
	}
	res, err := c.productServ.Delete(arg - 1)
	if err != nil {
		return false, errors.New("неверный индекс")
	}
	return res, nil
}
