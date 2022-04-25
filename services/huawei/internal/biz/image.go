package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ims/v2/model"
)

// ImageResponse is a Greater repo.
type ImageResponse interface {
	ListImages(context.Context, string, string, string, *model.GlanceListImagesRequest) (*model.GlanceListImagesResponse, error)
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
func (uc *ImageUseCase) ListImages(ctx context.Context, accessKey, secretKey, regionId string, request *model.GlanceListImagesRequest) (*model.GlanceListImagesResponse, error) {
	return uc.repo.ListImages(ctx, accessKey, secretKey, regionId, request)
}
