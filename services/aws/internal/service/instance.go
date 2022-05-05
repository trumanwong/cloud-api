package service

import (
	v1 "aws/api/instance/v1"
	"aws/pkg/util/convert"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// CreateInstances Create Instances
func (service *InstanceService) CreateInstances(ctx context.Context, request *v1.CreateInstancesRequest) (*v1.CreateInstancesResponse, error) {
	var blockDeviceMappings []types.BlockDeviceMapping
	if request.BlockDeviceMapping != nil {
		blockDeviceMappings = make([]types.BlockDeviceMapping, len(request.BlockDeviceMapping))
		for i, v := range request.BlockDeviceMapping {
			var ebs *types.EbsBlockDevice
			if v.Ebs != nil {
				ebs = &types.EbsBlockDevice{
					DeleteOnTermination: v.Ebs.DeleteOnTermination,
					Encrypted:           v.Ebs.Encrypted,
					Iops:                v.Ebs.Iops,
					KmsKeyId:            v.Ebs.KmsKeyId,
					OutpostArn:          v.Ebs.OutpostArn,
					SnapshotId:          v.Ebs.SnapshotId,
					Throughput:          v.Ebs.Throughput,
					VolumeSize:          v.Ebs.VolumeSize,
					VolumeType:          types.VolumeType(v.Ebs.VolumeType),
				}
			}
			blockDeviceMappings[i] = types.BlockDeviceMapping{
				DeviceName:  v.DeviceName,
				Ebs:         ebs,
				NoDevice:    v.NoDevice,
				VirtualName: v.VirtualName,
			}
		}
	}
	result, err := service.uc.CreateInstances(ctx, request.AccessKeyId, request.SecretKey, request.Region, &ec2.RunInstancesInput{
		ImageId:             request.ImageId,
		InstanceType:        types.InstanceType(request.InstanceType),
		BlockDeviceMappings: blockDeviceMappings,
		MinCount:            request.MinCount,
		MaxCount:            request.MaxCount,
		DryRun:              request.DryRun,
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
	result, err := service.uc.ListInstances(ctx, request.AccessKeyId, request.SecretKey, request.Region, &ec2.DescribeInstancesInput{
		MaxResults: request.MaxResults,
		NextToken:  request.NextToken,
		DryRun:     request.DryRun,
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
	result, err := service.uc.StartInstances(ctx, request.AccessKeyId, request.SecretKey, request.Region, &ec2.StartInstancesInput{
		InstanceIds: request.InstanceIds,
		DryRun:      request.DryRun,
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
	result, err := service.uc.StopInstances(ctx, request.AccessKeyId, request.SecretKey, request.Region, &ec2.StopInstancesInput{
		InstanceIds: request.InstanceIds,
		DryRun:      request.DryRun,
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

func (service *InstanceService) RebootInstance(ctx context.Context, request *v1.RebootInstancesRequest) (*v1.RebootInstancesResponse, error) {
	result, err := service.uc.RebootInstances(ctx, request.AccessKeyId, request.SecretKey, request.Region, &ec2.RebootInstancesInput{
		InstanceIds: request.InstanceIds,
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
	result, err := service.uc.DeleteInstances(ctx, request.AccessKeyId, request.SecretKey, request.Region, &ec2.TerminateInstancesInput{
		InstanceIds: request.InstanceIds,
		DryRun:      request.DryRun,
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
