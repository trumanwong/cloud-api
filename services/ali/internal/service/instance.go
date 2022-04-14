package service

import (
	v1 "ali/api/instance/v1"
	"ali/internal/biz"
	"context"
)

// CreateInstance Create Instance
func (service *InstanceService) CreateInstance(ctx context.Context, in *v1.CreateInstanceRequest) (*v1.CreateInstanceResponse, error) {
	g, err := service.uc.CreateInstance(ctx, &biz.CreateInstanceRequest{
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
