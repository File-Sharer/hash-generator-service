package service

import (
	"context"

	"github.com/File-Sharer/hash-generator-service/internal/model"
	"github.com/File-Sharer/hash-generator-service/internal/pb"
)

type Hasher interface {
	Hash(baseString string) (string, error)
	NewUID(userLogin string) (string, error)
	NewJWT(ctx context.Context, req *pb.NewJWTReq) (string, error)
	DecodeJWT(ctx context.Context, req *pb.DecodeJWTReq) (*model.User, error)
	GenerateJWTPair(ctx context.Context, req *pb.GenerateJWTPairReq) (*pb.GenerateJWTPairRes, error)
}

type Service struct {
	Hasher
}

func New() *Service {
	return &Service{
		Hasher: NewHasherService(),
	}
}
