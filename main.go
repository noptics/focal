package main

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/noptics/focal/config"
	"github.com/noptics/focal/registrygrpc"
	"github.com/noptics/focal/restapi"
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

	c := config.New()
	errChan := make(chan error)

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
	c.Set(config.Registry, rc)

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
	c.Set(config.Streamer, sc)

	restPort := os.Getenv("REST_PORT")
	if len(restPort) == 0 {
		restPort = "7766"
	}
	c.Set(config.RESTPort, restPort)

	rs := restapi.NewRestServer(errChan, l, c)

	// go until told to stop
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-sigs:
	case <-errChan:
		l.Infow("error", "error", err.Error())
	}

	l.Info("shutting down")

	err = rs.Stop()
	if err != nil {
		l.Infow("error shutting down rest server", "error", err.Error())
	}

	l.Info("finished")
}
