package service

import (
	"context"
	"github.com/linode/linodego"
	v1 "linode/api/instance/v1"
	"linode/pkg/util/convert"
)

func (service *InstanceService) ListInstanceTypes(ctx context.Context, request *v1.ListInstanceTypesRequest) (*v1.ListInstanceTypesResponse, error) {
	options := &linodego.ListOptions{
		PageSize: int(request.PageSize),
		PageOptions: &linodego.PageOptions{
			Page: int(request.Page),
		},
	}

	result, err := service.itc.ListInstanceTypes(ctx, request.AccessToken, options)
	if err != nil {
		return nil, err
	}

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.ListInstanceTypesResponse{Data: data}, nil
}
