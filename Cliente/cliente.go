// Creates a client that ask for a list of prime numbers to a server
package main

import (
	"../utils"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	var requestsPerSecond int
	var maxNumberRequest int
	pretty := false

	switch len(os.Args) {
	case 1, 2:
		fmt.Printf("Usage %s req/s [maxNumberRequest] [pretty printing]", os.Args[0])
		os.Exit(0)
	case 3:
		requestsPerSecond, _ = strconv.Atoi(os.Args[1])
		maxNumberRequest, _ = strconv.Atoi(os.Args[2])
	case 4:
		requestsPerSecond, _ = strconv.Atoi(os.Args[1])
		maxNumberRequest, _ = strconv.Atoi(os.Args[2])
		pretty = true
	}

	millisBetweenRequest := 1000 / int64(requestsPerSecond)

	for i := 0; i < maxNumberRequest; i++ {
		go makeRequest(i, requestsPerSecond, pretty)
		time.Sleep(time.Millisecond * time.Duration(millisBetweenRequest))
	}
}

// Makes a request to the remote server.
func makeRequest(counter int, petSec int, pretty bool) {
	start := time.Now()
	conn := connect()
	err := sendRequest(conn, utils.Request{ID: counter, Prime: utils.SIZE})
	if err == nil {
		utils.ReadPrimes(conn)
		elapsed := time.Now().Sub(start)
		if pretty {
			prettyPrint(elapsed)
		} else {
			fmt.Printf("%d,%d\n", elapsed.Milliseconds(), petSec)
		}
	}
}

// Gives a readable format a prints to std out.
func prettyPrint(elapsed time.Duration) {
	fmt.Printf("-------------------------\n"+
		"Time: %d ms\n", elapsed.Milliseconds())
	fmt.Printf("Overhead: %d ms\n", elapsed.Milliseconds()-utils.QOS)
	if elapsed.Milliseconds() > utils.QOS*2 {
		fmt.Printf("Mal QOS\n")
	}
}

// Serializes a Request
func sendRequest(conn net.Conn, request utils.Request) error {
	encoder := gob.NewEncoder(conn)
	return encoder.Encode(request)
}

// Connects to the remote host.
func connect() net.Conn {
	conn, err := net.Dial(utils.CONNECTION_TYPE, "localhost"+":"+utils.SERVER_PORT)
	utils.CheckError(err)
	return conn
}
