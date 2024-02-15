package app

import (
	"log"

	"github.com/Psakine/auth/internal/v1/api/user"
	"github.com/Psakine/auth/internal/v1/config"
)

// ServiceProvider ...
type ServiceProvider struct {
	grpcConfig config.GRPCConfig
	userServer *user.Server
}

// NewServiceProvider ...
func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

// GRPCConfig ...
func (s *ServiceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

// UserServer ...
func (s *ServiceProvider) UserServer() *user.Server {
	if s.userServer == nil {
		s.userServer = user.NewUserServer()
	}

	return s.userServer
}
