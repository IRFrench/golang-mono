package server

import (
	"golang-mono/godoc-test/pb"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	server *grpc.Server
}

func (g *GrpcServer) Run(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		slog.Error("Could not create listener", "Error", err)
		return
	}

	slog.Info("Started gRPC server", "Address", addr)

	if err := g.server.Serve(lis); err != nil {
		slog.Error("Failed to start server", "Error", err)
	}
}

func (g *GrpcServer) Shutdown() {
	g.server.GracefulStop()
}

func CreateServer() *GrpcServer {
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &greeterServer{})
	return &GrpcServer{server: s}
}
