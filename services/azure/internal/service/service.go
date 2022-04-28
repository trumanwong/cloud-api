package service

import (
	v1 "azure/api/instance/v1"
	"azure/internal/biz"
	"github.com/google/wire"
)

// InstanceService is a instance service.
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
