package redishandler

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

var Client *redis.Client

func Connect(){
	Client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	// Ping Test

	_, err := Client.Ping(Ctx).Result()

	if err != nil {
			panic("Redis Connection Failed: " + err.Error())
	}

	fmt.Println("Connected to Redis")
}
