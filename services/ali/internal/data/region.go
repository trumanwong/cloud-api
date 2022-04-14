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

func (r *regionResponse) ListAll(ctx context.Context, request *biz.ListRegionRequest) ([]*biz.Region, error) {
	client, err := createClient(
		tea.String(request.AccessKeyId),
		tea.String(request.AccessKeySecret),
	)
	if err != nil {
		return nil, err
	}
	describeRegionRequest := &ecs20140526.DescribeRegionsRequest{
		AcceptLanguage:     tea.String(request.AcceptLanguage),
		InstanceChargeType: tea.String(request.InstanceChargeType),
		ResourceType:       tea.String(request.ResourceType),
	}
	result, err := client.DescribeRegions(describeRegionRequest)
	if err != nil {
		return nil, err
	}
	regions := make([]*biz.Region, len(result.Body.Regions.Region))
	for i, region := range result.Body.Regions.Region {
		regions[i] = &biz.Region{
			RegionEndPoint: *region.RegionEndpoint,
			LocalName:      *region.LocalName,
			RegionId:       *region.RegionId,
		}
	}
	return regions, nil
}
