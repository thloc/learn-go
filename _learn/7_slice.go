package main

import "fmt"

// Array la 1 mang co dinh. Ko the them or bot thanh phan di dc
// Slice ban chat la Cau Truc Du Lieu

func main() {
	primes := [6]int{1, 2, 3, 4, 5, 6}

	var s []int = primes[:4] // slice la con tro -> mang

	fmt.Println(s) // 1 2 3 4
	s[0] = 200
	fmt.Println(primes) // 1 200 3 4 5 6

	fmt.Println("----------SLICE LENGTH------------")
	t := []int{1, 2, 3, 4, 5}
	printSlices(t)

	t = t[:0]
	printSlices(t)

	arr := make([]int, 10, 20)
	arr = append(arr, 1)
	printSlices(arr)
}

func printSlices(s []int) {
	// len: la do dai hien tai
	// cap: la do dai Toi Da
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
