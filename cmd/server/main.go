package main

import (
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/peterbourgon/ff"
	"github.com/tchaudhry91/wts-service/svc"
	"github.com/tchaudhry91/wts-service/svc/pb"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fs := flag.NewFlagSet("wts", flag.ExitOnError)
	var (
		grpcAddr = fs.String("grpc-addr", "127.0.0.1:7878", "The port to listen on")
		dbFile   = fs.String("db-file", "wts.db", "Local file to store the data in")
	)
	err := ff.Parse(fs, os.Args[1:],
		ff.WithEnvVarNoPrefix())

	if err != nil {
		panic(err)
	}

	var logger log.Logger
	{
		logger = log.NewJSONLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// Database
	var store svc.Store
	if *dbFile != "" {
		store, err = svc.NewBoltStore(*dbFile)
		if err != nil {
			panic(err)
		}
	}

	// Base Service
	service := svc.NewWTSService(store)

	// Logging Middleware
	service = svc.LoggingMiddleware(service, logger)

	// GRPC Transport
	grpcServer := svc.NewGRPCServer(svc.MakeServerEndpoints(service), logger)
	// The gRPC listener mounts the Go kit gRPC server we created.
	grpcListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}
	// we add the Go Kit gRPC Interceptor to our gRPC service as it is used by
	// the here demonstrated zipkin tracing middleware.
	baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	pb.RegisterWTSServer(baseServer, grpcServer)

	// Make Shutdown Channel
	shutdown := make(chan error, 1)
	// Make Interrupt Channel
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		logger.Log("transport", "gRPC", "addr", *grpcAddr)
		err = baseServer.Serve(grpcListener)
		shutdown <- err
	}()

	select {
	case signalKill := <-interrupt:
		logger.Log("msg", fmt.Sprintf("Stopping Server: %s", signalKill.String()))
	case err := <-shutdown:
		logger.Log("error", err)
	}

	err = grpcListener.Close()
	if err != nil {
		logger.Log("error", err)
	}
}
