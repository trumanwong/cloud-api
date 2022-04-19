package biz

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/go-kratos/kratos/v2/log"
)

// ImageResponse is a Greater repo.
type ImageResponse interface {
	ListImage(ctx context.Context, accessKeyId, secretAccessKey, region string) (*ec2.DescribeImagesOutput, error)
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
func (uc *ImageUseCase) ListImage(ctx context.Context, accessKeyId, secretAccessKey, region string) (*ec2.DescribeImagesOutput, error) {
	return uc.repo.ListImage(ctx, accessKeyId, secretAccessKey, region)
}
