package main

import (
	"fmt"
	"math/rand"
	"time"
)

// tra ve 1 channel
// <- chan: dung lay ra, khong them vao dc
// chan <-: dung them vao, khong lay ra dc
// func startSender(name string) <-chan string {
// 	c := make(chan string)
// 	go func() {
// 		for i := 1; i <= 5; i++ {
// 			c <- (name + " hello")
// 			time.Sleep(time.Second)
// 		}
// 	}()

// 	return c
// }

// func main() {
// 	sender := startSender("A")
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println(<-sender)
// 	}
// }

// -------------- C2 ---------
// chay 2 channel
// func startSender(name string) <-chan string {
// 	c := make(chan string)
// 	go func() {
// 		for i := 1; i <= 5; i++ {
// 			c <- (name + " hello")
// 			time.Sleep(time.Second)
// 		}
// 	}()

// 	return c
// }

// func main() {
// 	a := startSender("A")
// 	b := startSender("B")

// 	for {
// 		select {
// 		case msgA := <-a:
// 			fmt.Println(msgA)
// 		case msgB := <-b:
// 			fmt.Println(msgB)
// 		}
// 	}
// }

// -------------- C3* ---------
// Search nhieu services
// func fetchAPI(model string) string {
// 	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Microsecond)
// 	return model
// }

// func main() {
// 	responseChan := make(chan string)
// 	var results []string

// 	go func() { responseChan <- fetchAPI("users") }()
// 	go func() { responseChan <- fetchAPI("categories") }()
// 	go func() { responseChan <- fetchAPI("productions") }()

// 	for i := 1; i <= 3; i++ {
// 		results = append(results, <-responseChan)
// 	}

// 	fmt.Println(results)
// }

// -------------- C4 ---------
func query(url string) string {
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Microsecond)
	return url
}

// lay server nao tra ve nhanh nhat lay. con lai ko lay
func queryFirst(services ...string) <-chan string {
	c := make(chan string)
	for _, serv := range services {
		go func(s string) { c <- query(s) }(serv)
	}
	return c
}

func main() {
	results := queryFirst("Server 1", "Server 2", "Server 3")
	fmt.Println(<-results)
}
