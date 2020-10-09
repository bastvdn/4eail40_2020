package main

import ("fmt"
		"math"
)

// Implement types Rectangle, Circle and Shape
type Shape interface {
	Area() float64

}

type Circle struct {
	r float64
}

type Rectangle struct {
	Width float64
	Length float64

}

func (r Rectangle) Area() float64{

	return r.Width * r.Length
}

func (c Circle) Area() float64{
	

}

func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
		// Expected output:
		// Square of width 2 and length 3
		// Circle of radius 5
	}
}
