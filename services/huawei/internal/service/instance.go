package service

import (
	"context"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	v1 "huawei/api/instance/v1"
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
			DryRun: &request.DryRun,
			Server: &model.PrePaidServer{
				ImageRef:  request.ImageId,
				FlavorRef: request.InstanceType,
				Name:      request.Name,
				AdminPass: &request.Password,
				Vpcid:     request.VpcId,
				Nics:      nics,
				Count:     &request.Amount,
				RootVolume: &model.PrePaidServerRootVolume{
					Volumetype: volumeType,
					Size:       &request.RootVolume.Size,
				},
			},
		}})
	if err != nil {
		return nil, err
	}
	return &v1.CreateInstancesResponse{
		JobId:     *result.JobId,
		OrderId:   *result.OrderId,
		ServerIds: *result.ServerIds,
	}, nil
}

func (service *InstanceService) ListInstances(ctx context.Context, request *v1.ListInstancesRequest) (*v1.ListInstancesResponse, error) {
	var name *string
	if len(request.Name) > 0 {
		name = &request.Name
	}
	result, err := service.uc.ListInstances(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.ListServersDetailsRequest{
		Limit:  &request.Limit,
		Offset: &request.Offset,
		Name:   name,
	})
	if err != nil {
		return nil, err
	}

	instances := make([]*v1.ListInstancesResponse_Instance, 0)
	if result.Servers != nil {
		instances = make([]*v1.ListInstancesResponse_Instance, len(*result.Servers))
		for i, instance := range *result.Servers {
			var addresses map[string]*v1.ListInstancesResponse_Instance_ServerAddresses
			if instance.Addresses != nil {
				for k, arr := range instance.Addresses {
					serverAddresses := make([]*v1.ListInstancesResponse_Instance_ServerAddress, len(arr))
					for j, item := range arr {
						osExtIpsType, _ := item.OSEXTIPStype.MarshalJSON()
						serverAddresses[j] = &v1.ListInstancesResponse_Instance_ServerAddress{
							Version:            item.Version,
							Addr:               item.Addr,
							OsExtIpsType:       string(osExtIpsType),
							OsExtIpsMacMacAddr: *item.OSEXTIPSMACmacAddr,
							OsExtIpSportId:     *item.OSEXTIPSportId,
						}
					}
					addresses[k] = &v1.ListInstancesResponse_Instance_ServerAddresses{ServerAddresses: serverAddresses}
				}
			}
			var instanceType *v1.ListInstancesResponse_Instance_InstanceType
			if instance.Flavor != nil {
				instanceType = &v1.ListInstancesResponse_Instance_InstanceType{
					Id:    instance.Flavor.Id,
					Name:  instance.Flavor.Name,
					Disk:  instance.Flavor.Disk,
					Vcpus: instance.Flavor.Vcpus,
					Ram:   instance.Flavor.Ram,
				}
			}
			var securityGroups []*v1.ListInstancesResponse_Instance_SecurityGroup
			if instance.SecurityGroups != nil {
				securityGroups = make([]*v1.ListInstancesResponse_Instance_SecurityGroup, len(instance.SecurityGroups))
				for j, v := range instance.SecurityGroups {
					securityGroups[j] = &v1.ListInstancesResponse_Instance_SecurityGroup{
						Name: v.Name,
						Id:   v.Id,
					}
				}
			}
			var fault *v1.ListInstancesResponse_Instance_ServerFault
			if instance.Fault != nil {
				fault = &v1.ListInstancesResponse_Instance_ServerFault{
					Code:    *instance.Fault.Code,
					Created: *instance.Fault.Created,
					Message: *instance.Fault.Message,
					Details: *instance.Fault.Details,
				}
			}
			var osExtendedVolumesVolumesAttached []*v1.ListInstancesResponse_Instance_ServerExtendVolumeAttachment
			if instance.OsExtendedVolumesvolumesAttached != nil {
				osExtendedVolumesVolumesAttached = make([]*v1.ListInstancesResponse_Instance_ServerExtendVolumeAttachment, len(instance.OsExtendedVolumesvolumesAttached))
				for j, v := range instance.OsExtendedVolumesvolumesAttached {
					osExtendedVolumesVolumesAttached[j] = &v1.ListInstancesResponse_Instance_ServerExtendVolumeAttachment{
						Id:                  v.Id,
						DeleteOnTermination: v.DeleteOnTermination,
						BootIndex:           *v.BootIndex,
						Device:              v.Device,
					}
				}
			}
			var osSchedulerHints *v1.ListInstancesResponse_Instance_ServerSchedulerHints
			if instance.OsschedulerHints != nil {
				osSchedulerHints = &v1.ListInstancesResponse_Instance_ServerSchedulerHints{
					Group:           *instance.OsschedulerHints.Group,
					Tenancy:         *instance.OsschedulerHints.Tenancy,
					DedicatedHostId: *instance.OsschedulerHints.DedicatedHostId,
				}
			}
			var sysTags []*v1.ListInstancesResponse_Instance_ServerSystemTag
			if instance.SysTags != nil {
				sysTags = make([]*v1.ListInstancesResponse_Instance_ServerSystemTag, len(*instance.SysTags))
				for j, v := range *instance.SysTags {
					sysTags[j] = &v1.ListInstancesResponse_Instance_ServerSystemTag{
						Key:   *v.Key,
						Value: *v.Value,
					}
				}
			}
			var cpuOptions *v1.ListInstancesResponse_Instance_CpuOptions
			if instance.CpuOptions != nil {
				cpuOptions = &v1.ListInstancesResponse_Instance_CpuOptions{HwCpuThreads: *instance.CpuOptions.HwcpuThreads}
			}
			var hypervisor *v1.ListInstancesResponse_Instance_Hypervisor
			if instance.Hypervisor != nil {
				hypervisor = &v1.ListInstancesResponse_Instance_Hypervisor{
					HypervisorType: *instance.Hypervisor.HypervisorType,
					CsdHypervisor:  *instance.Hypervisor.CsdHypervisor,
				}
			}
			instances[i] = &v1.ListInstancesResponse_Instance{
				Status:                            instance.Status,
				Updated:                           instance.Updated,
				AutoTerminateTime:                 instance.AutoTerminateTime,
				HostId:                            instance.HostId,
				OsExtSrvAttrHost:                  instance.OSEXTSRVATTRhost,
				Addresses:                         addresses,
				KeyName:                           instance.KeyName,
				ImageId:                           instance.Image.Id,
				OsExtStsTaskState:                 instance.OSEXTSTStaskState,
				OsExtStsVmState:                   instance.OSEXTSTSvmState,
				OsExtSrvAttrInstanceName:          instance.OSEXTSRVATTRinstanceName,
				OsExtSrvAttrHypervisorHostname:    instance.OSEXTSRVATTRhypervisorHostname,
				InstanceType:                      instanceType,
				Id:                                instance.Id,
				SecurityGroups:                    securityGroups,
				OsExtAzAvailabilityZone:           instance.OSEXTAZavailabilityZone,
				UserId:                            instance.UserId,
				Name:                              instance.Name,
				Created:                           instance.Created,
				TenantId:                          instance.TenantId,
				OsDcfDiskConfig:                   *instance.OSDCFdiskConfig,
				AccessIpv4:                        instance.AccessIPv4,
				AccessIpv6:                        instance.AccessIPv6,
				Fault:                             fault,
				Progress:                          *instance.Progress,
				OsExtStsPowerState:                instance.OSEXTSTSpowerState,
				ConfigDrive:                       instance.ConfigDrive,
				Metadata:                          instance.Metadata,
				OsSrvUsgLaunchedAt:                instance.OSSRVUSGlaunchedAt,
				OsSrvUsgTerminatedAt:              instance.OSSRVUSGterminatedAt,
				OsExtendedVolumesVolumes_Attached: osExtendedVolumesVolumesAttached,
				Description:                       *instance.Description,
				HostStatus:                        instance.HostStatus,
				OsExtSrvAttrHostname:              instance.OSEXTSRVATTRhostname,
				OsExtSrvAttrReservationId:         *instance.OSEXTSRVATTRreservationId,
				OsExtSrvAttrLaunchIndex:           instance.OSEXTSRVATTRlaunchIndex,
				OsExtSrvAttrKernelId:              instance.OSEXTSRVATTRkernelId,
				OsExtSrvAttrRamdiskId:             instance.OSEXTSRVATTRramdiskId,
				OsExtSrvAttrRootDeviceName:        instance.OSEXTSRVATTRrootDeviceName,
				OsExtSrvAttrUserData:              *instance.OSEXTSRVATTRuserData,
				Locked:                            instance.Locked,
				Tags:                              *instance.Tags,
				OsSchedulerHints:                  osSchedulerHints,
				EnterpriseProjectId:               *instance.EnterpriseProjectId,
				SysTags:                           sysTags,
				CpuOptions:                        cpuOptions,
				Hypervisor:                        hypervisor,
			}
		}
	}
	listInstancesResponse := &v1.ListInstancesResponse{
		Count:     *result.Count,
		Instances: instances,
	}

	return listInstancesResponse, nil
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
	return &v1.StartInstancesResponse{
		JobId: *result.JobId,
	}, nil
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
	return &v1.StopInstancesResponse{
		JobId: *result.JobId,
	}, nil
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
	return &v1.RebootInstancesResponse{
		JobId: *result.JobId,
	}, nil
}

func (service *InstanceService) DeleteInstances(ctx context.Context, request *v1.DeleteInstancesRequest) (*v1.DeleteInstancesResponse, error) {
	servers := make([]model.ServerId, len(request.Servers))
	for i, v := range request.Servers {
		servers[i] = model.ServerId{Id: v.Id}
	}
	var deletePubicIp *bool
	if request.DeletePublicIp != false {
		deletePubicIp = &request.DeletePublicIp
	}
	var deleteVolume *bool
	if request.DeleteVolume != false {
		deleteVolume = &request.DeleteVolume
	}
	result, err := service.uc.DeleteInstances(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.DeleteServersRequest{Body: &model.DeleteServersRequestBody{
		DeletePublicip: deletePubicIp,
		DeleteVolume:   deleteVolume,
		Servers:        servers,
	}})
	if err != nil {
		return nil, err
	}
	return &v1.DeleteInstancesResponse{
		JobId: *result.JobId,
	}, nil
}

func (service *InstanceService) ShowJob(ctx context.Context, request *v1.ShowJobRequest) (*v1.ShowJobResponse, error) {
	result, err := service.uc.ShowJob(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.ShowJobRequest{JobId: request.JobId})
	if err != nil {
		return nil, err
	}
	status, _ := result.Status.MarshalJSON()
	return &v1.ShowJobResponse{
		BeginTime:  *result.BeginTime,
		Code:       *result.Code,
		EndTime:    *result.EndTime,
		ErrorCode:  *result.ErrorCode,
		FailReason: *result.FailReason,
		JobId:      *result.JobId,
		JobType:    *result.JobType,
		Message:    *result.Message,
		Status:     string(status),
	}, nil
}
