package service

import (
	v1 "ali/api/instance/v1"
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"os"
	"testing"
)

func TestInstanceService_ListInstanceType(t *testing.T) {
	accessKeyId := os.Getenv("accessKeyId")
	accessKeySecret := os.Getenv("accessKeySecret")
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	assert.Equal(t, err, nil)

	instanceClient := v1.NewInstanceClient(conn)
	_, err = instanceClient.ListInstanceType(context.Background(), &v1.ListInstanceTypeRequest{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	})
	assert.Equal(t, err, nil)
}