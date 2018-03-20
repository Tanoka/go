package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"bufio"
	"io"
)

func main() {
	//method 1
	dat, err := ioutil.ReadFile("/usr/test/helloworld.go")
        if err != nil {
		panic(err)
        }
	fmt.Print(string(dat)) //dat are bytes..translate to string.. utf8 and these stuff


	//method 2
	file, err := os.Open("usr/test/helloworld.go")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(file)
	fmt.Print(b) //print bytes!!
	fmt.Print(string(b))
	file.Close()


	//method 3
        file3, err := os.Open("usr/test/helloworld.go")
        if err != nil {
                panic(err)
        }
	scanner := bufio.NewScanner(file3)
        for scanner.Scan() {             // internally, it advances token based on sperator
               fmt.Println(scanner.Text())  // token in unicode-char
               fmt.Println(scanner.Bytes()) // token in bytes
        }
	file3.Close()


	//method 4 ..long files, etc..
	fmt.Println("Metodo 4")
	file4, err := os.Open("/usr/test/helloworld.go")
        if err != nil {
	        panic(err)
	}
	defer file4.Close()
        buf := make([]byte, 4) // define your buffer size here. 4 bytes each time...
        for {
	        n, err := file4.Read(buf)
	        if n > 0 {
			fmt.Print(buf[:n]) // your read buffer.
			fmt.Print(string(buf[:n]))
			fmt.Println("")
                }
                if err == io.EOF {
			fmt.Printf("End of file")
	                break
	        }
	        if err != nil {
			fmt.Printf("Error reading")
	                break
	        }
	}

}



