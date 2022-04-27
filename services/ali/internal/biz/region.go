package biz

import (
	"context"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/go-kratos/kratos/v2/log"
)

type RegionResponse interface {
	ListRegions(context.Context, *string, *string, *string, *ecs20140526.DescribeRegionsRequest) (*ecs20140526.DescribeRegionsResponse, error)
}

// RegionUseCase is a Region UseCase.
type RegionUseCase struct {
	repo RegionResponse
	log  *log.Helper
}

// NewRegionUseCase new a Region UseCase.
func NewRegionUseCase(repo RegionResponse, logger log.Logger) *RegionUseCase {
	return &RegionUseCase{repo: repo, log: log.NewHelper(logger)}
}

// ListRegions List All Regions
func (uc *RegionUseCase) ListRegions(ctx context.Context, accessKeyId, accessKeySecret, endpoint *string, g *ecs20140526.DescribeRegionsRequest) (*ecs20140526.DescribeRegionsResponse, error) {
	return uc.repo.ListRegions(ctx, accessKeyId, accessKeySecret, endpoint, g)
}
