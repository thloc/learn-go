package main

import "fmt"

func main() {
	var x float64
	x = 20.0
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)

	fmt.Println("----------------------")
	var a float64 = 20.0
	b := 42
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("x is of type %T\n", a)
	fmt.Printf("y is of type %T\n", b)

	fmt.Println("----------------------")
	k := 5
	fmt.Printf("k is of type %T\n", k)

	fmt.Println("----------------------")
	const Pi = 3.14 // const Pi := 3.14 ---> sai (:=)

	fmt.Println("----------------------")
	Big = 1 << 100 // mang 1 dich qua trai 100 bit
	// 1 << 2 = 100 (bit): tang them 00 la gia tri tang gap doi
	// 4 >> 2 = 0 (bit)
}

// Golang ko có public, private & static. Nếu muốn export phải viết Hoa
/* Kieu du lieu
 * bool
 * string -> dau ""
 * int int8 int16 int32 int64: Am -> Duong
 * unit unit8 unit16 uint32 uint64: 0 -> duong
 * byte // alias for uint8
 * rune // alias for int32 -> Dung trong lap ky tu trong chuoi
 * float32 float64
 */
