package service

import (
	"context"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
	v1 "huawei/api/instance/v1"
)

func (service *InstanceService) CreateSubnet(ctx context.Context, request *v1.CreateSubnetRequest) (*v1.CreateSubnetResponse, error) {
	var description *string
	if len(request.Description) > 0 {
		description = &request.Description
	}
	var ipv6Enable *bool
	if request.Ipv6Enable != true {
		ipv6Enable = &request.Ipv6Enable
	}
	var dhcpEnable *bool
	if request.DhcpEnable != true {
		dhcpEnable = &request.DhcpEnable
	}
	var primaryDns *string
	if len(request.PrimaryDns) > 0 {
		primaryDns = &request.PrimaryDns
	}
	var secondaryDns *string
	if len(request.SecondaryDns) > 0 {
		secondaryDns = &request.SecondaryDns
	}
	var dnsList *[]string
	if len(request.DnsList) > 0 {
		dnsList = &request.DnsList
	}
	var availabilityZone *string
	if len(request.AvailabilityZone) > 0 {
		availabilityZone = &request.AvailabilityZone
	}
	result, err := service.sc.CreateSubnet(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.CreateSubnetRequest{Body: &model.CreateSubnetRequestBody{Subnet: &model.CreateSubnetOption{
		Name:             request.Name,
		Description:      description,
		Cidr:             request.Cidr,
		VpcId:            request.VpcId,
		GatewayIp:        request.GatewayIp,
		Ipv6Enable:       ipv6Enable,
		DhcpEnable:       dhcpEnable,
		PrimaryDns:       primaryDns,
		SecondaryDns:     secondaryDns,
		DnsList:          dnsList,
		AvailabilityZone: availabilityZone,
		ExtraDhcpOpts:    nil,
	}}})

	if err != nil {
		return nil, err
	}
	status, _ := result.Subnet.Status.MarshalJSON()
	extraDhcpOpts := make([]*v1.ExtraDhcpOpts, len(result.Subnet.ExtraDhcpOpts))
	for i, v := range result.Subnet.ExtraDhcpOpts {
		optName, _ := v.OptName.MarshalJSON()
		extraDhcpOpts[i] = &v1.ExtraDhcpOpts{
			OptName:  string(optName),
			OptValue: *v.OptValue,
		}
	}
	return &v1.CreateSubnetResponse{Subnet: &v1.Subnet{
		Id:                result.Subnet.Id,
		Name:              result.Subnet.Name,
		Cidr:              result.Subnet.Cidr,
		Description:       result.Subnet.Description,
		Status:            string(status),
		GatewayIp:         result.Subnet.GatewayIp,
		Ipv6Enable:        result.Subnet.Ipv6Enable,
		CidrV6:            result.Subnet.CidrV6,
		GatewayIpV6:       result.Subnet.GatewayIpV6,
		DhcpEnable:        result.Subnet.DhcpEnable,
		PrimaryDns:        result.Subnet.PrimaryDns,
		SecondaryDns:      result.Subnet.SecondaryDns,
		DnsList:           result.Subnet.DnsList,
		AvailabilityZone:  result.Subnet.AvailabilityZone,
		VpcId:             result.Subnet.VpcId,
		NeutronNetworkId:  result.Subnet.NeutronNetworkId,
		NeutronSubnetId:   result.Subnet.NeutronSubnetId,
		NeutronSubnetIdV6: result.Subnet.NeutronSubnetIdV6,
		ExtraDhcpOpts:     extraDhcpOpts,
		Scope:             *result.Subnet.Scope,
	}}, nil
}

func (service *InstanceService) ListSubnets(ctx context.Context, request *v1.ListSubnetsRequest) (*v1.ListSubnetsResponse, error) {
	var vpcId *string
	if len(request.VpcId) > 0 {
		vpcId = &request.VpcId
	}
	result, err := service.sc.ListSubnets(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.ListSubnetsRequest{
		VpcId: vpcId,
	})

	if err != nil {
		return nil, err
	}
	subnets := make([]*v1.Subnet, 0)
	if result.Subnets != nil {
		subnets = make([]*v1.Subnet, len(*result.Subnets))
		for i, subnet := range *result.Subnets {
			status, _ := subnet.Status.MarshalJSON()
			extraDhcpOpts := make([]*v1.ExtraDhcpOpts, len(subnet.ExtraDhcpOpts))
			for i, v := range subnet.ExtraDhcpOpts {
				optName, _ := v.OptName.MarshalJSON()
				extraDhcpOpts[i] = &v1.ExtraDhcpOpts{
					OptName:  string(optName),
					OptValue: *v.OptValue,
				}
			}
			subnets[i] = &v1.Subnet{
				Id:                subnet.Id,
				Name:              subnet.Name,
				Cidr:              subnet.Cidr,
				Description:       subnet.Description,
				Status:            string(status),
				GatewayIp:         subnet.GatewayIp,
				Ipv6Enable:        subnet.Ipv6Enable,
				CidrV6:            subnet.CidrV6,
				GatewayIpV6:       subnet.GatewayIpV6,
				DhcpEnable:        subnet.DhcpEnable,
				PrimaryDns:        subnet.PrimaryDns,
				SecondaryDns:      subnet.SecondaryDns,
				DnsList:           subnet.DnsList,
				AvailabilityZone:  subnet.AvailabilityZone,
				VpcId:             subnet.VpcId,
				NeutronNetworkId:  subnet.NeutronNetworkId,
				NeutronSubnetId:   subnet.NeutronSubnetId,
				NeutronSubnetIdV6: subnet.NeutronSubnetIdV6,
				ExtraDhcpOpts:     extraDhcpOpts,
				Scope:             *subnet.Scope,
			}
		}
	}
	return &v1.ListSubnetsResponse{Subnets: subnets}, nil
}

func (service *InstanceService) DeleteSubnet(ctx context.Context, request *v1.DeleteSubnetRequest) (*v1.DeleteSubnetResponse, error) {
	_, err := service.sc.DeleteSubnet(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.DeleteSubnetRequest{
		SubnetId: request.SubnetId,
		VpcId:    request.VpcId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.DeleteSubnetResponse{}, nil
}

func (service *InstanceService) UpdateSubnet(ctx context.Context, request *v1.UpdateSubnetRequest) (*v1.UpdateSubnetResponse, error) {
	var description *string
	if len(request.Description) > 0 {
		description = &request.Description
	}
	var ipv6Enable *bool
	if request.Ipv6Enable != true {
		ipv6Enable = &request.Ipv6Enable
	}
	var dhcpEnable *bool
	if request.DhcpEnable != true {
		dhcpEnable = &request.DhcpEnable
	}
	var primaryDns *string
	if len(request.PrimaryDns) > 0 {
		primaryDns = &request.PrimaryDns
	}
	var secondaryDns *string
	if len(request.SecondaryDns) > 0 {
		secondaryDns = &request.SecondaryDns
	}
	var dnsList *[]string
	if len(request.DnsList) > 0 {
		dnsList = &request.DnsList
	}
	result, err := service.sc.UpdateSubnet(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.UpdateSubnetRequest{
		SubnetId: request.SubnetId,
		VpcId:    request.VpcId,
		Body: &model.UpdateSubnetRequestBody{Subnet: &model.UpdateSubnetOption{
			Name:          request.Name,
			Description:   description,
			Ipv6Enable:    ipv6Enable,
			DhcpEnable:    dhcpEnable,
			PrimaryDns:    primaryDns,
			SecondaryDns:  secondaryDns,
			DnsList:       dnsList,
			ExtraDhcpOpts: nil,
		}},
	})
	if err != nil {
		return nil, err
	}

	status, _ := result.Subnet.Status.MarshalJSON()
	return &v1.UpdateSubnetResponse{
		SubnetId: result.Subnet.Id,
		Status:   string(status),
	}, nil
}

func (service *InstanceService) ShowSubnet(ctx context.Context, request *v1.ShowSubnetRequest) (*v1.ShowSubnetResponse, error) {
	result, err := service.sc.ShowSubnet(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.ShowSubnetRequest{
		SubnetId: request.SubnetId,
	})
	if err != nil {
		return nil, err
	}
	var Subnet *v1.Subnet
	if result.Subnet != nil {
		status, _ := result.Subnet.Status.MarshalJSON()
		extraDhcpOpts := make([]*v1.ExtraDhcpOpts, len(result.Subnet.ExtraDhcpOpts))
		for i, v := range result.Subnet.ExtraDhcpOpts {
			optName, _ := v.OptName.MarshalJSON()
			extraDhcpOpts[i] = &v1.ExtraDhcpOpts{
				OptName:  string(optName),
				OptValue: *v.OptValue,
			}
		}
		Subnet = &v1.Subnet{
			Id:                result.Subnet.Id,
			Name:              result.Subnet.Name,
			Cidr:              result.Subnet.Cidr,
			Description:       result.Subnet.Description,
			Status:            string(status),
			GatewayIp:         result.Subnet.GatewayIp,
			Ipv6Enable:        result.Subnet.Ipv6Enable,
			CidrV6:            result.Subnet.CidrV6,
			GatewayIpV6:       result.Subnet.GatewayIpV6,
			DhcpEnable:        result.Subnet.DhcpEnable,
			PrimaryDns:        result.Subnet.PrimaryDns,
			SecondaryDns:      result.Subnet.SecondaryDns,
			DnsList:           result.Subnet.DnsList,
			AvailabilityZone:  result.Subnet.AvailabilityZone,
			VpcId:             result.Subnet.VpcId,
			NeutronNetworkId:  result.Subnet.NeutronNetworkId,
			NeutronSubnetId:   result.Subnet.NeutronSubnetId,
			NeutronSubnetIdV6: result.Subnet.NeutronSubnetIdV6,
			ExtraDhcpOpts:     extraDhcpOpts,
			Scope:             *result.Subnet.Scope,
		}
	}
	return &v1.ShowSubnetResponse{Subnet: Subnet}, nil
}
