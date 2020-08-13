package main

import (
	"fmt"
	"time"
)
func worker(id int, jobs <-chan int, results chan int) {
	for j := range jobs {

		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2

	}
}


//func main() {
//	jobs := make(chan int, 100)
//	results := make(chan int, 100)
//	// 开启3个goroutine
//	for w := 1; w <= 5; w++ {
//		go worker(w, jobs, results)
//	}
//	// 5个任务
//	for j := 1; j <= 5; j++ {
//		jobs <- j
//	}
//	close(jobs)
//	// 输出结果
//	var b int
//	for a := 1; a <= 5; a++ {
//		b=<-results
//		fmt.Println(b)
//	}
//
//}
func main() {
	ch := make(chan int,1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}