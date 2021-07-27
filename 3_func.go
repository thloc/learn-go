package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func slipt(sum int) (x, y int) {
	x = sum * 1 / 2
	y = sum - x
	return
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println("----------------------")

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println("----------------------")
	fmt.Println(slipt(15))

	fmt.Println("----------------------")
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
