package service

import (
	"context"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	v1 "tencent/api/instance/v1"
)

func (service *InstanceService) ListRegions(ctx context.Context, request *v1.ListRegionsRequest) (*v1.ListRegionsResponse, error) {
	response, err := service.rc.ListRegions(ctx, request.SecretId, request.SecretKey, cvm.NewDescribeRegionsRequest())
	if err != nil {
		return nil, err
	}

	listRegionResponse := &v1.ListRegionsResponse{
		TotalCount: response.Response.TotalCount,
		Regions:    make([]*v1.ListRegionsResponse_Region, len(response.Response.RegionSet)),
		RequestId:  response.Response.RequestId,
	}

	for i, region := range response.Response.RegionSet {
		listRegionResponse.Regions[i] = &v1.ListRegionsResponse_Region{
			Region:      region.Region,
			RegionName:  region.RegionName,
			RegionState: region.RegionState,
		}
	}
	return listRegionResponse, nil
}
