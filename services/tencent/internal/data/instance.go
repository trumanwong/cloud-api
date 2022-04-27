package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"tencent/internal/biz"
)

type instanceRepo struct {
	data *Data
	log  *log.Helper
}

// NewInstanceRepo .
func NewInstanceRepo(data *Data, logger log.Logger) biz.InstanceRepo {
	return &instanceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (i instanceRepo) CreateInstances(ctx context.Context, secretId, secretKey, region string, request *cvm.RunInstancesRequest) (*cvm.RunInstancesResponse, error) {
	client, err := getCvmClient(secretId, secretKey, region)
	if err != nil {
		return nil, err
	}
	response, err := client.RunInstances(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (i instanceRepo) ListInstances(ctx context.Context, secretId, secretKey, region string, request *cvm.DescribeInstancesRequest) (*cvm.DescribeInstancesResponse, error) {
	client, err := getCvmClient(secretId, secretKey, region)
	if err != nil {
		return nil, err
	}
	response, err := client.DescribeInstances(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (i instanceRepo) StartInstances(ctx context.Context, secretId, secretKey, region string, request *cvm.StartInstancesRequest) (*cvm.StartInstancesResponse, error) {
	client, err := getCvmClient(secretId, secretKey, region)
	if err != nil {
		return nil, err
	}
	response, err := client.StartInstances(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (i instanceRepo) StopInstances(ctx context.Context, secretId, secretKey, region string, request *cvm.StopInstancesRequest) (*cvm.StopInstancesResponse, error) {
	client, err := getCvmClient(secretId, secretKey, region)
	if err != nil {
		return nil, err
	}
	response, err := client.StopInstances(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (i instanceRepo) RebootInstances(ctx context.Context, secretId, secretKey, region string, request *cvm.RebootInstancesRequest) (*cvm.RebootInstancesResponse, error) {
	client, err := getCvmClient(secretId, secretKey, region)
	if err != nil {
		return nil, err
	}
	response, err := client.RebootInstances(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (i instanceRepo) DeleteInstances(ctx context.Context, secretId, secretKey, region string, request *cvm.TerminateInstancesRequest) (*cvm.TerminateInstancesResponse, error) {
	client, err := getCvmClient(secretId, secretKey, region)
	if err != nil {
		return nil, err
	}
	response, err := client.TerminateInstances(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
