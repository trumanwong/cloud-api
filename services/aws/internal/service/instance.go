package service

import (
	v1 "aws/api/instance/v1"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// CreateInstance Create Instance
func (service *InstanceService) CreateInstance(ctx context.Context, request *v1.CreateInstanceRequest) (*v1.CreateInstanceResponse, error) {
	result, err := service.uc.Create(ctx, request.AccessKeyId, request.SecretKey, request.Region, &ec2.RunInstancesInput{
		ImageId:      aws.String(request.ImageId),
		InstanceType: types.InstanceType(request.InstanceType),
		CpuOptions: &types.CpuOptionsRequest{
			CoreCount:      aws.Int32(request.CpuCoreCount),
			ThreadsPerCore: aws.Int32(request.ThreadsPerCore),
		},
		BlockDeviceMappings: []types.BlockDeviceMapping{
			{
				DeviceName: aws.String("/dev/sdh"),
				Ebs:        &types.EbsBlockDevice{VolumeSize: aws.Int32(request.SystemDiskSize)},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	instanceIdSets := make([]string, len(result.Instances))
	for i, v := range result.Instances {
		instanceIdSets[i] = *v.InstanceId
	}
	return &v1.CreateInstanceResponse{
		InstanceIdSets: instanceIdSets,
		OwnerId:        *result.OwnerId,
		RequesterId:    *result.RequesterId,
		ReservationId:  *result.ReservationId,
	}, nil
}

func (service *InstanceService) ListInstance(ctx context.Context, request *v1.ListInstanceRequest) (*v1.ListInstanceResponse, error) {
	var nextToken *string
	if len(request.NextToken) > 0 {
		nextToken = &request.NextToken
	}
	var maxResults *int32
	if request.MaxResults >= 5 && request.MaxResults <= 100 {
		maxResults = &request.MaxResults
	}
	result, err := service.uc.ListInstance(ctx, request.AccessKeyId, request.SecretKey, request.Region, &ec2.DescribeInstancesInput{
		MaxResults: maxResults,
		NextToken:  nextToken,
	})
	if err != nil {
		return nil, err
	}

	reservations := make([]*v1.ListInstanceResponse_Reservation, len(result.Reservations))
	for i, reservation := range result.Reservations {
		instancesSet := make([]*v1.InstanceObj, len(reservation.Instances))
		for j, v := range reservation.Instances {
			var cpuOptions *v1.InstanceObj_CpuOptions
			if v.CpuOptions != nil {
				cpuOptions = &v1.InstanceObj_CpuOptions{
					CoreCount:       *v.CpuOptions.CoreCount,
					ThreadsPerCount: *v.CpuOptions.ThreadsPerCore,
				}
			}
			var instanceStat *v1.InstanceState
			if v.State != nil {
				instanceStat = &v1.InstanceState{
					Code: *v.State.Code,
					Name: string(v.State.Name),
				}
			}
			var monitoring *v1.InstanceObj_Monitoring
			if v.Monitoring != nil {
				monitoring = &v1.InstanceObj_Monitoring{State: string(v.Monitoring.State)}
			}
			var privateDnsNameOptionsResponse *v1.InstanceObj_PrivateDnsNameOptionsResponse
			if v.PrivateDnsNameOptions != nil {
				privateDnsNameOptionsResponse = &v1.InstanceObj_PrivateDnsNameOptionsResponse{
					EnableResourceNameDnsAaaaRecord: *v.PrivateDnsNameOptions.EnableResourceNameDnsAAAARecord,
					EnableResourceNameDnsARecord:    *v.PrivateDnsNameOptions.EnableResourceNameDnsARecord,
					HostnameType:                    string(v.PrivateDnsNameOptions.HostnameType),
				}
			}
			instancesSet[j] = &v1.InstanceObj{
				AmiLaunchIndex:                *v.AmiLaunchIndex,
				Architecture:                  string(v.Architecture),
				BootMode:                      string(v.BootMode),
				CpuOptions:                    cpuOptions,
				EbsOptimized:                  *v.EbsOptimized,
				EnaSupport:                    *v.EnaSupport,
				Hypervisor:                    string(v.Hypervisor),
				ImageId:                       *v.ImageId,
				InstanceId:                    *v.InstanceId,
				InstanceLifecycle:             string(v.InstanceLifecycle),
				InstanceState:                 instanceStat,
				InstanceType:                  string(v.InstanceType),
				Ipv4Address:                   *v.PublicIpAddress,
				Ipv6Address:                   *v.Ipv6Address,
				KernelId:                      *v.KernelId,
				KeyName:                       *v.KeyName,
				LaunchTime:                    v.LaunchTime.String(),
				Monitoring:                    monitoring,
				Platform:                      string(v.Platform),
				PlatformDetails:               *v.PlatformDetails,
				PrivateDnsName:                *v.PrivateDnsName,
				PrivateDnsNameOptionsResponse: privateDnsNameOptionsResponse,
				PrivateIpAddress:              *v.PrivateIpAddress,
				RamdiskId:                     *v.RamdiskId,
				RootDeviceName:                *v.RootDeviceName,
				RootDeviceType:                string(v.RootDeviceType),
			}
		}
		groups := make([]*v1.GroupIdentifier, len(reservation.Groups))
		for j, v := range reservation.Groups {
			groups[j] = &v1.GroupIdentifier{
				GroupId:   *v.GroupId,
				GroupName: *v.GroupName,
			}
		}
		reservations[i] = &v1.ListInstanceResponse_Reservation{
			GroupSet:      groups,
			InstancesSet:  instancesSet,
			OwnerId:       *reservation.OwnerId,
			RequesterId:   *reservation.RequesterId,
			ReservationId: *reservation.ReservationId,
		}
	}
	listInstanceResponse := &v1.ListInstanceResponse{
		NextToken:      *result.NextToken,
		ReservationSet: reservations,
	}

	return listInstanceResponse, nil
}

func (service *InstanceService) StartInstance(ctx context.Context, request *v1.StartInstanceRequest) (*v1.StartInstanceResponse, error) {
	result, err := service.uc.StartInstance(ctx, request.AccessKeyId, request.SecretKey, request.Region, &ec2.StartInstancesInput{
		InstanceIds: request.InstanceIds,
	})
	if err != nil {
		return nil, err
	}

	instanceStateChanges := make([]*v1.InstanceStateChange, len(result.StartingInstances))
	for i, v := range result.StartingInstances {
		instanceStateChanges[i] = &v1.InstanceStateChange{
			CurrentState: &v1.InstanceState{
				Code: *v.CurrentState.Code,
				Name: string(v.CurrentState.Name),
			},
			InstanceId: *v.InstanceId,
			PreviousState: &v1.InstanceState{
				Code: *v.PreviousState.Code,
				Name: string(v.PreviousState.Name),
			},
		}
	}
	return &v1.StartInstanceResponse{
		InstanceStateChanges: instanceStateChanges,
	}, nil
}

func (service *InstanceService) StopInstance(ctx context.Context, request *v1.StopInstanceRequest) (*v1.StopInstanceResponse, error) {
	result, err := service.uc.StopInstance(ctx, request.AccessKeyId, request.SecretKey, request.Region, &ec2.StopInstancesInput{
		InstanceIds: request.InstanceIds,
	})
	if err != nil {
		return nil, err
	}
	instanceStateChanges := make([]*v1.InstanceStateChange, len(result.StoppingInstances))
	for i, v := range result.StoppingInstances {
		instanceStateChanges[i] = &v1.InstanceStateChange{
			CurrentState: &v1.InstanceState{
				Code: *v.CurrentState.Code,
				Name: string(v.CurrentState.Name),
			},
			InstanceId: *v.InstanceId,
			PreviousState: &v1.InstanceState{
				Code: *v.PreviousState.Code,
				Name: string(v.PreviousState.Name),
			},
		}
	}
	return &v1.StopInstanceResponse{
		InstanceStateChanges: instanceStateChanges,
	}, nil
}

func (service *InstanceService) RebootInstance(ctx context.Context, request *v1.RebootInstanceRequest) (*v1.RebootInstanceResponse, error) {
	_, err := service.uc.RebootInstance(ctx, request.AccessKeyId, request.SecretKey, request.Region, &ec2.RebootInstancesInput{
		InstanceIds: request.InstanceIds,
	})
	if err != nil {
		return nil, err
	}
	return &v1.RebootInstanceResponse{}, nil
}

func (service *InstanceService) DeleteInstance(ctx context.Context, request *v1.DeleteInstanceRequest) (*v1.DeleteInstanceResponse, error) {
	result, err := service.uc.DeleteInstance(ctx, request.AccessKeyId, request.SecretKey, request.Region, &ec2.TerminateInstancesInput{
		InstanceIds: request.InstanceIds,
	})
	if err != nil {
		return nil, err
	}
	instanceStateChanges := make([]*v1.InstanceStateChange, len(result.TerminatingInstances))
	for i, v := range result.TerminatingInstances {
		instanceStateChanges[i] = &v1.InstanceStateChange{
			CurrentState: &v1.InstanceState{
				Code: *v.CurrentState.Code,
				Name: string(v.CurrentState.Name),
			},
			InstanceId: *v.InstanceId,
			PreviousState: &v1.InstanceState{
				Code: *v.PreviousState.Code,
				Name: string(v.PreviousState.Name),
			},
		}
	}
	return &v1.DeleteInstanceResponse{
		InstanceStateChanges: instanceStateChanges,
	}, nil
}
