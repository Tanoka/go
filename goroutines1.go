//workers that keep results from jobs in global array. 
//////////////////////////////////////////////////////////
//Because array is accesed by id, there's no concurrency problem
//If we use a slice, append() will have concurrency issues, we need to lengh 
//slice in declaration, make([]int, 100) and use splice[pos] like with array,
//after jobs are done, if we have de total jobs done, we can slice[:totaljobs] to get
//the part or slice with the job done.
package main

import (
	"fmt"
	"time"
)
//import "time"

type forch struct {
	x int
	xx int
}

func bucle() {

	jobs  := 20000

	var res [20000]int  //HARCODED!!
	ch := make(chan forch)
	end := make(chan bool, jobs) //need a buffer, reading end is reached after a lot writing in end channel
	chr := make(chan forch)        //to result

	go func(chr chan forch, end chan bool) {
		for {
			select {

			case stru := <-chr:
				res[stru.x] = stru.xx
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
		go func(ch chan forch, chr chan forch, x int) {
			for stru := range ch {
				//sql, files...heavy process
				time.Sleep(15 * time.Millisecond)
			 	chr <- stru
//				fmt.Printf("+ job %d done by worker %d \n",stru.x, x)
//				fmt.Println("++ Sending end..")
			}
		}(ch, chr, x)
	}

	//Jobs
	for x := 0; x < jobs; x++ {
//		fmt.Println("- Sender ", x)
		ch <- forch{x, x}
	}
	fmt.Println("...........................................Waiting...")

	//waiting ends
	for x := 0; x < jobs; x++ {
		<-end
//		fmt.Println("* Received end")
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

