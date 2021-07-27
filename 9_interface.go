package main

import (
	"fmt"
	// "math"
)

// type Abser interface {
// 	// interface chi la interface co ham Abs(). Tuc la ca struct nao co method Abs() tra ve float thi goi la Abser
// 	Abs() float64
// }

// func main() {
// 	var a Abser
// 	f := MyFloat(-math.Sqrt2)
// 	v := Vertex{3, 4}

// 	a = f  // a MyFloat implements Abser
// 	a = &v // a *Vertex implements Abser

// 	// In the following line, v is a Vertex (not *Vertex)
// 	// and does NOT implement Abser.
// 	// a = v // -> err

// 	fmt.Println(a.Abs())
// }

// type MyFloat float64 // tu goi la Abser

// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// type Vertex struct {
// 	X, Y float64
// }

// func (v *Vertex) Abs() float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// ---------------------------------
type I interface {
	M()
}

type T struct {
	S string
}

// T implements the interface I
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
