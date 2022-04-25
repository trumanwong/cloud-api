package service

import (
	"context"
	v1 "huawei/api/instance/v1"
)

func (service *InstanceService) ListRegions(ctx context.Context, request *v1.ListRegionsRequest) (*v1.ListRegionsResponse, error) {
	regions := service.rc.ListRegions(ctx, request.RegionType)

	listRegionsResponse := &v1.ListRegionsResponse{
		Regions: make([]*v1.ListRegionsResponse_Region, len(regions)),
	}
	for i, region := range regions {
		listRegionsResponse.Regions[i] = &v1.ListRegionsResponse_Region{
			Id:       region.Id,
			EndPoint: region.Endpoint,
		}
	}
	return listRegionsResponse, nil
}
