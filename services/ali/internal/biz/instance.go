package biz

import (
	"context"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"

	"github.com/go-kratos/kratos/v2/log"
)

// InstanceResponse is a Instance repo.
type InstanceResponse interface {
	CreateInstances(context.Context, *string, *string, *string, *ecs20140526.RunInstancesRequest) (*ecs20140526.RunInstancesResponseBody, error)
	ListInstances(context.Context, *string, *string, *string, *ecs20140526.DescribeInstancesRequest) (*ecs20140526.DescribeInstancesResponse, error)
	StartInstances(context.Context, *string, *string, *string, *ecs20140526.StartInstancesRequest) (*ecs20140526.StartInstancesResponse, error)
	StopInstances(context.Context, *string, *string, *string, *ecs20140526.StopInstancesRequest) (*ecs20140526.StopInstancesResponse, error)
	RebootInstances(context.Context, *string, *string, *string, *ecs20140526.RebootInstancesRequest) (*ecs20140526.RebootInstancesResponse, error)
	DeleteInstances(context.Context, *string, *string, *string, *ecs20140526.DeleteInstancesRequest) (*ecs20140526.DeleteInstancesResponse, error)
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

// CreateInstances creates Instances, and returns the new Instances.
func (uc *InstanceUseCase) CreateInstances(ctx context.Context, accessKeyId, accessKeySecret, endpoint *string, request *ecs20140526.RunInstancesRequest) (*ecs20140526.RunInstancesResponseBody, error) {
	return uc.repo.CreateInstances(ctx, accessKeyId, accessKeySecret, endpoint, request)
}

func (uc *InstanceUseCase) ListInstances(ctx context.Context, accessKeyId, accessKeySecret, endpoint *string, request *ecs20140526.DescribeInstancesRequest) (*ecs20140526.DescribeInstancesResponse, error) {
	return uc.repo.ListInstances(ctx, accessKeyId, accessKeySecret, endpoint, request)
}

func (uc *InstanceUseCase) StartInstances(ctx context.Context, accessKeyId, accessKeySecret, endpoint *string, request *ecs20140526.StartInstancesRequest) (*ecs20140526.StartInstancesResponse, error) {
	return uc.repo.StartInstances(ctx, accessKeyId, accessKeySecret, endpoint, request)
}

func (uc *InstanceUseCase) StopInstances(ctx context.Context, accessKeyId, accessKeySecret, endpoint *string, request *ecs20140526.StopInstancesRequest) (*ecs20140526.StopInstancesResponse, error) {
	return uc.repo.StopInstances(ctx, accessKeyId, accessKeySecret, endpoint, request)
}

func (uc *InstanceUseCase) RebootInstances(ctx context.Context, accessKeyId, accessKeySecret, endpoint *string, request *ecs20140526.RebootInstancesRequest) (*ecs20140526.RebootInstancesResponse, error) {
	return uc.repo.RebootInstances(ctx, accessKeyId, accessKeySecret, endpoint, request)
}

func (uc *InstanceUseCase) DeleteInstances(ctx context.Context, accessKeyId, accessKeySecret, endpoint *string, request *ecs20140526.DeleteInstancesRequest) (*ecs20140526.DeleteInstancesResponse, error) {
	return uc.repo.DeleteInstances(ctx, accessKeyId, accessKeySecret, endpoint, request)
}
