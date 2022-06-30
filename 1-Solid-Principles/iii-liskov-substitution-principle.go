package main

import "fmt"

/*
Liskov Substitution Principle (LSP) states that objects of a superclass(Rectangle)
should be replaceable with objects of its subclasses (Square) without breaking the application.
https://stackoverflow.com/a/584732
https://stackoverflow.com/a/44913313
*/
//Not very rlevant to Go, as there's no inheritence

type Sized interface {
	getHeight() int
	setHeight(height int)
	getWidth() int
	setWidth(width int)
}

type Rectangle struct {
	width, height int
}

//Implementing Sized interface on Rectangle
func (r *Rectangle) getHeight() int {
	return r.height
}
func (r *Rectangle) setHeight(height int) {
	r.height = height
}
func (r *Rectangle) getWidth() int {
	return r.width
}
func (r *Rectangle) setWidth(width int) {
	r.width = width
}

func UseSize(sized Sized) {
	width := sized.getWidth()

	height := sized.getHeight()

	//This line breaks LSP for the square
	sized.setHeight(10)
	expectedArea := height * width
	actualArea := sized.getWidth() * sized.getHeight()

	fmt.Printf("Expexted:%v\nActual:%v\n", expectedArea, actualArea)
}

//Extending Rectangle to include Square
type Square Rectangle

func NewSquare(size int) *Square {
	sq := Square{}
	sq.height = size
	sq.width = size
	return &sq
}

// This breaks LSP, because we have to set both the width and height,
// to keep it a square
func (s *Square) setWidth(width int) {
	s.width = width
	s.height = width
}
func (s *Square) getWidth() int {
	return s.height
}
func (s *Square) setHeight(height int) {
	s.height = height
	s.width = height
}
func (s *Square) getHeight() int {
	return s.height
}

func main() {
	rectangle := &Rectangle{10, 20}
	UseSize(rectangle)

	square := NewSquare(5)
	UseSize(square)

}
