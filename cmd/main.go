package main

import (
	"context"
	"fmt"
	"gemini/cmd/protocol/grpc"
	"gemini/cmd/protocol/rest"
	"gemini/config"
	"gemini/handler"
	"gemini/storage"
	"os"
)

func main() {
	storage.Connect()
	if err := RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

// RunServer ..
func RunServer() error {
	port := config.GetPort()
	ctx := context.Background()

	go func() {
		withServerRest := []rest.OptionFunc{
			rest.WithEventServer(),
		}
		rest.RunServer(ctx, port.GRPCPort, port.HTTPPort, withServerRest...)
	}()

	withServerGRPC := []grpc.OptionFunc{
		grpc.WithEventServer(handler.NewEventHandler()),
	}
	return grpc.RunServer(ctx, port.GRPCPort, withServerGRPC...)
}
