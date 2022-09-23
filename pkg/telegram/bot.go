package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func BotStart() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Loop through each update.
	for update := range updates {
		switch update.Message.Text { // ignore any non-Message updates
		case "Да":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "да-да")
			bot.Send(msg)
		case "Нет":

		default:
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я ещё глупенький бот и понимаю только команды(\nМожешь написать /help, чтобы узнать о моих функциях")
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case "start":
			msg.Text = ""
		default:
			msg.Text = "Я пока не знаю такой команды"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
