package service

import (
	"context"
	"github.com/vultr/govultr/v2"
	v1 "vultr/api/instance/v1"
)

func (service *InstanceService) ListRegions(ctx context.Context, request *v1.ListRegionsRequest) (*v1.ListRegionsResponse, error) {
	regions, meta, err := service.rc.ListRegions(ctx, request.AccessToken, &govultr.ListOptions{
		PerPage: int(request.PerPage),
		Cursor:  request.Cursor,
	})
	if err != nil {
		return nil, err
	}
	var link *v1.Meta_Link
	if meta.Links != nil {
		link = &v1.Meta_Link{
			Next: meta.Links.Next,
			Prev: meta.Links.Prev,
		}
	}
	listRegionsResponse := &v1.ListRegionsResponse{
		Regions: make([]*v1.ListRegionsResponse_Region, len(regions)),
		Meta: &v1.Meta{
			Total: int32(meta.Total),
			Link:  link,
		},
	}
	for i, region := range regions {
		listRegionsResponse.Regions[i] = &v1.ListRegionsResponse_Region{
			Id:        region.ID,
			City:      region.City,
			Country:   region.Country,
			Continent: region.Continent,
			Options:   region.Options,
		}
	}
	return listRegionsResponse, nil
}
