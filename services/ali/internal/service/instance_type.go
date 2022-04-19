package service

import (
	v1 "ali/api/instance/v1"
	"context"
)

func (service *InstanceService) ListInstanceTypes(ctx context.Context, request *v1.ListInstanceTypesRequest) (*v1.ListInstanceTypesResponse, error) {
	instanceTypes, err := service.itc.ListAll(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint)
	if err != nil {
		return nil, err
	}

	listInstanceTypesResponse := &v1.ListInstanceTypesResponse{
		InstanceTypes: make([]*v1.ListInstanceTypesResponse_InstanceType, len(instanceTypes)),
	}
	for i, instanceType := range instanceTypes {
		listInstanceTypesResponse.InstanceTypes[i] = &v1.ListInstanceTypesResponse_InstanceType{
			InstanceTypeId:     *instanceType.InstanceTypeId,
			InstanceTypeFamily: *instanceType.InstanceTypeFamily,
			MemorySize:         *instanceType.MemorySize,
			CpuCoreCount:       *instanceType.CpuCoreCount,
			GpuSec:             *instanceType.GPUSpec,
		}
	}
	return listInstanceTypesResponse, nil
}
