package dao

import (
	"github.com/go-redis/redis"
	"strconv"
	"net"
)

func CreateRedisClient() *redis.Client {
	RdbClient := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort("localhost:", "6379"),
		Password: "",
		DB:       0,
	})
	_, err := RdbClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	return RdbClient
}

func generate_incr(client *redis.Client, incrName string) (incr_Key string) {
	result, err := client.Incr("incrname").Result()
	if err != nil {
		panic(err)
	}
	incr_Key = incrName + ":" + strconv.FormatInt(result, 10)
	return incr_Key
}

func get_lately_Incr(client *redis.Client, incrName string) (incr_Key string) {
	incr_Key, err := client.Get(incrName).Result()
	if err != nil {
		panic(err)
	}
	return incr_Key
}
