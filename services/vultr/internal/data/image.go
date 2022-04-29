package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/vultr/govultr/v2"
	"vultr/internal/biz"
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

func (r *imageResponse) ListImages(ctx context.Context, accessToken string, request *govultr.ListOptions) (*biz.ListImagesResponse, error) {
	client := getClient(ctx, accessToken)
	images, meta, err := client.OS.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return &biz.ListImagesResponse{Images: images, Meta: meta}, nil
}
