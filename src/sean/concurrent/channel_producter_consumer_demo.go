package main

import (
	"fmt"
	"time"
)

func worker2(id int, jobs <-chan int, results chan<- int) {
	// range会一直持续阻塞到其它协程close掉这个channel
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(3 * time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 启动3个goroutine完成任务
	for w := 1; w <= 3; w++ {
		go worker2(w, jobs, results)
	}

	// 发布任务
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 接受结果
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
