package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
	"huawei/internal/biz"
)

type vpcResponse struct {
	data *Data
	log  *log.Helper
}

// NewVpcRepo .
func NewVpcRepo(data *Data, logger log.Logger) biz.VpcResponse {
	return &vpcResponse{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *vpcResponse) CreateVpc(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.CreateVpcRequest) (*model.CreateVpcResponse, error) {
	client := getVpcClient(accessKey, secretKey, regionId, projectId)
	result, err := client.CreateVpc(request)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *vpcResponse) ListVpcs(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.ListVpcsRequest) (*model.ListVpcsResponse, error) {
	client := getVpcClient(accessKey, secretKey, regionId, projectId)
	result, err := client.ListVpcs(request)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *vpcResponse) DeleteVpc(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.DeleteVpcRequest) (*model.DeleteVpcResponse, error) {
	client := getVpcClient(accessKey, secretKey, regionId, projectId)
	result, err := client.DeleteVpc(request)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *vpcResponse) UpdateVpc(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.UpdateVpcRequest) (*model.UpdateVpcResponse, error) {
	client := getVpcClient(accessKey, secretKey, regionId, projectId)
	result, err := client.UpdateVpc(request)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *vpcResponse) ShowVpc(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.ShowVpcRequest) (*model.ShowVpcResponse, error) {
	client := getVpcClient(accessKey, secretKey, regionId, projectId)
	result, err := client.ShowVpc(request)
	if err != nil {
		return nil, err
	}
	return result, err
}
