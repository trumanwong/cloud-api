package service

import (
	"context"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	v1 "tencent/api/instance/v1"
	"tencent/pkg/util/convert"
)

func (s *InstanceService) CreateInstances(ctx context.Context, request *v1.CreateInstancesRequest) (*v1.CreateInstancesResponse, error) {
	runInstancesRequest := cvm.NewRunInstancesRequest()
	runInstancesRequest.LoginSettings = &cvm.LoginSettings{Password: request.Password}
	if request.SystemDisk != nil {
		runInstancesRequest.SystemDisk = &cvm.SystemDisk{
			DiskSize: request.SystemDisk.DiskSize,
			DiskType: request.SystemDisk.DiskType,
			CdcId:    request.SystemDisk.CdcId,
		}
	}
	runInstancesRequest.InstanceType = request.InstanceType
	runInstancesRequest.ImageId = request.ImageId
	runInstancesRequest.InstanceCount = request.InstanceCount
	runInstancesRequest.InstanceName = request.InstanceName
	runInstancesRequest.DryRun = request.DryRun
	result, err := s.uc.CreateInstances(ctx, request.SecretId, request.SecretKey, request.Region, runInstancesRequest)
	if err != nil {
		return nil, err
	}

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.CreateInstancesResponse{Data: data}, nil
}

func (s *InstanceService) ListInstances(ctx context.Context, request *v1.ListInstancesRequest) (*v1.ListInstancesResponse, error) {
	describeInstancesRequest := cvm.NewDescribeInstancesRequest()
	describeInstancesRequest.Offset = request.Offset
	describeInstancesRequest.Limit = request.Limit
	describeInstancesRequest.Filters = make([]*cvm.Filter, 0)
	if request.InstanceName != nil {
		filter := &cvm.Filter{
			Name:   common.StringPtr("instance-name"),
			Values: common.StringPtrs([]string{*request.InstanceName}),
		}
		describeInstancesRequest.Filters = append(describeInstancesRequest.Filters, filter)
	}
	result, err := s.uc.ListInstances(ctx, request.SecretId, request.SecretKey, request.Region, describeInstancesRequest)
	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.ListInstancesResponse{Data: data}, nil
}

func (s *InstanceService) StartInstances(ctx context.Context, request *v1.StartInstancesRequest) (*v1.StartInstancesResponse, error) {
	startInstanceRequest := cvm.NewStartInstancesRequest()
	startInstanceRequest.InstanceIds = common.StringPtrs(request.InstanceIds)
	result, err := s.uc.StartInstances(ctx, request.SecretId, request.SecretKey, request.Region, startInstanceRequest)
	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.StartInstancesResponse{
		Data: data,
	}, nil
}

func (s *InstanceService) StopInstances(ctx context.Context, request *v1.StopInstancesRequest) (*v1.StopInstancesResponse, error) {
	stopInstanceRequest := cvm.NewStopInstancesRequest()
	stopInstanceRequest.InstanceIds = common.StringPtrs(request.InstanceIds)
	stopInstanceRequest.StopType = request.StopType
	stopInstanceRequest.StoppedMode = request.StoppedMode
	result, err := s.uc.StopInstances(ctx, request.SecretId, request.SecretKey, request.Region, stopInstanceRequest)
	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.StopInstancesResponse{Data: data}, nil
}

func (s *InstanceService) RebootInstances(ctx context.Context, request *v1.RebootInstancesRequest) (*v1.RebootInstancesResponse, error) {
	rebootInstanceRequest := cvm.NewRebootInstancesRequest()
	rebootInstanceRequest.InstanceIds = common.StringPtrs(request.InstanceIds)
	rebootInstanceRequest.StopType = request.StopType
	result, err := s.uc.RebootInstances(ctx, request.SecretId, request.SecretKey, request.Region, rebootInstanceRequest)
	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.RebootInstancesResponse{Data: data}, nil
}

func (s *InstanceService) DeleteInstances(ctx context.Context, request *v1.DeleteInstancesRequest) (*v1.DeleteInstancesResponse, error) {
	terminateInstanceRequest := cvm.NewTerminateInstancesRequest()
	terminateInstanceRequest.InstanceIds = common.StringPtrs(request.InstanceIds)
	result, err := s.uc.DeleteInstances(ctx, request.SecretId, request.SecretKey, request.Region, terminateInstanceRequest)
	if err != nil {
		return nil, err
	}
	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteInstancesResponse{Data: data}, nil
}
