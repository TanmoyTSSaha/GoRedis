package client

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestNewClient1(t *testing.T) {

	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:5001",
		Password: "",
		DB:       0,
	})

	err := client.Set(ctx, "mykey", 42, 0).Err()
	if err != nil {
		panic(err)
	}

	c, err := New("localhost:5001")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		if err := c.Set(context.TODO(), "foo", "1"); err != nil {
			log.Fatal(err)
		}
		// val, werr := c.Get(context.TODO(), fmt.Sprintf("foo_%d", i))
		// if err != nil {
		// 	log.Fatal(err)
		// }
	}
}
