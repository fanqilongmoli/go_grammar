package main

import (
	"fmt"
	"time"
)

//使用Go的协程和通道实现一个工作池

func worker_worker(id int, jobs <-chan int, results chan <- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker_worker(w,jobs,results)
	}
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 9; a++ {
		<-results
	}
}
