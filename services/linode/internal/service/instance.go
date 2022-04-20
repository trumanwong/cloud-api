package service

import (
	"context"
	"github.com/vultr/govultr/v2"
	v1 "vultr/api/instance/v1"
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

	response := &v1.CreateInstanceResponse{}
	if result != nil {
		response.Instance = &v1.InstanceObj{
			Id:               result.ID,
			Os:               result.Os,
			Ram:              int32(result.RAM),
			Disk:             int32(result.Disk),
			MainIp:           result.MainIP,
			VcpuCount:        int32(result.VCPUCount),
			Region:           result.Region,
			DefaultPassword:  result.DefaultPassword,
			DateCreated:      result.DateCreated,
			Status:           result.Status,
			PowerStatus:      result.PowerStatus,
			ServerStatus:     result.ServerStatus,
			AllowedBandwidth: int32(result.AllowedBandwidth),
			NetmaskV4:        result.NetmaskV4,
			GatewayV4:        result.GatewayV4,
			V6Network:        result.V6Network,
			V6MainIp:         result.V6MainIP,
			V6NetworkSize:    int32(result.V6NetworkSize),
			Hostname:         result.Hostname,
			Label:            result.Label,
			Tag:              result.Tag,
			InternalIp:       result.InternalIP,
			Kvm:              result.KVM,
			OsId:             int32(result.OsID),
			AppId:            int32(result.AppID),
			ImageId:          result.ImageID,
			FirewallGroupId:  result.FirewallGroupID,
			Features:         result.Features,
			Plan:             result.Plan,
		}
	}
	return response, nil
}

func (service *InstanceService) ListInstances(ctx context.Context, request *v1.ListInstancesRequest) (*v1.ListInstancesResponse, error) {
	result, meta, err := service.uc.ListInstances(ctx, request.AccessToken, &govultr.ListOptions{
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

	var link *v1.Meta_Link
	if meta.Links != nil {
		link = &v1.Meta_Link{
			Next: meta.Links.Next,
			Prev: meta.Links.Prev,
		}
	}
	response := &v1.ListInstancesResponse{
		Instances: make([]*v1.InstanceObj, len(result)),
		Meta: &v1.Meta{
			Total: int32(meta.Total),
			Link:  link,
		},
	}
	for i, v := range result {
		response.Instances[i] = &v1.InstanceObj{
			Id:               v.ID,
			Os:               v.Os,
			Ram:              int32(v.RAM),
			Disk:             int32(v.Disk),
			MainIp:           v.MainIP,
			VcpuCount:        int32(v.VCPUCount),
			Region:           v.Region,
			DefaultPassword:  v.DefaultPassword,
			DateCreated:      v.DateCreated,
			Status:           v.Status,
			PowerStatus:      v.PowerStatus,
			ServerStatus:     v.ServerStatus,
			AllowedBandwidth: int32(v.AllowedBandwidth),
			NetmaskV4:        v.NetmaskV4,
			GatewayV4:        v.GatewayV4,
			V6Network:        v.V6Network,
			V6MainIp:         v.V6MainIP,
			V6NetworkSize:    int32(v.V6NetworkSize),
			Hostname:         v.Hostname,
			Label:            v.Label,
			Tag:              v.Tag,
			InternalIp:       v.InternalIP,
			Kvm:              v.KVM,
			OsId:             int32(v.OsID),
			AppId:            int32(v.AppID),
			ImageId:          v.ImageID,
			FirewallGroupId:  v.FirewallGroupID,
			Features:         v.Features,
			Plan:             v.Plan,
		}
	}
	return response, nil
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
