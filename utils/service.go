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
