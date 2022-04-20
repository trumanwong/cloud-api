package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/linode/linodego"
)

type PlanResponse interface {
	ListInstanceTypes(context.Context, string, *linodego.ListOptions) ([]linodego.LinodeType, error)
}

// PlanUseCase is a Plan UseCase.
type PlanUseCase struct {
	repo PlanResponse
	log  *log.Helper
}

// NewPlanUseCase new a Plan UseCase.
func NewPlanUseCase(repo PlanResponse, logger log.Logger) *PlanUseCase {
	return &PlanUseCase{repo: repo, log: log.NewHelper(logger)}
}

// ListInstanceTypes List All InstanceTypes
func (uc *PlanUseCase) ListInstanceTypes(ctx context.Context, accessToken string, request *linodego.ListOptions) ([]linodego.LinodeType, error) {
	return uc.repo.ListInstanceTypes(ctx, accessToken, request)
}
