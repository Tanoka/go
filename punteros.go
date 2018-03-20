package main

import "fmt"

type coord struct {
       x int
       y int
}


func add(x *int) int {
	*x++
	return *x
}

func dothing(c coord) int  {
	c.x++
	c.y++
	return c.x + c.y
}

func dothingpunt(c *coord) int  {
        c.x++
        c.y++
        return c.x + c.y
}


func main() {
	var x = 4
	var y *int = &x

	// var s *int --> just declaration ..can asign &var, but not *s = 3!.. use new for that
	// var y *int
	// *y = 10 	//DOESN'T WORK ...there's not space allocated!

	z := &x

	add(y)

	add(&x)

	x++

	*z++

	fmt.Printf("x:%d y:%d z:%d \n",x, *y, *z)


	c1 := coord{1,2}
	dothing(c1)
	fmt.Printf("c1:%v", c1)

        c2 := coord{1,2}
        dothingpunt(&c2)
        fmt.Printf("c2:%v", c2)

	c3 := new(coord) //var c3 *coord not works, there's not memory allocated! just to assign other coord type with &coordVar
	c3.x = 1 // == (*c3).x = 1
	c3.y = 2
	dothingpunt(c3)
	fmt.Printf("c3:%v", *c3)

}
