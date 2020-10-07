// Utility functions and constants used by servers and clients
package utils

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
)

const (
	ServerPort           = "5555"
	ConnectionType       = "tcp"
	Size                 = 60000
	QOS            int64 = 744
)

// Sends an array of primes through a connection.
func SendPrimes(conn io.Writer, primes []int) {
	encoder := gob.NewEncoder(conn)
	err := encoder.Encode(primes)
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error())
	}
}

// Reads an array of integers.
func ReadPrimes(conn io.Reader) []int {
	var primes []int
	decoder := gob.NewDecoder(conn)
	_ = decoder.Decode(&primes)
	return primes
}
