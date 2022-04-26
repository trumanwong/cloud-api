package data

import (
	"ali/internal/biz"
	"context"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
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

func (r *instanceTypeResponse) ListInstanceTypes(ctx context.Context, accessKeyId, accessKeySecret, endpoint string) (*ecs20140526.DescribeInstanceTypesResponse, error) {
	client, err := getClient(
		accessKeyId,
		accessKeySecret,
		endpoint,
	)
	if err != nil {
		return nil, err
	}
	describeInstanceTypesRequest := &ecs20140526.DescribeInstanceTypesRequest{}
	result, err := client.DescribeInstanceTypes(describeInstanceTypesRequest)
	if err != nil {
		return nil, err
	}
	return result, nil
}
