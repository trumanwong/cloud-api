package service

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func (service *InstanceService) ListRegion(ctx context.Context, request *v1.ListRegionRequest) (*v1.ListRegionResponse, error) {
	regions, err := service.rc.ListAll(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, &ec2.DescribeRegionsInput{
		AllRegions:  nil,
		DryRun:      nil,
		Filters:     nil,
		RegionNames: nil,
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
