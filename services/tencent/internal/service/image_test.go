package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"os"
	v1 "tencent/api/instance/v1"
	"testing"
)

func TestInstanceService_ListImages(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
	assert.Equal(t, err, nil)

	instanceClient := v1.NewInstanceClient(conn)
	_, err = instanceClient.ListImages(context.Background(), &v1.ListImagesRequest{
		SecretId:  os.Getenv("secretId"),
		SecretKey: os.Getenv("secretKey"),
		Region:    os.Getenv("region"),
	})
	assert.Equal(t, err, nil)
}
