package utils

import (
	"fmt"
	"github.com/streadway/amqp"
	"match-server/model"
)

var MQConn *amqp.Connection

func init()  {
	var config = model.SeeLogConfig.Amqp
	var url = fmt.Sprintf("amqp://%s:%s@%s:%s/",config.User_name,config.Password,config.Tcp,config.Port)
	conn, err := amqp.Dial(url)
	if err != nil {
		fmt.Errorf("mq connection fail %s",err)
	}

	MQConn = conn
}
