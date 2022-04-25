package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ims/v2/model"
	"huawei/internal/biz"
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

func (r *imageResponse) ListImages(ctx context.Context, accessKey, secretKey, regionId string, request *model.GlanceListImagesRequest) (*model.GlanceListImagesResponse, error) {
	client := getImsClient(accessKey, secretKey, regionId)
	result, err := client.GlanceListImages(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
