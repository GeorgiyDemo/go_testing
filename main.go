package main

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type TGUser struct {
	Tid         int
	Username    string
	FirstName   string
	LastName    string
	Counter     int
}

func main() {

	BOTTOKEN := ""
	mapUsers := make(map[int]TGUser)
	bot, err := tgbotapi.NewBotAPI(BOTTOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	//Цикл по каждому апдейту
	for update := range updates {
		if update.Message == nil { // Игнор не-сообщенек
			continue
		}

		if

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}