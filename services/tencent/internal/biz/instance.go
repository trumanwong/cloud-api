package biz

import (
	"context"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"

	"github.com/go-kratos/kratos/v2/log"
)

// InstanceRepo is a Instance repo.
type InstanceRepo interface {
	Create(ctx context.Context, secretId, secretKey, region string, request *cvm.RunInstancesRequest) (*cvm.RunInstancesResponse, error)
	ListInstance(ctx context.Context, secretId, secretKey, region string, request *cvm.DescribeInstancesRequest) (*cvm.DescribeInstancesResponse, error)
	StartInstance(ctx context.Context, secretId, secretKey, region string, request *cvm.StartInstancesRequest) (*cvm.StartInstancesResponse, error)
	StopInstance(ctx context.Context, secretId, secretKey, region string, request *cvm.StopInstancesRequest) (*cvm.StopInstancesResponse, error)
	RebootInstance(ctx context.Context, secretId, secretKey, region string, request *cvm.RebootInstancesRequest) (*cvm.RebootInstancesResponse, error)
	DeleteInstance(ctx context.Context, secretId, secretKey, region string, request *cvm.TerminateInstancesRequest) (*cvm.TerminateInstancesResponse, error)
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

func (uc *InstanceUseCase) Create(ctx context.Context, secretId, secretKey, region string, request *cvm.RunInstancesRequest) (*cvm.RunInstancesResponse, error) {
	return uc.repo.Create(ctx, secretId, secretKey, region, request)
}
func (uc *InstanceUseCase) ListInstance(ctx context.Context, secretId, secretKey, region string, request *cvm.DescribeInstancesRequest) (*cvm.DescribeInstancesResponse, error) {
	return uc.repo.ListInstance(ctx, secretId, secretKey, region, request)
}
func (uc *InstanceUseCase) StartInstance(ctx context.Context, secretId, secretKey, region string, request *cvm.StartInstancesRequest) (*cvm.StartInstancesResponse, error) {
	return uc.repo.StartInstance(ctx, secretId, secretKey, region, request)
}
func (uc *InstanceUseCase) StopInstance(ctx context.Context, secretId, secretKey, region string, request *cvm.StopInstancesRequest) (*cvm.StopInstancesResponse, error) {
	return uc.repo.StopInstance(ctx, secretId, secretKey, region, request)
}
func (uc *InstanceUseCase) RebootInstance(ctx context.Context, secretId, secretKey, region string, request *cvm.RebootInstancesRequest) (*cvm.RebootInstancesResponse, error) {
	return uc.repo.RebootInstance(ctx, secretId, secretKey, region, request)
}
func (uc *InstanceUseCase) DeleteInstance(ctx context.Context, secretId, secretKey, region string, request *cvm.TerminateInstancesRequest) (*cvm.TerminateInstancesResponse, error) {
	return uc.repo.DeleteInstance(ctx, secretId, secretKey, region, request)
}
