package main

import (
	"fmt"
	"net"
	"os"
)

const (
	RemoteHost = "ssdd.javierfuster.codes"
	PORT       = "5555"
	TYPE       = "tcp"
	SIZE       = 60000
)

type Request struct {
	ID    int
	Prime int
}

type Job struct {
	connection net.Conn
	request    Request
}

// Checks if there's an error
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func printError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr,err.Error())
	}
}