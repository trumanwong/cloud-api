package biz

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-12-01/compute"
	"github.com/go-kratos/kratos/v2/log"
)

type ListInstanceTypesResponse struct {
	InstanceTypes []compute.VirtualMachineSizeTypes `json:"instance_types"`
}

type InstanceTypeResponse interface {
	ListInstanceType(context.Context) (*ListInstanceTypesResponse, error)
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

// ListInstanceType List All InstanceTypes
func (uc *InstanceTypeUseCase) ListInstanceType(ctx context.Context) (*ListInstanceTypesResponse, error) {
	return uc.repo.ListInstanceType(ctx)
}
