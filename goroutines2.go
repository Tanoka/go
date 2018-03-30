//workers that keep results from jobs in global slice
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

	res := make([]int, jobs) //slice instead of array..but we need to know size 
	ch := make(chan forch)
	end := make(chan bool, jobs) //need a buffer, reading end is reached after a lot writing in end channel

	//Workers
	for x := 0; x < 10; x++ {
		go func(ch chan forch, bo chan bool, x int) {
			for stru := range ch {
				//sql, files...heavy process
				time.Sleep(15 * time.Millisecond)
				res[stru.x] = stru.xx
	//			fmt.Printf("+ job %d done by worker %d \n",stru.x, x)
				bo <- true  //to control end of goroutines
//				fmt.Println("++ Sending end..")
			}
		}(ch, end, x)
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

