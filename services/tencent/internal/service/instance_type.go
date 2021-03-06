package service

import (
	"context"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	v1 "tencent/api/instance/v1"
	"tencent/pkg/util/convert"
)

func (service *InstanceService) ListInstanceTypes(ctx context.Context, request *v1.ListInstanceTypesRequest) (*v1.ListInstanceTypesResponse, error) {
	result, err := service.itc.ListInstanceTypes(ctx, request.SecretId, request.SecretKey, request.Region, cvm.NewDescribeInstanceTypeConfigsRequest())
	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.ListInstanceTypesResponse{Data: data}, nil
}
