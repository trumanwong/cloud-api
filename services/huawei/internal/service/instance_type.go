package service

import (
	"context"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	v1 "huawei/api/instance/v1"
	"huawei/pkg/util/convert"
)

func (service *InstanceService) ListInstanceTypes(ctx context.Context, request *v1.ListInstanceTypesRequest) (*v1.ListInstanceTypesResponse, error) {
	result, err := service.itc.ListInstanceTypes(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.ListFlavorsRequest{})
	if err != nil {
		return nil, err
	}

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.ListInstanceTypesResponse{Data: data}, nil
}
