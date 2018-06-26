package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()

	if err != nil {
		fmt.Println(pong)
		fmt.Println(err)
		return
	}

	doEvery(3*time.Second, checkRedis, client)
}

func doEvery(d time.Duration, f func(*redis.Client), c *redis.Client) {
	for range time.Tick(d) {
		f(c)
	}
}

func checkRedis(c *redis.Client) {
	fmt.Printf("%v: Hello, World!\n")
}
