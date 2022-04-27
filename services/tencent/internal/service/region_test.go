package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"os"
	v1 "tencent/api/instance/v1"
	"testing"
)

func TestInstanceService_ListRegions(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
	assert.Equal(t, err, nil)

	instanceClient := v1.NewInstanceClient(conn)
	_, err = instanceClient.ListRegions(context.Background(), &v1.ListRegionsRequest{
		SecretId:  os.Getenv("secretId"),
		SecretKey: os.Getenv("secretKey"),
	})
	assert.Equal(t, err, nil)
}
