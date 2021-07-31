package main

import (
	"builder"
	"fmt"
)

func main() {
	normalBuilder := builder.GetBuilder("normal")
	iglooBuilder := builder.GetBuilder("igloo")

	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal Housue Door Type: %s\n", normalHouse.GetDoorType())
	fmt.Printf("Normal Housue Window Type: %s\n", normalHouse.GetWindowType())
	fmt.Printf("Normal Housue Num Floors: %s\n", normalHouse.GetNumFloor())

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.GetHouse()
	fmt.Println("Igloo House Door Type: %s\n", iglooHouse.GetDoorType())
	fmt.Println("Igloo House Window Type: %s\n", iglooHouse.GetWindowType())
	fmt.Println("Igloo House Num Floor: %s\n", iglooHouse.GetNumFloor())
}
