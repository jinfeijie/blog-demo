package main
import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/jinfeijie/blog-demo/post-835/pool/pool"
	"github.com/labstack/gommon/random"
	"os"
	"strconv"
	"time"
)

var (
	msg chan string
	num int
)

func init() {
	msg = make(chan string)
	num, _ = strconv.Atoi(os.Args[1])
}
func main() {
	cnt := 0
	t := time.Now()
	go fetch()
	for {
		select {
		case _ = <-msg:
			//fmt.Println(d)
			cnt++
			if cnt%100 == 0 {
				fmt.Println(time.Now().Sub(t), cnt)
				//t = time.Now()
			}
		}
	}
}

func fetch() {
	p := pool.NewRedisPool(num)
	for i := 0; i < 1010; i++ {
		go do(p)
	}
}

func do(rdsP *pool.RedisPool) {
	c, err := rdsP.Get()
	if err != nil {
		panic(err.Error())
	}
	defer rdsP.Release(c)
	rds := c.(redis.Conn)
	key := random.String(10)
	_, _ = rds.Do("set", key, time.Now().UnixNano())
	rly, err := rds.Do("get", key)
	_, _ = rds.Do("del", key)
	if err != nil {
		panic(err.Error())
	}
	s, _ := redis.String(rly, err)
	msg <- s
}
