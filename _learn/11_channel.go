package main

import "fmt"

// Xuat hien khi co 2 hoac nhieu Goroutine trao doi su lieu voi nhau
// chay dong thoi. 1 input vao thi phai co 1 output ra neu khong bi block
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	// make(chan int): -> kieu chap nhan la "int", "c" la 1 channel
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x*y)
}

// VD2:
// func main() {
// 	// G1
// 	fmt.Println("Hello, playground")

// 	c := make(chan int)
// 	// c <- 1 // -> loi~

// 	// Fix
// 	go func() {
// 		// G khac. G2
// 		time.Sleep(time.Second * 3)
// 		c <- 1
// 	}() // Khai bao va chay luon ()

// 	data := <-c

// 	fmt.Println(data)
// }
