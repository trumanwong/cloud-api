package service

import (
	"context"
	"github.com/linode/linodego"
	v1 "linode/api/instance/v1"
	"linode/pkg/util/convert"
)

// CreateInstance Create Instance
func (service *InstanceService) CreateInstance(ctx context.Context, request *v1.CreateInstanceRequest) (*v1.CreateInstanceResponse, error) {
	result, err := service.uc.CreateInstance(ctx, request.AccessToken, linodego.InstanceCreateOptions{
		Region:   request.Region,
		Type:     request.InstanceType,
		Label:    request.Label,
		RootPass: request.Password,
		Image:    request.Image,
	})
	if err != nil {
		return nil, err
	}

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.CreateInstanceResponse{Data: data}, nil
}

func (service *InstanceService) ListInstances(ctx context.Context, request *v1.ListInstancesRequest) (*v1.ListInstancesResponse, error) {
	options := &linodego.ListOptions{
		PageSize: int(request.PageSize),
		PageOptions: &linodego.PageOptions{
			Page: int(request.Page),
		},
	}
	result, err := service.uc.ListInstances(ctx, request.AccessToken, options)
	if err != nil {
		return nil, err
	}

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.ListInstancesResponse{Data: data}, nil
}

func (service *InstanceService) StartInstance(ctx context.Context, request *v1.StartInstanceRequest) (*v1.StartInstanceResponse, error) {
	err := service.uc.StartInstance(ctx, request.AccessToken, int(request.InstanceId), int(request.ConfigId))
	if err != nil {
		return nil, err
	}
	return &v1.StartInstanceResponse{}, nil
}

func (service *InstanceService) StopInstance(ctx context.Context, request *v1.StopInstanceRequest) (*v1.StopInstanceResponse, error) {
	err := service.uc.StopInstance(ctx, request.AccessToken, int(request.InstanceId))
	if err != nil {
		return nil, err
	}
	return &v1.StopInstanceResponse{}, nil
}

func (service *InstanceService) RebootInstance(ctx context.Context, request *v1.RebootInstanceRequest) (*v1.RebootInstanceResponse, error) {
	err := service.uc.RebootInstance(ctx, request.AccessToken, int(request.InstanceId), int(request.ConfigId))
	if err != nil {
		return nil, err
	}
	return &v1.RebootInstanceResponse{}, nil
}

func (service *InstanceService) DeleteInstance(ctx context.Context, request *v1.DeleteInstanceRequest) (*v1.DeleteInstanceResponse, error) {
	err := service.uc.DeleteInstance(ctx, request.AccessToken, int(request.InstanceId))
	if err != nil {
		return nil, err
	}
	return &v1.DeleteInstanceResponse{}, nil
}
