package mqtt

import (
	"sendip/config"
	"log"

	"github.com/eclipse/paho.mqtt.golang"
)

var mqttClient mqtt.Client

func Init() {
	log.Println("MQTT Initializing...")
	options := mqtt.NewClientOptions()
	options.AddBroker(config.MqttURL)
	options.SetAutoReconnect(true)
	options.SetClientID(config.MqttClientID)
	options.SetUsername(config.MqttClientID)

	mqttClient = mqtt.NewClient(options)
	t := mqttClient.Connect()
	if t.Wait() {
		if t.Error() != nil {
			log.Println("Mqtt Error", t.Error())
		} else {
			log.Println("MQTT Connected")
		}
	}
}

func publish(topic string, qos byte, retained bool, payload interface{}) {
	if !config.IsProductionMode() {
		return
	}
	mqttClient.Publish(topic, qos, retained, payload)
}
