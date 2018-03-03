/*
author: zjjmlrs
date: 2018/1/28
*/
package GoRedisLock

import (
	"fmt"
	"github.com/bsm/redis-lock"
	"github.com/go-redis/redis"
	"time"
)

// redis-lock 简单的使用例子
func GetRedisLock() {

	// redis 客户端
	client := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 1})
	key := "testRedisLock"
	// 新建锁
	l := lock.New(client, key, nil)
	// 申请锁 默认5s
	l.Lock()

	// 别的Locker拿不到相同key的锁
	l2 := lock.New(client, key, nil)
	fmt.Println(l2.Lock()) // false <nil>

	// 自己可以刷新
	fmt.Println(l.Lock()) // true <nil>

	// 释放后别的Locker可以拿到锁
	l.Unlock()
	l3, _ := lock.Obtain(client, key, nil)
	fmt.Println("Locker", l3)

	// 等待锁自动释放
	time.Sleep(time.Second * 6)
	lock.Run(client, key, nil, run)

}

func run() error {
	fmt.Println("GetLock")

	return nil
}
