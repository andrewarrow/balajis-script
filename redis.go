package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func AddToRedisStream(rdb *redis.Client, values map[string]interface{}) {

	id, err := rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: "stream",
		ID:     "1-0",
		Values: values,
	}).Result()
	fmt.Println(id, err)
}
