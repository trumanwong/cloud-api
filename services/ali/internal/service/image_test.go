package service

import (
	v1 "ali/api/instance/v1"
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"os"
	"testing"
)

func TestInstanceService_ListImages(t *testing.T) {
	accessKeyId := os.Getenv("accessKeyId")
	accessKeySecret := os.Getenv("accessKeySecret")
	endpoint := os.Getenv("endpoint")
	regionId := os.Getenv("regionId")
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	assert.Equal(t, err, nil)

	instanceClient := v1.NewInstanceClient(conn)
	_, err = instanceClient.ListImages(context.Background(), &v1.ListImagesRequest{
		AccessKeyId:     &accessKeyId,
		AccessKeySecret: &accessKeySecret,
		Endpoint:        &endpoint,
		RegionId:        &regionId,
	})
	assert.Equal(t, err, nil)
}
