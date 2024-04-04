package controllers

import (
	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func InitializeRedisClient() {
	// Buat connect ke Redis
	client = redis.NewClient(&redis.Options{
		Addr:     "redis-14879.c100.us-east-1-4.ec2.cloud.redislabs.com:14879", // Redis server address
		Password: "JzQRabYn7F354Kc4drYqV92nz8SsLMqc",                           // Redis password
		// Addr:     "redis-12396.c13.us-east-1-3.ec2.cloud.redislabs.com:12396",
		// Password: "Z4igN5jmj6tsc9vINPMIGwvXZDeaN6eV",
	})
}
