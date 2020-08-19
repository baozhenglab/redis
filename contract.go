package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis"
)

type RedisService interface {
	Context() context.Context
	WithContext(ctx context.Context) *redis.Client
	Options() *redis.Options
	PoolStats() *redis.PoolStats
	Pipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error)
	TxPipelined(fn func(redis.Pipeliner) error) ([]redis.Cmder, error)
	TxPipeline() redis.Pipeliner
	Subscribe(channels ...string) *redis.PubSub
	PSubscribe(channels ...string) *redis.PubSub
	Conn(ctx context.Context) *redis.Conn

	Get(key string) *redis.StringCmd
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}
