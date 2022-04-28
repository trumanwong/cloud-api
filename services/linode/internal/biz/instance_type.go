package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/linode/linodego"
)

type PlanResponse interface {
	ListInstanceTypes(context.Context, string, *linodego.ListOptions) (*ListInstanceTypesResponse, error)
}

type ListInstanceTypesResponse struct {
	InstanceTypes []linodego.LinodeType
}

// InstanceTypeUseCase is a Plan UseCase.
type InstanceTypeUseCase struct {
	repo PlanResponse
	log  *log.Helper
}

// NewInstanceTypeUseCase new a Plan UseCase.
func NewInstanceTypeUseCase(repo PlanResponse, logger log.Logger) *InstanceTypeUseCase {
	return &InstanceTypeUseCase{repo: repo, log: log.NewHelper(logger)}
}

// ListInstanceTypes List All InstanceTypes
func (uc *InstanceTypeUseCase) ListInstanceTypes(ctx context.Context, accessToken string, request *linodego.ListOptions) (*ListInstanceTypesResponse, error) {
	return uc.repo.ListInstanceTypes(ctx, accessToken, request)
}
