package service

import (
	"context"
	"github.com/vultr/govultr/v2"
	v1 "vultr/api/instance/v1"
)

func (service *InstanceService) ListPlans(ctx context.Context, request *v1.ListPlansRequest) (*v1.ListPlansResponse, error) {
	result, meta, err := service.pc.ListPlans(ctx, request.AccessToken, request.PlanType, &govultr.ListOptions{
		PerPage: int(request.PerPage),
		Cursor:  request.Cursor,
	})
	if err != nil {
		return nil, err
	}

	var link *v1.Meta_Link
	if meta.Links != nil {
		link = &v1.Meta_Link{
			Next: meta.Links.Next,
			Prev: meta.Links.Prev,
		}
	}
	listPlansResponse := &v1.ListPlansResponse{
		Plans: make([]*v1.ListPlansResponse_Plan, len(result)),
		Meta: &v1.Meta{
			Total: int32(meta.Total),
			Link:  link,
		},
	}
	for i, plan := range result {
		listPlansResponse.Plans[i] = &v1.ListPlansResponse_Plan{
			Id:          plan.ID,
			VcpuCount:   int32(plan.VCPUCount),
			Ram:         int32(plan.RAM),
			Disk:        int32(plan.Disk),
			Bandwidth:   int32(plan.Bandwidth),
			MonthlyCost: int32(plan.MonthlyCost),
			PlanType:    plan.Type,
			Locations:   plan.Locations,
			DiskCount:   int32(plan.DiskCount),
		}
	}
	return listPlansResponse, nil
}
