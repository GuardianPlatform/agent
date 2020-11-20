package mqtt

import (
	"fmt"
	mq "github.com/eclipse/paho.mqtt.golang"
)

var MqttMessageHandler mq.MessageHandler = func(client mq.Client, message mq.Message) {
	fmt.Println(client, string(message.Payload()))
}

