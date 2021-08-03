package main

import "fmt"

// func main() {
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

// -------------RANGE--------------
// range: lay gia tri Buffer trong C gan cho i.
// dung range thuong cho vong lap.
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // dong channel
}

func main() {
	c := make(chan int, 10) // buffered channel
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

	// Giong tren, nhung phai biet gioi han so la` 10
	// for i := 0; i < 10; i++{
	// 	fmt.Println(i)
	// }
}
