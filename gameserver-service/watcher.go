package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

var currentPort int

// TODO can we just cut redis out? Can we use docker as a store of state? Do we need to use redis for messaging? Can we just hold this state in memory and just operate on Docker's state?

func main() {
	currentPort = 10000
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()

	if err != nil {
		fmt.Println(err)
		fmt.Println("exiting")
		return
	}

	fmt.Println("not exiting")
	doEvery(3*time.Second, checkRedis, client)
}

func doEvery(d time.Duration, f func(*redis.Client), c *redis.Client) {
	for range time.Tick(d) {
		f(c)
	}
}

func addGameServer(c *redis.Client) {
	err := c.Set(strconv.Itoa(currentPort), "idle", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	currentPort++
}

func checkRedis(c *redis.Client) {
	idleCount := 0
	keys, _ := c.Keys("*").Result()

	for _, key := range keys {
		status, _ := c.Get(key).Result()
		if status == "idle" {
			idleCount++
		}
	}

	if idleCount < 2 {
		addGameServer(c)
	}
}
