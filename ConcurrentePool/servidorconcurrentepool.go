package main

import (
	"../utils"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Printf("Starting server\n")
	jobsBuffer := make(chan utils.Job, 10)
	listener, err := net.Listen(utils.ConnectionType, ":"+utils.ServerPort)
	utils.CheckError(err)
	initializeGoRoutinePool(jobsBuffer)

	petitionId := 0
	fmt.Printf("Accepting petitions on port %s\n", utils.ServerPort)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		request := receiveRequest(conn)
		if err != nil {
			continue
		}

		jobsBuffer <- utils.Job{Connection: conn, Request: request}
		petitionId++
		fmt.Printf("[%d] Request from %s queued\n", petitionId, conn.RemoteAddr().String())
	}

}

// Initializes a pool of requestHandler functions
func initializeGoRoutinePool(buffer chan utils.Job) {
	for i := 0; i < 3; i++ {
		go requestHandler(i, buffer)
	}
}

// Gets a request from th buffer and processes it
func requestHandler(workerId int, buffer chan utils.Job) {
	var job utils.Job
	for {
		job = <-buffer
		prime := job.Request.Prime
		fmt.Printf("Worker [%d] Serving Request: %d \n", workerId, prime)
		primes := utils.FindPrimes(prime)
		sendPrimes(job.Connection, primes)
		fmt.Printf("Worker [%d] Sending primes: %d \n", workerId, prime)
		err := job.Connection.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, err.Error()+"\n")
		}
	}
}

// Sends the primes
func sendPrimes(connection net.Conn, primes []int) {
	encoder := gob.NewEncoder(connection)
	err := encoder.Encode(primes)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error()+"\n")
	}
}

// Receives data from a TCP connection
func receiveRequest(connection net.Conn) utils.Request {
	var data utils.Request
	decoder := gob.NewDecoder(connection)
	err := decoder.Decode(&data)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error()+"\n")
	}
	return data
}
