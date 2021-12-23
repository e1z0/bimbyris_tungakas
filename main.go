package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func init() {
CmdLine()
LoadSettings()
ReadMessages()
}

func main() {
	bot, err := tgbotapi.NewBotAPI(settings.Bot_token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
                if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if update.Message != nil { // If we got a message
                        DummyMessages(bot,update.Message.Chat.ID, update.Message.Text)
		}

                if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			// Extract the command from the Message.
			switch update.Message.Command() {
			case "help":
				msg.Text = "I understand /sayhi and /status."
			case "sayhi":
				msg.Text = "Hi :)"
			case "status":
				msg.Text = "I'm ok."
                        case "sensors":
                                msg.Text = sensors()
                        case "news":
                                msg.ParseMode = "HTML"
                                msg.Text = news()
			default:
				msg.Text = "I don't know that command"
			}

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
                }


	}
}
