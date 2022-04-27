package service

import (
	"context"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	v1 "tencent/api/instance/v1"
)

func (s *InstanceService) CreateInstances(ctx context.Context, request *v1.CreateInstancesRequest) (*v1.CreateInstancesResponse, error) {
	runInstancesRequest := cvm.NewRunInstancesRequest()
	runInstancesRequest.LoginSettings = &cvm.LoginSettings{Password: request.Password}
	runInstancesRequest.SystemDisk = &cvm.SystemDisk{
		DiskSize: request.SystemDisk.DiskSize,
		DiskType: request.SystemDisk.DiskType,
		CdcId:    request.SystemDisk.CdcId,
	}
	runInstancesRequest.InstanceType = request.InstanceType
	runInstancesRequest.ImageId = request.ImageId
	runInstancesRequest.InstanceCount = request.InstanceCount
	runInstancesRequest.InstanceName = request.Name
	runInstancesRequest.DryRun = request.DryRun
	response, err := s.uc.CreateInstances(ctx, request.SecretId, request.SecretKey, request.Region, runInstancesRequest)
	if err != nil {
		return nil, err
	}

	instanceIdSets := make([]string, len(response.Response.InstanceIdSet))
	for i, v := range response.Response.InstanceIdSet {
		instanceIdSets[i] = *v
	}

	createInstancesResponse := &v1.CreateInstancesResponse{
		RequestId:      response.Response.RequestId,
		InstanceIdSets: instanceIdSets,
	}
	return createInstancesResponse, nil
}

func (s *InstanceService) ListInstances(ctx context.Context, request *v1.ListInstancesRequest) (*v1.ListInstancesResponse, error) {
	describeInstancesRequest := cvm.NewDescribeInstancesRequest()
	describeInstancesRequest.Offset = request.Offset
	describeInstancesRequest.Limit = request.Limit
	describeInstancesRequest.Filters = make([]*cvm.Filter, 0)
	if request.InstanceName != nil {
		describeInstancesRequest.Filters = append(describeInstancesRequest.Filters, &cvm.Filter{
			Name: request.InstanceName,
		})
	}
	response, err := s.uc.ListInstances(ctx, request.SecretId, request.SecretKey, request.Region, describeInstancesRequest)
	if err != nil {
		return nil, err
	}
	instances := make([]*v1.ListInstancesResponse_Instance, len(response.Response.InstanceSet))
	for i, instance := range response.Response.InstanceSet {
		var placement *v1.ListInstancesResponse_Instance_Placement
		if instance.Placement != nil {
			var hostIds, hostIps []string
			if instance.Placement.HostIds != nil {
				hostIds = make([]string, len(instance.Placement.HostIds))
				for j, v := range instance.Placement.HostIds {
					hostIds[j] = *v
				}
			}
			if instance.Placement.HostIps != nil {
				hostIps = make([]string, len(instance.Placement.HostIps))
				for j, v := range instance.Placement.HostIps {
					hostIps[j] = *v
				}
			}
			placement = &v1.ListInstancesResponse_Instance_Placement{
				Zone:      instance.Placement.Zone,
				ProjectId: instance.Placement.ProjectId,
				HostIds:   hostIds,
				HostIps:   hostIps,
				HostId:    instance.Placement.HostId,
			}
		}
		var systemDisk *v1.SystemDisk
		if instance.SystemDisk != nil {
			systemDisk = &v1.SystemDisk{
				DiskType: instance.SystemDisk.DiskType,
				DiskSize: instance.SystemDisk.DiskSize,
				CdcId:    instance.SystemDisk.CdcId,
			}
		}
		dataDisks := make([]*v1.ListInstancesResponse_Instance_DataDisk, len(instance.DataDisks))
		for j, v := range instance.DataDisks {
			dataDisks[j] = &v1.ListInstancesResponse_Instance_DataDisk{
				DiskSize:              v.DiskSize,
				DiskType:              v.DiskType,
				CdcId:                 v.CdcId,
				ThroughputPerformance: v.ThroughputPerformance,
			}
		}
		privateIpAddresses, publicIpAddresses := make([]string, len(instance.PrivateIpAddresses)), make([]string, len(instance.PublicIpAddresses))
		for j, v := range instance.PrivateIpAddresses {
			privateIpAddresses[j] = *v
		}
		for j, v := range instance.PrivateIpAddresses {
			publicIpAddresses[j] = *v
		}
		var internetAccessible *v1.ListInstancesResponse_Instance_InternetAccessible
		var virtualPrivateCloud *v1.ListInstancesResponse_Instance_VirtualPrivateCloud
		if instance.InternetAccessible != nil {
			internetAccessible = &v1.ListInstancesResponse_Instance_InternetAccessible{
				InternetChargeType:      instance.InternetAccessible.InternetChargeType,
				InternetMaxBandwidthOut: instance.InternetAccessible.InternetMaxBandwidthOut,
			}
		}
		if instance.VirtualPrivateCloud != nil {
			virtualPrivateIpAddresses := make([]string, len(instance.VirtualPrivateCloud.PrivateIpAddresses))
			for j, v := range instance.VirtualPrivateCloud.PrivateIpAddresses {
				virtualPrivateIpAddresses[j] = *v
			}
			virtualPrivateCloud = &v1.ListInstancesResponse_Instance_VirtualPrivateCloud{
				VpcId:              instance.VirtualPrivateCloud.VpcId,
				SubnetId:           instance.VirtualPrivateCloud.SubnetId,
				AsVpcGateway:       instance.VirtualPrivateCloud.AsVpcGateway,
				PrivateIpAddresses: virtualPrivateIpAddresses,
				Ipv6AddressCount:   instance.VirtualPrivateCloud.Ipv6AddressCount,
			}
		}
		var securityGroupIds, rdmaIpAddresses []string
		if instance.SecurityGroupIds != nil {
			securityGroupIds = make([]string, len(instance.SecurityGroupIds))
			for j, v := range instance.SecurityGroupIds {
				securityGroupIds[j] = *v
			}
		}
		if instance.RdmaIpAddresses != nil {
			rdmaIpAddresses = make([]string, len(instance.RdmaIpAddresses))
			for j, v := range instance.RdmaIpAddresses {
				rdmaIpAddresses[j] = *v
			}
		}
		var tags []*v1.ListInstancesResponse_Instance_Tag
		if instance.Tags != nil {
			tags = make([]*v1.ListInstancesResponse_Instance_Tag, len(instance.Tags))
			for j, v := range instance.Tags {
				tags[j] = &v1.ListInstancesResponse_Instance_Tag{
					Key:   v.Key,
					Value: v.Value,
				}
			}
		}
		var ipv6Addresses []string
		if instance.IPv6Addresses != nil {
			ipv6Addresses = make([]string, len(instance.IPv6Addresses))
			for j, v := range instance.IPv6Addresses {
				ipv6Addresses[j] = *v
			}
		}
		var gpuInfo *v1.ListInstancesResponse_Instance_GPUInfo
		if instance.GPUInfo != nil {
			gpuId := make([]string, len(instance.GPUInfo.GPUId))
			for j, v := range instance.GPUInfo.GPUId {
				gpuId[j] = *v
			}
			var gpuCount float32
			if instance.GPUInfo.GPUCount != nil {
				gpuCount = float32(*instance.GPUInfo.GPUCount)
			}
			gpuInfo = &v1.ListInstancesResponse_Instance_GPUInfo{
				GpuCount: &gpuCount,
				GpuId:    gpuId,
				GpuType:  instance.GPUInfo.GPUType,
			}
		}
		instances[i] = &v1.ListInstancesResponse_Instance{
			Placement:                placement,
			InstanceId:               instance.InstanceId,
			InstanceType:             instance.InstanceType,
			Cpu:                      instance.CPU,
			Memory:                   instance.Memory,
			RestrictState:            instance.RestrictState,
			InstanceName:             instance.InstanceName,
			InstanceChargeType:       instance.InstanceChargeType,
			SystemDisk:               systemDisk,
			DataDisks:                dataDisks,
			PrivateIpAddresses:       privateIpAddresses,
			PublicIpAddresses:        publicIpAddresses,
			InternetAccessible:       internetAccessible,
			VirtualPrivateCloud:      virtualPrivateCloud,
			ImageId:                  instance.ImageId,
			RenewFlag:                instance.RenewFlag,
			CreatedTime:              instance.CreatedTime,
			ExpiredTime:              instance.ExpiredTime,
			OsName:                   instance.OsName,
			SecurityGroupIds:         securityGroupIds,
			InstanceState:            instance.InstanceState,
			Tags:                     tags,
			StopChargingMode:         instance.StopChargingMode,
			Uuid:                     instance.Uuid,
			LatestOperation:          instance.LatestOperation,
			LatestOperationState:     instance.LatestOperationState,
			LatestOperationRequestId: instance.LatestOperationRequestId,
			DisasterRecoverGroupId:   instance.DisasterRecoverGroupId,
			Ipv6Addresses:            ipv6Addresses,
			CamRoleName:              instance.CamRoleName,
			HpcClusterId:             instance.HpcClusterId,
			RdmaIpAddresses:          rdmaIpAddresses,
			IsolatedSource:           instance.IsolatedSource,
			GpuInfo:                  gpuInfo,
		}
	}
	listInstancesResponse := &v1.ListInstancesResponse{
		RequestId:  response.Response.RequestId,
		TotalCount: response.Response.TotalCount,
		Instances:  instances,
	}
	return listInstancesResponse, nil
}

func (s *InstanceService) StartInstances(ctx context.Context, request *v1.StartInstancesRequest) (*v1.StartInstancesResponse, error) {
	startInstanceRequest := cvm.NewStartInstancesRequest()
	startInstanceRequest.InstanceIds = common.StringPtrs(request.InstanceIds)
	response, err := s.uc.StartInstances(ctx, request.SecretId, request.SecretKey, request.Region, startInstanceRequest)
	if err != nil {
		return nil, err
	}
	return &v1.StartInstancesResponse{
		RequestId: response.Response.RequestId,
	}, nil
}

func (s *InstanceService) StopInstances(ctx context.Context, request *v1.StopInstancesRequest) (*v1.StopInstancesResponse, error) {
	stopInstanceRequest := cvm.NewStopInstancesRequest()
	stopInstanceRequest.InstanceIds = common.StringPtrs(request.InstanceIds)
	stopInstanceRequest.StopType = request.StopType
	stopInstanceRequest.StoppedMode = request.StoppedMode
	response, err := s.uc.StopInstances(ctx, request.SecretId, request.SecretKey, request.Region, stopInstanceRequest)
	if err != nil {
		return nil, err
	}
	return &v1.StopInstancesResponse{RequestId: response.Response.RequestId}, nil
}

func (s *InstanceService) RebootInstances(ctx context.Context, request *v1.RebootInstancesRequest) (*v1.RebootInstancesResponse, error) {
	rebootInstanceRequest := cvm.NewRebootInstancesRequest()
	rebootInstanceRequest.InstanceIds = common.StringPtrs(request.InstanceIds)
	rebootInstanceRequest.StopType = request.StopType
	response, err := s.uc.RebootInstances(ctx, request.SecretId, request.SecretKey, request.Region, rebootInstanceRequest)
	if err != nil {
		return nil, err
	}
	return &v1.RebootInstancesResponse{RequestId: response.Response.RequestId}, nil
}

func (s *InstanceService) DeleteInstances(ctx context.Context, request *v1.DeleteInstancesRequest) (*v1.DeleteInstancesResponse, error) {
	terminateInstanceRequest := cvm.NewTerminateInstancesRequest()
	terminateInstanceRequest.InstanceIds = common.StringPtrs(request.InstanceIds)
	response, err := s.uc.DeleteInstances(ctx, request.SecretId, request.SecretKey, request.Region, terminateInstanceRequest)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteInstancesResponse{RequestId: response.Response.RequestId}, nil
}
