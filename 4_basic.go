package main

import (
	"fmt"
	"math"
	"runtime"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println("----------FOR------------")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum) // 45

	fmt.Println("----------------------")
	sum2 := 1
	for sum2 < 10 {
		sum2 += sum2
	}
	fmt.Println(sum2) // 16

	fmt.Println("----------IF------------")
	fmt.Println(
		pow(3, 2, 10), // 9
		pow(3, 3, 20), // 20
	)

	fmt.Println("----------SWITCH------------")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println(os)
	}

	fmt.Println("----------DEFER------------")
	defer fmt.Println("world")
	fmt.Println("hello")

	// defer chay khi ham ket thuc
	// Mang tinh xep trong len nhau: 1,2,3,4,5 -> 5,4,3,2,1
}
