package main

import (
	"encoding/json"
	"fmt"
	"github.com/cihub/seelog"
	"match-server/model"
	"match-server/service"
	"match-server/utils"
	"os"
	"strconv"
	"testing"
	"time"
)

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
			seelog.Info("读取",temp)
		}(chan1)
	}

	time.Sleep(1 *time.Second)
}

func Test_config(test *testing.T)  {
	file,err1 :=os.Open("config/config.json");
	defer file.Close()
	if err1 !=nil {
		seelog.Error("读取配置文件错误", err1)
	}
	decoder := json.NewDecoder(file)
	config := model.Config{}
	err2:= decoder.Decode(&config)
	if err2 !=nil{
		seelog.Error("数据绑定错误",err2)
	}
	fmt.Println("参数详情",config)
}

func Test_pub(test *testing.T)  {
	// Publish a message.
	err := utils.Client.Publish("mychannel1", "hello").Err()
	if err != nil {
		panic(err)
	}


}


func Test_sub(test *testing.T)  {
	pubsub := utils.Client.Subscribe("mychannel1")

	// Wait for confirmation that subscription is created before publishing anything.
	_, err := pubsub.Receive()
	if err != nil {
		panic(err)
	}

	// Go channel which receives messages.
	ch := pubsub.Channel()
	//after 30s close
	time.AfterFunc(30 *time.Second, func() {
		// When pubsub is closed channel is closed too.
		_ = pubsub.Close()
	})

	// Consume messages.
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}

}

func TestMQ(test *testing.T)  {
	var connection = utils.LoadMQConn()
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

func Test(test *testing.T)  {
	age,_ := utils.Client.Do("get","age").Int()
	fmt.Println(age)
	name,_ := utils.Client.Do("get","name").String()
	fmt.Println(name)
	var resultList,_ = utils.DBExchange().Query("select * from account")
	fmt.Println(string(resultList[0]["id"]))

	balance := service.SearchBalance(127001,2160001,false)
	fmt.Println("数据库资产=",balance)

}

