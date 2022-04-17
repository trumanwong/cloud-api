package service

import (
	"context"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	v1 "tencent/api/instance/v1"
)

func (s *InstanceService) CreateInstance(ctx context.Context, request *v1.CreateInstanceRequest) (*v1.CreateInstanceResponse, error) {
	var loginSettings *cvm.LoginSettings
	{
	}
	if len(request.Password) > 0 {
		loginSettings = &cvm.LoginSettings{Password: common.StringPtr(request.Password)}
	}
	response, err := s.uc.Create(ctx, request.SecretId, request.SecretKey, request.Region, &cvm.RunInstancesRequest{
		InstanceType: common.StringPtr(request.InstanceType),
		ImageId:      common.StringPtr(request.ImageId),
		SystemDisk: &cvm.SystemDisk{
			DiskSize: common.Int64Ptr(request.SystemDiskSize),
		},
		InstanceCount: common.Int64Ptr(request.Amount),
		InstanceName:  common.StringPtr(request.Name),
		LoginSettings: loginSettings,
	})
	if err != nil {
		return nil, err
	}

	instanceIdSets := make([]string, len(response.Response.InstanceIdSet))
	for i, v := range response.Response.InstanceIdSet {
		instanceIdSets[i] = *v
	}

	createInstanceResponse := &v1.CreateInstanceResponse{
		RequestId:      *response.Response.RequestId,
		InstanceIdSets: instanceIdSets,
	}
	return createInstanceResponse, nil
}

func (s *InstanceService) ListInstance(ctx context.Context, request *v1.ListInstanceRequest) (*v1.ListInstanceResponse, error) {
	response, err := s.uc.ListInstance(ctx, request.SecretId, request.SecretKey, request.Region, &cvm.DescribeInstancesRequest{
		Offset: common.Int64Ptr(request.PageNumber - 1),
		Limit:  common.Int64Ptr(request.PageSize),
	})
	if err != nil {
		return nil, err
	}
	instances := make([]*v1.ListInstanceResponse_Instance, len(response.Response.InstanceSet))
	for i, instance := range response.Response.InstanceSet {
		var gpuInfo *v1.ListInstanceResponse_Instance_GpuInfo
		if instance.GPUInfo != nil {
			gpuId := make([]string, len(instance.GPUInfo.GPUId))
			for j, v := range instance.GPUInfo.GPUId {
				gpuId[j] = *v
			}

			gpuInfo = &v1.ListInstanceResponse_Instance_GpuInfo{
				GpuCount: float32(*instance.GPUInfo.GPUCount),
				GpuId:    gpuId,
				GpuType:  *instance.GPUInfo.GPUType,
			}
		}
		var systemDisk *v1.ListInstanceResponse_Instance_SystemDisk
		if instance.SystemDisk != nil {
			systemDisk = &v1.ListInstanceResponse_Instance_SystemDisk{
				DiskType: *instance.SystemDisk.DiskType,
				DiskId:   *instance.SystemDisk.DiskId,
				DiskSize: *instance.SystemDisk.DiskSize,
				CdCid:    *instance.SystemDisk.CdcId,
			}
		}
		dataDisks := make([]*v1.ListInstanceResponse_Instance_DataDisk, len(instance.DataDisks))
		for j, v := range instance.DataDisks {
			dataDisks[j] = &v1.ListInstanceResponse_Instance_DataDisk{
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
		var internetAccessible *v1.ListInstanceResponse_Instance_InternetAccessible
		var virtualPrivateCloud *v1.ListInstanceResponse_Instance_VirtualPrivateCloud
		if instance.InternetAccessible != nil {
			internetAccessible = &v1.ListInstanceResponse_Instance_InternetAccessible{
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
			virtualPrivateCloud = &v1.ListInstanceResponse_Instance_VirtualPrivateCloud{
				VpcId:              *instance.VirtualPrivateCloud.VpcId,
				SubnetId:           *instance.VirtualPrivateCloud.SubnetId,
				AsVpcGateway:       *instance.VirtualPrivateCloud.AsVpcGateway,
				PrivateIpAddresses: virtualPrivateIpAddresses,
				Ipv6AddressCount:   *instance.VirtualPrivateCloud.Ipv6AddressCount,
			}
		}
		instances[i] = &v1.ListInstanceResponse_Instance{
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
	listInstanceResponse := &v1.ListInstanceResponse{
		RequestId:  *response.Response.RequestId,
		TotalCount: *response.Response.TotalCount,
		Instances:  instances,
	}
	return listInstanceResponse, nil
}

func (s *InstanceService) StartInstance(ctx context.Context, request *v1.StartInstanceRequest) (*v1.StartInstanceResponse, error) {
	response, err := s.uc.StartInstance(ctx, request.SecretId, request.SecretKey, request.Region, &cvm.StartInstancesRequest{
		InstanceIds: common.StringPtrs(request.InstanceIds),
	})
	if err != nil {
		return nil, err
	}
	return &v1.StartInstanceResponse{RequestId: *response.Response.RequestId}, nil
}

func (s *InstanceService) StopInstance(ctx context.Context, request *v1.StopInstanceRequest) (*v1.StopInstanceResponse, error) {
	response, err := s.uc.StopInstance(ctx, request.SecretId, request.SecretKey, request.Region, &cvm.StopInstancesRequest{
		InstanceIds: common.StringPtrs(request.InstanceIds),
	})
	if err != nil {
		return nil, err
	}
	return &v1.StopInstanceResponse{RequestId: *response.Response.RequestId}, nil
}

func (s *InstanceService) RebootInstance(ctx context.Context, request *v1.RebootInstanceRequest) (*v1.RebootInstanceResponse, error) {
	response, err := s.uc.RebootInstance(ctx, request.SecretId, request.SecretKey, request.Region, &cvm.RebootInstancesRequest{
		InstanceIds: common.StringPtrs(request.InstanceIds),
	})
	if err != nil {
		return nil, err
	}
	return &v1.RebootInstanceResponse{RequestId: *response.Response.RequestId}, nil
}

func (s *InstanceService) DeleteInstance(ctx context.Context, request *v1.DeleteInstanceRequest) (*v1.DeleteInstanceResponse, error) {
	response, err := s.uc.DeleteInstance(ctx, request.SecretId, request.SecretKey, request.Region, &cvm.TerminateInstancesRequest{
		InstanceIds: common.StringPtrs(request.InstanceIds),
	})
	if err != nil {
		return nil, err
	}
	return &v1.DeleteInstanceResponse{RequestId: *response.Response.RequestId}, nil
}
