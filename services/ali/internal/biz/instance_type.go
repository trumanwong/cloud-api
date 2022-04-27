package biz

import (
	"context"
	"github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/go-kratos/kratos/v2/log"
)

type InstanceTypeResponse interface {
	ListInstanceTypes(context.Context, *string, *string, *string) (*client.DescribeInstanceTypesResponse, error)
}

// InstanceTypeUseCase is a InstanceType UseCase.
type InstanceTypeUseCase struct {
	repo InstanceTypeResponse
	log  *log.Helper
}

// NewInstanceTypeUseCase new a InstanceType UseCase.
func NewInstanceTypeUseCase(repo InstanceTypeResponse, logger log.Logger) *InstanceTypeUseCase {
	return &InstanceTypeUseCase{repo: repo, log: log.NewHelper(logger)}
}

// ListInstanceTypes List All InstanceTypes
func (uc *InstanceTypeUseCase) ListInstanceTypes(ctx context.Context, accessKeyId, accessKeySecret, endpoint *string) (*client.DescribeInstanceTypesResponse, error) {
	return uc.repo.ListInstanceTypes(ctx, accessKeyId, accessKeySecret, endpoint)
}
