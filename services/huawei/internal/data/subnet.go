package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
	"huawei/internal/biz"
)

type subnetResponse struct {
	data *Data
	log  *log.Helper
}

// NewSubnetRepo .
func NewSubnetRepo(data *Data, logger log.Logger) biz.SubnetResponse {
	return &subnetResponse{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *subnetResponse) CreateSubnet(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.CreateSubnetRequest) (*model.CreateSubnetResponse, error) {
	client := getVpcClient(accessKey, secretKey, regionId, projectId)
	result, err := client.CreateSubnet(request)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *subnetResponse) ListSubnets(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.ListSubnetsRequest) (*model.ListSubnetsResponse, error) {
	client := getVpcClient(accessKey, secretKey, regionId, projectId)
	result, err := client.ListSubnets(request)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *subnetResponse) DeleteSubnet(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.DeleteSubnetRequest) (*model.DeleteSubnetResponse, error) {
	client := getVpcClient(accessKey, secretKey, regionId, projectId)
	result, err := client.DeleteSubnet(request)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *subnetResponse) UpdateSubnet(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.UpdateSubnetRequest) (*model.UpdateSubnetResponse, error) {
	client := getVpcClient(accessKey, secretKey, regionId, projectId)
	result, err := client.UpdateSubnet(request)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *subnetResponse) ShowSubnet(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.ShowSubnetRequest) (*model.ShowSubnetResponse, error) {
	client := getVpcClient(accessKey, secretKey, regionId, projectId)
	result, err := client.ShowSubnet(request)
	if err != nil {
		return nil, err
	}
	return result, err
}
