package service

import (
	"context"
	"github.com/vultr/govultr/v2"
	v1 "vultr/api/instance/v1"
	"vultr/pkg/util/convert"
)

// CreateInstance Create Instance
func (service *InstanceService) CreateInstance(ctx context.Context, request *v1.CreateInstanceRequest) (*v1.CreateInstanceResponse, error) {
	result, err := service.uc.CreateInstance(ctx, request.AccessToken, &govultr.InstanceCreateReq{
		Region:   request.Region,
		Plan:     request.Plan,
		ImageID:  request.ImageId,
		Hostname: request.Hostname,
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
	result, err := service.uc.ListInstances(ctx, request.AccessToken, &govultr.ListOptions{
		PerPage: int(request.PerPage),
		Cursor:  request.Cursor,
		MainIP:  request.MainIp,
		Label:   request.Label,
		Tag:     request.Tag,
		Region:  request.Region,
	})
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
	err := service.uc.StartInstance(ctx, request.AccessToken, request.InstanceId)
	if err != nil {
		return nil, err
	}
	return &v1.StartInstanceResponse{}, nil
}

func (service *InstanceService) StopInstance(ctx context.Context, request *v1.StopInstanceRequest) (*v1.StopInstanceResponse, error) {
	err := service.uc.StopInstance(ctx, request.AccessToken, request.InstanceId)
	if err != nil {
		return nil, err
	}
	return &v1.StopInstanceResponse{}, nil
}

func (service *InstanceService) RebootInstance(ctx context.Context, request *v1.RebootInstanceRequest) (*v1.RebootInstanceResponse, error) {
	err := service.uc.RebootInstance(ctx, request.AccessToken, request.InstanceId)
	if err != nil {
		return nil, err
	}
	return &v1.RebootInstanceResponse{}, nil
}

func (service *InstanceService) DeleteInstance(ctx context.Context, request *v1.DeleteInstanceRequest) (*v1.DeleteInstanceResponse, error) {
	err := service.uc.DeleteInstance(ctx, request.AccessToken, request.InstanceId)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteInstanceResponse{}, nil
}
