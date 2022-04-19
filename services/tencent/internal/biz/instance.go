package biz

import (
	"context"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"

	"github.com/go-kratos/kratos/v2/log"
)

// InstanceRepo is a Instance repo.
type InstanceRepo interface {
	CreateInstances(ctx context.Context, secretId, secretKey, region string, request *cvm.RunInstancesRequest) (*cvm.RunInstancesResponse, error)
	ListInstances(ctx context.Context, secretId, secretKey, region string, request *cvm.DescribeInstancesRequest) (*cvm.DescribeInstancesResponse, error)
	StartInstances(ctx context.Context, secretId, secretKey, region string, request *cvm.StartInstancesRequest) (*cvm.StartInstancesResponse, error)
	StopInstances(ctx context.Context, secretId, secretKey, region string, request *cvm.StopInstancesRequest) (*cvm.StopInstancesResponse, error)
	RebootInstances(ctx context.Context, secretId, secretKey, region string, request *cvm.RebootInstancesRequest) (*cvm.RebootInstancesResponse, error)
	DeleteInstances(ctx context.Context, secretId, secretKey, region string, request *cvm.TerminateInstancesRequest) (*cvm.TerminateInstancesResponse, error)
}

// InstanceUseCase is a Instance usecase.
type InstanceUseCase struct {
	repo InstanceRepo
	log  *log.Helper
}

// NewInstanceUseCase new a Instance usecase.
func NewInstanceUseCase(repo InstanceRepo, logger log.Logger) *InstanceUseCase {
	return &InstanceUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *InstanceUseCase) CreateInstances(ctx context.Context, secretId, secretKey, region string, request *cvm.RunInstancesRequest) (*cvm.RunInstancesResponse, error) {
	return uc.repo.CreateInstances(ctx, secretId, secretKey, region, request)
}
func (uc *InstanceUseCase) ListInstances(ctx context.Context, secretId, secretKey, region string, request *cvm.DescribeInstancesRequest) (*cvm.DescribeInstancesResponse, error) {
	return uc.repo.ListInstances(ctx, secretId, secretKey, region, request)
}
func (uc *InstanceUseCase) StartInstances(ctx context.Context, secretId, secretKey, region string, request *cvm.StartInstancesRequest) (*cvm.StartInstancesResponse, error) {
	return uc.repo.StartInstances(ctx, secretId, secretKey, region, request)
}
func (uc *InstanceUseCase) StopInstances(ctx context.Context, secretId, secretKey, region string, request *cvm.StopInstancesRequest) (*cvm.StopInstancesResponse, error) {
	return uc.repo.StopInstances(ctx, secretId, secretKey, region, request)
}
func (uc *InstanceUseCase) RebootInstances(ctx context.Context, secretId, secretKey, region string, request *cvm.RebootInstancesRequest) (*cvm.RebootInstancesResponse, error) {
	return uc.repo.RebootInstances(ctx, secretId, secretKey, region, request)
}
func (uc *InstanceUseCase) DeleteInstances(ctx context.Context, secretId, secretKey, region string, request *cvm.TerminateInstancesRequest) (*cvm.TerminateInstancesResponse, error) {
	return uc.repo.DeleteInstances(ctx, secretId, secretKey, region, request)
}
