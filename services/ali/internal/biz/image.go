package biz

import (
	"context"
	"github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/go-kratos/kratos/v2/log"
)

// ImageResponse is a Greater repo.
type ImageResponse interface {
	ListImage(context context.Context, accessKeyId, accessKeySecret, regionId string) ([]*client.DescribeImagesResponseBodyImagesImage, error)
}

// ImageUseCase is a Image UseCase.
type ImageUseCase struct {
	repo ImageResponse
	log  *log.Helper
}

// NewImageUseCase new a Image UseCase.
func NewImageUseCase(repo ImageResponse, logger log.Logger) *ImageUseCase {
	return &ImageUseCase{repo: repo, log: log.NewHelper(logger)}
}

// ListImage List Image
func (uc *ImageUseCase) ListImage(ctx context.Context, accessKeyId, accessKeySecret, regionId string) ([]*client.DescribeImagesResponseBodyImagesImage, error) {
	return uc.repo.ListImage(ctx, accessKeyId, accessKeySecret, regionId)
}
