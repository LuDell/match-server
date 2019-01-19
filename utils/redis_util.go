package utils

import (
	"github.com/go-redis/redis"
)

var Client *redis.ClusterClient

func init()  {
	var config = &seeLogConfig.Redis
	Client = redis.NewClusterClient(config)
}