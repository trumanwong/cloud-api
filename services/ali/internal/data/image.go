package data

import (
	"ali/internal/biz"
	"context"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
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

func (r *imageResponse) ListImages(ctx context.Context, accessKeyId, accessKeySecret, endpoint, regionId *string) (*ecs20140526.DescribeImagesResponse, error) {
	client, err := getClient(
		accessKeyId,
		accessKeySecret,
		endpoint,
	)
	if err != nil {
		return nil, err
	}
	describeImagesRequest := &ecs20140526.DescribeImagesRequest{
		RegionId: regionId,
	}
	result, err := client.DescribeImages(describeImagesRequest)
	if err != nil {
		return nil, err
	}
	return result, nil
}
