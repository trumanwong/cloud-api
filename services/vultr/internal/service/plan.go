package service

import (
	"context"
	"github.com/vultr/govultr/v2"
	v1 "vultr/api/instance/v1"
	"vultr/pkg/util/convert"
)

func (service *InstanceService) ListPlans(ctx context.Context, request *v1.ListPlansRequest) (*v1.ListPlansResponse, error) {
	result, err := service.pc.ListPlans(ctx, request.AccessToken, request.PlanType, &govultr.ListOptions{
		PerPage: int(request.PerPage),
		Cursor:  request.Cursor,
	})
	if err != nil {
		return nil, err
	}

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.ListPlansResponse{Data: data}, nil
}
