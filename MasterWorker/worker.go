package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

// Args
// Arg 1 -> Port to listen
// Arg 2 -> Master's host
// Arg 3 -> Master's port

func main() {
	port := os.Args[1]
	listener, err := net.Listen(CONNECTION_TYPE, ":"+port)
	checkError(err)

	conn, err := listener.Accept()
	for {
		if err != nil {
			continue
		}
		fmt.Printf("Job Accepted\n")
		decoder := gob.NewDecoder(conn)
		var n int
		decoder.Decode(&n)
		fmt.Printf("\tCalc petition -> %d\n", n)

		primes := FindPrimes(n)
		fmt.Printf("\tSending primes\n")
		sendPrimes(conn, primes)
	}
}
