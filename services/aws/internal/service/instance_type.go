package service

import (
	v1 "aws/api/instance/v1"
	"aws/pkg/util/convert"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func (service *InstanceService) ListInstanceTypes(ctx context.Context, request *v1.ListInstanceTypesRequest) (*v1.ListInstanceTypesResponse, error) {
	result, err := service.itc.ListInstanceTypes(ctx, request.AccessKeyId, request.SecretKey, request.Region, &ec2.DescribeInstanceTypesInput{
		MaxResults: request.MaxResults,
		NextToken:  request.NextToken,
	})
	if err != nil {
		return nil, err
	}

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.ListInstanceTypesResponse{Data: data}, nil
}
