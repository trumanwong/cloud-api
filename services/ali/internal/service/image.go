package service

import (
	v1 "ali/api/instance/v1"
	"ali/pkg/util/convert"
	"context"
)

func (service *InstanceService) ListImages(ctx context.Context, request *v1.ListImagesRequest) (*v1.ListImagesResponse, error) {
	result, err := service.ic.ListImages(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, request.RegionId)
	if err != nil {
		return nil, err
	}

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.ListImagesResponse{Data: data}, nil
}
