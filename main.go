package main

import (
	"fmt"
	"time"
)

func main () {

	sequence := 0
	runningTime := time.Now()
	
	for i := 0; i <= 10000; i++ {

		if time.Since(runningTime) > 1000000 { //nanoseconds
			fmt.Println("changing stuff")
			ResetSequenceNumber(&sequence)
			runningTime = time.Now()
		}

		IncrementSequence(&sequence)
		fmt.Println(sequence)
	}

}


func IncrementSequence(iptr *int) {
    *iptr++
}

func ResetSequenceNumber(iptr *int) {
	*iptr = 0
}










