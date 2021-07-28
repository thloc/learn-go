package main

import (
	"fmt"
)

// select: chon lua nhieu channel.
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}

// ------------vd2----------
// func main() {
// 	c1 := make(chan int)
// 	c2 := make(chan int)

// 	go func() {
// 		time.Sleep(time.Second)
// 		c1 <- 5
// 	}()

// 	go func() {
// 		time.Sleep(time.Second * 2)
// 		c2 <- 10
// 	}()

// 	// // chay tuan tu
// 	// fmt.Println(<-c1)
// 	// fmt.Println(<-c2)

// 	select {
// 	case x := <-c1:
// 		fmt.Println(x)
// 	case y := <-c2:
// 		fmt.Println(y)
// 	}
// }
