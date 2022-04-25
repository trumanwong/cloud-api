package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	"huawei/internal/biz"
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

func (r *instanceTypeResponse) ListInstanceTypes(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.ListFlavorsRequest) (*model.ListFlavorsResponse, error) {
	client := getEcsClient(accessKey, secretKey, regionId, projectId)
	result, err := client.ListFlavors(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
