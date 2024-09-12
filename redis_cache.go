package main

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(c *redis.Client) *RedisCache {
	return &RedisCache{
		client: c,
	}
}

func (c *RedisCache) Get(key int) (string, error) {
	ctx := context.Background()
	keyStr := strconv.Itoa(key)
	return c.client.Get(ctx, keyStr).Result()
}

func (c *RedisCache) Set(key int, val string) error {
	ctx := context.Background()
	keyStr := strconv.Itoa(key)
	_, err := c.client.Set(ctx, keyStr, val, 0).Result()
	return err
}

func (c *RedisCache) Remove(key int) error {
	ctx := context.Background()
	keyStr := strconv.Itoa(key)
	_, err := c.client.Del(ctx, keyStr).Result()
	return err
}
