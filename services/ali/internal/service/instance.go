package service

import (
	"context"

	v1 "ali/api/instance/v1"
	"ali/internal/biz"
)

// InstanceService is a instance service.
type InstanceService struct {
	v1.UnimplementedInstanceServer

	uc *biz.InstanceUseCase
}

// NewInstanceService new a Instance service.
func NewInstanceService(uc *biz.InstanceUseCase) *InstanceService {
	return &InstanceService{uc: uc}
}

// Create Instance
func (s *InstanceService) Create(ctx context.Context, in *v1.CreateInstanceRequest) (*v1.CreateInstanceResponse, error) {
	g, err := s.uc.CreateInstance(ctx, &biz.CreateInstanceRequest{
		RegionId:       in.RegionId,
		ImageId:        in.ImageId,
		Name:           in.Name,
		InstanceType:   in.InstanceType,
		SystemDiskSize: in.SystemDiskSize,
		UniqueSuffix:   in.UniqueSuffix,
		Amount:         in.Amount,
		Password:       in.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateInstanceResponse{
		RequestId:      g.RequestId,
		OrderId:        g.OrderId,
		TradePrice:     g.TradePrice,
		InstanceIdSets: g.InstanceIdSets,
	}, nil
}
