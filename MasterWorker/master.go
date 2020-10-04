package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os/exec"
	"time"
)

func main() {
	fmt.Printf("Starting server\n")
	jobsBuffer := make(chan Job, 100)

	// Slice of workers to start
	workers := []Service{{
		name: "Worker1",
		host: "",
		port: "5556",
	}, {
		name: "Worker2",
		host: "",
		port: "5557",
	}, {
		name: "Worker3",
		host: "",
		port: "5558",
	}, {
		name: "Worker4",
		host: "",
		port: "5559",
	}, {
		name: "Worker5",
		host: "",
		port: "5560",
	}}
	initializeWorkers(workers)
	initializeHandlerPool(workers, jobsBuffer)

	petitionsReceived := 0
	listener, err := net.Listen(CONNECTION_TYPE, ":"+PORT)
	checkError(err)
	fmt.Printf("Accepting petitions on port %s\n", PORT)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		request := receivePetition(conn)
		if err != nil {
			continue
		}
		jobsBuffer <- Job{conn, request}
		petitionsReceived++
	}
}

// Counts 10 petitions and returns an estimated rate
func petitionsPerSecond(notifier chan byte, out chan float64) {
	for {
		//Blocks
		<-notifier
		start := time.Now()
		for i := 0; i < 9; i++ {
			<-notifier
		}
		elapsed := time.Now().Sub(start)
		out <- elapsed.Seconds() / 10.0
	}
}

// Starts the workers processes
func initializeWorkers(workers []Service) {
	for _, worker := range workers {
		startWorker(worker.port)
	}
}

// Executes the worker's binary
func startWorker(port string) {
	command := exec.Command("./worker", port)
	err := command.Start()
	command.Stdout = nil
	command.Stderr = nil
	checkError(err)
}

// Initializes a pool of petitionHandler functions
func initializeHandlerPool(workers []Service, buffer chan Job) {
	for _, worker := range workers {
		go petitionHandler(worker, buffer)
	}
}

// Gets a petition from the buffer and passes it to a worker
func petitionHandler(worker Service, buffer chan Job) {
	var job Job
	var workerConn net.Conn
	var err error
	connected := false
	// Connect with the worker
	for !connected {
		workerConn, err = net.Dial(CONNECTION_TYPE, worker.address())
		if err == nil {
			connected = true
		}
	}
	fmt.Printf("Worker %s is listening on %s:%s\n", worker.name, worker.host, worker.port)
	for {
		job = <-buffer
		fmt.Printf("Worker [%s] Serving Petition: %d \n", worker.name, job.petition.Prime)

		// ----------------------
		primes := calcPrimes(workerConn, job.petition.Prime)
		// ----------------------

		sendPrimes(job.connection, primes)
		fmt.Printf("Worker [%s] Sending primes: %d \n", worker.name, job.petition.Prime)
		err := job.connection.Close()
		if err != nil {
			printError(err)
		}
	}
}

// Ask the worker for the calculation
func calcPrimes(conn net.Conn, n int) []int {
	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)
	encoder.Encode(n)
	var primes []int
	err := decoder.Decode(&primes)
	checkError(err)
	return primes
}

// Receives a petition from a client
func receivePetition(connection net.Conn) Petition {
	var petition Petition
	decoder := gob.NewDecoder(connection)
	err := decoder.Decode(&petition)
	if err != nil {
		printError(err)

	}
	return petition
}
