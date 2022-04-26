package data

import (
	"aws/internal/biz"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
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

func (r *instanceTypeResponse) ListAll(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.DescribeInstanceTypesInput) (*ec2.DescribeInstanceTypesOutput, error) {
	client, err := getClient(
		accessKeyId,
		secretAccessKey,
		region,
	)
	if err != nil {
		return nil, err
	}
	result, err := client.DescribeInstanceTypes(ctx, request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
