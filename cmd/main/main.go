package main

import (
	"log"
	"os"

	"github.com/YevgeniyMakkaveev/gobot/internal/app/commands"
	"github.com/YevgeniyMakkaveev/gobot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}
	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {

		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}
	productService := product.NewService()
	updates := bot.GetUpdatesChan(u)
	commander := commands.NewCommander(bot, productService)
	for update := range updates {
		commander.HandleUpdate(update)
	}
}
