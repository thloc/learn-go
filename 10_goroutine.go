package main

import (
	"fmt"
	"sync"
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

	fmt.Println("---------- wait group ------------")
	var wg = sync.WaitGroup{}
	wg.Add(2) // cho 2 goroutine

	go func() {
		count("Sheep")
		wg.Done()
	}()

	go func() {
		count("Fish")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Done")
}

// wait group
func count(name string) {
	for i := 1; i < 5; i++ {
		fmt.Println(name, i)
		time.Sleep(time.Second)
	}
}
