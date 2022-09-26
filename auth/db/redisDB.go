package db

import (
	"log"

	"github.com/go-redis/redis"
)

var Rdb *redis.Client

func init() {
	log.Println("redis init")
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := Rdb.Set("key", "value", 0).Err()
	if err != nil {
		log.Println("err", err)
		panic(err)
	}
}

func Set(key string, value string) {
	err := Rdb.Set(key, value, 0).Err()
	if err != nil {
		log.Println("err", err)
		panic(err)
	}
}

func Get(key string) string {
	log.Println("key", key)
	val, err := Rdb.Get(key).Result()
	if err != nil {
		log.Println("err retrieveing key")
	}
	log.Println(val)
	return val
}
