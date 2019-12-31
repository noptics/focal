package main

import (
	"os"
	"runtime"

	"github.com/noptics/focal/config"
	"github.com/noptics/focal/registrygrpc"
	"github.com/noptics/focal/streamergrpc"
	"github.com/noptics/golog"
	"google.golang.org/grpc"
)

func main() {
	var l golog.Logger
	if len(os.Getenv("DEBUG")) != 0 {
		l = golog.StdOut(golog.LEVEL_DEBUG)
	} else {
		l = golog.StdOut(golog.LEVEL_ERROR)
	}

	l.Init()
	defer l.Finish()

	l.Infow("starting focal", "version", os.Getenv("VERSION"), "commit", os.Getenv("COMMIT"), "go", runtime.Version())

	s := config.New()

	// Connect to the registry
	registryAddress := os.Getenv("REGISTRY_SERVICE")
	if len(registryAddress) == 0 {
		l.Errorw("must provide REGISTRY_SERVICE with the grpc address/port of the registry", "example", `REGISTRY_SERVICE=127.0.0.1:7775`)
		os.Exit(1)
	}

	l.Infow("connecting to registry service", "addr", registryAddress)
	registryConn, err := grpc.Dial(registryAddress, grpc.WithInsecure())
	if err != nil {
		l.Errorw("error connecting to registry service", "error", err.Error())
		os.Exit(1)
	}
	defer registryConn.Close()

	rc := registrygrpc.NewProtoRegistryClient(registryConn)

	l.Info("successfully connected to the registry service")
	s.Set(config.Registry, rc)

	// Connect to streamer
	streamerAddress := os.Getenv("STREAMER_SERVICE")
	if len(streamerAddress) == 0 {
		l.Errorw("must provide STREAMER_SERVICE with the grpc address/port of the streamer", "example", `STREAMER_SERVICE=127.0.0.1:7785`)
		os.Exit(1)
	}

	streamerConn, err := grpc.Dial(streamerAddress, grpc.WithInsecure())
	if err != nil {
		l.Errorw("error connecting to streamer service", "error", err.Error())
		os.Exit(1)
	}

	sc := streamergrpc.NewMessagesClient(streamerConn)
	l.Info("successfully connected to the streaer service")
	s.Set(config.Streamer, sc)
}
