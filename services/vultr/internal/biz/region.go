package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/vultr/govultr/v2"
)

type ListRegionsResponse struct {
	Regions []govultr.Region `json:"regions"`
	Meta    *govultr.Meta    `json:"meta"`
}

type RegionResponse interface {
	ListRegions(context.Context, string, *govultr.ListOptions) (*ListRegionsResponse, error)
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

// ListRegions List Regions
func (uc *RegionUseCase) ListRegions(ctx context.Context, accessToken string, request *govultr.ListOptions) (*ListRegionsResponse, error) {
	return uc.repo.ListRegions(ctx, accessToken, request)
}
