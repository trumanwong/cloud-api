package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/linode/linodego"
	"linode/internal/biz"
)

type instanceTypeResponse struct {
	data *Data
	log  *log.Helper
}

// NewInstanceTypeRepo .
func NewInstanceTypeRepo(data *Data, logger log.Logger) biz.PlanResponse {
	return &instanceTypeResponse{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *instanceTypeResponse) ListInstanceTypes(ctx context.Context, accessToken string, request *linodego.ListOptions) ([]linodego.LinodeType, error) {
	client := newClient(accessToken)
	result, err := client.ListTypes(ctx, request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
