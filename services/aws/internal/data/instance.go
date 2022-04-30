package data

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"

	"aws/internal/biz"
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

func (r *instanceResponse) CreateInstances(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.RunInstancesInput) (*ec2.RunInstancesOutput, error) {
	client, err := getClient(
		ctx,
		accessKeyId,
		secretAccessKey,
		region,
	)
	if err != nil {
		return nil, err
	}
	result, err := client.RunInstances(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *instanceResponse) ListInstances(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	client, err := getClient(
		ctx,
		accessKeyId,
		secretAccessKey,
		region,
	)
	if err != nil {
		return nil, err
	}
	result, err := client.DescribeInstances(ctx, request)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *instanceResponse) StartInstances(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error) {
	client, err := getClient(
		ctx,
		accessKeyId,
		secretAccessKey,
		region,
	)
	if err != nil {
		return nil, err
	}
	result, err := client.StartInstances(ctx, request)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *instanceResponse) StopInstances(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error) {
	client, err := getClient(
		ctx,
		accessKeyId,
		secretAccessKey,
		region,
	)
	if err != nil {
		return nil, err
	}
	result, err := client.StopInstances(ctx, request)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *instanceResponse) RebootInstances(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.RebootInstancesInput) (*ec2.RebootInstancesOutput, error) {
	client, err := getClient(
		ctx,
		accessKeyId,
		secretAccessKey,
		region,
	)
	if err != nil {
		return nil, err
	}
	result, err := client.RebootInstances(ctx, request)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *instanceResponse) DeleteInstances(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	client, err := getClient(
		ctx,
		accessKeyId,
		secretAccessKey,
		region,
	)
	if err != nil {
		return nil, err
	}
	result, err := client.TerminateInstances(ctx, request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
