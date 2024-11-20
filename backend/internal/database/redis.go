package database

import (
    "fmt"
    "os"

    "github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func ConnectToRedisDB() {
    host := os.Getenv("REDIS_HOST")
    port := os.Getenv("REDIS_PORT")
    password := os.Getenv("REDIS_PASSWORD")

    rdb = redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%s", host, port),
        Password: password, 
    })
}

func GetRedisDB() *redis.Client {
    return rdb
}
