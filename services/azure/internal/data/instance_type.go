package data

import (
	"azure/internal/biz"
	"context"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/subscriptions"
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

func (r *instanceTypeResponse) ListInstanceType(ctx context.Context, tenantID, clientID, clientSecret, subscriptionId string, includeExtendedLocations bool) (*[]subscriptions.Location, error) {
	client, err := getSubscriptionsClient(tenantID, clientID, clientSecret, subscriptionId)
	if err != nil {
		return nil, err
	}
	result, err := client.ListLocations(ctx, subscriptionId, &includeExtendedLocations)
	if err != nil {
		return nil, err
	}
	return result.Value, nil
}
