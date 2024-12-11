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

func (s *GRPCServer) NewJWT(ctx context.Context, req *pb.NewJWTReq) (*pb.NewJWTRes, error) {
	token, err := s.services.Hasher.NewJWT(ctx, req)
	if err != nil {
		return &pb.NewJWTRes{Ok: false}, err
	}

	return &pb.NewJWTRes{Ok: true, Token: token}, nil
}

func (s *GRPCServer) DecodeJWT(ctx context.Context, req *pb.DecodeJWTReq) (*pb.DecodeJWTRes, error) {
	user, err := s.services.Hasher.DecodeJWT(ctx, req)
	if err != nil {
		return &pb.DecodeJWTRes{Ok: false}, err
	}

	return &pb.DecodeJWTRes{Ok: true, UserId: user.ID, Role: user.Role}, nil
}

func (s *GRPCServer) GenerateJWTPair(ctx context.Context, req *pb.GenerateJWTPairReq) (*pb.GenerateJWTPairRes, error) {
	res, err := s.services.Hasher.GenerateJWTPair(ctx, req)
	if err != nil {
		return &pb.GenerateJWTPairRes{Ok: false}, err
	}

	return res, nil
}
