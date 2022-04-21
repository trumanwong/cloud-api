package biz

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	"github.com/go-kratos/kratos/v2/log"
)

// ImageResponse is a Greater repo.
type ImageResponse interface {
	ListImage(context.Context, string, string, string, string) (*compute.ImageListResultPage, error)
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
func (uc *ImageUseCase) ListImage(ctx context.Context, tenantID, clientID, clientSecret, subscriptionId string) (*compute.ImageListResultPage, error) {
	return uc.repo.ListImage(ctx, tenantID, clientID, clientSecret, subscriptionId)
}
