package service

import (
	v1 "ali/api/instance/v1"
	"context"
)

func (service *InstanceService) ListInstanceType(ctx context.Context, request *v1.ListInstanceTypeRequest) (*v1.ListInstanceTypeResponse, error) {
	instanceTypes, err := service.itc.ListAll(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint)
	if err != nil {
		return nil, err
	}

	listInstanceTypeResponse := &v1.ListInstanceTypeResponse{
		InstanceTypes: make([]*v1.ListInstanceTypeResponse_InstanceType, len(instanceTypes)),
	}
	for i, instanceType := range instanceTypes {
		listInstanceTypeResponse.InstanceTypes[i] = &v1.ListInstanceTypeResponse_InstanceType{
			InstanceTypeId:     *instanceType.InstanceTypeId,
			InstanceTypeFamily: *instanceType.InstanceTypeFamily,
			MemorySize:         *instanceType.MemorySize,
			CpuCoreCount:       *instanceType.CpuCoreCount,
			GpuSec:             *instanceType.GPUSpec,
		}
	}
	return listInstanceTypeResponse, nil
}
