package gateway

import (
	"encoding/json"
	"net/http"

	"gateway/pkg/pb"
)

type Service struct {
	hashingClient pb.HashingServiceClient
}

func NewService(hashingClient pb.HashingServiceClient) *Service {
	return &Service{
		hashingClient: hashingClient,
	}
}

func (s *Service) CheckHashHandler(w http.ResponseWriter, r *http.Request) {
	var req CheckHashRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := s.hashingClient.CheckHash(r.Context(), &pb.CheckHashRequest{Payload: req.Payload})
	if err != nil {
		http.Error(w, "Error calling CheckHash: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := CheckHashResponse{Exists: response.Exists}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Service) GetHashHandler(w http.ResponseWriter, r *http.Request) {
	var req pb.GetHashRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := s.hashingClient.GetHash(r.Context(), &req)
	if err != nil {
		http.Error(w, "Error calling GetHash: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := pb.GetHashResponse{Hash: response.Hash}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Service) CreateHashHandler(w http.ResponseWriter, r *http.Request) {
	var req pb.CreateHashRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := s.hashingClient.CreateHash(r.Context(), &req)
	if err != nil {
		http.Error(w, "Error calling CreateHash: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := pb.CreateHashResponse{Hash: response.Hash}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
