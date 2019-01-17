package config

import (
	"os"
	"sync"
)

// Port ..
type Port struct {
	HTTPPort string
	GRPCPort string
}

var confPort *Port
var portInit sync.Once

// GetPort ..
func GetPort() *Port {
	portInit.Do(func() {
		httpPort := os.Getenv("HTTP_PORT")
		grpcPort := os.Getenv("GRPC_PORT")
		confPort = &Port{
			HTTPPort: httpPort,
			GRPCPort: grpcPort,
		}
	})
	return confPort
}
