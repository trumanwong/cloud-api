package data

import (
	"ali/internal/biz"
	"context"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/go-kratos/kratos/v2/log"
)

type instanceTypeResponse struct {
	data *Data
	log  *log.Helper
}

// NewInstanceTypeRepo .
func NewInstanceTypeRepo(data *Data, logger log.Logger) biz.InstanceTypeResponse {
	return &instanceTypeResponse{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *instanceTypeResponse) ListAll(ctx context.Context, AccessKeyId, AccessKeySecret string) ([]*ecs20140526.DescribeInstanceTypesResponseBodyInstanceTypesInstanceType, error) {
	client, err := createClient(
		tea.String(AccessKeyId),
		tea.String(AccessKeySecret),
	)
	if err != nil {
		return nil, err
	}
	describeInstanceTypesRequest := &ecs20140526.DescribeInstanceTypesRequest{}
	result, err := client.DescribeInstanceTypes(describeInstanceTypesRequest)
	if err != nil {
		return nil, err
	}
	return result.Body.InstanceTypes.InstanceType, nil
}
