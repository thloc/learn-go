package main

import (
	"fmt"
	"reflect"
)

type Vertex struct {
	X int // int = 4 byte
	Y int // int = 4 byte
} // La 1 kieu du lieu moi. Cau truc du lieu nay se nam trong 1 o du lieu, do lon 8 byte

type People struct {
	name string
	age  int
}

type Student struct {
	People
	number  int
	subject []string
}

// Tag
type Demo struct {
	number int `Maximum can't over 10`
}

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

	fmt.Println("----------------------")
	student := Student{}
	student.name = "A"
	student.age = 19
	student.number = 8
	student.subject = []string{"Math", "English", "Computer"}
	fmt.Println(student)

	fmt.Println("----------------------")
	demo := Demo{}
	demo.number = 6
	t := reflect.TypeOf(demo)
	field, _ := t.FieldByName("number")
	fmt.Println(field)
}
