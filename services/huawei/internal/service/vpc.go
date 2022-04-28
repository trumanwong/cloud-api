package service

import (
	"context"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
	v1 "huawei/api/instance/v1"
	"huawei/pkg/util/convert"
)

func (service *InstanceService) CreateVpc(ctx context.Context, request *v1.CreateVpcRequest) (*v1.CreateVpcResponse, error) {
	result, err := service.vc.CreateVpc(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.CreateVpcRequest{Body: &model.CreateVpcRequestBody{Vpc: &model.CreateVpcOption{
		Cidr:                request.Cidr,
		Name:                request.Name,
		Description:         request.Description,
		EnterpriseProjectId: request.EnterpriseProjectId,
	}}})

	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.CreateVpcResponse{Data: data}, nil
}

func (service *InstanceService) ListVpcs(ctx context.Context, request *v1.ListVpcsRequest) (*v1.ListVpcsResponse, error) {
	result, err := service.vc.ListVpcs(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.ListVpcsRequest{
		Id:                  request.Id,
		EnterpriseProjectId: request.EnterpriseProjectId,
	})

	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.ListVpcsResponse{Data: data}, nil
}

func (service *InstanceService) DeleteVpc(ctx context.Context, request *v1.DeleteVpcRequest) (*v1.DeleteVpcResponse, error) {
	result, err := service.vc.DeleteVpc(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.DeleteVpcRequest{
		VpcId: request.VpcId,
	})
	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteVpcResponse{Data: data}, nil
}

func (service *InstanceService) UpdateVpc(ctx context.Context, request *v1.UpdateVpcRequest) (*v1.UpdateVpcResponse, error) {
	var routes *[]model.Route
	if len(request.Routes) > 0 {
		temp := make([]model.Route, len(request.Routes))
		for i, route := range request.Routes {
			temp[i] = model.Route{
				Destination: route.Destination,
				Nexthop:     route.NextHop,
			}
		}
		routes = &temp
	}
	result, err := service.vc.UpdateVpc(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.UpdateVpcRequest{
		VpcId: request.VpcId,
		Body: &model.UpdateVpcRequestBody{Vpc: &model.UpdateVpcOption{
			Name:        request.Name,
			Description: request.Description,
			Cidr:        request.Cidr,
			Routes:      routes,
		}},
	})
	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateVpcResponse{Data: data}, nil
}

func (service *InstanceService) ShowVpc(ctx context.Context, request *v1.ShowVpcRequest) (*v1.ShowVpcResponse, error) {
	result, err := service.vc.ShowVpc(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.ShowVpcRequest{
		VpcId: request.VpcId,
	})
	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.ShowVpcResponse{Data: data}, nil
}
