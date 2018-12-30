package utils

import (
	"github.com/go-redis/redis"
	"match-server/model"
)

var Client *redis.ClusterClient

func init()  {
	var config = &model.SeeLogConfig.Redis
	Client = redis.NewClusterClient(config)
}