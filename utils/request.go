// AUTORES: Javier Fuster Trallero / Javier Herrer Torres
// NIAs: 626901 / 776609
// FICHERO: request.go
// FECHA: 04-oct-2020
// TIEMPO: 5'
// DESCRIPCIÓN: Describe las peticiones que hacen los clientes al servidor

package utils

// Format of the request made by the clients.
type Request struct {
	ID    int
	Prime int
}
