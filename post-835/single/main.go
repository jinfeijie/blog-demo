package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/gommon/random"
	"sync"
	"time"
)

var (
	instance       sync.Once
	redisSingleton redis.Conn
)

func main() {
	t := time.Now()
	for i := 0; i < 1000; i++ {
		do()
	}
	fmt.Println(time.Now().Sub(t))
}

func do() {
	rds := getInstance()
	key := random.String(10)
	_, _ = rds.Do("set", key, time.Now().UnixNano())
	rly, err := rds.Do("get", key)
	_, _ = rds.Do("del", key)
	if err != nil {
		panic(err.Error())
	}
	_, _ = redis.String(rly, err)
}

func getInstance() redis.Conn {
	instance.Do(func() {
		var err error
		redisSingleton, err = redis.Dial("tcp", "127.0.0.1:6379")
		if err != nil {
			panic(err.Error())
		}
	})
	return redisSingleton
}
