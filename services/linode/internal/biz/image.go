package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/linode/linodego"
)

// ImageResponse is a Greater repo.
type ImageResponse interface {
	ListImages(context.Context, string, *linodego.ListOptions) (*ListImagesResponse, error)
}

type ListImagesResponse struct {
	Images []linodego.Image `json:"images"`
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

// ListImages List Image
func (uc *ImageUseCase) ListImages(ctx context.Context, accessToken string, request *linodego.ListOptions) (*ListImagesResponse, error) {
	return uc.repo.ListImages(ctx, accessToken, request)
}
