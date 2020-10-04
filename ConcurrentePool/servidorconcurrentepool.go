package main

import (
	"encoding/gob"
	"fmt"
	"net"
)


func main() {
	fmt.Printf("Starting server\n")
	jobsBuffer := make(chan Job, 10)
	listener, err := net.Listen(TYPE,":" + PORT)
	checkError(err)
	initializeGoRoutinePool(jobsBuffer)

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
		jobsBuffer <- Job{conn, request}
		petitionId++
		fmt.Printf("[%d] Request from %s queued\n", petitionId, conn.RemoteAddr().String())
	}
}

// Initializes a pool of requestHandler functions
func initializeGoRoutinePool(buffer chan Job) {
	for i := 0; i < 3; i++ {
		go requestHandler(i, buffer)
	}
}

// Gets a request from th buffer and processes it
func requestHandler(workerId int, buffer chan Job) {
	var job Job
	for {
		job  = <- buffer
		fmt.Printf("Worker [%d] Serving Request: %d \n", workerId, job.request.Prime)
		primes := FindPrimes(job.request.Prime)
		sendPrimes(job.connection, primes)
		fmt.Printf("Worker [%d] Sending primes: %d \n", workerId, job.request.Prime)
		err := job.connection.Close()
		if err != nil {
			printError(err)
		}
	}
}

// Sends the primes
func sendPrimes(connection net.Conn, primes []int) {
	encoder := gob.NewEncoder(connection)
	err := encoder.Encode(primes)
	printError(err)
}

// Receives data from a TCP connection
func receiveRequest(connection net.Conn) Request{
	var data Request
	decoder := gob.NewDecoder(connection)
	error := decoder.Decode(&data)
	printError(error)
	return data
}

