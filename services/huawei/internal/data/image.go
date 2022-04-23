package data

import (
	"ali/internal/biz"
	"context"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
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

func (r *imageResponse) ListImage(ctx context.Context, accessKeyId, accessKeySecret, endpoint, regionId string) ([]*ecs20140526.DescribeImagesResponseBodyImagesImage, error) {
	client, err := getClient(
		accessKeyId,
		accessKeySecret,
		endpoint,
	)
	if err != nil {
		return nil, err
	}
	describeImagesRequest := &ecs20140526.DescribeImagesRequest{
		RegionId: tea.String(regionId),
	}
	result, err := client.DescribeImages(describeImagesRequest)
	if err != nil {
		return nil, err
	}
	return result.Body.Images.Image, nil
}
