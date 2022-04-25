package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
)

type InstanceTypeResponse interface {
	ListInstanceTypes(context.Context, string, string, string, string, *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error)
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
func (uc *InstanceTypeUseCase) ListInstanceTypes(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	return uc.repo.ListInstanceTypes(ctx, accessKey, secretKey, regionId, projectId, request)
}
