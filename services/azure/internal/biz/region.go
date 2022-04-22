package biz

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/subscriptions"
	"github.com/go-kratos/kratos/v2/log"
)

type RegionResponse interface {
	ListRegions(context.Context, string, string, string, string, bool) (*[]subscriptions.Location, error)
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
func (uc *RegionUseCase) ListRegions(ctx context.Context, tenantID, clientID, clientSecret, subscriptionId string, includeExtendedLocations bool) (*[]subscriptions.Location, error) {
	return uc.repo.ListRegions(ctx, tenantID, clientID, clientSecret, subscriptionId, includeExtendedLocations)
}
