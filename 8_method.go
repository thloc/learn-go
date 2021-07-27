package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 { // ham cua struct
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//----------
type MyFloat float64 // MyFloat is alias of float64

func (f MyFloat) Abs2() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	fmt.Println("----------------------")
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs2())
}
