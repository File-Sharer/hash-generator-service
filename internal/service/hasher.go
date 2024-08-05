package service

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
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
