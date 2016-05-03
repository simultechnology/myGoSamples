package main

import (
	"fmt"
	"net/http"
	"io"
	"gopkg.in/redis.v3"
	"log"
)

var redisDB *redis.Client = NewRedisClient()

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, "hello world")
	io.WriteString(w, "hello world")
}

func secretMessageHandler(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "42")
}

func requireKey(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		userID, err := redisDB.Get("auth:" + key).Result()
		if key == "" || err != nil {
			log.Printf("%s\n", err)
			http.Error(w, "bad key", http.StatusForbidden)
			return
		}
		log.Println("user", userID, "viewed message")
		h(w, r)
	}
}

func requireKeyOrSession(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		if session := r.FormValue("session"); session != "" {
			var err error
			if key, err = redisDB.Get("session:" + session).Result(); err != nil {
				log.Printf("%s\n", err)
				http.Error(w, "bad session", http.StatusForbidden)
				return
			}
		}
		userID, err := redisDB.Get("auth:" + key).Result()
		if key == "" || err != nil {
			log.Printf("%s\n", err)
			http.Error(w, "bad key", http.StatusForbidden)
			return
		}
		log.Println("user", userID, "viewed message")
		h(w, r)
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/secret/message", requireKeyOrSession(secretMessageHandler))
	http.ListenAndServe(":13000", nil)
}

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>

	return client
}
