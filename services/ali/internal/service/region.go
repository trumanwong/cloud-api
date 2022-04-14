package service

import (
	v1 "ali/api/instance/v1"
	"ali/internal/biz"
	"context"
)

func (service *InstanceService) ListRegion(ctx context.Context, request *v1.ListRegionRequest) (*v1.ListRegionResponse, error) {
	regions, err := service.rc.ListAll(ctx, &biz.ListRegionRequest{
		AccessKeyId:        request.AccessKeyId,
		AccessKeySecret:    request.AccessKeySecret,
		InstanceChargeType: request.InstanceChargeType,
		ResourceType:       request.ResourceType,
		AcceptLanguage:     request.AcceptLanguage,
	})
	if err != nil {
		return nil, err
	}

	listRegionResponse := &v1.ListRegionResponse{
		Regions: make([]*v1.ListRegionResponse_Region, len(regions)),
	}
	for i, region := range regions {
		listRegionResponse.Regions[i] = &v1.ListRegionResponse_Region{
			RegionEndPoint: region.RegionEndPoint,
			LocalName:      region.LocalName,
			RegionId:       region.RegionId,
		}
	}
	return listRegionResponse, nil
}
