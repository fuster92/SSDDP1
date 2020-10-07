package main

import (
	"../utils"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

// Args
// Arg 1 -> Port to listen
// Arg 2 -> Master's Host
// Arg 3 -> Master's Port

func main() {
	port := os.Args[1]
	listener, err := net.Listen(utils.ConnectionType, ":"+port)
	utils.CheckError(err)

	conn, err := listener.Accept()
	for {
		if err != nil {
			continue
		}
		fmt.Printf("Job Accepted\n")
		decoder := gob.NewDecoder(conn)
		var n int
		err := decoder.Decode(&n)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, err.Error()+"\n")
		}
		fmt.Printf("\tCalc Petition -> %d\n", n)

		primes := utils.FindPrimes(n)
		fmt.Printf("\tSending primes\n")
		utils.SendPrimes(conn, primes)
	}
}
