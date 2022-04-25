package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
)

type SubnetResponse interface {
	CreateSubnet(context.Context, string, string, string, string, *model.CreateSubnetRequest) (*model.CreateSubnetResponse, error)
	ListSubnets(context.Context, string, string, string, string, *model.ListSubnetsRequest) (*model.ListSubnetsResponse, error)
	DeleteSubnet(context.Context, string, string, string, string, *model.DeleteSubnetRequest) (*model.DeleteSubnetResponse, error)
	UpdateSubnet(context.Context, string, string, string, string, *model.UpdateSubnetRequest) (*model.UpdateSubnetResponse, error)
	ShowSubnet(context.Context, string, string, string, string, *model.ShowSubnetRequest) (*model.ShowSubnetResponse, error)
}

// SubnetUseCase is a Subnet UseCase.
type SubnetUseCase struct {
	repo SubnetResponse
	log  *log.Helper
}

// NewSubnetUseCase new a Subnet UseCase.
func NewSubnetUseCase(repo RegionResponse, logger log.Logger) *RegionUseCase {
	return &RegionUseCase{repo: repo, log: log.NewHelper(logger)}
}

// CreateSubnet create Subnet
func (uc *SubnetUseCase) CreateSubnet(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.CreateSubnetRequest) (*model.CreateSubnetResponse, error) {
	return uc.repo.CreateSubnet(ctx, accessKey, secretKey, regionId, projectId, request)
}

// ListSubnets List all Subnets
func (uc *SubnetUseCase) ListSubnets(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.ListSubnetsRequest) (*model.ListSubnetsResponse, error) {
	return uc.repo.ListSubnets(ctx, accessKey, secretKey, regionId, projectId, request)
}

// DeleteSubnet delete Subnet
func (uc *SubnetUseCase) DeleteSubnet(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.DeleteSubnetRequest) (*model.DeleteSubnetResponse, error) {
	return uc.repo.DeleteSubnet(ctx, accessKey, secretKey, regionId, projectId, request)
}

// UpdateSubnet update Subnet
func (uc *SubnetUseCase) UpdateSubnet(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.UpdateSubnetRequest) (*model.UpdateSubnetResponse, error) {
	return uc.repo.UpdateSubnet(ctx, accessKey, secretKey, regionId, projectId, request)
}

// ShowSubnet show a Subnet
func (uc *SubnetUseCase) ShowSubnet(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.ShowSubnetRequest) (*model.ShowSubnetResponse, error) {
	return uc.repo.ShowSubnet(ctx, accessKey, secretKey, regionId, projectId, request)
}
