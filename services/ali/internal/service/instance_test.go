package service

import (
	v1 "ali/api/instance/v1"
	"context"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"os"
	"strings"
	"testing"
)

func TestInstanceService_CreateInstances(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	assert.Equal(t, err, nil)

	instanceClient := v1.NewInstanceClient(conn)
	_, err = instanceClient.CreateInstances(context.Background(), &v1.CreateInstancesRequest{
		AccessKeyId:     tea.String(os.Getenv("accessKeyId")),
		AccessKeySecret: tea.String(os.Getenv("accessKeySecret")),
		Endpoint:        tea.String(os.Getenv("endpoint")),
		RegionId:        tea.String(os.Getenv("regionId")),
		ImageId:         tea.String(os.Getenv("imageId")),
		Name:            tea.String(os.Getenv("name")),
		InstanceType:    tea.String(os.Getenv("instanceType")),
		SystemDiskSize:  tea.String(os.Getenv("systemDiskSize")),
		UniqueSuffix:    tea.Bool(false),
		Amount:          tea.Int32(1),
		Password:        tea.String(os.Getenv("password")),
		DryRun:          tea.Bool(true),
	})
	assert.Equal(t, err, nil)
}

func TestInstanceService_ListInstances(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	assert.Equal(t, err, nil)

	instanceClient := v1.NewInstanceClient(conn)
	_, err = instanceClient.ListInstances(context.Background(), &v1.ListInstancesRequest{
		AccessKeyId:     tea.String(os.Getenv("accessKeyId")),
		AccessKeySecret: tea.String(os.Getenv("accessKeySecret")),
		Endpoint:        tea.String(os.Getenv("endpoint")),
		RegionId:        tea.String(os.Getenv("regionId")),
		DryRun:          tea.Bool(true),
	})
	assert.Equal(t, err, nil)
}

func TestInstanceService_StartInstances(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	assert.Equal(t, err, nil)

	instanceClient := v1.NewInstanceClient(conn)
	instanceIds := strings.Split(os.Getenv("instance_ids"), ",")
	_, err = instanceClient.StartInstances(context.Background(), &v1.StartInstancesRequest{
		AccessKeyId:     tea.String(os.Getenv("accessKeyId")),
		AccessKeySecret: tea.String(os.Getenv("accessKeySecret")),
		Endpoint:        tea.String(os.Getenv("endpoint")),
		RegionId:        tea.String(os.Getenv("regionId")),
		InstanceIds:     instanceIds,
		DryRun:          tea.Bool(true),
	})
	assert.Equal(t, err, nil)
}

func TestInstanceService_StopInstances(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	assert.Equal(t, err, nil)

	instanceClient := v1.NewInstanceClient(conn)
	instanceIds := strings.Split(os.Getenv("instance_ids"), ",")
	_, err = instanceClient.StopInstances(context.Background(), &v1.StopInstancesRequest{
		AccessKeyId:     tea.String(os.Getenv("accessKeyId")),
		AccessKeySecret: tea.String(os.Getenv("accessKeySecret")),
		Endpoint:        tea.String(os.Getenv("endpoint")),
		RegionId:        tea.String(os.Getenv("regionId")),
		ForceStop:       tea.Bool(false),
		InstanceIds:     instanceIds,
		DryRun:          tea.Bool(true),
	})
	assert.Equal(t, err, nil)
}

func TestInstanceService_RebootInstances(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	assert.Equal(t, err, nil)

	instanceClient := v1.NewInstanceClient(conn)
	instanceIds := strings.Split(os.Getenv("instance_ids"), ",")
	_, err = instanceClient.RebootInstances(context.Background(), &v1.RebootInstancesRequest{
		AccessKeyId:     tea.String(os.Getenv("accessKeyId")),
		AccessKeySecret: tea.String(os.Getenv("accessKeySecret")),
		Endpoint:        tea.String(os.Getenv("endpoint")),
		RegionId:        tea.String(os.Getenv("regionId")),
		ForceReboot:     tea.Bool(false),
		InstanceIds:     instanceIds,
		DryRun:          tea.Bool(true),
	})
	assert.Equal(t, err, nil)
}

func TestInstanceService_DeleteInstances(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	assert.Equal(t, err, nil)

	instanceClient := v1.NewInstanceClient(conn)
	instanceIds := strings.Split(os.Getenv("instance_ids"), ",")
	_, err = instanceClient.DeleteInstances(context.Background(), &v1.DeleteInstancesRequest{
		AccessKeyId:     tea.String(os.Getenv("accessKeyId")),
		AccessKeySecret: tea.String(os.Getenv("accessKeySecret")),
		Endpoint:        tea.String(os.Getenv("endpoint")),
		RegionId:        tea.String(os.Getenv("regionId")),
		InstanceIds:     instanceIds,
		DryRun:          tea.Bool(true),
	})
	assert.Equal(t, err, nil)
}
