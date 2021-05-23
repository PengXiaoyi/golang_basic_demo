package main

import (
	"fmt"
	"time"
)

// 限速是控制资源利用率和保持服务质量的重要机制。 Go可以通过goroutine，通道和代码来优雅地支持速率限制。
func main() {
	// demo1
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	//通过time.Tick来限制range的执行速度
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	//demo2 类似漏桶模式的限流
	fmt.Println("demo2 类似漏桶模式的限流")
	burstyLimiter := make(chan time.Time, 3)

	// 事先往里面放三个令牌
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 相当于简单的漏桶，每隔200ms就往burstyLimiter里面放一个令牌
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		// 从burstyLimiter的桶里取令牌，如果取完了会被阻塞
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}