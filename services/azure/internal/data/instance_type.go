package data

import (
	"azure/internal/biz"
	"context"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-12-01/compute"
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

func (r *instanceTypeResponse) ListInstanceTypes(ctx context.Context, tenantID, clientID, clientSecret, subscriptionId, filter, includeExtendedLocations string) (*compute.ResourceSkusResultPage, error) {
	client, err := getResourceSkusClient(tenantID, clientID, clientSecret, subscriptionId)
	if err != nil {
		return nil, err
	}
	result, err := client.List(ctx, filter, includeExtendedLocations)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
