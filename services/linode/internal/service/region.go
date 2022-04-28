package service

import (
	"context"
	"github.com/linode/linodego"
	v1 "linode/api/instance/v1"
	"linode/pkg/util/convert"
)

func (service *InstanceService) ListRegions(ctx context.Context, request *v1.ListRegionsRequest) (*v1.ListRegionsResponse, error) {
	options := &linodego.ListOptions{
		PageSize: int(request.PageSize),
		PageOptions: &linodego.PageOptions{
			Page: int(request.Page),
		},
	}
	result, err := service.rc.ListRegions(ctx, request.AccessToken, options)
	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.ListRegionsResponse{Data: data}, nil
}
