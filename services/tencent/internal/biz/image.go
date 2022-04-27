package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

type ImageResponse interface {
	ListImages(context.Context, string, string, string, *cvm.DescribeImagesRequest) (*cvm.DescribeImagesResponse, error)
}

// ImageUseCase is Image Region UseCase.
type ImageUseCase struct {
	repo ImageResponse
	log  *log.Helper
}

// NewImageUseCase new a Image UseCase.
func NewImageUseCase(repo ImageResponse, logger log.Logger) *ImageUseCase {
	return &ImageUseCase{repo: repo, log: log.NewHelper(logger)}
}

// ListImages List All Images
func (uc *ImageUseCase) ListImages(ctx context.Context, secretId, secretKey, region string, request *cvm.DescribeImagesRequest) (*cvm.DescribeImagesResponse, error) {
	return uc.repo.ListImages(ctx, secretId, secretKey, region, request)
}
