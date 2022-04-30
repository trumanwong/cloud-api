package service

import (
	"context"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	v1 "huawei/api/instance/v1"
	"huawei/pkg/util/convert"
)

// CreateInstances Create Instances
func (service *InstanceService) CreateInstances(ctx context.Context, request *v1.CreateInstancesRequest) (*v1.CreateInstancesResponse, error) {
	nics := make([]model.PrePaidServerNic, len(request.Nics))
	for i, v := range request.Nics {
		nics[i] = model.PrePaidServerNic{
			SubnetId:      v.SubnetId,
			IpAddress:     nil,
			Ipv6Enable:    nil,
			Ipv6Bandwidth: nil,
		}
	}
	var volumeType model.PrePaidServerRootVolumeVolumetype
	_ = volumeType.UnmarshalJSON([]byte(request.RootVolume.VolumeType))
	result, err := service.uc.CreateInstances(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.CreateServersRequest{
		Body: &model.CreateServersRequestBody{
			DryRun: request.DryRun,
			Server: &model.PrePaidServer{
				ImageRef:  request.ImageId,
				FlavorRef: request.InstanceType,
				Name:      request.Name,
				AdminPass: request.Password,
				Vpcid:     request.VpcId,
				Nics:      nics,
				Count:     request.Count,
				RootVolume: &model.PrePaidServerRootVolume{
					Volumetype: volumeType,
					Size:       request.RootVolume.Size,
				},
			},
		}})
	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.CreateInstancesResponse{Data: data}, nil
}

func (service *InstanceService) ListInstances(ctx context.Context, request *v1.ListInstancesRequest) (*v1.ListInstancesResponse, error) {
	result, err := service.uc.ListInstances(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.ListServersDetailsRequest{
		Limit:  request.Limit,
		Offset: request.Offset,
		Name:   request.Name,
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

func (service *InstanceService) StartInstances(ctx context.Context, request *v1.StartInstancesRequest) (*v1.StartInstancesResponse, error) {
	servers := make([]model.ServerId, len(request.Servers))
	for i, v := range request.Servers {
		servers[i] = model.ServerId{Id: v.Id}
	}
	result, err := service.uc.StartInstances(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.BatchStartServersRequest{
		Body: &model.BatchStartServersRequestBody{
			OsStart: &model.BatchStartServersOption{
				Servers: servers,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.StartInstancesResponse{Data: data}, nil
}

func (service *InstanceService) StopInstances(ctx context.Context, request *v1.StopInstancesRequest) (*v1.StopInstancesResponse, error) {
	servers := make([]model.ServerId, len(request.Servers))
	for i, v := range request.Servers {
		servers[i] = model.ServerId{Id: v.Id}
	}
	var stopType model.BatchStopServersOptionType
	_ = stopType.UnmarshalJSON([]byte(request.StopType))
	result, err := service.uc.StopInstances(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.BatchStopServersRequest{
		Body: &model.BatchStopServersRequestBody{
			OsStop: &model.BatchStopServersOption{
				Servers: servers,
				Type:    &stopType,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.StopInstancesResponse{Data: data}, nil
}

func (service *InstanceService) RebootInstances(ctx context.Context, request *v1.RebootInstancesRequest) (*v1.RebootInstancesResponse, error) {
	servers := make([]model.ServerId, len(request.Servers))
	for i, v := range request.Servers {
		servers[i] = model.ServerId{Id: v.Id}
	}
	var stopType model.BatchRebootSeversOptionType
	_ = stopType.UnmarshalJSON([]byte(request.RestartType))
	result, err := service.uc.RebootInstances(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.BatchRebootServersRequest{
		Body: &model.BatchRebootServersRequestBody{
			Reboot: &model.BatchRebootSeversOption{
				Servers: servers,
				Type:    stopType,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.RebootInstancesResponse{Data: data}, nil
}

func (service *InstanceService) DeleteInstances(ctx context.Context, request *v1.DeleteInstancesRequest) (*v1.DeleteInstancesResponse, error) {
	servers := make([]model.ServerId, len(request.Servers))
	for i, v := range request.Servers {
		servers[i] = model.ServerId{Id: v.Id}
	}
	result, err := service.uc.DeleteInstances(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.DeleteServersRequest{Body: &model.DeleteServersRequestBody{
		DeletePublicip: request.DeletePublicIp,
		DeleteVolume:   request.DeleteVolume,
		Servers:        servers,
	}})
	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteInstancesResponse{Data: data}, nil
}

func (service *InstanceService) ShowJob(ctx context.Context, request *v1.ShowJobRequest) (*v1.ShowJobResponse, error) {
	result, err := service.uc.ShowJob(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.ShowJobRequest{JobId: request.JobId})
	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.ShowJobResponse{Data: data}, nil
}
