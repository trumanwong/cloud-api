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

func (r *instanceResponse) Create(ctx context.Context, accessKeyId, accessKeySecret, endpoint string, request *ecs20140526.RunInstancesRequest) (*ecs20140526.RunInstancesResponseBody, error) {
	client, err := createClient(
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

func (r *instanceResponse) ListInstance(ctx context.Context, accessKeyId, accessKeySecret, endpoint string, request *ecs20140526.DescribeInstancesRequest) (*ecs20140526.DescribeInstancesResponseBody, error) {
	client, err := createClient(
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
	return result.Body, nil
}

func (r *instanceResponse) StartInstance(ctx context.Context, accessKeyId, accessKeySecret, endpoint string, request *ecs20140526.StartInstanceRequest) (*string, error) {
	client, err := createClient(
		accessKeyId,
		accessKeySecret,
		endpoint,
	)
	if err != nil {
		return nil, err
	}
	result, err := client.StartInstance(request)
	if err != nil {
		return nil, err
	}
	return result.Body.RequestId, nil
}

func (r *instanceResponse) StopInstance(ctx context.Context, accessKeyId, accessKeySecret, endpoint string, request *ecs20140526.StopInstanceRequest) (*string, error) {
	client, err := createClient(
		accessKeyId,
		accessKeySecret,
		endpoint,
	)
	if err != nil {
		return nil, err
	}
	result, err := client.StopInstance(request)
	if err != nil {
		return nil, err
	}
	return result.Body.RequestId, nil
}

func (r *instanceResponse) RebootInstance(ctx context.Context, accessKeyId, accessKeySecret, endpoint string, request *ecs20140526.RebootInstanceRequest) (*string, error) {
	client, err := createClient(
		accessKeyId,
		accessKeySecret,
		endpoint,
	)
	if err != nil {
		return nil, err
	}
	result, err := client.RebootInstance(request)
	if err != nil {
		return nil, err
	}
	return result.Body.RequestId, nil
}

func (r *instanceResponse) DeleteInstance(ctx context.Context, accessKeyId, accessKeySecret, endpoint string, request *ecs20140526.DeleteInstanceRequest) (*string, error) {
	client, err := createClient(
		accessKeyId,
		accessKeySecret,
		endpoint,
	)
	if err != nil {
		return nil, err
	}
	result, err := client.DeleteInstance(request)
	if err != nil {
		return nil, err
	}
	return result.Body.RequestId, nil
}
