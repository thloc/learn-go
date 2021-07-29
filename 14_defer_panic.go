package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start")
	panicker()
	fmt.Println("End")
}

func panicker() {
	fmt.Println("Create panic")
	defer func() {
		if err := recover(); err != nil { // goi panic
			fmt.Println("Error: ", err)
		}
	}()
	panic("Chia cho mot so la 0")
	fmt.Println("End panic")
}
