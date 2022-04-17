package service

import (
	"context"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	v1 "tencent/api/instance/v1"
)

func (service *InstanceService) ListRegion(ctx context.Context, request *v1.ListRegionRequest) (*v1.ListRegionResponse, error) {
	response, err := service.rc.ListAll(ctx, request.SecretId, request.SecretKey, cvm.NewDescribeRegionsRequest())
	if err != nil {
		return nil, err
	}

	listRegionResponse := &v1.ListRegionResponse{
		Regions: make([]*v1.ListRegionResponse_Region, len(response.Response.RegionSet)),
	}

	for i, region := range response.Response.RegionSet {
		listRegionResponse.Regions[i] = &v1.ListRegionResponse_Region{
			Region:      *region.Region,
			RegionName:  *region.RegionName,
			RegionState: *region.RegionState,
		}
	}
	return listRegionResponse, nil
}
