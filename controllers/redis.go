package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var client *redis.Client

func InitializeRedisClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis-14879.c100.us-east-1-4.ec2.cloud.redislabs.com:14879", // Redis server address
		Password: "JzQRabYn7F354Kc4drYqV92nz8SsLMqc",                          // Redis password
	})
}

func Token() {
	p, err := client.Ping(context.Background()).Result()
	fmt.Println("lol")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(p)

	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/connect", handleConnect)

	log.Fatal(http.ListenAndServe(":14879", nil))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	switch r.Method {
	case "POST":
		// Generate Token
		token := uuid.New().String()

		// Store token untuk 1 jam (3600 detik)
		err := client.Set(ctx, token, "authenticated", 3600).Err()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, token)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func handleConnect(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	switch r.Method {
	case "GET":
		token := r.URL.Query().Get("token")
		val, err := client.Get(ctx, token).Result()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if val != "authenticated" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Perform connection to external service (e.g., Gmail, Line, WhatsApp)
		fmt.Fprint(w, "Connected to external service")
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
