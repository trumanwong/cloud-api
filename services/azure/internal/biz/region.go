package biz

import (
	"context"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/go-kratos/kratos/v2/log"
)

type RegionResponse interface {
	ListAll(ctx context.Context, accessKeyId, accessKeySecret, endpoint string, request *ecs20140526.DescribeRegionsRequest) ([]*ecs20140526.DescribeRegionsResponseBodyRegionsRegion, error)
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

// ListAll List All Regions
func (uc *RegionUseCase) ListAll(ctx context.Context, accessKeyId, accessKeySecret, endpoint string, g *ecs20140526.DescribeRegionsRequest) ([]*ecs20140526.DescribeRegionsResponseBodyRegionsRegion, error) {
	return uc.repo.ListAll(ctx, accessKeyId, accessKeySecret, endpoint, g)
}
