package gredis

import (
	"encoding/json"
	"log"
	"go-gin-jwt-authorization-example/configs"
	"time"
)

func Set(key string, value interface{}, expiration int64) error {
	data, err := json.Marshal(value)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	if _, err := configs.RedisClient.Set(key, data, time.Duration(expiration)).Result(); err != nil {
		return err
	}

	return nil
}

func Get(key string, dest interface{}) error {
	data, err := configs.RedisClient.Get(key).Result()
	if err != nil {
		log.Print(err.Error())
		return err
	}
	return json.Unmarshal([]byte(data), dest)
}

func Delete(key string) error {
	deleted, err := configs.RedisClient.Del(key).Result()
	if err != nil || deleted == 0 {
		log.Print(err.Error())
		return err
	}
	return nil
}
