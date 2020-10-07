// AUTORES: Javier Fuster Trallero / Javier Herrer Torres
// NIAs: 626901 / 776609
// FICHERO: service.go
// FECHA: 04-oct-2020
// TIEMPO: 10'
// DESCRIPCIÓN: Describe un struct con un trabajo a realizar

package utils

import "net"

// Job inside the master process.
type Job struct {
	Connection net.Conn
	Request    Request
}
