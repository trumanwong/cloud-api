package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"tencent/internal/biz"
)

type regionResponse struct {
	data *Data
	log  *log.Helper
}

// NewRegionRepo .
func NewRegionRepo(data *Data, logger log.Logger) biz.RegionResponse {
	return &regionResponse{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *regionResponse) ListAll(ctx context.Context, secretId, secretKey string, request *cvm.DescribeRegionsRequest) (*cvm.DescribeRegionsResponse, error) {
	client, err := NewCvmClient(secretId, secretKey, "")
	if err != nil {
		return nil, err
	}
	response, err := client.DescribeRegions(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
