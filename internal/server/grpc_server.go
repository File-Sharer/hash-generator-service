package server

import (
	"net"

	"github.com/File-Sharer/hash-generator-service/internal/config"
	"github.com/File-Sharer/hash-generator-service/internal/pb"
	"github.com/File-Sharer/hash-generator-service/internal/service"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedHasherServer

	services *service.Service
}

func RunGRPCServer(cfg *config.GRPCServerConfig, services *service.Service) error {
	lis, err := net.Listen(cfg.Network, cfg.Addr)
	if err != nil {
		return err
	}
	defer lis.Close()

	s := grpc.NewServer()
	pb.RegisterHasherServer(s, &GRPCServer{services: services})
	return s.Serve(lis)
}
