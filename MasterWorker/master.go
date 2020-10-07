// AUTORES: Javier Fuster Trallero / Javier Herrer Torres
// NIAs: 626901 / 776609
// FICHERO: master.go
// FECHA: 05-oct-2020
// TIEMPO: 1h
// DESCRIPCIÓN: Implementa un servidor master dentro de una arquitectura master-worker

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
	jobsBuffer := make(chan utils.Job, 100)

	// Slice of workers to start
	workers := []utils.Service{{
		Name: "Worker1",
		Host: "",
		Port: "5556",
	}, {
		Name: "Worker2",
		Host: "",
		Port: "5557",
	}, {
		Name: "Worker3",
		Host: "",
		Port: "5558",
	}, {
		Name: "Worker4",
		Host: "",
		Port: "5559",
	}, {
		Name: "Worker5",
		Host: "",
		Port: "5560",
	}}

	initializeHandlerPool(workers, jobsBuffer)

	petitionsReceived := 0
	listener, err := net.Listen(utils.ConnectionType, ":"+utils.ServerPort)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Printf("Accepting petitions on Port %s\n", utils.ServerPort)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		request, err := receivePetition(conn)
		if err != nil {
			continue
		}
		jobsBuffer <- utils.Job{Connection: conn, Request: request}
		petitionsReceived++
	}
}

// Receives a petition from a client
func receivePetition(connection net.Conn) (utils.Request, error) {
	var petition utils.Request
	decoder := gob.NewDecoder(connection)
	err := decoder.Decode(&petition)
	if err != nil {

		return petition, err
	}
	return petition, nil
}

// Initializes a pool of petitionHandler functions
func initializeHandlerPool(workers []utils.Service, buffer chan utils.Job) {
	for _, worker := range workers {
		go petitionHandler(worker, buffer)
	}
}

// Gets a Petition from the buffer and passes it to a worker
func petitionHandler(worker utils.Service, buffer chan utils.Job) {
	var (
		job        utils.Job
		workerConn net.Conn
		err        error
	)
	connected := false

	// Connect with the worker
	for !connected {
		workerConn, err = net.Dial(utils.ConnectionType, worker.Address())
		if err == nil {
			connected = true
		}
	}
	fmt.Printf("Worker %s is listening on %s:%s\n", worker.Name, worker.Host, worker.Port)
	for {
		job = <-buffer
		fmt.Printf("Worker [%s] Assigned Request: %d \n", worker.Name, job.Request.Prime)

		// ----------------------
		primes, err := calcPrimes(workerConn, job.Request.Prime)
		// ----------------------
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, err.Error()+"\n")
			continue
		}
		utils.SendPrimes(job.Connection, primes)
		fmt.Printf("Worker [%s] Sending primes: %d \n", worker.Name, job.Request.Prime)
		if job.Connection.Close() != nil {
			_, _ = fmt.Fprintf(os.Stderr, err.Error()+"\n")
		}

	}
}

// Ask the worker for the calculation
func calcPrimes(conn net.Conn, n int) ([]int, error) {
	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)
	err := encoder.Encode(n)
	if err != nil {
		return nil, err
	}
	var primes []int
	err = decoder.Decode(&primes)
	if err != nil {
		return nil, err
	}
	return primes, err
}
