package reg

import (
	"cheque-04/common/api"
)

var serviceServersRegistry []api.ServiceServer

func RegisterServiceServers(data []api.ServiceServer) {
	serviceServersRegistry = append(serviceServersRegistry, data...)
}

func GetServiceServersRegistry() []api.ServiceServer {
	return serviceServersRegistry
}
