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
		RegionId:     tea.String(request.RegionId),
		PageNumber:   tea.Int32(int32(request.PageNumber)),
		PageSize:     tea.Int32(int32(request.PageSize)),
		NextToken:    tea.String(request.NextToken),
		InstanceName: tea.String(request.InstanceName),
	})
	if err != nil {
		return nil, err
	}

	instances := make([]v1.ListInstancesResponse_Instance, len(result.Instances.Instance))
	for i, instance := range result.Instances.Instance {
		publicIpAddress, innerIpAddress, rdmaIpAddress, securityGroupIds := make([]string, 0), make([]string, 0), make([]string, 0), make([]string, 0)
		if instance.PublicIpAddress != nil {
			for _, v := range instance.PublicIpAddress.IpAddress {
				publicIpAddress = append(publicIpAddress, *v)
			}
		}
		if instance.InnerIpAddress != nil {
			for _, v := range instance.InnerIpAddress.IpAddress {
				innerIpAddress = append(innerIpAddress, *v)
			}
		}
		if instance.RdmaIpAddress != nil {
			for _, v := range instance.RdmaIpAddress.IpAddress {
				rdmaIpAddress = append(rdmaIpAddress, *v)
			}
		}
		if instance.SecurityGroupIds != nil {
			for _, v := range instance.SecurityGroupIds.SecurityGroupId {
				securityGroupIds = append(securityGroupIds, *v)
			}
		}
		var eIpAddress *v1.ListInstancesResponse_Instance_EipAddress
		if instance.EipAddress != nil {
			eIpAddress = &v1.ListInstancesResponse_Instance_EipAddress{
				AllocationId:         *instance.EipAddress.AllocationId,
				IsSupportUnAssociate: *instance.EipAddress.IsSupportUnassociate,
				InternalChargeType:   *instance.EipAddress.InternetChargeType,
				IpAddress:            *instance.EipAddress.IpAddress,
				Bandwidth:            uint32(*instance.EipAddress.Bandwidth),
			}
		}
		var networkInterfaces []*v1.ListInstancesResponse_Instance_NetworkInterface
		if instance.NetworkInterfaces != nil {
			networkInterfaces = make([]*v1.ListInstancesResponse_Instance_NetworkInterface, len(instance.NetworkInterfaces.NetworkInterface))
			for j, v := range instance.NetworkInterfaces.NetworkInterface {
				var privateIpSets []*v1.ListInstancesResponse_Instance_NetworkInterface_PrivateIpSet
				if v.PrivateIpSets != nil {
					privateIpSets = make([]*v1.ListInstancesResponse_Instance_NetworkInterface_PrivateIpSet, len(v.PrivateIpSets.PrivateIpSet))
					for k, privateIpSet := range v.PrivateIpSets.PrivateIpSet {
						privateIpSets[k] = &v1.ListInstancesResponse_Instance_NetworkInterface_PrivateIpSet{
							PrivateIpAddress: *privateIpSet.PrivateIpAddress,
							Primary:          *privateIpSet.Primary,
						}
					}
				}
				var ipv6Sets []*v1.ListInstancesResponse_Instance_NetworkInterface_Ipv6Set
				if v.Ipv6Sets != nil {
					ipv6Sets = make([]*v1.ListInstancesResponse_Instance_NetworkInterface_Ipv6Set, len(v.Ipv6Sets.Ipv6Set))
					for k, ipv6Set := range v.Ipv6Sets.Ipv6Set {
						ipv6Sets[k] = &v1.ListInstancesResponse_Instance_NetworkInterface_Ipv6Set{Ipv6Address: *ipv6Set.Ipv6Address}
					}
				}
				networkInterfaces[j] = &v1.ListInstancesResponse_Instance_NetworkInterface{
					Type:               *v.Type,
					MacAddress:         *v.MacAddress,
					PrimaryIpAddress:   *v.PrimaryIpAddress,
					NetworkInterfaceId: *v.NetworkInterfaceId,
					PrivateIpSets:      privateIpSets,
					Ipv6Sets:           ipv6Sets,
				}
			}
		}
		var operationLocks []*v1.ListInstancesResponse_Instance_OperationLock
		if instance.OperationLocks != nil {
			operationLocks = make([]*v1.ListInstancesResponse_Instance_OperationLock, len(instance.OperationLocks.LockReason))
			for j, v := range instance.OperationLocks.LockReason {
				operationLocks[j] = &v1.ListInstancesResponse_Instance_OperationLock{
					LockMsg:    *v.LockMsg,
					LockReason: *v.LockReason,
				}
			}
		}
		var tags []*v1.ListInstancesResponse_Instance_Tag
		if instance.Tags != nil {
			tags = make([]*v1.ListInstancesResponse_Instance_Tag, len(instance.Tags.Tag))
			for j, v := range instance.Tags.Tag {
				tags[j] = &v1.ListInstancesResponse_Instance_Tag{
					TagValue: *v.TagValue,
					TagKey:   *v.TagKey,
				}
			}
		}
		instances[i] = v1.ListInstancesResponse_Instance{
			CreationTime:               *instance.CreationTime,
			SerialNumber:               *instance.SerialNumber,
			Status:                     *instance.Status,
			DeploymentSetId:            *instance.DeploymentSetId,
			KeyPairName:                *instance.KeyPairName,
			SaleCycle:                  *instance.SaleCycle,
			DeviceAvailable:            *instance.DeviceAvailable,
			LocalStorageCapacity:       *instance.LocalStorageCapacity,
			Description:                *instance.Description,
			SpotDuration:               *instance.SpotDuration,
			InstanceNetworkType:        *instance.InstanceNetworkType,
			InstanceName:               *instance.InstanceName,
			OsNameEn:                   *instance.OSNameEn,
			HpcClusterId:               *instance.HpcClusterId,
			Memory:                     *instance.Memory,
			OsName:                     *instance.OSName,
			DeploymentSetGroupNo:       *instance.DeploymentSetGroupNo,
			ImageId:                    *instance.ImageId,
			GpuSpec:                    *instance.GPUSpec,
			AutoReleaseTime:            *instance.AutoReleaseTime,
			DeletionProtection:         *instance.DeletionProtection,
			StoppedMode:                *instance.StoppedMode,
			GpuAmount:                  *instance.GPUAmount,
			HostName:                   *instance.HostName,
			InstanceId:                 *instance.InstanceId,
			InternetMaxBandwidthOut:    *instance.InternetMaxBandwidthOut,
			InternetMaxBandwidthIn:     *instance.InternetMaxBandwidthIn,
			InstanceType:               *instance.InstanceType,
			InstanceChargeType:         *instance.InstanceChargeType,
			RegionId:                   *instance.RegionId,
			IoOptimized:                *instance.IoOptimized,
			StartTime:                  *instance.StartTime,
			Cpu:                        *instance.Cpu,
			LocalStorageAmount:         *instance.LocalStorageAmount,
			ExpiredTime:                *instance.ExpiredTime,
			ResourceGroupId:            *instance.ResourceGroupId,
			InternetChargeType:         *instance.InternetChargeType,
			ZoneId:                     *instance.ZoneId,
			Recyclable:                 *instance.Recyclable,
			CreditSpecification:        *instance.CreditSpecification,
			InstanceTypeFamily:         *instance.InstanceTypeFamily,
			OsType:                     *instance.OSType,
			NetworkInterfaces:          networkInterfaces,
			OperationLocks:             operationLocks,
			Tags:                       tags,
			RdmaIpAddress:              rdmaIpAddress,
			SecurityGroupIds:           securityGroupIds,
			PublicIpAddress:            publicIpAddress,
			InnerIpAddress:             innerIpAddress,
			VpcAttributes:              nil,
			EipAddress:                 eIpAddress,
			DedicatedInstanceAttribute: nil,
			CpuOptions:                 nil,
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
