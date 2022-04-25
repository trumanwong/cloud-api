package service

import (
	"github.com/google/wire"
	v1 "huawei/api/instance/v1"
	"huawei/internal/biz"
)

// InstanceService is a instance service.
type InstanceService struct {
	v1.UnimplementedInstanceServer

	uc  *biz.InstanceUseCase
	rc  *biz.RegionUseCase
	ic  *biz.ImageUseCase
	itc *biz.InstanceTypeUseCase
	vc  *biz.VpcUseCase
}

// NewInstanceService new a Instance service.
func NewInstanceService(uc *biz.InstanceUseCase, rc *biz.RegionUseCase, ic *biz.ImageUseCase, itc *biz.InstanceTypeUseCase, vc *biz.VpcUseCase) *InstanceService {
	return &InstanceService{uc: uc, rc: rc, ic: ic, itc: itc, vc: vc}
}

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewInstanceService)
