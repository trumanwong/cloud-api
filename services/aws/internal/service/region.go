package service

import (
	v1 "aws/api/instance/v1"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func (service *InstanceService) ListRegion(ctx context.Context, request *v1.ListRegionRequest) (*v1.ListRegionResponse, error) {
	regions, err := service.rc.ListAll(ctx, request.AccessKeyId, request.SecretKey, &ec2.DescribeRegionsInput{})
	if err != nil {
		return nil, err
	}

	listRegionResponse := &v1.ListRegionResponse{
		Regions: make([]*v1.ListRegionResponse_Region, len(regions)),
	}
	for i, region := range regions {
		listRegionResponse.Regions[i] = &v1.ListRegionResponse_Region{
			OptInStatus:    *region.OptInStatus,
			RegionEndPoint: *region.Endpoint,
			RegionName:     *region.RegionName,
		}
	}
	return listRegionResponse, nil
}
