package data

import (
	"aws/internal/biz"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/go-kratos/kratos/v2/log"
)

type imageResponse struct {
	data *Data
	log  *log.Helper
}

// NewImageRepo .
func NewImageRepo(data *Data, logger log.Logger) biz.ImageResponse {
	return &imageResponse{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *imageResponse) ListImage(ctx context.Context, accessKeyId, secretAccessKey, region string) (*ec2.DescribeImagesOutput, error) {
	client, err := getClient(
		accessKeyId,
		secretAccessKey,
		region,
	)
	if err != nil {
		return nil, err
	}
	describeImagesRequest := &ec2.DescribeImagesInput{}
	result, err := client.DescribeImages(ctx, describeImagesRequest)
	if err != nil {
		return nil, err
	}
	return result, nil
}
