package main

import "fmt"
import "math/rand"

func consume(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch) // Blocks until there's something to read in ch
	}
}

func main() {
	ch := make(chan int)
	go consume(ch) // Runs in background

	// Produce 10 random numbers
	for i := 0; i < 10; i++ {
		ch <- rand.Intn(10)
	}
}
