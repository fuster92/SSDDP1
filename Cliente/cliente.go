package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	var counter int
	var petPerSec int64

	var petitions int
	var maxNumberRequest int
	pretty := false
	switch len(os.Args) {
	case 1,2:
		fmt.Printf("Usage %s req/s [maxNumberRequest] [pretty]", os.Args[0])
		os.Exit(0)
	case 3:
		petitions, _ = strconv.Atoi(os.Args[1])
		maxNumberRequest, _ = strconv.Atoi(os.Args[2])
	case 4:
		petitions, _ = strconv.Atoi(os.Args[1])
		maxNumberRequest, _ = strconv.Atoi(os.Args[2])
		pretty = true
	}

	petPerSec = 1000 / int64(petitions)

	for i := 0; i < maxNumberRequest; i++ {
		go makePetition(counter, petitions, pretty)
		time.Sleep(time.Millisecond * time.Duration(petPerSec))
	}
}

func makePetition(counter int, petSec int, pretty bool) {
	start := time.Now()
	conn := connect()
	sendRequest(conn, Petition{int(counter), SIZE})

	readPrimes(conn)
	elapsed := time.Now().Sub(start)
	if pretty {
		prettyPrint(elapsed)
	} else {
		fmt.Printf("%d,%d\n", elapsed.Milliseconds(), petSec)
	}
}

func prettyPrint(elapsed time.Duration) {
	fmt.Printf("-------------------------\n"+
		"Time: %d ms\n", elapsed.Milliseconds())
	fmt.Printf("Overhead: %d ms\n", elapsed.Milliseconds()-QOS)
	if elapsed.Milliseconds() > QOS*2 {
		fmt.Printf("Mal QOS\n")
	}
}

// Formats an integer as string and sends it
func sendRequest(conn net.Conn, request Petition){
	encoder := gob.NewEncoder(conn)
	encoder.Encode(request)
}

// Connects to the remote HOST
func connect() net.Conn {
	conn, err := net.Dial(CONNECTION_TYPE, "localhost"+ ":" + PORT)
	checkError(err)
	return conn
}
