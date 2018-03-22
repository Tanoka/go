package main

import "fmt"

type natural int32
type coord struct {
	x int
	y int
}

func test(a *int, b *int) int {
	*a++
	*b++
	return 	 3
}


func main() {
	var x natural = 4
	punto := coord{1,2}

	punto2 := new(coord)
	punto2.x = 3 //punto2 es un puntero, en este caso (*punto2).x == punto2.x
	punto2.y = 4

	//Se puede pasar solo una propiedad de un struct por referencia si se quiere.
	test(&punto.x, &punto.y)

	fmt.Printf("x %d, punto %v, punto2 memory!:%v", x, punto, punto2);
}
