package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func main() {
	fmt.Printf("Starting server\n")
	listener, err := net.Listen(TYPE, ":"+PORT)
	checkError(err)

	petitionId := 0
	fmt.Printf("Accepting petitions on port %s\n", PORT)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		request := receiveRequest(conn)
		if err != nil {
			continue
		}
		go requestHandler(Job{conn, request})
		petitionId++
		fmt.Printf("[%d] Request from %s queued\n", petitionId, conn.RemoteAddr().String())
	}
}

// Gets a request from th buffer and processes it
func requestHandler(job Job) {
    primes := FindPrimes(job.request.Prime)
		sendPrimes(job.connection, primes)
		err := job.connection.Close()
		if err != nil {
			printError(err)
		}
}

// Sends the primes
func sendPrimes(connection net.Conn, primes []int) {
	encoder := gob.NewEncoder(connection)
	err := encoder.Encode(primes)
	printError(err)
}

// Receives data from a TCP connection
func receiveRequest(connection net.Conn) Request {
	var data Request
	decoder := gob.NewDecoder(connection)
	error := decoder.Decode(&data)
	printError(error)
	return data
}
