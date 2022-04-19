package service

import (
	v1 "ali/api/instance/v1"
	"context"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

// CreateInstances Create Instances
func (service *InstanceService) CreateInstances(ctx context.Context, request *v1.CreateInstancesRequest) (*v1.CreateInstancesResponse, error) {
	result, err := service.uc.CreateInstances(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, &ecs20140526.RunInstancesRequest{
		RegionId:     tea.String(request.RegionId),
		ImageId:      tea.String(request.ImageId),
		InstanceName: tea.String(request.Name),
		InstanceType: tea.String(request.InstanceType),
		SystemDisk: &ecs20140526.RunInstancesRequestSystemDisk{
			Size: tea.String(request.SystemDiskSize),
		},
		UniqueSuffix: tea.Bool(request.UniqueSuffix),
		Amount:       tea.Int32(int32(request.Amount)),
		Password:     tea.String(request.Password),
		DryRun:       tea.Bool(request.DryRun),
	})
	if err != nil {
		return nil, err
	}
	instanceIdsSets := make([]string, len(result.InstanceIdSets.InstanceIdSet))
	for i, v := range result.InstanceIdSets.InstanceIdSet {
		instanceIdsSets[i] = *v
	}
	return &v1.CreateInstancesResponse{
		RequestId:      *result.RequestId,
		OrderId:        *result.OrderId,
		TradePrice:     *result.TradePrice,
		InstanceIdSets: instanceIdsSets,
	}, nil
}

func (service *InstanceService) ListInstances(ctx context.Context, request *v1.ListInstancesRequest) (*v1.ListInstancesResponse, error) {
	result, err := service.uc.ListInstances(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, &ecs20140526.DescribeInstancesRequest{
		RegionId:   tea.String(request.RegionId),
		PageNumber: tea.Int32(int32(request.PageNumber)),
		PageSize:   tea.Int32(int32(request.PageSize)),
		NextToken:  tea.String(request.NextToken),
	})
	if err != nil {
		return nil, err
	}

	instances := make([]v1.ListInstancesResponse_Instance, len(result.Instances.Instance))
	for i, v := range result.Instances.Instance {
		publicIpAddress, innerIpAddress, rdmaIpAddress, securityGroupIds := make([]string, 0), make([]string, 0), make([]string, 0), make([]string, 0)
		if v.PublicIpAddress != nil {
			for _, v := range v.PublicIpAddress.IpAddress {
				publicIpAddress = append(publicIpAddress, *v)
			}
		}
		if v.InnerIpAddress != nil {
			for _, v := range v.InnerIpAddress.IpAddress {
				innerIpAddress = append(innerIpAddress, *v)
			}
		}
		if v.RdmaIpAddress != nil {
			for _, v := range v.RdmaIpAddress.IpAddress {
				rdmaIpAddress = append(rdmaIpAddress, *v)
			}
		}
		if v.SecurityGroupIds != nil {
			for _, v := range v.SecurityGroupIds.SecurityGroupId {
				securityGroupIds = append(securityGroupIds, *v)
			}
		}
		var eIpAddress *v1.ListInstancesResponse_Instance_EipAddress
		if v.EipAddress != nil {
			eIpAddress = &v1.ListInstancesResponse_Instance_EipAddress{
				AllocationId:         *v.EipAddress.AllocationId,
				IsSupportUnAssociate: *v.EipAddress.IsSupportUnassociate,
				InternalChargeType:   *v.EipAddress.InternetChargeType,
				IpAddress:            *v.EipAddress.IpAddress,
				Bandwidth:            uint32(*v.EipAddress.Bandwidth),
			}
		}
		instances[i] = v1.ListInstancesResponse_Instance{
			InstanceId:           *v.InstanceId,
			InstanceName:         *v.InstanceName,
			InstanceType:         *v.InstanceType,
			InstanceChargeType:   *v.InstanceChargeType,
			RegionId:             *v.RegionId,
			Cpu:                  uint32(*v.Cpu),
			OsNameEn:             *v.OSNameEn,
			Status:               *v.Status,
			CreationTime:         *v.CreationTime,
			InstanceNetworkType:  *v.InstanceNetworkType,
			Memory:               uint32(*v.Memory),
			GpuSpec:              *v.GPUSpec,
			AutoReleaseTime:      *v.AutoReleaseTime,
			GpuAmount:            uint32(*v.GPUAmount),
			OsType:               *v.OSType,
			ExpiredTime:          *v.ExpiredTime,
			StartTime:            *v.StartTime,
			PublicIpAddress:      publicIpAddress,
			InnerIpAddress:       innerIpAddress,
			RdmaIpAddress:        rdmaIpAddress,
			SecurityGroupIds:     securityGroupIds,
			EipAddress:           eIpAddress,
			DeviceAvailable:      *v.DeviceAvailable,
			LocalStorageCapacity: uint64(*v.LocalStorageCapacity),
			LocalStorageAmount:   uint32(*v.LocalStorageAmount),
			ImageId:              *v.ImageId,
		}
	}
	listInstancesResponse := &v1.ListInstancesResponse{
		RequestId:  *result.RequestId,
		PageNumber: uint32(*result.PageNumber),
		PageSize:   uint32(*result.PageSize),
		TotalCount: uint32(*result.TotalCount),
		NextToken:  *result.NextToken,
		Instances:  nil,
	}

	return listInstancesResponse, nil
}

func (service *InstanceService) StartInstances(ctx context.Context, request *v1.StartInstancesRequest) (*v1.StartInstancesResponse, error) {
	instanceIds := make([]*string, len(request.InstanceIds))
	for i, v := range request.InstanceIds {
		instanceIds[i] = tea.String(v)
	}
	result, err := service.uc.StartInstances(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, &ecs20140526.StartInstancesRequest{
		InstanceId: instanceIds,
		DryRun:     tea.Bool(request.DryRun),
	})
	if err != nil {
		return nil, err
	}
	return &v1.StartInstancesResponse{
		RequestId: *result,
	}, nil
}

func (service *InstanceService) StopInstances(ctx context.Context, request *v1.StopInstancesRequest) (*v1.StopInstancesResponse, error) {
	instanceIds := make([]*string, len(request.InstanceIds))
	for i, v := range request.InstanceIds {
		instanceIds[i] = tea.String(v)
	}
	result, err := service.uc.StopInstances(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, &ecs20140526.StopInstancesRequest{
		InstanceId: instanceIds,
		DryRun:     tea.Bool(request.DryRun),
	})
	if err != nil {
		return nil, err
	}
	return &v1.StopInstancesResponse{
		RequestId: *result,
	}, nil
}

func (service *InstanceService) RebootInstances(ctx context.Context, request *v1.RebootInstancesRequest) (*v1.RebootInstancesResponse, error) {
	instanceIds := make([]*string, len(request.InstanceIds))
	for i, v := range request.InstanceIds {
		instanceIds[i] = tea.String(v)
	}
	result, err := service.uc.RebootInstances(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, &ecs20140526.RebootInstancesRequest{
		InstanceId: instanceIds,
		DryRun:     tea.Bool(request.DryRun),
	})
	if err != nil {
		return nil, err
	}
	return &v1.RebootInstancesResponse{
		RequestId: *result,
	}, nil
}

func (service *InstanceService) DeleteInstances(ctx context.Context, request *v1.DeleteInstancesRequest) (*v1.DeleteInstancesResponse, error) {
	instanceIds := make([]*string, len(request.InstanceIds))
	for i, v := range request.InstanceIds {
		instanceIds[i] = tea.String(v)
	}
	result, err := service.uc.DeleteInstances(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, &ecs20140526.DeleteInstancesRequest{
		InstanceId: instanceIds,
		DryRun:     tea.Bool(request.DryRun),
	})
	if err != nil {
		return nil, err
	}
	return &v1.DeleteInstancesResponse{
		RequestId: *result,
	}, nil
}
