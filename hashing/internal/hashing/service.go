package hashing

import (
	"context"
	"crypto/sha256"
	"fmt"

	"hashing/pkg/pb"
)

type HashingService struct {
	store map[string]string
	pb.UnimplementedHashingServiceServer
}

func NewHashingService() *HashingService {
	return &HashingService{
		store: make(map[string]string),
	}
}

func (s *HashingService) CheckHash(ctx context.Context, req *pb.CheckHashRequest) (*pb.CheckHashResponse, error) {
	_, exists := s.store[req.Payload]
	return &pb.CheckHashResponse{Exists: exists}, nil
}

func (s *HashingService) GetHash(ctx context.Context, req *pb.GetHashRequest) (*pb.GetHashResponse, error) {
	hash, exists := s.store[req.Payload]
	if !exists {
		return nil, fmt.Errorf("hash not found for payload: %s", req.Payload)
	}
	return &pb.GetHashResponse{Hash: hash}, nil
}

func (s *HashingService) CreateHash(ctx context.Context, req *pb.CreateHashRequest) (*pb.CreateHashResponse, error) {
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(req.Payload)))
	s.store[req.Payload] = hash
	return &pb.CreateHashResponse{Hash: hash}, nil
}
