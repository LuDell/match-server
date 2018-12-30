package model

import (
	"encoding/json"
	"github.com/cihub/seelog"
	"github.com/go-redis/redis"
	"os"
	"time"
)

var SeeLogConfig *Config

func init()  {
	SeeLogConfig = initConfig()
}

func initConfig() *Config {
	file,err1 := os.Open("config/config.json")
	defer file.Close()
	if err1 !=nil {
		seelog.Error("读取配置文件错误", err1)
	}
	decoder := json.NewDecoder(file)
	config := Config{}
	err2:= decoder.Decode(&config)
	if err2 !=nil {
		seelog.Error("数据绑定错误",err2)
	}
	seelog.Info("config参数",config)
	return &config
}

type Config struct {
	Redis redis.ClusterOptions
	Mongo struct {
		Tcp string
		Port string
		Database string
		User_name string
		Password string
		Timeout time.Duration
	}
	Amqp struct{
		Tcp string
		Port string
		User_name string
		Password string
	}
	Mysql struct{
		Tcp string
		Port string
		User_name string
		Password string
		Database string
	}
}
