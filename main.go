package main

import (
	"Guardian/docker"
	"Guardian/mqtt"
	"github.com/docker/docker/client"
	mq "github.com/eclipse/paho.mqtt.golang"
	"github.com/robfig/cron"
	"sync"
	"time"
)
type RegularReport struct {
	ch chan<- interface{}
	docker client.Client
}

func (that *RegularReport) Run() {
	// 向ch 里写入数据
	imgChan := make(chan interface{})
	go docker.GetImageList(imgChan, &that.docker)
	img := <- imgChan
	that.ch <- img
}
func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	opts := mq.NewClientOptions().AddBroker("tcp://47.104.19.38:1883").SetClientID("123123")
	opts.SetKeepAlive(60 * time.Second)
	opts.SetDefaultPublishHandler(mqtt.MqttMessageHandler)
	opts.SetPingTimeout(1 * time.Second)
	mqc := mq.NewClient(opts)
	if token := mqc.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	if token := mqc.Subscribe("agent/test",0,mqtt.MqttMessageHandler); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	/*
		docker sdk
	*/
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.40"))
	if err != nil {
		panic(err)
	}
	defer cli.Close()


	/*
		启动push协程用于发送数据
	*/
	ch := make(chan interface{},10)
	go mqtt.Push(mqc,ch,&wg)
	//启动定时任务
	c := cron.New()
	spec := "*/5 * * * * ?"
	c.AddJob(spec, &RegularReport{ch: ch,docker: *cli})
	c.Start()
	defer c.Stop()
	select {}
}
