package service

import (
	v1 "ali/api/instance/v1"
	"context"
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
		AccessKeyId:     os.Getenv("accessKeyId"),
		AccessKeySecret: os.Getenv("accessKeySecret"),
		Endpoint:        os.Getenv("endpoint"),
		RegionId:        os.Getenv("regionId"),
		ImageId:         os.Getenv("imageId"),
		Name:            os.Getenv("name"),
		InstanceType:    os.Getenv("instanceType"),
		SystemDiskSize:  os.Getenv("systemDiskSize"),
		UniqueSuffix:    false,
		Amount:          1,
		Password:        os.Getenv("password"),
		DryRun:          true,
	})
	assert.Equal(t, err, nil)
}

func TestInstanceService_ListInstances(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	assert.Equal(t, err, nil)

	instanceClient := v1.NewInstanceClient(conn)
	_, err = instanceClient.ListInstances(context.Background(), &v1.ListInstancesRequest{
		AccessKeyId:     os.Getenv("accessKeyId"),
		AccessKeySecret: os.Getenv("accessKeySecret"),
		Endpoint:        os.Getenv("endpoint"),
		RegionId:        os.Getenv("regionId"),
		PageNumber:      0,
		PageSize:        0,
		NextToken:       "",
		InstanceName:    "",
		DryRun:          true,
	})
	assert.Equal(t, err, nil)
}

func TestInstanceService_StartInstances(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	assert.Equal(t, err, nil)

	instanceClient := v1.NewInstanceClient(conn)
	instanceIds := strings.Split(os.Getenv("instance_ids"), ",")
	_, err = instanceClient.StartInstances(context.Background(), &v1.StartInstancesRequest{
		AccessKeyId:       os.Getenv("accessKeyId"),
		AccessKeySecret:   os.Getenv("accessKeySecret"),
		Endpoint:          os.Getenv("endpoint"),
		RegionId:          os.Getenv("regionId"),
		InstanceIds:       instanceIds,
		BatchOptimization: "",
		DryRun:            true,
	})
	assert.Equal(t, err, nil)
}

func TestInstanceService_StopInstances(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	assert.Equal(t, err, nil)

	instanceClient := v1.NewInstanceClient(conn)
	instanceIds := strings.Split(os.Getenv("instance_ids"), ",")
	_, err = instanceClient.StopInstances(context.Background(), &v1.StopInstancesRequest{
		AccessKeyId:       os.Getenv("accessKeyId"),
		AccessKeySecret:   os.Getenv("accessKeySecret"),
		Endpoint:          os.Getenv("endpoint"),
		RegionId:          os.Getenv("regionId"),
		ForceStop:         false,
		StoppedMode:       "",
		BatchOptimization: "",
		InstanceIds:       instanceIds,
		DryRun:            true,
	})
	assert.Equal(t, err, nil)
}

func TestInstanceService_RebootInstances(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	assert.Equal(t, err, nil)

	instanceClient := v1.NewInstanceClient(conn)
	instanceIds := strings.Split(os.Getenv("instance_ids"), ",")
	_, err = instanceClient.RebootInstances(context.Background(), &v1.RebootInstancesRequest{
		AccessKeyId:       os.Getenv("accessKeyId"),
		AccessKeySecret:   os.Getenv("accessKeySecret"),
		Endpoint:          os.Getenv("endpoint"),
		RegionId:          os.Getenv("regionId"),
		ForceReboot:       false,
		BatchOptimization: "",
		InstanceIds:       instanceIds,
		DryRun:            true,
	})
	assert.Equal(t, err, nil)
}

func TestInstanceService_DeleteInstances(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	assert.Equal(t, err, nil)

	instanceClient := v1.NewInstanceClient(conn)
	instanceIds := strings.Split(os.Getenv("instance_ids"), ",")
	_, err = instanceClient.DeleteInstances(context.Background(), &v1.DeleteInstancesRequest{
		AccessKeyId:           os.Getenv("accessKeyId"),
		AccessKeySecret:       os.Getenv("accessKeySecret"),
		Endpoint:              os.Getenv("endpoint"),
		RegionId:              os.Getenv("regionId"),
		Force:                 false,
		TerminateSubscription: false,
		InstanceIds:           instanceIds,
		DryRun:                true,
	})
	assert.Equal(t, err, nil)
}
