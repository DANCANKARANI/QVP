package database

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/go-redis/redis/v8"
)

//connecting to RedisClient
func RedisClient()*redis.Client{
	rdb :=redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return rdb
}
func StartRedisServer() error {
    cmd := exec.Command("redis-server") // Assumes redis-server is in your PATH
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err := cmd.Start()
    if err != nil {
        return fmt.Errorf("failed to start Redis server: %v", err)
    }

    // Give the server some time to start
    time.Sleep(2 * time.Second)
    return nil
}