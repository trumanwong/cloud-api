package service

import (
	"context"
	v1 "huawei/api/instance/v1"
	"huawei/pkg/util/convert"
)

func (service *InstanceService) ListRegions(ctx context.Context, request *v1.ListRegionsRequest) (*v1.ListRegionsResponse, error) {
	result := service.rc.ListRegions(ctx, request.RegionType)

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.ListRegionsResponse{Data: data}, nil
}
