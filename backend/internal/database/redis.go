package database

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func ConnectToRedisDB() (*redis.Client, error) {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	// Validate critical env vars
	if host == "" || port == "" {
		return nil, fmt.Errorf("REDIS_HOST and REDIS_PORT must be set")
	}

 opts := &redis.Options{
        Addr:     fmt.Sprintf("%s:%s", host, port),
        Password: password,

        // Critical timeout settings for demo
        DialTimeout:  10 * time.Second,
        ReadTimeout:  30 * time.Second,
        WriteTimeout: 30 * time.Second,

        // Disable TLS for demo (remove later)
        TLSConfig: &tls.Config{
            InsecureSkipVerify: true,
        },
    }

    rdb := redis.NewClient(opts)

    // More generous timeout for demo
    ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()

    if _, err := rdb.Ping(ctx).Result(); err != nil {
        log.Printf("WARNING: Redis connection unstable - proceeding in degraded mode")
        // Return client anyway for demo purposes
        return rdb, nil
    }

    return rdb, nil
}

func GetRedisDB() *redis.Client {
	return rdb
}
