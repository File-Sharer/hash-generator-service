package server

import (
	"context"
	"strings"

	"github.com/File-Sharer/hash-generator-service/internal/pb"
)

func (s *GRPCServer) Hash(ctx context.Context, req *pb.HashReq) (*pb.HashRes, error) {
	hash, err := s.services.Hasher.Hash(strings.TrimSpace(req.GetBaseString()))
	if err != nil {
		return nil, err
	}

	return &pb.HashRes{
		Ok: true,
		Hash: hash,
	}, nil
}

func (s *GRPCServer) NewUID(ctx context.Context, req *pb.NewUIDReq) (*pb.NewUIDRes, error) {
	hash, err := s.services.Hasher.NewUID(strings.TrimSpace(req.GetUserLogin()))
	if err != nil {
		return nil, err
	}

	return &pb.NewUIDRes{
		Ok: true,
		Uid: hash,
	}, nil
}
