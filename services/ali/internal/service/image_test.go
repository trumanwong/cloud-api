package service

import (
	v1 "ali/api/instance/v1"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"os"
	"testing"
)

func TestInstanceService_ListImage(t *testing.T) {
	accessKeyId := os.Getenv("accessKeyId")
	accessKeySecret := os.Getenv("accessKeySecret")
	endpoint := os.Getenv("endpoint")
	regionId := os.Getenv("regionId")
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	assert.Equal(t, err, nil)

	instanceClient := v1.NewInstanceClient(conn)
	_, err = instanceClient.ListImage(context.Background(), &v1.ListImageRequest{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Endpoint:        endpoint,
		RegionId:        regionId,
	})
	fmt.Println(err)
	assert.Equal(t, err, nil)
}
