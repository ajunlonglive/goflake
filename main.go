package main

import (
	"fmt"
	"github.com/hart87/goflake/generator"
)


func main () {

	for i := 0; i <= 100; i++ {
		fmt.Println(generator.GenerateIdentifier())
	}

}













