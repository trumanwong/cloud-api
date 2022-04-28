package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/linode/linodego"
	"linode/internal/biz"
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

func (r *regionResponse) ListRegions(ctx context.Context, accessToken string, request *linodego.ListOptions) (*biz.ListRegionResponse, error) {
	client := newClient(accessToken)
	regions, err := client.ListRegions(ctx, request)
	if err != nil {
		return nil, err
	}
	return &biz.ListRegionResponse{Regions: regions}, nil
}
