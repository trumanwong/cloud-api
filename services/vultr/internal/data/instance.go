package data

import (
	"context"
	"github.com/vultr/govultr/v2"

	"github.com/go-kratos/kratos/v2/log"
	"vultr/internal/biz"
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

func (r *instanceResponse) CreateInstance(ctx context.Context, accessToken string, request *govultr.InstanceCreateReq) (*govultr.Instance, error) {
	client := getClient(accessToken)
	result, err := client.Instance.Create(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *instanceResponse) ListInstances(ctx context.Context, accessToken string, request *govultr.ListOptions) (*biz.ListInstancesResponse, error) {
	client := getClient(accessToken)
	instances, meta, err := client.Instance.List(ctx, request)
	if err != nil {
		return nil, err
	}
	return &biz.ListInstancesResponse{
		Instances: instances,
		Meta:      meta,
	}, nil
}

func (r *instanceResponse) StartInstance(ctx context.Context, accessToken, instanceId string) error {
	return getClient(accessToken).Instance.Start(ctx, instanceId)
}

func (r *instanceResponse) StopInstance(ctx context.Context, accessToken, instanceId string) error {
	return getClient(accessToken).Instance.Halt(ctx, instanceId)
}

func (r *instanceResponse) RebootInstance(ctx context.Context, accessToken, instanceId string) error {
	return getClient(accessToken).Instance.Reboot(ctx, instanceId)
}

func (r *instanceResponse) DeleteInstance(ctx context.Context, accessToken, instanceId string) error {
	return getClient(accessToken).Instance.Delete(ctx, instanceId)
}
