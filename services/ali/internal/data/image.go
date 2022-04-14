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

func (r *imageResponse) ListImage(ctx context.Context, request *biz.ListImageRequest) ([]*biz.Image, error) {
	client, err := createClient(
		tea.String(request.AccessKeyId),
		tea.String(request.AccessKeySecret),
	)
	if err != nil {
		return nil, err
	}
	describeImagesRequest := &ecs20140526.DescribeImagesRequest{
		RegionId: tea.String(request.RegionId),
	}
	result, err := client.DescribeImages(describeImagesRequest)
	if err != nil {
		return nil, err
	}
	images := make([]*biz.Image, len(result.Body.Images.Image))
	for i, region := range result.Body.Images.Image {
		images[i] = &biz.Image{
			ImageName:    *region.ImageName,
			OsName:       *region.OSName,
			OsNameEn:     *region.OSNameEn,
			Architecture: *region.Architecture,
			OsType:       *region.OSType,
		}
	}
	return images, nil
}
