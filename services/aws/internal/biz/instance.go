package biz

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"

	"github.com/go-kratos/kratos/v2/log"
)

// InstanceResponse is a Instance repo.
type InstanceResponse interface {
	Create(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.RunInstancesInput) (*ec2.RunInstancesOutput, error)
	ListInstance(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error)
	StartInstance(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error)
	StopInstance(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error)
	RebootInstance(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.RebootInstancesInput) (*ec2.RebootInstancesOutput, error)
	DeleteInstance(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error)
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

// Create Instances, and returns the new Instances.
func (uc *InstanceUseCase) Create(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.RunInstancesInput) (*ec2.RunInstancesOutput, error) {
	return uc.repo.Create(ctx, accessKeyId, secretAccessKey, region, request)
}

func (uc *InstanceUseCase) ListInstance(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	return uc.repo.ListInstance(ctx, accessKeyId, secretAccessKey, region, request)
}

func (uc *InstanceUseCase) StartInstance(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.StartInstancesInput) (*ec2.StartInstancesOutput, error) {
	return uc.repo.StartInstance(ctx, accessKeyId, secretAccessKey, region, request)
}

func (uc *InstanceUseCase) StopInstance(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.StopInstancesInput) (*ec2.StopInstancesOutput, error) {
	return uc.repo.StopInstance(ctx, accessKeyId, secretAccessKey, region, request)
}

func (uc *InstanceUseCase) RebootInstance(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.RebootInstancesInput) (*ec2.RebootInstancesOutput, error) {
	return uc.repo.RebootInstance(ctx, accessKeyId, secretAccessKey, region, request)
}

func (uc *InstanceUseCase) DeleteInstance(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	return uc.repo.DeleteInstance(ctx, accessKeyId, secretAccessKey, region, request)
}
