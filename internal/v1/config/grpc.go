package config

import (
	"errors"
	"net"
	"os"
)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

// GRPCConfig describes config interface
type GRPCConfig interface {
	Address() string
}

type grpcConfig struct {
	host string
	port string
}

// NewGRPCConfig creates config struct
func NewGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)

	if len(host) == 0 {
		return nil, errors.New("grpc host not defined")
	}

	port := os.Getenv(grpcPortEnvName)

	if len(port) == 0 {
		return nil, errors.New("grpc port not defined")
	}

	return &grpcConfig{
		host,
		port,
	}, nil
}

func (cfg *grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
