package main
//import "runtime"

func wait(ch chan int) {
	var i int64
	for ;i < 1e10; i++ {} // Just count to 10 billion
	ch <- 1
}

func main() {
	//runtime.GOMAXPROCS(4) // Use 4 cpus
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ { // Start 10 go routines
		go wait(ch)
	}

	// Wait for everyone to finish
	for i := 0; i < 10; i++ {
		<- ch
	}
}