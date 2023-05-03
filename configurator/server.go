package configurator

import (
	context "context"
	"log"
)

// ConfigurationAgentServer is the struct that implements the gRPC procedures.
type ConfigurationAgentServer struct {
	UnimplementedConfigurationAgentServiceServer
}

// NewServer returns a new ConfigurationAgentServer object.
func NewServer() *ConfigurationAgentServer {
	return &ConfigurationAgentServer{}
}

// AssignIPToInterface receives an IP and a Interface name and it configure the given IP on the given Interface.
func (c *ConfigurationAgentServer) AssignIPToInterface(_ context.Context, assignment *IPAssignment) (*Result, error) {
	log.Printf("AssignIP Called with %+v as input", assignment)
	return nil, nil
}

// ConfigureNat configures nat with the given information.
func (c *ConfigurationAgentServer) ConfigureNat(_ context.Context, natConfiguration *NatConfiguration) (*Result, error) {
	log.Printf("ConfigureNat Called with %+v as input", natConfiguration)
	return nil, nil
}

// AddRoute receives a destination network and a nexthop and configure a route.
func (c *ConfigurationAgentServer) AddRoute(_ context.Context, route *Route) (*Result, error) {
	log.Printf("AddRoute Called with %+v as input", route)
	return nil, nil
}
