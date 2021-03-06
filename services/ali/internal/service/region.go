package service

import (
	v1 "ali/api/instance/v1"
	"ali/pkg/util/convert"
	"context"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
)

func (service *InstanceService) ListRegions(ctx context.Context, request *v1.ListRegionsRequest) (*v1.ListRegionsResponse, error) {
	result, err := service.rc.ListRegions(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, &ecs20140526.DescribeRegionsRequest{
		InstanceChargeType: request.InstanceChargeType,
		ResourceType:       request.ResourceType,
		AcceptLanguage:     request.AcceptLanguage,
	})
	if err != nil {
		return nil, err
	}

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.ListRegionsResponse{Data: data}, nil
}
