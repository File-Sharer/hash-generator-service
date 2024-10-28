package service

import (
	"context"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/File-Sharer/hash-generator-service/internal/model"
	"github.com/File-Sharer/hash-generator-service/internal/pb"
	"github.com/golang-jwt/jwt/v5"
)

type HasherService struct {}

func NewHasherService() *HasherService {
	return &HasherService{}
}

func (s *HasherService) Hash(baseString string) (string, error) {
	hash := sha1.New()
	_, err := hash.Write([]byte(baseString))
	if err != nil {
		return "", err
	}

	timestamp := strconv.Itoa(time.Now().Nanosecond())
	_, err = hash.Write([]byte(timestamp))
	if err != nil {
		return "", err
	}

	hashBytes := hash.Sum([]byte{})
	hashString := hex.EncodeToString(hashBytes)
	return hashString[:32], nil
}

func (s *HasherService) NewUID(userLogin string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(userLogin))
	if err != nil {
		return "", err
	}

	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString[:16], nil
}

func (s *HasherService) NewJWT(ctx context.Context, req *pb.NewJWTReq) (string, error) {
	if req.Secret != os.Getenv("SECRET") {
		return "", errNoAccess
	}

	userID := strings.TrimSpace(req.UserId)
	userRole := strings.TrimSpace(req.Role)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"role": userRole,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_JWT")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *HasherService) DecodeJWT(ctx context.Context, req *pb.DecodeJWTReq) (*model.User, error) {
	if req.Secret != os.Getenv("SECRET") {
		return nil, errNoAccess
	}

	parsedToken, err := jwt.Parse(req.Jwt, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_JWT")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, errors.New("token is not valid")
	}

	return &model.User{
		ID: claims["sub"].(string),
		Role: claims["role"].(string),
	}, nil
}
