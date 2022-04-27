package biz

import (
	"context"
	"github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/go-kratos/kratos/v2/log"
)

// ImageResponse is a Image repo.
type ImageResponse interface {
	ListImages(context.Context, *string, *string, *string, *string) (*client.DescribeImagesResponse, error)
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

// ListImages List all Images
func (uc *ImageUseCase) ListImages(ctx context.Context, accessKeyId, accessKeySecret, endpoint, regionId *string) (*client.DescribeImagesResponse, error) {
	return uc.repo.ListImages(ctx, accessKeyId, accessKeySecret, endpoint, regionId)
}
