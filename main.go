package main

import (
	"fmt"
	"time"
)

//Get epoch time since 01/01/2021 @ 00:00:01 hours
//get MAC Address and convert last 2 to Decimal
//Add the checkdigit
//Add the Service Number
//Return the ID as a string or type of int ?


func main () {

	sequence := 0
	runningTime := time.Now()
	
	for i := 0; i <= 100000; i++ {

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










