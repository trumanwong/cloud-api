package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/vultr/govultr/v2"
	"vultr/internal/biz"
)

type regionResponse struct {
	data *Data
	log  *log.Helper
}

// NewRegionRepo .
func NewRegionRepo(data *Data, logger log.Logger) biz.RegionResponse {
	return &regionResponse{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *regionResponse) ListRegions(ctx context.Context, accessToken string, request *govultr.ListOptions) (*biz.ListRegionsResponse, error) {
	client := getClient(accessToken)
	regions, meta, err := client.Region.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return &biz.ListRegionsResponse{Regions: regions, Meta: meta}, nil
}
