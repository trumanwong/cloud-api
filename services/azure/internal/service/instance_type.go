package service

import (
	v1 "azure/api/instance/v1"
	"azure/pkg/util/convert"
	"context"
)

func (service *InstanceService) ListInstanceTypes(ctx context.Context, request *v1.ListInstanceTypesRequest) (*v1.ListInstanceTypesResponse, error) {
	result, err := service.itc.ListInstanceTypes(ctx, request.TenantID, request.ClientID, request.ClientSecret, request.SubscriptionId, request.Filter, request.IncludeExtendedLocations)
	if err != nil {
		return nil, err
	}

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}

	return &v1.ListInstanceTypesResponse{Data: data}, nil
}
