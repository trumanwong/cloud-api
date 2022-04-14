package service

import (
	v1 "ali/api/instance/v1"
	"context"
	"github.com/go-playground/assert/v2"
	"google.golang.org/grpc"
	"os"
	"testing"
)

func TestListRegion(t *testing.T) {
	accessKeyId := os.Getenv("accessKeyId")
	accessKeySecret := os.Getenv("accessKeySecret")
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	assert.Equal(t, err, nil)

	instanceClient := v1.NewInstanceClient(conn)
	_, err = instanceClient.ListRegion(context.Background(), &v1.ListRegionRequest{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	})
	assert.Equal(t, err, nil)
}
