package service

import (
	v1 "ali/api/instance/v1"
	"ali/pkg/util/convert"
	"context"
	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

// CreateInstances Create Instances
func (service *InstanceService) CreateInstances(ctx context.Context, request *v1.CreateInstancesRequest) (*v1.CreateInstancesResponse, error) {
	result, err := service.uc.CreateInstances(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, &ecs20140526.RunInstancesRequest{
		RegionId:     request.RegionId,
		ImageId:      request.ImageId,
		InstanceName: request.InstanceName,
		InstanceType: request.InstanceType,
		SystemDisk: &ecs20140526.RunInstancesRequestSystemDisk{
			Size: request.SystemDiskSize,
		},
		UniqueSuffix: request.UniqueSuffix,
		Amount:       request.Amount,
		Password:     request.Password,
		DryRun:       request.DryRun,
	})
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
	result, err := service.uc.ListInstances(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, &ecs20140526.DescribeInstancesRequest{
		RegionId:     request.RegionId,
		PageNumber:   request.PageNumber,
		PageSize:     request.PageSize,
		NextToken:    request.NextToken,
		InstanceName: request.InstanceName,
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
	instanceIds := make([]*string, len(request.InstanceIds))
	for i, v := range request.InstanceIds {
		instanceIds[i] = tea.String(v)
	}
	result, err := service.uc.StartInstances(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, &ecs20140526.StartInstancesRequest{
		RegionId:          request.RegionId,
		InstanceId:        instanceIds,
		DryRun:            request.DryRun,
		BatchOptimization: request.BatchOptimization,
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
	instanceIds := make([]*string, len(request.InstanceIds))
	for i, v := range request.InstanceIds {
		instanceIds[i] = tea.String(v)
	}
	result, err := service.uc.StopInstances(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, &ecs20140526.StopInstancesRequest{
		RegionId:          request.RegionId,
		InstanceId:        instanceIds,
		DryRun:            request.DryRun,
		ForceStop:         request.ForceStop,
		BatchOptimization: request.BatchOptimization,
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
	instanceIds := make([]*string, len(request.InstanceIds))
	for i, v := range request.InstanceIds {
		instanceIds[i] = tea.String(v)
	}
	result, err := service.uc.RebootInstances(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, &ecs20140526.RebootInstancesRequest{
		RegionId:          request.RegionId,
		InstanceId:        instanceIds,
		DryRun:            request.DryRun,
		ForceReboot:       request.ForceReboot,
		BatchOptimization: request.BatchOptimization,
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
	instanceIds := make([]*string, len(request.InstanceIds))
	for i, v := range request.InstanceIds {
		instanceIds[i] = tea.String(v)
	}
	result, err := service.uc.DeleteInstances(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, &ecs20140526.DeleteInstancesRequest{
		InstanceId: instanceIds,
		DryRun:     request.DryRun,
	})
	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteInstancesResponse{Data: data}, nil
}
