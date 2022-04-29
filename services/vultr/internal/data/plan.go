package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/vultr/govultr/v2"
	"vultr/internal/biz"
)

type planResponse struct {
	data *Data
	log  *log.Helper
}

// NewPlanRepo .
func NewPlanRepo(data *Data, logger log.Logger) biz.PlanResponse {
	return &planResponse{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *planResponse) ListPlans(ctx context.Context, accessToken, planType string, request *govultr.ListOptions) (*biz.ListPlansResponse, error) {
	client := getClient(ctx, accessToken)
	plans, meta, err := client.Plan.List(ctx, planType, request)
	if err != nil {
		return nil, err
	}
	return &biz.ListPlansResponse{
		Plans: plans,
		Meta:  meta,
	}, nil
}
