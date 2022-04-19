package service

import (
	"context"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	v1 "tencent/api/instance/v1"
)

func (s *InstanceService) CreateInstances(ctx context.Context, request *v1.CreateInstancesRequest) (*v1.CreateInstancesResponse, error) {
	var loginSettings *cvm.LoginSettings
	if len(request.Password) > 0 {
		loginSettings = &cvm.LoginSettings{Password: common.StringPtr(request.Password)}
	}
	response, err := s.uc.CreateInstances(ctx, request.SecretId, request.SecretKey, request.Region, &cvm.RunInstancesRequest{
		InstanceType: common.StringPtr(request.InstanceType),
		ImageId:      common.StringPtr(request.ImageId),
		SystemDisk: &cvm.SystemDisk{
			DiskSize: common.Int64Ptr(request.SystemDiskSize),
		},
		InstanceCount: common.Int64Ptr(request.Amount),
		InstanceName:  common.StringPtr(request.Name),
		LoginSettings: loginSettings,
		DryRun:        common.BoolPtr(request.DryRun),
	})
	if err != nil {
		return nil, err
	}

	instanceIdSets := make([]string, len(response.Response.InstanceIdSet))
	for i, v := range response.Response.InstanceIdSet {
		instanceIdSets[i] = *v
	}

	createInstancesResponse := &v1.CreateInstancesResponse{
		RequestId:      *response.Response.RequestId,
		InstanceIdSets: instanceIdSets,
	}
	return createInstancesResponse, nil
}

func (s *InstanceService) ListInstances(ctx context.Context, request *v1.ListInstancesRequest) (*v1.ListInstancesResponse, error) {
	response, err := s.uc.ListInstances(ctx, request.SecretId, request.SecretKey, request.Region, &cvm.DescribeInstancesRequest{
		Offset: common.Int64Ptr(request.PageNumber - 1),
		Limit:  common.Int64Ptr(request.PageSize),
	})
	if err != nil {
		return nil, err
	}
	instances := make([]*v1.ListInstancesResponse_Instance, len(response.Response.InstanceSet))
	for i, instance := range response.Response.InstanceSet {
		var gpuInfo *v1.ListInstancesResponse_Instance_GpuInfo
		if instance.GPUInfo != nil {
			gpuId := make([]string, len(instance.GPUInfo.GPUId))
			for j, v := range instance.GPUInfo.GPUId {
				gpuId[j] = *v
			}

			gpuInfo = &v1.ListInstancesResponse_Instance_GpuInfo{
				GpuCount: float32(*instance.GPUInfo.GPUCount),
				GpuId:    gpuId,
				GpuType:  *instance.GPUInfo.GPUType,
			}
		}
		var systemDisk *v1.ListInstancesResponse_Instance_SystemDisk
		if instance.SystemDisk != nil {
			systemDisk = &v1.ListInstancesResponse_Instance_SystemDisk{
				DiskType: *instance.SystemDisk.DiskType,
				DiskId:   *instance.SystemDisk.DiskId,
				DiskSize: *instance.SystemDisk.DiskSize,
				CdCid:    *instance.SystemDisk.CdcId,
			}
		}
		dataDisks := make([]*v1.ListInstancesResponse_Instance_DataDisk, len(instance.DataDisks))
		for j, v := range instance.DataDisks {
			dataDisks[j] = &v1.ListInstancesResponse_Instance_DataDisk{
				DiskSize: *v.DiskSize,
				DiskType: *v.DiskType,
				DiskId:   *v.DiskId,
				CdCid:    *v.CdcId,
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
				InternetChargeType:      *instance.InternetAccessible.InternetChargeType,
				InternetMaxBandwidthOut: *instance.InternetAccessible.InternetMaxBandwidthOut,
				PublicIpAssigned:        *instance.InternetAccessible.PublicIpAssigned,
				BandwidthPackageId:      *instance.InternetAccessible.BandwidthPackageId,
			}
		}
		if instance.VirtualPrivateCloud != nil {
			virtualPrivateIpAddresses := make([]string, len(instance.VirtualPrivateCloud.PrivateIpAddresses))
			for j, v := range instance.VirtualPrivateCloud.PrivateIpAddresses {
				virtualPrivateIpAddresses[j] = *v
			}
			virtualPrivateCloud = &v1.ListInstancesResponse_Instance_VirtualPrivateCloud{
				VpcId:              *instance.VirtualPrivateCloud.VpcId,
				SubnetId:           *instance.VirtualPrivateCloud.SubnetId,
				AsVpcGateway:       *instance.VirtualPrivateCloud.AsVpcGateway,
				PrivateIpAddresses: virtualPrivateIpAddresses,
				Ipv6AddressCount:   *instance.VirtualPrivateCloud.Ipv6AddressCount,
			}
		}
		instances[i] = &v1.ListInstancesResponse_Instance{
			InstanceId:          *instance.InstanceId,
			InstanceName:        *instance.InstanceName,
			InstanceType:        *instance.InstanceType,
			InstanceChargeType:  *instance.InstanceChargeType,
			Cpu:                 *instance.CPU,
			OsName:              *instance.OsName,
			Status:              *instance.InstanceState,
			CreatedTime:         *instance.CreatedTime,
			Memory:              *instance.Memory,
			GpuInfo:             gpuInfo,
			ExpiredTime:         *instance.ExpiredTime,
			SystemDisk:          systemDisk,
			DataDisks:           dataDisks,
			PrivateIpAddresses:  privateIpAddresses,
			PublicIpAddresses:   publicIpAddresses,
			InternetAccessible:  internetAccessible,
			VirtualPrivateCloud: virtualPrivateCloud,
			ImageId:             *instance.ImageId,
		}
	}
	listInstancesResponse := &v1.ListInstancesResponse{
		RequestId:  *response.Response.RequestId,
		TotalCount: *response.Response.TotalCount,
		Instances:  instances,
	}
	return listInstancesResponse, nil
}

func (s *InstanceService) StartInstances(ctx context.Context, request *v1.StartInstancesRequest) (*v1.StartInstancesResponse, error) {
	response, err := s.uc.StartInstances(ctx, request.SecretId, request.SecretKey, request.Region, &cvm.StartInstancesRequest{
		InstanceIds: common.StringPtrs(request.InstanceIds),
	})
	if err != nil {
		return nil, err
	}
	return &v1.StartInstancesResponse{RequestId: *response.Response.RequestId}, nil
}

func (s *InstanceService) StopInstances(ctx context.Context, request *v1.StopInstancesRequest) (*v1.StopInstancesResponse, error) {
	response, err := s.uc.StopInstances(ctx, request.SecretId, request.SecretKey, request.Region, &cvm.StopInstancesRequest{
		InstanceIds: common.StringPtrs(request.InstanceIds),
	})
	if err != nil {
		return nil, err
	}
	return &v1.StopInstancesResponse{RequestId: *response.Response.RequestId}, nil
}

func (s *InstanceService) RebootInstances(ctx context.Context, request *v1.RebootInstancesRequest) (*v1.RebootInstancesResponse, error) {
	response, err := s.uc.RebootInstances(ctx, request.SecretId, request.SecretKey, request.Region, &cvm.RebootInstancesRequest{
		InstanceIds: common.StringPtrs(request.InstanceIds),
	})
	if err != nil {
		return nil, err
	}
	return &v1.RebootInstancesResponse{RequestId: *response.Response.RequestId}, nil
}

func (s *InstanceService) DeleteInstances(ctx context.Context, request *v1.DeleteInstancesRequest) (*v1.DeleteInstancesResponse, error) {
	response, err := s.uc.DeleteInstances(ctx, request.SecretId, request.SecretKey, request.Region, &cvm.TerminateInstancesRequest{
		InstanceIds: common.StringPtrs(request.InstanceIds),
	})
	if err != nil {
		return nil, err
	}
	return &v1.DeleteInstancesResponse{RequestId: *response.Response.RequestId}, nil
}
