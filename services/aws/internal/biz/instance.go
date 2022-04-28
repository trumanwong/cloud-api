package biz

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"

	"github.com/go-kratos/kratos/v2/log"
)

// InstanceResponse is a Instance repo.
type InstanceResponse interface {
	CreateInstances(context.Context, string, string, string, *ec2.RunInstancesInput) (*ec2.RunInstancesOutput, error)
	ListInstances(context.Context, string, string, string, *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error)
	StartInstances(context.Context, string, string, string, *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error)
	StopInstances(context.Context, string, string, string, *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error)
	RebootInstances(context.Context, string, string, string, *ec2.RebootInstancesInput) (*ec2.RebootInstancesOutput, error)
	DeleteInstances(context.Context, string, string, string, *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error)
}

// InstanceUseCase is a Instance UseCase.
type InstanceUseCase struct {
	repo InstanceResponse
	log  *log.Helper
}

// NewInstanceUseCase new a Instance UseCase.
func NewInstanceUseCase(repo InstanceResponse, logger log.Logger) *InstanceUseCase {
	return &InstanceUseCase{repo: repo, log: log.NewHelper(logger)}
}

// CreateInstances and returns the new Instances.
func (uc *InstanceUseCase) CreateInstances(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.RunInstancesInput) (*ec2.RunInstancesOutput, error) {
	return uc.repo.CreateInstances(ctx, accessKeyId, secretAccessKey, region, request)
}

func (uc *InstanceUseCase) ListInstances(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	return uc.repo.ListInstances(ctx, accessKeyId, secretAccessKey, region, request)
}

func (uc *InstanceUseCase) StartInstances(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error) {
	return uc.repo.StartInstances(ctx, accessKeyId, secretAccessKey, region, request)
}

func (uc *InstanceUseCase) StopInstances(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error) {
	return uc.repo.StopInstances(ctx, accessKeyId, secretAccessKey, region, request)
}

func (uc *InstanceUseCase) RebootInstances(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.RebootInstancesInput) (*ec2.RebootInstancesOutput, error) {
	return uc.repo.RebootInstances(ctx, accessKeyId, secretAccessKey, region, request)
}

func (uc *InstanceUseCase) DeleteInstances(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	return uc.repo.DeleteInstances(ctx, accessKeyId, secretAccessKey, region, request)
}
