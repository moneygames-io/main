package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
)

var currentPort int

func main() {
	currentPort = 10000
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()

	if err != nil {
		fmt.Println("Gameserver-service could not connect to redis")
		fmt.Println(err)
		return
	}

	// TODO  redis has pubsub which might be better than polling
	doEvery(3*time.Second, checkRedis, client)
}

func doEvery(d time.Duration, f func(*redis.Client), c *redis.Client) {
	for range time.Tick(d) {
		f(c)
	}
}

func makeSpec(image string, externPort int) swarm.ServiceSpec {
	max := uint64(1)

	spec := swarm.ServiceSpec{
		Annotations: swarm.Annotations{
			Name: "sneks_gameserver_" + strconv.Itoa(externPort),
			Labels: map[string]string{
				"com.docker.stack.image":     "parthmehrotra/gameserver",
				"com.docker.stack.namespace": "sneks",
			},
		},
		TaskTemplate: swarm.TaskSpec{
			RestartPolicy: &swarm.RestartPolicy{
				MaxAttempts: &max,
				Condition:   swarm.RestartPolicyConditionNone,
			},
			ContainerSpec: swarm.ContainerSpec{
				Image: image,
				Env:   []string{"GSPORT=" + strconv.Itoa(externPort)},
				Labels: map[string]string{
					"com.docker.stack.image":     "parthmehrotra/gameserver",
					"com.docker.stack.namespace": "sneks",
				},
			},
			Networks: []swarm.NetworkAttachmentConfig{
				swarm.NetworkAttachmentConfig{
					Target: "sneks_default",
				},
			},
		},
		EndpointSpec: &swarm.EndpointSpec{
			Ports: []swarm.PortConfig{
				swarm.PortConfig{
					TargetPort:    uint32(10000),
					PublishedPort: uint32(externPort),
				},
			},
		},
	}
	return spec
}

func makeOpts() types.ServiceCreateOptions {
	authConfig := types.AuthConfig{
		Username: "parthmehrotra",
		Password: PASSWORD,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	return types.ServiceCreateOptions{
		EncodedRegistryAuth: authStr,
	}
}

func addGameServer(redisClient *redis.Client) {
	dockerClient, dockerErr := client.NewEnvClient()
	if dockerErr != nil {
		fmt.Println("DOCKER ERROR")
		fmt.Println(dockerErr)
		return
	}

	createResponse, serviceErr :=
		dockerClient.ServiceCreate(
			context.Background(),
			makeSpec("parthmehrotra/gameserver", currentPort),
			makeOpts())

	fmt.Println(createResponse)
	if serviceErr != nil {
		fmt.Println("Service ERROR")
		fmt.Println(serviceErr)
		return
	}

	redisErr := redisClient.Set(strconv.Itoa(currentPort), "idle", 0).Err()
	if redisErr != nil {
		fmt.Println("REDDIS ERROR")
		fmt.Println(serviceErr)
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
