package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Image struct {
	ImageName    string
	OsName       string
	OsNameEn     string
	Architecture string
	OsType       string
}

type ListImageRequest struct {
	AccessKeyId     string
	AccessKeySecret string
	RegionId        string
}

// ImageResponse is a Greater repo.
type ImageResponse interface {
	ListImage(context.Context, *ListImageRequest) ([]*Image, error)
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
func (uc *ImageUseCase) ListImage(ctx context.Context, g *ListImageRequest) ([]*Image, error) {
	return uc.repo.ListImage(ctx, g)
}
