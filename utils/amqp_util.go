package utils

import (
	"fmt"
	"github.com/streadway/amqp"
)

func LoadMQConn() *amqp.Connection {
	var config = seeLogConfig.Amqp
	var url = fmt.Sprintf("amqp://%s:%s@%s:%s/",config.User_name,config.Password,config.Tcp,config.Port)
	conn, err := amqp.Dial(url)
	if err != nil {
		fmt.Errorf("mq connection fail %s",err)
	}
	return conn
}
