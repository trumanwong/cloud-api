package service

import (
	"context"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	v1 "tencent/api/instance/v1"
)

func (service *InstanceService) ListInstanceType(ctx context.Context, request *v1.ListInstanceTypeRequest) (*v1.ListInstanceTypeResponse, error) {
	response, err := service.itc.ListAll(ctx, request.SecretId, request.SecretKey, request.Region, cvm.NewDescribeInstanceTypeConfigsRequest())
	if err != nil {
		return nil, err
	}

	listInstanceTypeResponse := &v1.ListInstanceTypeResponse{
		InstanceTypes: make([]*v1.ListInstanceTypeResponse_InstanceType, len(response.Response.InstanceTypeConfigSet)),
	}

	for i, instanceType := range response.Response.InstanceTypeConfigSet {
		listInstanceTypeResponse.InstanceTypes[i] = &v1.ListInstanceTypeResponse_InstanceType{
			InstanceTypeId:     *instanceType.InstanceType,
			InstanceTypeFamily: *instanceType.InstanceFamily,
			MemorySize:         *instanceType.Memory,
			CpuCoreCount:       *instanceType.CPU,
			GpuSec:             *instanceType.GPU,
		}
	}
	return listInstanceTypeResponse, nil
}
