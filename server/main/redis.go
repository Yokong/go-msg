package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

func initPool(maxIdle, maxActive int, addr string, timeOut time.Duration) {
	pool = &redis.Pool{
		MaxIdle: maxIdle,
		MaxActive: maxActive,
		IdleTimeout: timeOut,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}
}
