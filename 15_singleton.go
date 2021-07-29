package main

import "fmt"

// Signleton interface
type Signleton interface {
	AddOne() int
}

type signleton struct {
	count int
}

var instance *signleton

func init() {
	instance = &signleton{count: 100}
}

/*
*	GetInstance()
* return object
 */
func GetInstance() Signleton {
	return instance
}

func (s *signleton) AddOne() int {
	s.count++
	return s.count
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1.AddOne())
	fmt.Println(s2.AddOne())

	// So giong nhau do tro ve cung bo nho.
	fmt.Println(s1)
	fmt.Println(s2)
}
