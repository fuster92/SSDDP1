package utils

import "net"

// Job inside the master process.
type Job struct {
	Connection net.Conn
	Request    Request
}
