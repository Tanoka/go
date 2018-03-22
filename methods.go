package main

import "fmt"
import "math"

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x, y float64
}

func (c *Circle) area() float64 {  //could be (c Circle) but then we couldn't change properties of c.
	return  math.Pi * 2 * c.r
}

func (r *Rectangle) area() float64 {
	return r.x * r.y
}

type Shape interface {
	area() float64
}

//https://medium.com/@agileseeker/go-interfaces-pointers-4d1d98d5c9c6
func area(shape Shape) float64 { //Shape is an interface, can't *shape  - shape.area undefined (type *Shape is pointer to interface, not interface)
	return shape.area()
}


func main() {
	c1 := Circle{1,1,2}

	r1 := new(Rectangle)
	r1.x = 2
	r1.y = 2

	var r2 Rectangle
	r2.x = 3
	r2.y = 2.2

	fmt.Println(c1.area(), r1.area(), r2.area())
	fmt.Println(area(&c1), area(r1), area(&r2))  //if area(c1): cannot use c1 (type Circle) as type Shape in argument to area: 	Circle does not implement Shape (area method has pointer receiver)
}
