package main

import (
	"bytes"
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
// ------ thong qua struct ---------
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

	fmt.Println("---------- ko thong qua struct ------------")
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 1; i < 5; i++ {
		fmt.Println(inc.Increment())
	}

	fmt.Println("---------- Nhieu interface ------------")
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello world"))
	wc.Close()
}

// -------------------------------
// --- ko thong qua struct ----
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// ------------- Nhieu interface ------------------
type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)

	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)

	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)

		if err != nil {
			return 0, err
		}

		_, err = fmt.Println(string(v))

		if err != nil {
			return 0, err
		}
	}

	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))

		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
