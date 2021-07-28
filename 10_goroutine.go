package main

import (
	"fmt"
	"time"
)

// goroutine la 1 function khong nam main thread. Chay rieng cac tac vu gan nhu dong thoi.
// -> go f(x, y, z) // chay goroutine khac. 1 Thread quan ly nhieu Goroutine. nen G nhe hon.
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world") // chay dong thoi voi say("hello")
	say("hello")
}
