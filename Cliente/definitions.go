package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

const (
	PORT            = "5555"
	CONNECTION_TYPE = "tcp"
	SIZE            = 60000
	QOS	int64 			= 744
)

// Describes a service in a distributed architecture
type Service struct {
	name string
	host string
	port string
}

func (service Service) address() string {
	return service.host + ":" + service.port
}

// Format of the petitions made by the clients
type Petition struct {
	ID    int
	Prime int
}

// Job inside the master process
type Job struct {
	connection net.Conn
	petition   Petition
}

// Sends the primes
func sendPrimes(connection net.Conn, primes []int) {
	encoder := gob.NewEncoder(connection)
	err := encoder.Encode(primes)
	printError(err)
}

// Reads an array of ints
func readPrimes(conn net.Conn) []int{
	var primes []int
	decoder := gob.NewDecoder(conn)
	decoder.Decode(&primes)
	return primes
}

// Checks if there's an error
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

// Prints the error instead of exiting the program
func printError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr,err.Error() + "\n")
	}
}