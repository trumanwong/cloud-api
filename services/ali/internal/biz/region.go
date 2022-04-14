package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Region struct {
	RegionEndPoint string
	LocalName      string
	RegionId       string
}

type ListRegionRequest struct {
	AccessKeyId        string
	AccessKeySecret    string
	InstanceChargeType string
	ResourceType       string
	AcceptLanguage     string
}

type RegionResponse interface {
	ListAll(ctx context.Context, request *ListRegionRequest) ([]*Region, error)
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
func (uc *RegionUseCase) ListAll(ctx context.Context, g *ListRegionRequest) ([]*Region, error) {
	return uc.repo.ListAll(ctx, g)
}
