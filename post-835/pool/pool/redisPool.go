package pool

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type RedisPool Pool
func NewRedisPool(size int) *RedisPool {
	p := new(RedisPool)
	p.res = make(chan redis.Conn, size)
	for i := 0; i < size; i++ {
		p.res.(chan redis.Conn) <- p.Factory().(redis.Conn)
	}
	fmt.Println(size)
	return p
}

func (p *RedisPool) Factory() interface{} {
	c, err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		panic(err.Error())
	}
	return c
}

func (p *RedisPool) Get() (interface{}, error) {
	select {
	case r, ok := <-p.res.(chan redis.Conn):
		if !ok {
			return nil, errors.New("pool is close")
		}

		if p.close {
			return nil, errors.New("pool is close")
		}
		
		return r, nil
	}
}

func (p *RedisPool) Release(client interface{}) {
	if p.close {
		return
	}

	select {
	default:
		p.res.(chan redis.Conn) <- client.(redis.Conn)
	}
}

func (p *RedisPool) ForceCloseAll() {
	p.Lock()
	defer p.Unlock()
	p.close = true
	close(p.res.(chan redis.Conn))
}
