// AUTORES: Javier Fuster Trallero / Javier Herrer Torres
// NIAs: 626901 / 776609
// FICHERO: servidorconcurrente.go
// FECHA: 04-oct-2020
// TIEMPO: 30'
// DESCRIPCIÓN: Implementa un servidor que atiende concurrentemente las peticiones que le llegan mediante la
// creación de goroutines

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
	listener, err := net.Listen(utils.ConnectionType, ":"+utils.ServerPort)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	petitionID := 0
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
		go requestHandler(utils.Job{Connection: conn, Request: request})
		petitionID++
		fmt.Printf("[%d] Request from %s queued\n", petitionID, conn.RemoteAddr().String())
	}
}

// Gets a request from th buffer and processes it
func requestHandler(job utils.Job) {
	primes := utils.FindPrimes(job.Request.Prime)
	sendPrimes(job.Connection, primes)
	err := job.Connection.Close()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error()+"\n")
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
