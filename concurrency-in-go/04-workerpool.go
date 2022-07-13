package main

import "fmt"

func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 1; i <= 100; i++ {
		jobs <- i
	}

	close(jobs)

	for i := 1; i <= 100; i++ {
		fmt.Println(<-results)
	}

}

func worker(jobs chan int, result chan int) {

	for i := range jobs {
		result <- fib(i)
	}

}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)

}
