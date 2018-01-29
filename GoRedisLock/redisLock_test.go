/*
author: ZengJJ
date: 2018/1/28  
*/
package GoRedisLock

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis"

	"github.com/bsm/redis-lock"
)

func TestRedisLock(t *testing.T) {
	client := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 1})
	l := lock.New(client, "testRedisLock", nil)
	fmt.Println(l)
	fmt.Println(l.Lock())

	time.Sleep(time.Second * 10)

	fmt.Println(l.IsLocked())
}
