package service

import (
	"context"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
	v1 "huawei/api/instance/v1"
	"huawei/pkg/util/convert"
)

func (service *InstanceService) CreateSubnet(ctx context.Context, request *v1.CreateSubnetRequest) (*v1.CreateSubnetResponse, error) {
	var dnsList *[]string
	if len(request.DnsList) > 0 {
		dnsList = &request.DnsList
	}
	result, err := service.sc.CreateSubnet(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.CreateSubnetRequest{Body: &model.CreateSubnetRequestBody{Subnet: &model.CreateSubnetOption{
		Name:             request.Name,
		Description:      request.Description,
		Cidr:             request.Cidr,
		VpcId:            request.VpcId,
		GatewayIp:        request.GatewayIp,
		Ipv6Enable:       request.Ipv6Enable,
		DhcpEnable:       request.DhcpEnable,
		PrimaryDns:       request.PrimaryDns,
		SecondaryDns:     request.SecondaryDns,
		DnsList:          dnsList,
		AvailabilityZone: request.AvailabilityZone,
		ExtraDhcpOpts:    nil,
	}}})

	if err != nil {
		return nil, err
	}

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.CreateSubnetResponse{Data: data}, nil
}

func (service *InstanceService) ListSubnets(ctx context.Context, request *v1.ListSubnetsRequest) (*v1.ListSubnetsResponse, error) {
	result, err := service.sc.ListSubnets(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.ListSubnetsRequest{
		VpcId: request.VpcId,
	})

	if err != nil {
		return nil, err
	}

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.ListSubnetsResponse{Data: data}, nil
}

func (service *InstanceService) DeleteSubnet(ctx context.Context, request *v1.DeleteSubnetRequest) (*v1.DeleteSubnetResponse, error) {
	result, err := service.sc.DeleteSubnet(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.DeleteSubnetRequest{
		SubnetId: request.SubnetId,
		VpcId:    request.VpcId,
	})
	if err != nil {
		return nil, err
	}

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteSubnetResponse{Data: data}, nil
}

func (service *InstanceService) UpdateSubnet(ctx context.Context, request *v1.UpdateSubnetRequest) (*v1.UpdateSubnetResponse, error) {
	var dnsList *[]string
	if len(request.DnsList) > 0 {
		dnsList = &request.DnsList
	}
	result, err := service.sc.UpdateSubnet(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.UpdateSubnetRequest{
		SubnetId: request.SubnetId,
		VpcId:    request.VpcId,
		Body: &model.UpdateSubnetRequestBody{Subnet: &model.UpdateSubnetOption{
			Name:          request.Name,
			Description:   request.Description,
			Ipv6Enable:    request.Ipv6Enable,
			DhcpEnable:    request.DhcpEnable,
			PrimaryDns:    request.PrimaryDns,
			SecondaryDns:  request.SecondaryDns,
			DnsList:       dnsList,
			ExtraDhcpOpts: nil,
		}},
	})
	if err != nil {
		return nil, err
	}

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.UpdateSubnetResponse{Data: data}, nil
}

func (service *InstanceService) ShowSubnet(ctx context.Context, request *v1.ShowSubnetRequest) (*v1.ShowSubnetResponse, error) {
	result, err := service.sc.ShowSubnet(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.ShowSubnetRequest{
		SubnetId: request.SubnetId,
	})
	if err != nil {
		return nil, err
	}

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.ShowSubnetResponse{Data: data}, nil
}
