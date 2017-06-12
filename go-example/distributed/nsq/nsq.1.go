package main

import (
	"log"
	"time"
	"github.com/nsqio/go-nsq"
)

func main(){
	go startConsumer()
	startProducter()
}

//生产者
func startProducter() {
	cfg := nsq.NewConfig()

	producter, err := nsq.NewProducer("127.0.0.1:4150",cfg)

	if err != nil {
		log.Fatal(err)
	}

	//发布消息
	for {
		if err := producter.Publish("test",[]byte("test message")); err != nil {
			log.Fatal("publish error: "+ err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}

//消费者
func startConsumer() {
	cfg := nsq.NewConfig()

	consumer, err := nsq.NewConsumer("test","xingcuntian001",cfg)
	if err != nil {
		log.Fatal(err)
	}

	//设置消息处理函数
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println(string(message.Body))
		return nil
	}))

	//连接到单列nsqd
	if err := consumer.ConnectToNSQD("127.0.0.1:4150"); err != nil {
		log.Fatal(err)
	}
	<-consumer.StopChan

}