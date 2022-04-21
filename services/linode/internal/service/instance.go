package service

import (
	"context"
	"github.com/linode/linodego"
	v1 "linode/api/instance/v1"
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

	response := &v1.CreateInstanceResponse{}
	if result != nil {
		var alerts *v1.InstanceObj_Alert
		if result.Alerts != nil {
			alerts = &v1.InstanceObj_Alert{
				Cpu:           int32(result.Alerts.CPU),
				Io:            int32(result.Alerts.CPU),
				NetworkIn:     int32(result.Alerts.NetworkIn),
				NetworkOut:    int32(result.Alerts.NetworkOut),
				TransferQuota: int32(result.Alerts.TransferQuota),
			}
		}
		var backups *v1.InstanceObj_Backup
		if result.Backups != nil {
			backups = &v1.InstanceObj_Backup{
				Enabled:        result.Backups.Enabled,
				LastSuccessful: "",
				Schedule: &v1.InstanceObj_Backup_Schedule{
					Day:    result.Backups.Schedule.Day,
					Window: result.Backups.Schedule.Window,
				},
			}
		}
		ipv4 := make([]string, len(result.IPv4))
		for i, v := range result.IPv4 {
			ipv4[i] = v.String()
		}
		var specs *v1.InstanceObj_Specs
		if result.Specs != nil {
			specs = &v1.InstanceObj_Specs{
				Disk:     int32(result.Specs.Disk),
				Memory:   int32(result.Specs.Memory),
				Transfer: int32(result.Specs.Transfer),
				Vcpus:    int32(result.Specs.VCPUs),
			}
		}
		response.Instance = &v1.InstanceObj{
			Alters:          alerts,
			Backups:         backups,
			Created:         result.Created.Format("2006-01-02 15:04:05"),
			Hypervisor:      result.Hypervisor,
			Id:              int32(result.ID),
			Image:           result.Image,
			Ipv4:            ipv4,
			Ipv6:            result.IPv6,
			Label:           result.Label,
			Region:          result.Region,
			Specs:           specs,
			Status:          string(result.Status),
			Tags:            result.Tags,
			InstanceType:    result.Type,
			Updated:         result.Updated.Format("2006-01-02 15:04:05"),
			WatchdogEnabled: result.WatchdogEnabled,
		}
	}
	return response, nil
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

	response := &v1.ListInstancesResponse{
		Instances: make([]*v1.InstanceObj, len(result)),
		Page:      int32(options.Page),
		Pages:     int32(options.Pages),
		Results:   int32(options.Results),
	}
	for i, instance := range result {
		var alerts *v1.InstanceObj_Alert
		if instance.Alerts != nil {
			alerts = &v1.InstanceObj_Alert{
				Cpu:           int32(instance.Alerts.CPU),
				Io:            int32(instance.Alerts.CPU),
				NetworkIn:     int32(instance.Alerts.NetworkIn),
				NetworkOut:    int32(instance.Alerts.NetworkOut),
				TransferQuota: int32(instance.Alerts.TransferQuota),
			}
		}
		var backups *v1.InstanceObj_Backup
		if instance.Backups != nil {
			backups = &v1.InstanceObj_Backup{
				Enabled:        instance.Backups.Enabled,
				LastSuccessful: "",
				Schedule: &v1.InstanceObj_Backup_Schedule{
					Day:    instance.Backups.Schedule.Day,
					Window: instance.Backups.Schedule.Window,
				},
			}
		}
		ipv4 := make([]string, len(instance.IPv4))
		for i, v := range instance.IPv4 {
			ipv4[i] = v.String()
		}
		var specs *v1.InstanceObj_Specs
		if instance.Specs != nil {
			specs = &v1.InstanceObj_Specs{
				Disk:     int32(instance.Specs.Disk),
				Memory:   int32(instance.Specs.Memory),
				Transfer: int32(instance.Specs.Transfer),
				Vcpus:    int32(instance.Specs.VCPUs),
			}
		}
		response.Instances[i] = &v1.InstanceObj{
			Alters:          alerts,
			Backups:         backups,
			Created:         instance.Created.Format("2006-01-02 15:04:05"),
			Hypervisor:      instance.Hypervisor,
			Id:              int32(instance.ID),
			Image:           instance.Image,
			Ipv4:            ipv4,
			Ipv6:            instance.IPv6,
			Label:           instance.Label,
			Region:          instance.Region,
			Specs:           specs,
			Status:          string(instance.Status),
			Tags:            instance.Tags,
			InstanceType:    instance.Type,
			Updated:         instance.Updated.Format("2006-01-02 15:04:05"),
			WatchdogEnabled: instance.WatchdogEnabled,
		}
	}
	return response, nil
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
