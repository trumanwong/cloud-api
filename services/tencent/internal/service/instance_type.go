package service

import (
	"context"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	v1 "tencent/api/instance/v1"
)

func (service *InstanceService) ListInstanceTypes(ctx context.Context, request *v1.ListInstanceTypesRequest) (*v1.ListInstanceTypesResponse, error) {
	response, err := service.itc.ListInstanceTypes(ctx, request.SecretId, request.SecretKey, request.Region, cvm.NewDescribeInstanceTypeConfigsRequest())
	if err != nil {
		return nil, err
	}

	listInstanceTypeResponse := &v1.ListInstanceTypesResponse{
		InstanceTypes: make([]*v1.ListInstanceTypesResponse_InstanceType, len(response.Response.InstanceTypeConfigSet)),
		RequestId:     response.Response.RequestId,
	}

	for i, instanceType := range response.Response.InstanceTypeConfigSet {
		listInstanceTypeResponse.InstanceTypes[i] = &v1.ListInstanceTypesResponse_InstanceType{
			Zone:           instanceType.Zone,
			InstanceType:   instanceType.InstanceType,
			InstanceFamily: instanceType.InstanceFamily,
			Gpu:            instanceType.GPU,
			Cpu:            instanceType.CPU,
			Memory:         instanceType.Memory,
			Fpga:           instanceType.FPGA,
		}
	}
	return listInstanceTypeResponse, nil
}
