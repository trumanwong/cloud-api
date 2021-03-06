package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/linode/linodego"
)

type ListRegionResponse struct {
	Regions []linodego.Region
}

type RegionResponse interface {
	ListRegions(context.Context, string, *linodego.ListOptions) (*ListRegionResponse, error)
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
func (uc *RegionUseCase) ListRegions(ctx context.Context, accessToken string, request *linodego.ListOptions) (*ListRegionResponse, error) {
	return uc.repo.ListRegions(ctx, accessToken, request)
}
