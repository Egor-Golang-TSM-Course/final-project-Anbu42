package gateway_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	"gateway/internal/gateway"
	"gateway/pkg/pb"
)

type mockHashingServiceClient struct{}

func (m *mockHashingServiceClient) CheckHash(ctx context.Context, in *pb.CheckHashRequest, opts ...grpc.CallOption) (*pb.CheckHashResponse, error) {
	return &pb.CheckHashResponse{Exists: true}, nil
}

func (m *mockHashingServiceClient) GetHash(ctx context.Context, in *pb.GetHashRequest, opts ...grpc.CallOption) (*pb.GetHashResponse, error) {
	return &pb.GetHashResponse{Hash: "mock-hash"}, nil
}

func (m *mockHashingServiceClient) CreateHash(ctx context.Context, in *pb.CreateHashRequest, opts ...grpc.CallOption) (*pb.CreateHashResponse, error) {
	return &pb.CreateHashResponse{Hash: "mock-hash"}, nil
}

func TestCheckHashHandler(t *testing.T) {
	svc := gateway.NewService(&mockHashingServiceClient{})

	reqBody := `{"payload": "test-payload"}`
	req := httptest.NewRequest("POST", "/checkhash", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	svc.CheckHashHandler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp gateway.CheckHashResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.True(t, resp.Exists)
}

func TestGetHashHandler(t *testing.T) {
	svc := gateway.NewService(&mockHashingServiceClient{})

	reqBody := `{"payload": "test-payload"}`
	req := httptest.NewRequest("POST", "/gethash", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	svc.GetHashHandler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp pb.GetHashResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "mock-hash", resp.Hash)
}

func TestCreateHashHandler(t *testing.T) {
	svc := gateway.NewService(&mockHashingServiceClient{})

	reqBody := `{"payload": "test-payload"}`
	req := httptest.NewRequest("POST", "/createhash", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	svc.CreateHashHandler(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp pb.CreateHashResponse
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "mock-hash", resp.Hash)
}
