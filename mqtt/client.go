package mqtt

import (
	"encoding/json"
	"fmt"
	mq "github.com/eclipse/paho.mqtt.golang"
	"log"
	"sync"
)

var MqttMessageHandler mq.MessageHandler = func(client mq.Client, message mq.Message) {
	client.Publish("testtopic",0,false,"123")
	log.Println(string(message.Payload()))
}

func Push(client mq.Client,  ch <-chan  interface{},wg *sync.WaitGroup)  {
	defer wg.Done()
	for i := range ch {
		fmt.Println(i)
		imglist, err := json.Marshal(i)
		if err != nil {
			panic(err.Error())
		}
		client.Publish("testtopic",0,false,imglist)
	}
}