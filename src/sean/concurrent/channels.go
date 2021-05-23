package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {

	messages := make(chan string)

	go func() {
		fmt.Println("start sleep")
		time.Sleep(5 * time.Second)
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)

	// 带缓冲队列的channel
	// 下面的demo证明声明的channel是不带缓冲的
	messages2 := make(chan string)
	go func() {
		fmt.Println("send msg1")
		messages2 <- "msg1"
		fmt.Println("send msg2")
		messages2 <- "msg2"
		fmt.Println("send msg done")
	}()

	time.Sleep(5 * time.Second)
	fmt.Println(<-messages2)
	fmt.Println(<-messages2)

	// 缓冲大小为2的channel
	messages3 := make(chan string, 2)

	messages3 <- "buffered"
	messages3 <- "channel"

	fmt.Println(<-messages3)
	fmt.Println(<-messages3)

	// 使用channel来实现多协程同步
	done := make(chan bool, 1)
	go worker(done)
	<-done

	// 关闭channel
	jobs := make(chan int, 5)
	done2 := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done2 <- true
				return
			}
		}
	}()

	for j := 1; j <= 5; j++ {
		jobs <- j
		fmt.Println("send Job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done2

	// 遍历channel
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}

}
