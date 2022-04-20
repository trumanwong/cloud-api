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

func (r *imageResponse) ListImages(ctx context.Context, accessToken string, request *linodego.ListOptions) ([]linodego.Image, error) {
	client := newClient(accessToken)
	result, err := client.ListImages(ctx, request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
