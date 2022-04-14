package service

import (
	v1 "ali/api/instance/v1"
	"context"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

// CreateInstance Create Instance
func (service *InstanceService) CreateInstance(ctx context.Context, request *v1.CreateInstanceRequest) (*v1.CreateInstanceResponse, error) {
	result, err := service.uc.CreateInstance(ctx, request.AccessKeyId, request.AccessKeySecret, &ecs20140526.RunInstancesRequest{
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
	})
	if err != nil {
		return nil, err
	}
	instanceIdsSets := make([]string, len(result.InstanceIdSets.InstanceIdSet))
	for i, v := range result.InstanceIdSets.InstanceIdSet {
		instanceIdsSets[i] = *v
	}
	return &v1.CreateInstanceResponse{
		RequestId:      *result.RequestId,
		OrderId:        *result.OrderId,
		TradePrice:     *result.TradePrice,
		InstanceIdSets: instanceIdsSets,
	}, nil
}

func (service *InstanceService) ListInstance(ctx context.Context, request *v1.ListInstanceRequest) (*v1.ListInstanceResponse, error) {
	result, err := service.uc.ListInstance(ctx, request.AccessKeyId, request.AccessKeySecret, &ecs20140526.DescribeInstancesRequest{
		RegionId:   tea.String(request.RegionId),
		PageNumber: tea.Int32(int32(request.PageNumber)),
		PageSize:   tea.Int32(int32(request.PageSize)),
		NextToken:  tea.String(request.NextToken),
	})
	if err != nil {
		return nil, err
	}

	instances := make([]v1.ListInstanceResponse_Instance, len(result.Instances.Instance))
	for i, v := range result.Instances.Instance {
		publicIpAddress, innerIpAddress, rdmaIpAddress := make([]string, 0), make([]string, 0), make([]string, 0)
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
		instances[i] = v1.ListInstanceResponse_Instance{
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
			SecurityGroupIds:     *v.SecurityGroupIds,
			EipAddress:           nil,
			DeviceAvailable:      *v.DeviceAvailable,
			LocalStorageCapacity: uint64(*v.LocalStorageCapacity),
			LocalStorageAmount:   uint32(*v.LocalStorageAmount),
			ImageId:              *v.ImageId,
		}
	}
	listInstanceResponse := &v1.ListInstanceResponse{
		RequestId:  *result.RequestId,
		PageNumber: uint32(*result.PageNumber),
		PageSize:   uint32(*result.PageSize),
		TotalCount: uint32(*result.TotalCount),
		NextToken:  *result.NextToken,
		Instances:  nil,
	}

	return listInstanceResponse, nil
}

func (service *InstanceService) StartInstance(ctx context.Context, request *v1.StartInstanceRequest) (*v1.StartInstanceResponse, error) {
	result, err := service.uc.StartInstance(ctx, request.AccessKeyId, request.AccessKeySecret, &ecs20140526.StartInstanceRequest{
		InstanceId: tea.String(request.InstanceId),
	})
	if err != nil {
		return nil, err
	}
	return &v1.StartInstanceResponse{
		RequestId: *result,
	}, nil
}

func (service *InstanceService) StopInstance(ctx context.Context, request *v1.StopInstanceRequest) (*v1.StopInstanceResponse, error) {
	result, err := service.uc.StopInstance(ctx, request.AccessKeyId, request.AccessKeySecret, &ecs20140526.StopInstanceRequest{
		InstanceId: tea.String(request.InstanceId),
	})
	if err != nil {
		return nil, err
	}
	return &v1.StopInstanceResponse{
		RequestId: *result,
	}, nil
}

func (service *InstanceService) RebootInstance(ctx context.Context, request *v1.RebootInstanceRequest) (*v1.RebootInstanceResponse, error) {
	result, err := service.uc.RebootInstance(ctx, request.AccessKeyId, request.AccessKeySecret, &ecs20140526.RebootInstanceRequest{
		InstanceId: tea.String(request.InstanceId),
	})
	if err != nil {
		return nil, err
	}
	return &v1.RebootInstanceResponse{
		RequestId: *result,
	}, nil
}

func (service *InstanceService) DeleteInstance(ctx context.Context, request *v1.DeleteInstanceRequest) (*v1.DeleteInstanceResponse, error) {
	result, err := service.uc.DeleteInstance(ctx, request.AccessKeyId, request.AccessKeySecret, &ecs20140526.DeleteInstanceRequest{
		InstanceId: tea.String(request.InstanceId),
	})
	if err != nil {
		return nil, err
	}
	return &v1.DeleteInstanceResponse{
		RequestId: *result,
	}, nil
}
