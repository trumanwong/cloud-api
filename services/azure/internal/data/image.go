package data

import (
	"azure/internal/biz"
	"context"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-12-01/compute"
	"github.com/go-kratos/kratos/v2/log"
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

func (r *imageResponse) ListImages(ctx context.Context, tenantID, clientID, clientSecret, subscriptionId string) (*compute.ImageListResultPage, error) {
	client, err := getImagesClient(tenantID, clientID, clientSecret, subscriptionId)
	if err != nil {
		return nil, err
	}

	result, err := client.List(ctx)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
