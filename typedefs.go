package main

import "fmt"

type natural int32
type coord struct {
	x int
	y int
}


func main() {
	var x natural = 4
	punto := coord{1,2}

	punto2 := new(coord)
	punto2.x = 3 //punto2 es un puntero, en este caso (*punto2).x == punto2.x
	punto2.y = 4

	fmt.Printf("x %d, punto %v, punto2 memory!:%v", x, punto, punto2);
}
