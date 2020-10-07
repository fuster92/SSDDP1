// AUTORES: Javier Fuster Trallero / Javier Herrer Torres
// NIAs: 626901 / 776609
// FICHERO: service.go
// FECHA: 04-oct-2020
// TIEMPO: 15'
// DESCRIPCIÓN: Describe un struct con un servicio en red

package utils

// Describes a service with a name and an address.
type Service struct {
	Name string
	Host string
	Port string
}

// Returns an string representing the address of the service
func (service Service) Address() string {
	return service.Host + ":" + service.Port
}
