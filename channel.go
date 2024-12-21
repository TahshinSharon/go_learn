package main

import "fmt"

func worker(done chan bool) {
	fmt.Println("worker is running")
	done <- true
}
func main() {
	done := make(chan bool)

	go worker(done)
	<-done
	buffer()
	fmt.Println("Worker is finished")
}
func buffer() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
