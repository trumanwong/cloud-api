package data

import (
	"azure/internal/biz"
	"context"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/go-kratos/kratos/v2/log"
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

func (r *instanceTypeResponse) ListAll(ctx context.Context, tenantID, clientID, clientSecret, subscriptionId string) ([]*ecs20140526.DescribeInstanceTypesResponseBodyInstanceTypesInstanceType, error) {
	client, err := getSubscriptionsClient(tenantID, clientID, clientSecret, subscriptionId)
	if err != nil {
		return nil, err
	}
	client.L
	if err != nil {
		return nil, err
	}
	return result.Body.InstanceTypes.InstanceType, nil
}
