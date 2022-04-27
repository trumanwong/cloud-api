package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"tencent/internal/biz"
)

type instanceTypeResponse struct {
	data *Data
	log  *log.Helper
}

// NewInstanceTypeRepo .
func NewInstanceTypeRepo(data *Data, logger log.Logger) biz.InstanceTypeResponse {
	return &instanceTypeResponse{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *instanceTypeResponse) ListInstanceTypes(ctx context.Context, secretId, secretKey, region string, request *cvm.DescribeInstanceTypeConfigsRequest) (*cvm.DescribeInstanceTypeConfigsResponse, error) {
	client, err := getCvmClient(secretId, secretKey, region)
	if err != nil {
		return nil, err
	}
	response, err := client.DescribeInstanceTypeConfigs(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
