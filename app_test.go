package main

import (
	"encoding/json"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
	"match-server/model"
	"match-server/utils"
	"strconv"
	"testing"
	"time"
)

var logger = utils.Logger

func Test_run(t *testing.T)  {
	var chan1 = make(chan string,3)
	for i :=0; i<10; i++ {
		go func(num int) {
			var temp= "chan" + strconv.FormatInt(int64(num),10)
			chan1 <- temp
		}(i)
	}

	for  i := 0 ; i<10; i++ {
		go func(chanTemp chan string) {
			var temp = <-chan1
			logger.Info("读取",temp)
		}(chan1)
	}

	time.Sleep(1 *time.Second)
}

func Test_config(test *testing.T)  {
	file,err1 :=os.Open("config/config.json");
	defer file.Close()
	if err1 !=nil {
		logger.Error("读取配置文件错误", err1)
	}
	decoder := json.NewDecoder(file)
	config := model.Config{}
	err2:= decoder.Decode(&config)
	if err2 !=nil{
		logger.Error("数据绑定错误",err2)
	}
	fmt.Println("参数详情",config)
}

func Test_pub(test *testing.T)  {
	con := utils.RedisPool.Get()
	psc := redis.PubSubConn{con}

	psc.Subscribe("hello")

	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Printf(v.Error())
		}
	}
}


func Test_sub(test *testing.T)  {
	con := utils.RedisPool.Get()

	con.Send("SUBSCRIBE","hello")
	con.Flush()

	for {
		reply, err := redis.Values(con.Receive())
		if err != nil {
			log.Println("Receive failed:", err)
		}

		log.Println("reply:", reply)

		for i, v := range reply {

			switch vv := v.(type) {
			case int64:
				log.Println(i, ":", vv)
			case []byte:
				log.Println(i, ":", string(vv))
			case string:
				log.Println(i, ":", vv)
			default:
				log.Println("unknown:", v)
			}
		}
	}

}

func TestMQ(test *testing.T)  {
	var connection = utils.MQConn
	defer connection.Close()

	var channel,_ = connection.Channel()
	defer channel.Close()
	//channel.ExchangeDeclare("okay","topic",false,false,false,false,nil)
	var queue,_ = channel.QueueDeclare("okay_queue",false,false,false,false,nil)
	//channel.QueueBind("okay_queue","info","okay",false,nil)

	msgs, _ := channel.Consume(
		 queue.Name,
		"smile-by",
		false,
		false,
		false,
		false,
		nil)
	forever := make(chan bool)

	 go func() {
	         for d := range msgs {
	         	header := d.Headers
	         	for k,v := range header{
	         		seelog.Info("header参数k = ",k,",v = ",v)
				}
	         	seelog.Info("Received a message: ", string(d.Body))
	         	d.Ack(false)
	         	}
	     }()

	seelog.Info(" [*] Waiting for messages. To exit press CTRL+C")

	 <-forever
}

