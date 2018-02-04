/*
author: zjjmlrs
date: 2018/2/2  
*/
package core

import (
	"fmt"
	"testing"
)

//ch := make(<-chan int) //创建只读channel   用作函数参数
//ch := make(chan<- int) //创建只写channel

func TestDeadLock(t *testing.T) {
	ch := make(chan int) // 无缓冲
	ch <- 1              // 无缓冲直接写入数据会死锁  all goroutines are asleep - deadlock!
	go func() {
		<-ch
	}()

	/*
	写的时候必须有另一个地方在读，不然会阻塞
	读的时候如果没有数据可读也会阻塞
		1. 增加缓冲区  ch :=make(chan int,1),注意缓冲区满了也一样会阻塞
		2. ch <- 1 放后面，先go func 取读取信道，会阻塞在 <-ch上，但是不是在主线程，所以会继续执行写入信道
	这个特性可以利用：由子程序向信道里面写数据，来通知主程序已完成任务，如下：
			go func() {
				// 其他代码
				ch <- 1
			}
			<- ch   // 收到数据，则子线程已跑完  , 替代time.Sleep
	另一个是利用读取管道时阻塞，但是如果管道关闭后读取则返回0(不阻塞)，加上select来实现：
		主程序通过关闭管道，通知子程序该结束运行了，保证资源的释放
	*/
}

func TestAnotherDeadLock(t *testing.T) {
	c, quit := make(chan int), make(chan int) // 解决方法：设置成缓冲的信道
	go func() {
		quit <- 0
		c <- 1
	}()
	<-c // 取走c的数据！
	<-quit

	/*
	go func 中先写入quit，但是主程序中却先读取c
	造成了双方都阻塞着，等待对方写入或读取数据
	*/
}

func TestClosed(t *testing.T) {
	ch := make(chan int, 3) // 有缓冲
	ch <- 1
	close(ch)         //关闭信道, 会禁止数据流入, 是只读的
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 关闭后读到的是0
}

func TestRange(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	close(ch) //关闭信道, 会禁止数据流入, 是只读的
	for v := range ch {
		fmt.Println(v) //死锁，因为range不等到信道关闭是不会结束读取的
	}

	// 如果没有close(ch), 则 all goroutines are asleep - deadlock!
	// 注意用range 读取ch的数据时， 一定要保证ch最终是关闭的

}
