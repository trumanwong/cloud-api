package biz

import (
	"context"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"

	"github.com/go-kratos/kratos/v2/log"
)

// InstanceResponse is a Instance repo.
type InstanceResponse interface {
	Create(ctx context.Context, accessKeyId, accessKeySecret string, request *ecs20140526.RunInstancesRequest) (*ecs20140526.RunInstancesResponseBody, error)
	ListInstance(ctx context.Context, accessKeyId, accessKeySecret string, request *ecs20140526.DescribeInstancesRequest) (*ecs20140526.DescribeInstancesResponseBody, error)
	StartInstance(ctx context.Context, accessKeyId, accessKeySecret string, request *ecs20140526.StartInstanceRequest) (*string, error)
	StopInstance(ctx context.Context, accessKeyId, accessKeySecret string, request *ecs20140526.StopInstanceRequest) (*string, error)
	RebootInstance(ctx context.Context, accessKeyId, accessKeySecret string, request *ecs20140526.RebootInstanceRequest) (*string, error)
	DeleteInstance(ctx context.Context, accessKeyId, accessKeySecret string, request *ecs20140526.DeleteInstanceRequest) (*string, error)
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

// CreateInstance creates a Instance, and returns the new Instance.
func (uc *InstanceUseCase) CreateInstance(ctx context.Context, accessKeyId, accessKeySecret string, request *ecs20140526.RunInstancesRequest) (*ecs20140526.RunInstancesResponseBody, error) {
	return uc.repo.Create(ctx, accessKeyId, accessKeySecret, request)
}

func (uc *InstanceUseCase) ListInstance(ctx context.Context, accessKeyId, accessKeySecret string, request *ecs20140526.DescribeInstancesRequest) (*ecs20140526.DescribeInstancesResponseBody, error) {
	return uc.repo.ListInstance(ctx, accessKeyId, accessKeySecret, request)
}

func (uc *InstanceUseCase) StartInstance(ctx context.Context, accessKeyId, accessKeySecret string, request *ecs20140526.StartInstanceRequest) (*string, error) {
	return uc.repo.StartInstance(ctx, accessKeyId, accessKeySecret, request)
}

func (uc *InstanceUseCase) StopInstance(ctx context.Context, accessKeyId, accessKeySecret string, request *ecs20140526.StopInstanceRequest) (*string, error) {
	return uc.repo.StopInstance(ctx, accessKeyId, accessKeySecret, request)
}

func (uc *InstanceUseCase) RebootInstance(ctx context.Context, accessKeyId, accessKeySecret string, request *ecs20140526.RebootInstanceRequest) (*string, error) {
	return uc.repo.RebootInstance(ctx, accessKeyId, accessKeySecret, request)
}

func (uc *InstanceUseCase) DeleteInstance(ctx context.Context, accessKeyId, accessKeySecret string, request *ecs20140526.DeleteInstanceRequest) (*string, error) {
	return uc.repo.DeleteInstance(ctx, accessKeyId, accessKeySecret, request)
}
