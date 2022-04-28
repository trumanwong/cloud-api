package data

import (
	"aws/internal/biz"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/go-kratos/kratos/v2/log"
)

type regionResponse struct {
	data *Data
	log  *log.Helper
}

// NewRegionRepo .
func NewRegionRepo(data *Data, logger log.Logger) biz.RegionResponse {
	return &regionResponse{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *regionResponse) ListRegions(ctx context.Context, accessKeyId, secretAccessKey string, request *ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error) {
	client, err := getClient(accessKeyId, secretAccessKey, "")
	if err != nil {
		return nil, err
	}
	result, err := client.DescribeRegions(ctx, request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
