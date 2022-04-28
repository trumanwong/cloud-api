package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/linode/linodego"
	"linode/internal/biz"
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

func (r *imageResponse) ListImages(ctx context.Context, accessToken string, request *linodego.ListOptions) (*biz.ListImagesResponse, error) {
	client := newClient(accessToken)
	images, err := client.ListImages(ctx, request)
	if err != nil {
		return nil, err
	}
	return &biz.ListImagesResponse{Images: images}, nil
}
