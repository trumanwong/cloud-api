package service

import (
	"context"
	"github.com/linode/linodego"
	v1 "linode/api/instance/v1"
)

func (service *InstanceService) ListRegions(ctx context.Context, request *v1.ListRegionsRequest) (*v1.ListRegionsResponse, error) {
	options := &linodego.ListOptions{
		PageSize: int(request.PageSize),
		PageOptions: &linodego.PageOptions{
			Page: int(request.Page),
		},
	}
	regions, err := service.rc.ListRegions(ctx, request.AccessToken, options)
	if err != nil {
		return nil, err
	}
	listRegionsResponse := &v1.ListRegionsResponse{
		Regions: make([]*v1.ListRegionsResponse_Region, len(regions)),
		Page:    int32(options.Page),
		Pages:   int32(options.Pages),
		Results: int32(options.Results),
	}
	for i, region := range regions {
		listRegionsResponse.Regions[i] = &v1.ListRegionsResponse_Region{
			Capabilities: region.Capabilities,
			Country:      region.Country,
			Id:           region.ID,
			Resolvers: &v1.ListRegionsResponse_Region_Resolvers{
				Ipv4: region.Resolvers.IPv4,
				Ipv6: region.Resolvers.IPv6,
			},
			Status: region.Status,
		}
	}
	return listRegionsResponse, nil
}
