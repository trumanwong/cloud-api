package service

import (
	"context"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
	v1 "huawei/api/instance/v1"
)

func (service *InstanceService) CreateVpc(ctx context.Context, request *v1.CreateVpcRequest) (*v1.CreateVpcResponse, error) {
	result, err := service.vc.CreateVpc(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.CreateVpcRequest{Body: &model.CreateVpcRequestBody{Vpc: &model.CreateVpcOption{
		Cidr:                &request.Cidr,
		Name:                &request.Name,
		Description:         &request.Description,
		EnterpriseProjectId: &request.EnterpriseProjectId,
	}}})

	if err != nil {
		return nil, err
	}
	routes := make([]*v1.VpcRoute, len(result.Vpc.Routes))
	for i, v := range result.Vpc.Routes {
		routes[i] = &v1.VpcRoute{
			Destination: *v.Destination,
			NextHop:     *v.Nexthop,
		}
	}
	status, _ := result.Vpc.Status.MarshalJSON()
	return &v1.CreateVpcResponse{Vpc: &v1.Vpc{
		Id:                  result.Vpc.Id,
		Name:                result.Vpc.Name,
		Cidr:                result.Vpc.Cidr,
		Description:         result.Vpc.Description,
		Routes:              routes,
		Status:              string(status),
		EnterpriseProjectId: result.Vpc.EnterpriseProjectId,
	}}, nil
}

func (service *InstanceService) ListVpcs(ctx context.Context, request *v1.ListVpcsRequest) (*v1.ListVpcsResponse, error) {
	result, err := service.vc.ListVpcs(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.ListVpcsRequest{
		Id:                  &request.Id,
		EnterpriseProjectId: &request.EnterpriseProjectId,
	})

	if err != nil {
		return nil, err
	}
	vpcs := make([]*v1.Vpc, 0)
	if result.Vpcs != nil {
		vpcs = make([]*v1.Vpc, len(*result.Vpcs))
		for i, vpc := range *result.Vpcs {
			routes := make([]*v1.VpcRoute, len(vpc.Routes))
			for j, route := range vpc.Routes {
				routes[j] = &v1.VpcRoute{
					Destination: *route.Destination,
					NextHop:     *route.Nexthop,
				}
			}
			status, _ := vpc.Status.MarshalJSON()
			vpcs[i] = &v1.Vpc{
				Id:                  vpc.Id,
				Name:                vpc.Name,
				Cidr:                vpc.Cidr,
				Description:         vpc.Description,
				Routes:              routes,
				Status:              string(status),
				EnterpriseProjectId: vpc.EnterpriseProjectId,
			}
		}
	}
	return &v1.ListVpcsResponse{Vpcs: vpcs}, nil
}

func (service *InstanceService) DeleteVpc(ctx context.Context, request *v1.DeleteVpcRequest) (*v1.DeleteVpcResponse, error) {
	_, err := service.vc.DeleteVpc(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.DeleteVpcRequest{
		VpcId: request.VpcId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.DeleteVpcResponse{}, nil
}

func (service *InstanceService) UpdateVpc(ctx context.Context, request *v1.UpdateVpcRequest) (*v1.UpdateVpcResponse, error) {
	var name *string
	if len(request.Name) > 0 {
		name = &request.Name
	}
	var description *string
	if len(request.Description) > 0 {
		description = &request.Description
	}
	var cidr *string
	if len(request.Cidr) > 0 {
		cidr = &request.Cidr
	}
	var routes *[]model.Route
	if len(request.Routes) > 0 {
		temp := make([]model.Route, len(request.Routes))
		for i, route := range request.Routes {
			temp[i] = model.Route{
				Destination: &route.Destination,
				Nexthop:     &route.NextHop,
			}
		}
		routes = &temp
	}
	result, err := service.vc.UpdateVpc(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.UpdateVpcRequest{
		VpcId: request.VpcId,
		Body: &model.UpdateVpcRequestBody{Vpc: &model.UpdateVpcOption{
			Name:        name,
			Description: description,
			Cidr:        cidr,
			Routes:      routes,
		}},
	})
	if err != nil {
		return nil, err
	}
	var vpc *v1.Vpc
	if result.Vpc != nil {
		routes := make([]*v1.VpcRoute, len(result.Vpc.Routes))
		for j, route := range result.Vpc.Routes {
			routes[j] = &v1.VpcRoute{
				Destination: *route.Destination,
				NextHop:     *route.Nexthop,
			}
		}
		status, _ := result.Vpc.Status.MarshalJSON()
		vpc = &v1.Vpc{
			Id:                  result.Vpc.Id,
			Name:                result.Vpc.Name,
			Cidr:                result.Vpc.Cidr,
			Description:         result.Vpc.Description,
			Routes:              routes,
			Status:              string(status),
			EnterpriseProjectId: result.Vpc.EnterpriseProjectId,
		}
	}
	return &v1.UpdateVpcResponse{Vpc: vpc}, nil
}

func (service *InstanceService) ShowVpc(ctx context.Context, request *v1.ShowVpcRequest) (*v1.ShowVpcResponse, error) {
	result, err := service.vc.ShowVpc(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.ShowVpcRequest{
		VpcId: request.VpcId,
	})
	if err != nil {
		return nil, err
	}
	var vpc *v1.Vpc
	if result.Vpc != nil {
		routes := make([]*v1.VpcRoute, len(result.Vpc.Routes))
		for j, route := range result.Vpc.Routes {
			routes[j] = &v1.VpcRoute{
				Destination: *route.Destination,
				NextHop:     *route.Nexthop,
			}
		}
		status, _ := result.Vpc.Status.MarshalJSON()
		vpc = &v1.Vpc{
			Id:                  result.Vpc.Id,
			Name:                result.Vpc.Name,
			Cidr:                result.Vpc.Cidr,
			Description:         result.Vpc.Description,
			Routes:              routes,
			Status:              string(status),
			EnterpriseProjectId: result.Vpc.EnterpriseProjectId,
		}
	}
	return &v1.ShowVpcResponse{Vpc: vpc}, nil
}
