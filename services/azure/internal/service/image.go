package service

import (
	v1 "azure/api/instance/v1"
	"azure/pkg/util/convert"
	"context"
)

func (service *InstanceService) ListImages(ctx context.Context, request *v1.ListImagesRequest) (*v1.ListImagesResponse, error) {
	result, err := service.ic.ListImages(ctx, request.TenantID, request.ClientID, request.ClientSecret, request.SubscriptionId)
	if err != nil {
		return nil, err
	}

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}

	return &v1.ListImagesResponse{Data: data}, nil
}
