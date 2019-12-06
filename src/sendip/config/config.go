package config

import (
	"log"
	"os"
)

const (
	ModeProduction  = "Production"
	ModeDevelopment = "Development"
)

var Mode = ModeDevelopment

var MqttURL = ""
var MqttClientID = ""
var TelegramBotToken = ""
var TelegramChatID = ""

func init() {
	MqttURL = os.Getenv("SENDIP_MQTT_URL")
	MqttClientID = os.Getenv("SENDIP_MQTT_CLIENT_ID")
	TelegramBotToken = os.Getenv("SENDIP_TELEGRAM_TOKEN")
	TelegramChatID = os.Getenv("SENDIP_TELEGRAM_CHATID")

	Mode = os.Getenv("SENDIP_MODE")
	if Mode == ModeProduction {
		log.Println("Running SENDIP in Production Mode")
	} else {
		log.Println("Running SENDIP in Development Mode")
	}

	// Disabling temporary...
	//if MqttURL == "" {
	//	log.Fatal("SENDIP_MQTT_URL missing")
	//}
	//if MqttClientID == "" {
	//	log.Fatal("SENDIP_MQTT_CLIENT_ID missing")
	//}

	if TelegramBotToken == "" {
		log.Fatalln("SENDIP_TELEGRAM_TOKEN missing")
	}
	if TelegramChatID == "" {
		log.Fatalln("SENDIP_TELEGRAM_CHATID missing")
	}
}

func IsProductionMode() bool {
	return Mode == ModeProduction
}
