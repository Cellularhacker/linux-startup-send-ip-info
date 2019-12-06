package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sendip/config"
	//"sendip/service/mqtt"
	"sendip/service/telegram"

	"github.com/chyeh/pubip"
)

func init() {
	telegram.Init()
	//mqtt.Init()
}

func main() {
	hostname, _ := os.Hostname()
	localIP := GetOutboundIP().String()
	pubips, _ := pubip.Get()
	pubIP := pubips.String()
	telegram.SendMessage(fmt.Sprintf("Server started successfully\nHostname:%s\nLocal IP:%s\nPublic IP:%s\nID:%s\n", hostname, localIP, pubIP, config.MqttClientID))

	return
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}
