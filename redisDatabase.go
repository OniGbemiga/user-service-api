package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func connectToRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Printf("Could not connect to Redis because: %v\n", err)
		return nil
	}
	fmt.Printf("Connected to Redis %v\n", pong)

	return client
}
