package service

import (
	v1 "aws/api/instance/v1"
	"aws/pkg/util/convert"
	"context"
)

func (service *InstanceService) ListImages(ctx context.Context, request *v1.ListImagesRequest) (*v1.ListImagesResponse, error) {
	result, err := service.ic.ListImages(ctx, request.AccessKeyId, request.SecretKey, request.Region)
	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.ListImagesResponse{Data: data}, nil
}
