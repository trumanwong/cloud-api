package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Region struct {
	Id       string `json:"id"`
	Endpoint string `json:"endpoint"`
}

type RegionResponse interface {
	ListRegions(context.Context, string) []*Region
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
func (uc *RegionUseCase) ListRegions(ctx context.Context, regionType string) []*Region {
	return uc.repo.ListRegions(ctx, regionType)
}
