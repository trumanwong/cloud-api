package biz

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/subscriptions"
	"github.com/go-kratos/kratos/v2/log"
)

type InstanceTypeResponse interface {
	ListInstanceType(context.Context, string, string, string, string, bool) (*[]subscriptions.Location, error)
}

// InstanceTypeUseCase is a InstanceType UseCase.
type InstanceTypeUseCase struct {
	repo InstanceTypeResponse
	log  *log.Helper
}

// NewInstanceTypeUseCase new a InstanceType UseCase.
func NewInstanceTypeUseCase(repo InstanceTypeResponse, logger log.Logger) *InstanceTypeUseCase {
	return &InstanceTypeUseCase{repo: repo, log: log.NewHelper(logger)}
}

// ListInstanceType List All Regions
func (uc *InstanceTypeUseCase) ListInstanceType(ctx context.Context, tenantID, clientID, clientSecret, subscriptionId string, includeExtendedLocations bool) (*[]subscriptions.Location, error) {
	return uc.repo.ListInstanceType(ctx, tenantID, clientID, clientSecret, subscriptionId, includeExtendedLocations)
}
