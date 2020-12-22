package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	err2 := initRedis()
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	ret, err := redisDb.Do("LPUSH", "list", "a", "b", "c")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret)
}

var redisDb redis.Conn

func initRedis() (err error) {

	redisDb, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return err
	}
	return

}



