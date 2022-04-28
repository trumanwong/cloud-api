package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
)

type RegionResponse interface {
	ListRegions(context.Context, string) *ListRegionResponse
}

// RegionUseCase is a Region UseCase.
type RegionUseCase struct {
	repo RegionResponse
	log  *log.Helper
}

type ListRegionResponse struct {
	Regions []*region.Region `json:"regions"`
}

// NewRegionUseCase new a Region UseCase.
func NewRegionUseCase(repo RegionResponse, logger log.Logger) *RegionUseCase {
	return &RegionUseCase{repo: repo, log: log.NewHelper(logger)}
}

// ListRegions List All Regions
func (uc *RegionUseCase) ListRegions(ctx context.Context, regionType string) *ListRegionResponse {
	return uc.repo.ListRegions(ctx, regionType)
}
