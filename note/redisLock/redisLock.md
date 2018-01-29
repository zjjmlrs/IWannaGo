### redis-lock

#### 用法及功能介绍
* todo

#### 基础了解
* ```import "github.com/bsm/redis-lock"```
* 需要 redis client ```import "github.com/go-redis/redis"```
* todo

#### 源码分析
###### 结构体定义
- Locker
![Locker](Locker.png)
 > 显然，Locker需要一个 redis client 来进行加锁，使用的是redis的```SetNX```命令
 >
 > key 是加锁用的字符串，相同字符串即请求相同的锁，会相互排斥，可以理解为锁的名字， 注意此处的相互排斥指的是不同的Locker之间，同一个Locker取申请锁，则会刷新这个锁
 >
 > opts 是锁的相关设置
 >
 > token 相当于一个标志，如果是空字符串则表示该Locker未加锁，非字符串表示已加锁，但是此锁可能已过期，同时根据这个字段判断申请锁的时候是要刷新还是创造
 >
 > mutex 暂时不是很了解这个模块，感觉是对代码加锁，防止多个线程同时操作
 >
 
- RedisClient
![RedisClient](RedisClient.png)
 > ```SetNx```命令是加锁最核心的命令，正是由于这个命令，才有的redis-lock
 >
 > 接下来四个方法不是很懂，程序中好像也没用到，不知道SetNx方法中有没有用到 todo 
 >
 > 注意，这是一个接口，那是不是意味着只要client实现了这几个接口就可以使用redis-lock，不一定非要要redis？

- Opts 
![Options](Options.png)
> 图片中的注释很清楚了，但是它还有一个```normalize```方法实现默认赋值，这可以学习
![normalize](normalize.png)

###### 程序流程
*  ```New(client RedisClient, key string, opts *Options) *Locker```
    > 初始化锁，如无Opts,设置默认配置
* ```(l *Locker) Lock() (bool, error)```
    > 进行加锁 使用默认context， context 可以在任何时刻结束阻塞等待锁的线程,这里使用默认context预示着我们不打算自己去结束阻塞的线程，只能等到配置的等待时间到达
* ```(l *Locker) LockWithContext(ctx context.Context) (bool, error)```
    > 使用自己的context加锁，可以在任何时刻主动结束线程，注意，返回值bool指示加锁是否成功
* ```(l *Locker) Unlock() error```
    > 解锁
* ```(l *Locker) IsLocked() bool```
    > 查看是否被解锁，此处是通过判断token是否为空字符串来判断的，注意token为非空时，说明未被解锁，但是可能此锁已过期
* ```Obtain(client RedisClient, key string, opts *Options) (*Locker, error)```
    > 初始化锁，并加锁
* ```Run(client RedisClient, key string, opts *Options, handler func() error)```
    > 初始化锁并加锁，成功后直接执行传入的方法
##### 总结 
#### SetNx 详解
* todo 


#### 其他学习点
* mutex
* context
* time.Timer
* redis.NewScript -- lua 脚本







