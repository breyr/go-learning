package structs

import "math"

// Cannot overload functions in Go!
// func (recieverName RecieverType) MethodName (args)
// convention to make the recieverName the first letter of its type

// interface resolution is implicit, passing a type that matches what the interface is asking for it will compile
type Shape interface {
	Area() float64
}

type Reactangle struct {
	Width  float64
	Height float64
}

func (r Reactangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Reactangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
