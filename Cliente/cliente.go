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
	var millisBetweenRequest int
	//var maxNumberRequest int
	//var wg sync.WaitGroup
	pretty := false

	switch len(os.Args) {
	case 1:
		fmt.Printf("Usage %s ms [pretty printing]", os.Args[0])
		os.Exit(0)
	case 2:
		millisBetweenRequest, _ = strconv.Atoi(os.Args[1])
		pretty = false
	case 3:
		millisBetweenRequest, _ = strconv.Atoi(os.Args[1])
		//maxNumberRequest, _ = strconv.Atoi(os.Args[2])
		pretty = true
	}

	for i := 0; ; i++ {
		//wg.Add(1)
		go makeRequest(i, millisBetweenRequest, pretty)
		time.Sleep(time.Millisecond * time.Duration(millisBetweenRequest))
	}
	//wg.Wait()
}

// Makes a request to the remote server.
func makeRequest(counter int, petSec int, pretty bool) {
	start := time.Now()
	conn := connect()
	err := sendRequest(conn, utils.Request{ID: counter, Prime: utils.Size})
	if err == nil {
		utils.ReadPrimes(conn)
		elapsed := time.Since(start)
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
	conn, err := net.Dial(utils.ConnectionType, utils.ServerHost+":"+utils.ServerPort)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error())
	}
	return conn
}
