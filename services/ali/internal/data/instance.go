package data

import (
	"ali/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
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

func (r *instanceResponse) Create(ctx context.Context, request *biz.CreateInstanceRequest) (*biz.CreateInstanceResponse, error) {
	return nil, nil
}

func (r *instanceResponse) ListAll(context.Context, *biz.Instance) ([]*biz.Instance, error) {
	return nil, nil
}
