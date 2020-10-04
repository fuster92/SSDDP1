package main

import (
	"fmt"
	"time"
)

func main() {
	for  {
		for i := 0; ; i++{
			//Blocks
			if i % 10 == 0 {
				fmt.Printf("%d\n",i)
				i = 0
				time.Sleep(time.Millisecond * 500)
			}
		}
	}
}
