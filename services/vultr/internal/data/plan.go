package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/vultr/govultr/v2"
	"vultr/internal/biz"
)

type instanceTypeResponse struct {
	data *Data
	log  *log.Helper
}

// NewInstanceTypeRepo .
func NewInstanceTypeRepo(data *Data, logger log.Logger) biz.PlanResponse {
	return &instanceTypeResponse{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *instanceTypeResponse) ListPlans(ctx context.Context, accessToken, planType string, request *govultr.ListOptions) ([]govultr.Plan, *govultr.Meta, error) {
	client := newClient(
		accessToken,
	)
	result, meta, err := client.Plan.List(ctx, planType, request)
	if err != nil {
		return nil, nil, err
	}
	return result, meta, nil
}
