package service

import (
	v1 "ali/api/instance/v1"
	"context"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

func (service *InstanceService) ListRegion(ctx context.Context, request *v1.ListRegionRequest) (*v1.ListRegionResponse, error) {
	regions, err := service.rc.ListAll(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, &ecs20140526.DescribeRegionsRequest{
		InstanceChargeType: tea.String(request.InstanceChargeType),
		ResourceType:       tea.String(request.ResourceType),
		AcceptLanguage:     tea.String(request.AcceptLanguage),
	})
	if err != nil {
		return nil, err
	}

	listRegionResponse := &v1.ListRegionResponse{
		Regions: make([]*v1.ListRegionResponse_Region, len(regions)),
	}
	for i, region := range regions {
		listRegionResponse.Regions[i] = &v1.ListRegionResponse_Region{
			RegionEndPoint: *region.RegionEndpoint,
			LocalName:      *region.LocalName,
			RegionId:       *region.RegionId,
		}
	}
	return listRegionResponse, nil
}
