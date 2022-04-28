package service

import (
	"context"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	v1 "tencent/api/instance/v1"
	"tencent/pkg/util/convert"
)

func (service *InstanceService) ListImages(ctx context.Context, request *v1.ListImagesRequest) (*v1.ListImagesResponse, error) {
	describeImagesRequest := cvm.NewDescribeImagesRequest()
	describeImagesRequest.Offset = request.Offset
	describeImagesRequest.Limit = request.Limit
	result, err := service.ic.ListImages(ctx, request.SecretId, request.SecretKey, request.Region, describeImagesRequest)
	if err != nil {
		return nil, err
	}

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.ListImagesResponse{Data: data}, nil
}
