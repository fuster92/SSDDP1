package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i:=0 ; i < 10; i++ {
		FindPrimes(60000)
	}
	elapsed := time.Now().Sub(start)
	fmt.Print(elapsed.Milliseconds()/10)
}
