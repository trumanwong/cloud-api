package service

import (
	v1 "aws/api/instance/v1"
	"aws/internal/biz"
	"github.com/google/wire"
)

// InstanceService is a greeter service.
type InstanceService struct {
	v1.UnimplementedInstanceServer

	uc  *biz.InstanceUseCase
	rc  *biz.RegionUseCase
	ic  *biz.ImageUseCase
	itc *biz.InstanceTypeUseCase
}

// NewInstanceService new a Instance service.
func NewInstanceService(uc *biz.InstanceUseCase, rc *biz.RegionUseCase, ic *biz.ImageUseCase, itc *biz.InstanceTypeUseCase) *InstanceService {
	return &InstanceService{uc: uc, rc: rc, ic: ic, itc: itc}
}

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewInstanceService)
