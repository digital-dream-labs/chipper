package main

import (
	pb "github.com/digital-dream-labs/api/go/chipperpb"
	"github.com/digital-dream-labs/chipper/pkg/server"
	"github.com/digital-dream-labs/chipper/pkg/voice_processors/noop"
	grpclog "github.com/digital-dream-labs/hugh/grpc/interceptors/log"
	grpcserver "github.com/digital-dream-labs/hugh/grpc/server"
	"github.com/digital-dream-labs/hugh/log"
)

func main() {
	log.SetJSONFormat("2006-01-02 15:04:05")
	startServer()
}

func startServer() {
	srv, err := grpcserver.New(
		grpcserver.WithViper(),
		grpcserver.WithLogger(log.Base()),
		grpcserver.WithReflectionService(),

		grpcserver.WithUnaryServerInterceptors(
			grpclog.UnaryServerInterceptor(),
		),

		grpcserver.WithStreamServerInterceptors(
			grpclog.StreamServerInterceptor(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	p, err := noop.New()
	if err != nil {
		log.Fatal(err)
	}

	s, _ := server.New(
		server.WithLogger(log.Base()),
		server.WithIntentProcessor(p),
		server.WithKnowledgeGraphProcessor(p),
	)

	pb.RegisterChipperGrpcServer(srv.Transport(), s)

	srv.Start()

	<-srv.Notify(grpcserver.Stopped)
}
