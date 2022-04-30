package service

import (
	v1 "aws/api/instance/v1"
	"aws/pkg/util/convert"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func (service *InstanceService) ListRegions(ctx context.Context, request *v1.ListRegionsRequest) (*v1.ListRegionsResponse, error) {
	result, err := service.rc.ListRegions(ctx, request.AccessKeyId, request.SecretKey, &ec2.DescribeRegionsInput{})
	if err != nil {
		return nil, err
	}

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.ListRegionsResponse{Data: data}, nil
}
