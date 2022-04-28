package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/vultr/govultr/v2"
)

type ListImagesResponse struct {
	Images []govultr.OS  `json:"images"`
	Meta   *govultr.Meta `json:"meta"`
}

// ImageResponse is a Greater repo.
type ImageResponse interface {
	ListImages(context.Context, string, *govultr.ListOptions) (*ListImagesResponse, error)
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

// ListImages List Images
func (uc *ImageUseCase) ListImages(ctx context.Context, accessToken string, request *govultr.ListOptions) (*ListImagesResponse, error) {
	return uc.repo.ListImages(ctx, accessToken, request)
}
