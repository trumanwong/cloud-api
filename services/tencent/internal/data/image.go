package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"tencent/internal/biz"
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

func (r *imageResponse) ListImages(ctx context.Context, secretId, secretKey, region string, request *cvm.DescribeImagesRequest) (*cvm.DescribeImagesResponse, error) {
	client, err := getCvmClient(secretId, secretKey, region)
	if err != nil {
		return nil, err
	}
	response, err := client.DescribeImages(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
