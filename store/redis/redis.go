package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/spinel/gophermart/config"
)

// RedisDB is a shortcut structure to a Redis.
type RedisDB struct {
	redis.Conn
}

// Dial creates new connection to redis.
func Dial(cfg *config.Config) (*RedisDB, error) {
	// Initialize the redis connection to a redis instance running on your local machine
	redisConn, err := redis.DialURL(cfg.REDIS_URL)
	if err != nil {
		return nil, err
	}

	return &RedisDB{redisConn}, nil
}
