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
		ExtraDhcpOpts:     nil,
		Scope:             *result.Subnet.Scope,
	}}, nil
}

func (service *InstanceService) ListSubnets(ctx context.Context, request *v1.ListSubnetsRequest) (*v1.ListSubnetsResponse, error) {
	result, err := service.sc.ListSubnets(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.ListSubnetsRequest{
		Id:                  &request.Id,
		EnterpriseProjectId: &request.EnterpriseProjectId,
	})

	if err != nil {
		return nil, err
	}
	Subnets := make([]*v1.Subnet, 0)
	if result.Subnets != nil {
		Subnets = make([]*v1.Subnet, len(*result.Subnets))
		for i, Subnet := range *result.Subnets {
			routes := make([]*v1.SubnetRoute, len(Subnet.Routes))
			for j, route := range Subnet.Routes {
				routes[j] = &v1.SubnetRoute{
					Destination: *route.Destination,
					NextHop:     *route.Nexthop,
				}
			}
			status, _ := Subnet.Status.MarshalJSON()
			Subnets[i] = &v1.Subnet{
				Id:                  Subnet.Id,
				Name:                Subnet.Name,
				Cidr:                Subnet.Cidr,
				Description:         Subnet.Description,
				Routes:              routes,
				Status:              string(status),
				EnterpriseProjectId: Subnet.EnterpriseProjectId,
			}
		}
	}
	return &v1.ListSubnetsResponse{Subnets: Subnets}, nil
}

func (service *InstanceService) DeleteSubnet(ctx context.Context, request *v1.DeleteSubnetRequest) (*v1.DeleteSubnetResponse, error) {
	_, err := service.sc.DeleteSubnet(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.DeleteSubnetRequest{
		SubnetId: request.SubnetId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.DeleteSubnetResponse{}, nil
}

func (service *InstanceService) UpdateSubnet(ctx context.Context, request *v1.UpdateSubnetRequest) (*v1.UpdateSubnetResponse, error) {
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
	result, err := service.sc.UpdateSubnet(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.UpdateSubnetRequest{
		SubnetId: request.SubnetId,
		Body: &model.UpdateSubnetRequestBody{Subnet: &model.UpdateSubnetOption{
			Name:        name,
			Description: description,
			Cidr:        cidr,
			Routes:      routes,
		}},
	})
	if err != nil {
		return nil, err
	}
	var Subnet *v1.Subnet
	if result.Subnet != nil {
		routes := make([]*v1.SubnetRoute, len(result.Subnet.Routes))
		for j, route := range result.Subnet.Routes {
			routes[j] = &v1.SubnetRoute{
				Destination: *route.Destination,
				NextHop:     *route.Nexthop,
			}
		}
		status, _ := result.Subnet.Status.MarshalJSON()
		Subnet = &v1.Subnet{
			Id:                  result.Subnet.Id,
			Name:                result.Subnet.Name,
			Cidr:                result.Subnet.Cidr,
			Description:         result.Subnet.Description,
			Routes:              routes,
			Status:              string(status),
			EnterpriseProjectId: result.Subnet.EnterpriseProjectId,
		}
	}
	return &v1.UpdateSubnetResponse{Subnet: Subnet}, nil
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
		routes := make([]*v1.SubnetRoute, len(result.Subnet.Routes))
		for j, route := range result.Subnet.Routes {
			routes[j] = &v1.SubnetRoute{
				Destination: *route.Destination,
				NextHop:     *route.Nexthop,
			}
		}
		status, _ := result.Subnet.Status.MarshalJSON()
		Subnet = &v1.Subnet{
			Id:                  result.Subnet.Id,
			Name:                result.Subnet.Name,
			Cidr:                result.Subnet.Cidr,
			Description:         result.Subnet.Description,
			Routes:              routes,
			Status:              string(status),
			EnterpriseProjectId: result.Subnet.EnterpriseProjectId,
		}
	}
	return &v1.ShowSubnetResponse{Subnet: Subnet}, nil
}
