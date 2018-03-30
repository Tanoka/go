//workers that keep results from jobs in global slice calling a goroutine thats add elements with append
//append is only called from 1 place to avoid concurrency issues
//////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"time"
)

type forch struct {
	x int
	xx int
}


func bucle() {

	jobs := 20000

	res := make([]int, 0)           //slice
	ch := make(chan forch)         //job number
	end := make(chan bool, jobs) //need a buffer, reading end is reached after a lot writing in end channel
	chr := make(chan int)        //to result

	//just can add element to res from one routine, this avoid concurrency issues
	go func(chr chan int, end chan bool) {
		for {
			select {

			case r := <-chr:
				res = append(res, r)
				end <- true //jobs done after add element to slice
			case <-time.After(15 * time.Second): //if 1 second without response, all jobs done
				fmt.Println("end populate done--------------")
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
		ch <- forch{x,x}
	}
	fmt.Println("...........................................Waiting...")

// Not works, infinite loop, some issue with gorutine recolector, has another select...
//	z :=0
//	for {
//		select {
//		case <-end:
//			fmt.Println("zzzzzzz ",z)
//			z++
//		case <-time.After(2 * time.Second):
//			fmt.Println("end control done.........")
//			break
//		}
//	}
	//waiting ends
		for x := 0; x < jobs; x++ {
			<-end
//			fmt.Println("* Received end")
		}

	fmt.Println("Result: ", len(res))
//	fmt.Println("Result: ", res)
}

func main() {
	start := time.Now()
	bucle()
	end := time.Since(start)
	fmt.Printf("End. %s \n", end)
}
