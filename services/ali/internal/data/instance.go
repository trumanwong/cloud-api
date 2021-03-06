package data

import (
	"ali/internal/biz"
	"context"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/go-kratos/kratos/v2/log"
)

type instanceResponse struct {
	data *Data
	log  *log.Helper
}

// NewInstanceRepo .
func NewInstanceRepo(data *Data, logger log.Logger) biz.InstanceResponse {
	return &instanceResponse{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *instanceResponse) CreateInstances(ctx context.Context, accessKeyId, accessKeySecret, endpoint *string, request *ecs20140526.RunInstancesRequest) (*ecs20140526.RunInstancesResponseBody, error) {
	client, err := getClient(
		accessKeyId,
		accessKeySecret,
		endpoint,
	)
	if err != nil {
		return nil, err
	}
	result, err := client.RunInstances(request)
	if err != nil {
		return nil, err
	}

	return result.Body, nil
}

func (r *instanceResponse) ListInstances(ctx context.Context, accessKeyId, accessKeySecret, endpoint *string, request *ecs20140526.DescribeInstancesRequest) (*ecs20140526.DescribeInstancesResponse, error) {
	client, err := getClient(
		accessKeyId,
		accessKeySecret,
		endpoint,
	)
	if err != nil {
		return nil, err
	}
	result, err := client.DescribeInstances(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *instanceResponse) StartInstances(ctx context.Context, accessKeyId, accessKeySecret, endpoint *string, request *ecs20140526.StartInstancesRequest) (*ecs20140526.StartInstancesResponse, error) {
	client, err := getClient(
		accessKeyId,
		accessKeySecret,
		endpoint,
	)
	if err != nil {
		return nil, err
	}
	result, err := client.StartInstances(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *instanceResponse) StopInstances(ctx context.Context, accessKeyId, accessKeySecret, endpoint *string, request *ecs20140526.StopInstancesRequest) (*ecs20140526.StopInstancesResponse, error) {
	client, err := getClient(
		accessKeyId,
		accessKeySecret,
		endpoint,
	)
	if err != nil {
		return nil, err
	}
	result, err := client.StopInstances(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *instanceResponse) RebootInstances(ctx context.Context, accessKeyId, accessKeySecret, endpoint *string, request *ecs20140526.RebootInstancesRequest) (*ecs20140526.RebootInstancesResponse, error) {
	client, err := getClient(
		accessKeyId,
		accessKeySecret,
		endpoint,
	)
	if err != nil {
		return nil, err
	}
	result, err := client.RebootInstances(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *instanceResponse) DeleteInstances(ctx context.Context, accessKeyId, accessKeySecret, endpoint *string, request *ecs20140526.DeleteInstancesRequest) (*ecs20140526.DeleteInstancesResponse, error) {
	client, err := getClient(
		accessKeyId,
		accessKeySecret,
		endpoint,
	)
	if err != nil {
		return nil, err
	}
	result, err := client.DeleteInstances(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
