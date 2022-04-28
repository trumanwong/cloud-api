package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
)

type VpcResponse interface {
	CreateVpc(context.Context, string, string, string, string, *model.CreateVpcRequest) (*model.CreateVpcResponse, error)
	ListVpcs(context.Context, string, string, string, string, *model.ListVpcsRequest) (*model.ListVpcsResponse, error)
	DeleteVpc(context.Context, string, string, string, string, *model.DeleteVpcRequest) (*model.DeleteVpcResponse, error)
	UpdateVpc(context.Context, string, string, string, string, *model.UpdateVpcRequest) (*model.UpdateVpcResponse, error)
	ShowVpc(context.Context, string, string, string, string, *model.ShowVpcRequest) (*model.ShowVpcResponse, error)
}

// VpcUseCase is a Vpc UseCase.
type VpcUseCase struct {
	repo VpcResponse
	log  *log.Helper
}

// NewVpcUseCase new a Vpc UseCase.
func NewVpcUseCase(repo VpcResponse, logger log.Logger) *VpcUseCase {
	return &VpcUseCase{repo: repo, log: log.NewHelper(logger)}
}

// CreateVpc create vpc
func (uc *VpcUseCase) CreateVpc(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.CreateVpcRequest) (*model.CreateVpcResponse, error) {
	return uc.repo.CreateVpc(ctx, accessKey, secretKey, regionId, projectId, request)
}

// ListVpcs List all Vpcs
func (uc *VpcUseCase) ListVpcs(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.ListVpcsRequest) (*model.ListVpcsResponse, error) {
	return uc.repo.ListVpcs(ctx, accessKey, secretKey, regionId, projectId, request)
}

// DeleteVpc delete vpc
func (uc *VpcUseCase) DeleteVpc(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.DeleteVpcRequest) (*model.DeleteVpcResponse, error) {
	return uc.repo.DeleteVpc(ctx, accessKey, secretKey, regionId, projectId, request)
}

// UpdateVpc update vpc
func (uc *VpcUseCase) UpdateVpc(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.UpdateVpcRequest) (*model.UpdateVpcResponse, error) {
	return uc.repo.UpdateVpc(ctx, accessKey, secretKey, regionId, projectId, request)
}

// ShowVpc show a vpc
func (uc *VpcUseCase) ShowVpc(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.ShowVpcRequest) (*model.ShowVpcResponse, error) {
	return uc.repo.ShowVpc(ctx, accessKey, secretKey, regionId, projectId, request)
}
