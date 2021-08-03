package main

import (
	"abstract"
	"fmt"
)

func main() {
	adidasFactory := abstract.GetSportFactory("adidas")
	adidasShoe := adidasFactory.MakeShoe()
	printShoeDetails(adidasShoe)
	adidasShort := adidasFactory.MakeShort()
	printShortDetails(adidasShort)

	nikeFactory := abstract.GetSportFactory("nike")
	nikeShoe := nikeFactory.MakeShoe()
	printShoeDetails(nikeShoe)
	nikeShort := nikeFactory.MakeShort()
	printShortDetails(nikeShort)
}

func printShoeDetails(s abstract.IShoe) {
	fmt.Println("Logo: %s\n", s.GetLogo())
	fmt.Println("Size: %s\n", s.GetSize())
}

func printShortDetails(s abstract.IShort) {
	fmt.Println("Logo: %s\n", s.GetLogo())
	fmt.Println("Size: %s\n", s.GetSize())
}
