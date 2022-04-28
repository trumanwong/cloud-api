package service

import (
	"context"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ims/v2/model"
	v1 "huawei/api/instance/v1"
	"huawei/pkg/util/convert"
)

func (service *InstanceService) ListImages(ctx context.Context, request *v1.ListImagesRequest) (*v1.ListImagesResponse, error) {
	result, err := service.ic.ListImages(ctx, request.AccessKey, request.SecretKey, request.RegionId, &model.GlanceListImagesRequest{
		Limit:  request.Limit,
		Marker: request.Marker,
	})
	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.ListImagesResponse{Data: data}, nil
}
