package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

type InstanceTypeResponse interface {
	ListInstanceTypes(context.Context, string, string, string, *cvm.DescribeInstanceTypeConfigsRequest) (*cvm.DescribeInstanceTypeConfigsResponse, error)
}

// InstanceTypeUseCase is Image Region UseCase.
type InstanceTypeUseCase struct {
	repo InstanceTypeResponse
	log  *log.Helper
}

// NewInstanceTypeUseCase new a Image UseCase.
func NewInstanceTypeUseCase(repo InstanceTypeResponse, logger log.Logger) *InstanceTypeUseCase {
	return &InstanceTypeUseCase{repo: repo, log: log.NewHelper(logger)}
}

// ListInstanceTypes List All InstanceType
func (uc *InstanceTypeUseCase) ListInstanceTypes(ctx context.Context, secretId, secretKey, region string, request *cvm.DescribeInstanceTypeConfigsRequest) (*cvm.DescribeInstanceTypeConfigsResponse, error) {
	return uc.repo.ListInstanceTypes(ctx, secretId, secretKey, region, request)
}
