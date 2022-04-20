package service

import (
	"github.com/google/wire"
	v1 "vultr/api/instance/v1"
	"vultr/internal/biz"
)

// InstanceService is a instance service.
type InstanceService struct {
	v1.UnimplementedInstanceServer

	uc *biz.InstanceUseCase
	rc *biz.RegionUseCase
	ic *biz.ImageUseCase
	pc *biz.PlanUseCase
}

// NewInstanceService new a Instance service.
func NewInstanceService(uc *biz.InstanceUseCase, rc *biz.RegionUseCase, ic *biz.ImageUseCase, pc *biz.PlanUseCase) *InstanceService {
	return &InstanceService{uc: uc, rc: rc, ic: ic, pc: pc}
}

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewInstanceService)
