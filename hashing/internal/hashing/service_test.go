package hashing

import (
	"context"
	"testing"

	"hashing/pkg/pb"

	"github.com/stretchr/testify/assert"
)

func TestCheckHash(t *testing.T) {
	service := NewHashingService()
	service.store["test_payload"] = "test_hash"

	req := &pb.CheckHashRequest{Payload: "test_payload"}
	resp, err := service.CheckHash(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Exists)

	req = &pb.CheckHashRequest{Payload: "non_existent_payload"}
	resp, err = service.CheckHash(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.False(t, resp.Exists)
}

func TestGetHash(t *testing.T) {
	service := NewHashingService()
	service.store["test_payload"] = "test_hash"

	req := &pb.GetHashRequest{Payload: "test_payload"}
	resp, err := service.GetHash(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "test_hash", resp.Hash)

	req = &pb.GetHashRequest{Payload: "non_existent_payload"}
	resp, err = service.GetHash(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestCreateHash(t *testing.T) {
	service := NewHashingService()

	req := &pb.CreateHashRequest{Payload: "new_payload"}
	resp, err := service.CreateHash(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Hash)
	assert.Equal(t, resp.Hash, service.store["new_payload"])
}
