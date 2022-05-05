package service

import (
	v1 "azure/api/instance/v1"
	"azure/pkg/util/convert"
	"context"
)

func (service *InstanceService) ListRegions(ctx context.Context, request *v1.ListRegionsRequest) (*v1.ListRegionsResponse, error) {
	result, err := service.rc.ListRegions(ctx, request.TenantId, request.ClientId, request.ClientSecret, request.SubscriptionId, request.IncludeExtendedLocations)
	if err != nil {
		return nil, err
	}

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}

	return &v1.ListRegionsResponse{Data: data}, nil
}
