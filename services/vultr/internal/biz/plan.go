package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/vultr/govultr/v2"
)

type PlanResponse interface {
	ListPlans(context.Context, string, string, *govultr.ListOptions) ([]govultr.Plan, *govultr.Meta, error)
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

// ListAll List All Regions
func (uc *PlanUseCase) ListPlans(ctx context.Context, accessToken, planType string, request *govultr.ListOptions) ([]govultr.Plan, *govultr.Meta, error) {
	return uc.repo.ListPlans(ctx, accessToken, planType, request)
}
