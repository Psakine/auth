package app

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Psakine/auth/pkg/user_v1"
)

// App ...
type App struct {
	serviceProvider *ServiceProvider
	grpcServer      *grpc.Server
}

// NewApp ...
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)

	if err != nil {
		return nil, err
	}

	return a, nil
}

// Run ...
func (a *App) Run() error {
	return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, init := range inits {
		err := init(ctx)

		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = NewServiceProvider()

	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer()
	reflection.Register(a.grpcServer)

	user_v1.RegisterUserV1Server(a.grpcServer, a.serviceProvider.UserServer())

	return nil
}

func (a *App) runGRPCServer() error {
	address := a.serviceProvider.GRPCConfig().Address()

	log.Printf("GRPC server is running on %s", address)

	listener, err := net.Listen("tcp", address)

	if err != nil {
		return err
	}

	return a.grpcServer.Serve(listener)
}
