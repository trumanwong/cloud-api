package service

import (
	v1 "ali/api/instance/v1"
	"ali/internal/biz"
	"github.com/google/wire"
)

// InstanceService is a instance service.
type InstanceService struct {
	v1.UnimplementedInstanceServer

	uc *biz.InstanceUseCase
	rc *biz.RegionUseCase
	ic *biz.ImageUseCase
}

// NewInstanceService new a Instance service.
func NewInstanceService(uc *biz.InstanceUseCase, rc *biz.RegionUseCase, ic *biz.ImageUseCase) *InstanceService {
	return &InstanceService{uc: uc, rc: rc, ic: ic}
}

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewInstanceService)
