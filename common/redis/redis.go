package redis

import (
	"github.com/go-redis/redis"
)

var client *redis.Client

func LoadRedisClient(DNS string) error {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr: DNS,
		})
	}

	if _, err := client.Ping().Result(); err != nil {
		return err
	}

	return nil
}

func GetClient() *redis.Client {
	return client
}
