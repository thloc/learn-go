package main

import (
	"fmt"
	"sync"
	"time"
)

// Data racing bi khi cung write 1 du lieu.
// "new" day tuong duong &[la con tro].
// Lock() se block tat ca cac goroutine den sau cho toi khi Unlock().
// RWLock: block tat ca G con lai du dang la Read hay Write.
// RLock: block tat ca G write, cho phep cac G Read dc phep truy xuat (Shared lock).

// Channel la kenh giao tiep nhieu hon G. Day du lieu tu G qua G kia.
// Lock share bien o ngoai va 2 G dung 1 bien do va thay doi
func main() {
	var count int = 0
	lock := new(sync.RWMutex)

	for i := 1; i < 5; i++ {
		go func() {
			for j := 1; j < 1000; j++ {
				lock.Lock()
				count += 1
				lock.Unlock()
			}
		}()
	}

	time.Sleep(time.Second * 7)
	fmt.Println("Count:", count)
}
