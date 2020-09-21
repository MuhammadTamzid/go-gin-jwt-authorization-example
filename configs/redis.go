package configs

import (
	"fmt"
	"log"

	"github.com/go-redis/redis/v7"
)

var RedisClient *redis.Client

func InitRedis() {
	redisEnv := EnvVariables.Redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisEnv.Host, redisEnv.Port),
		Password: redisEnv.Password,
		DB:       0,
	})

	checkRedisConnection(RedisClient)
}

func checkRedisConnection(client *redis.Client) {
	if _, err := client.Ping().Result(); err != nil {
		log.Printf("Redis connection error: %s", err)
	}
}
