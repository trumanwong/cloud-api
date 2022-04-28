package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/linode/linodego"
	"linode/internal/biz"
)

type instanceResponse struct {
	data *Data
	log  *log.Helper
}

// NewInstanceRepo .
func NewInstanceRepo(data *Data, logger log.Logger) biz.InstanceResponse {
	return &instanceResponse{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *instanceResponse) CreateInstance(ctx context.Context, accessToken string, request linodego.InstanceCreateOptions) (*linodego.Instance, error) {
	client := newClient(accessToken)
	result, err := client.CreateInstance(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *instanceResponse) ListInstances(ctx context.Context, accessToken string, request *linodego.ListOptions) (*biz.ListInstancesResponse, error) {
	client := newClient(accessToken)
	instances, err := client.ListInstances(ctx, request)
	if err != nil {
		return nil, err
	}
	return &biz.ListInstancesResponse{Instances: instances}, nil
}

func (r *instanceResponse) StartInstance(ctx context.Context, accessToken string, instanceId, configId int) error {
	client := newClient(accessToken)
	return client.BootInstance(ctx, instanceId, configId)
}

func (r *instanceResponse) StopInstance(ctx context.Context, accessToken string, instanceId int) error {
	client := newClient(accessToken)
	return client.ShutdownInstance(ctx, instanceId)
}

func (r *instanceResponse) RebootInstance(ctx context.Context, accessToken string, instanceId, configId int) error {
	client := newClient(accessToken)
	return client.RebootInstance(ctx, instanceId, configId)
}

func (r *instanceResponse) DeleteInstance(ctx context.Context, accessToken string, instanceId int) error {
	client := newClient(accessToken)
	return client.DeleteInstance(ctx, instanceId)
}
