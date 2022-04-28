package biz

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-12-01/compute"
	"github.com/go-kratos/kratos/v2/log"
)

type InstanceTypeResponse interface {
	ListInstanceTypes(context.Context, string, string, string, string, string, string) (*compute.ResourceSkusResultPage, error)
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

// ListInstanceTypes List All InstanceTypes
func (uc *InstanceTypeUseCase) ListInstanceTypes(ctx context.Context, tenantID, clientID, clientSecret, subscriptionId, location, includeExtendedLocations string) (*compute.ResourceSkusResultPage, error) {
	return uc.repo.ListInstanceTypes(ctx, tenantID, clientID, clientSecret, subscriptionId, location, includeExtendedLocations)
}
