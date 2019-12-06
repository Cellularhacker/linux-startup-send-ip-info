package telegram

import (
	"fmt"
	"log"
	"sendip/config"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var to *MonitorRoom

type MonitorRoom struct{}

func (MonitorRoom) Recipient() string {
	return config.TelegramChatID
}

var bot *tb.Bot

func Init() {
	log.Println("Initializing telegram bot..")
	var err error
	bot, err = tb.NewBot(tb.Settings{
		Token:  config.TelegramBotToken,
		Poller: &tb.LongPoller{Timeout: 5 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	to = &MonitorRoom{}
}

func SendMessage(message string) {
	if !config.IsProductionMode() {
		return
	}
	msg := fmt.Sprintf("%s\n%s", message, time.Now().Format(time.RFC822))
	_, _ = bot.Send(to, msg)
}
