package service

import (
	v1 "ali/api/instance/v1"
	"context"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

func (service *InstanceService) ListRegions(ctx context.Context, request *v1.ListRegionsRequest) (*v1.ListRegionsResponse, error) {
	regions, err := service.rc.ListAll(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, &ecs20140526.DescribeRegionsRequest{
		InstanceChargeType: tea.String(request.InstanceChargeType),
		ResourceType:       tea.String(request.ResourceType),
		AcceptLanguage:     tea.String(request.AcceptLanguage),
	})
	if err != nil {
		return nil, err
	}

	listRegionsResponse := &v1.ListRegionsResponse{
		Regions: make([]*v1.ListRegionsResponse_Region, len(regions)),
	}
	for i, region := range regions {
		listRegionsResponse.Regions[i] = &v1.ListRegionsResponse_Region{
			RegionEndPoint: *region.RegionEndpoint,
			LocalName:      *region.LocalName,
			RegionId:       *region.RegionId,
		}
	}
	return listRegionsResponse, nil
}
