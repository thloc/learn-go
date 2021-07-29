package main

import "fmt"

func main() {
	var a [2]string // la mang String va co 2 phan tu
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes)

	test := [...]int{7, 8, 9}
	fmt.Println(test)
	// test[4] = 6 // -> error
}
