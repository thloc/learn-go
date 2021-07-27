package main

import "fmt"

type Vertex struct {
	X int // int = 4 byte
	Y int // int = 4 byte
} // La 1 kieu du lieu moi. Cau truc du lieu nay se nam trong 1 o du lieu, do lon 8 byte

func main() {
	// Pointers:
	i, j := 42, 250
	// k = i // k copy value from i. K change value, i not change

	p := &i
	fmt.Println(*p) // read i thong qua pointer
	*p = 21         // set i la 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	fmt.Println("----------STRUCT------------")
	a := Vertex{1, 2}
	b := a // b copy value from a

	b.X = 10
	fmt.Println(a.X, a.Y)

	fmt.Println("----------------------")
	m := Vertex{X: 1, Y: 2}
	fmt.Println(m)

	fmt.Println("----------------------")
	v := Vertex{1, 2}
	n := &v
	n.X = 1e9
	fmt.Println(v)
}
