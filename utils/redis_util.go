package utils

import (
	"github.com/go-redis/redis"
	_ "match-server/model"
)

var Client *redis.ClusterClient

func init()  {

	Client = redis.NewClusterClient(
		&redis.ClusterOptions{
			Addrs: []string{"207.246.71.35:7001", "207.246.71.35:7002", "207.246.71.35:7003","207.246.71.35:7004","207.246.71.35:7005","207.246.71.35:7006"},
			Password: "1q2w3e4r",
		})
}