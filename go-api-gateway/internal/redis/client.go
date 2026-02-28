package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func Init() error {
	RDB = redis.NewClient(&redis.Options{
		Addr:        "localhost:6379",
		Password:    "", // no password set
		DB:          0,
		DialTimeout: 2 * time.Second,
	})
	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := RDB.Ping(ctx).Err(); err != nil {
		return err
	}
	return nil
}
