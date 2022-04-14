package data

import (
	"ali/internal/biz"
	"context"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/go-kratos/kratos/v2/log"
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

func (r *regionResponse) ListAll(ctx context.Context, accessKeyId, accessKeySecret string, request *ecs20140526.DescribeRegionsRequest) ([]*ecs20140526.DescribeRegionsResponseBodyRegionsRegion, error) {
	client, err := createClient(
		tea.String(accessKeyId),
		tea.String(accessKeySecret),
	)
	if err != nil {
		return nil, err
	}
	result, err := client.DescribeRegions(request)
	if err != nil {
		return nil, err
	}
	return result.Body.Regions.Region, nil
}
