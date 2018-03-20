package main

import (
	"fmt"
	"os"
)

func main() {
	// args := os.Args[1:]

	var cadena string;

	for x := 1; x < len(os.Args); x++ {
		cadena += " valor: " + os.Args[x]
	}
	fmt.Println(cadena)

}
