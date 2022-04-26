package service

import (
	v1 "ali/api/instance/v1"
	"context"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

func (service *InstanceService) ListRegions(ctx context.Context, request *v1.ListRegionsRequest) (*v1.ListRegionsResponse, error) {
	result, err := service.rc.ListRegions(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, &ecs20140526.DescribeRegionsRequest{
		InstanceChargeType: tea.String(request.InstanceChargeType),
		ResourceType:       tea.String(request.ResourceType),
		AcceptLanguage:     tea.String(request.AcceptLanguage),
	})
	if err != nil {
		return nil, err
	}

	var regions []*v1.ListRegionsResponse_Region
	if result.Body.Regions != nil {
		regions = make([]*v1.ListRegionsResponse_Region, len(regions))
		for i, v := range result.Body.Regions.Region {
			regions[i] = &v1.ListRegionsResponse_Region{
				RegionEndPoint: *v.RegionEndpoint,
				LocalName:      *v.LocalName,
				RegionId:       *v.RegionId,
				Status:         *v.Status,
			}
		}
	}
	return &v1.ListRegionsResponse{Regions: regions}, nil
}
