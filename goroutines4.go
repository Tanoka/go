// Use timeout and a channel to ensure goroutines are done
//////////////////////////////////////
package main

import (
	"fmt"
	"time"
)

type forch struct {
	x  int
	xx int
}

func bucle() {

	jobs := 20000

	res := make([]int, 0)        //slice
	ch := make(chan forch)       //job number
	end := make(chan bool) //need a buffer, reading end is reached after a lot writing in end channel
	chr := make(chan int)        //to result

	//just can add element to res from one routine, this avoid concurrency issues
	go func(chr chan int, end chan bool) {
		for {
			select {
			case r := <-chr:
				res = append(res, r)
			case <-time.After(1 * time.Second): //if 1 second without response, all jobs done
				fmt.Println("end populate done--------------")
				end <- true
				break
			}
		}
		fmt.Println("saliendo go!!")
	}(chr, end)

	//Workers
	for x := 0; x < 10; x++ {
		go func(ch chan forch, chr chan int, x int) {
			for r := range ch {
				//sql, files...heavy process
				time.Sleep(15 * time.Millisecond)
				chr <- r.xx
//				fmt.Printf("+ job %d done by worker %d \n", r, x)
			}
		}(ch, chr, x)
	}

	//Jobs
	for x := 0; x < jobs; x++ {
		ch <- forch{x, x}
	}
	fmt.Println("...........................................Waiting...")

	//waiting end
	_ = <-end

	fmt.Println("Result: ", len(res))
	//	fmt.Println("Result: ", res)
}

func main() {
	start := time.Now()
	bucle()
	end := time.Since(start)
	fmt.Printf("End. %s \n", end)
}
