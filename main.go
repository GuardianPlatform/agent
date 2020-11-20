package main

import (
	"Guardian/mqtt"
	"Guardian/utils"
	"fmt"
	mq "github.com/eclipse/paho.mqtt.golang"
	"time"
)
func main() {
	ip, err := utils.GetExternalIP()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(ip)
	opts := mq.NewClientOptions().AddBroker("tcp://47.104.19.38:1883").SetClientID(ip)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetDefaultPublishHandler(mqtt.MqttMessageHandler)
	opts.SetPingTimeout(1 * time.Second)
	mqc := mq.NewClient(opts)
	if token := mqc.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	if token := mqc.Subscribe("agent/test",0,mqtt.MqttMessageHandler); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}

	select {

	}
	//cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.40"))
	//if err != nil {
	//	panic(err)
	//}
	//defer cli.Close()
	//imgChan := make(chan interface{})
	//ConaChan := make(chan interface{})
	//go docker.GetImageList(imgChan, cli)
	//go docker.GetContainersList(ConaChan, cli)
	//res := docker.CreateContainer(cli)
	//fmt.Printf(res)
	//docker.StartContainer(res,cli)
	//img,cona := <-imgChan,<-ConaChan
	//fmt.Println(img,cona)
}
