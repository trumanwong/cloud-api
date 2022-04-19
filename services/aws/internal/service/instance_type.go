package service

import (
	v1 "aws/api/instance/v1"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func (service *InstanceService) ListInstanceType(ctx context.Context, request *v1.ListInstanceTypeRequest) (*v1.ListInstanceTypeResponse, error) {
	var nextToken *string
	if len(request.NextToken) > 0 {
		nextToken = &request.NextToken
	}
	var maxResults *int32
	if request.MaxResults >= 5 && request.MaxResults <= 100 {
		maxResults = &request.MaxResults
	}
	result, err := service.itc.ListAll(ctx, request.AccessKeyId, request.SecretKey, request.Region, &ec2.DescribeInstanceTypesInput{
		MaxResults: maxResults,
		NextToken:  nextToken,
	})
	if err != nil {
		return nil, err
	}

	listInstanceTypeResponse := &v1.ListInstanceTypeResponse{
		InstanceTypes: make([]*v1.ListInstanceTypeResponse_InstanceType, len(result.InstanceTypes)),
	}
	for i, instanceType := range result.InstanceTypes {
		var inferenceAcceleratorInfo *v1.ListInstanceTypeResponse_InstanceType_InferenceAcceleratorInfo
		if instanceType.InferenceAcceleratorInfo != nil {
			accelerators := make([]*v1.ListInstanceTypeResponse_InstanceType_InferenceAcceleratorInfo_InferenceDeviceInfo, len(instanceType.InferenceAcceleratorInfo.Accelerators))
			for j, accelerator := range instanceType.InferenceAcceleratorInfo.Accelerators {
				accelerators[j] = &v1.ListInstanceTypeResponse_InstanceType_InferenceAcceleratorInfo_InferenceDeviceInfo{
					Count:        *accelerator.Count,
					Manufacturer: *accelerator.Manufacturer,
					Name:         *accelerator.Name,
				}
			}
			inferenceAcceleratorInfo = &v1.ListInstanceTypeResponse_InstanceType_InferenceAcceleratorInfo{
				Accelerators: accelerators,
			}
		}
		var instanceStorageInfo *v1.ListInstanceTypeResponse_InstanceType_InstanceStorageInfo
		if instanceType.InstanceStorageInfo != nil {
			disks := make([]*v1.ListInstanceTypeResponse_InstanceType_InstanceStorageInfo_DiskInfo, len(instanceType.InstanceStorageInfo.Disks))
			for i, disk := range instanceType.InstanceStorageInfo.Disks {
				disks[i] = &v1.ListInstanceTypeResponse_InstanceType_InstanceStorageInfo_DiskInfo{
					Count:    *disk.Count,
					SizeInGb: *disk.SizeInGB,
					Type:     string(disk.Type),
				}
			}
			instanceStorageInfo = &v1.ListInstanceTypeResponse_InstanceType_InstanceStorageInfo{
				Disks:             disks,
				EncryptionSupport: string(instanceType.InstanceStorageInfo.EncryptionSupport),
				NvmeSupport:       string(instanceType.InstanceStorageInfo.NvmeSupport),
				TotalSizeInGb:     *instanceType.InstanceStorageInfo.TotalSizeInGB,
			}
		}
		var memoryInfo *v1.ListInstanceTypeResponse_InstanceType_MemoryInfo
		if instanceType.MemoryInfo != nil {
			memoryInfo = &v1.ListInstanceTypeResponse_InstanceType_MemoryInfo{SizeInMib: *instanceType.MemoryInfo.SizeInMiB}
		}
		listInstanceTypeResponse.InstanceTypes[i] = &v1.ListInstanceTypeResponse_InstanceType{
			AutoRecoverySupported:         *instanceType.AutoRecoverySupported,
			BareMetal:                     *instanceType.BareMetal,
			BurstablePerformanceSupported: *instanceType.BurstablePerformanceSupported,
			CurrentGeneration:             *instanceType.CurrentGeneration,
			DedicatedHostsSupported:       *instanceType.DedicatedHostsSupported,
			HibernationSupported:          *instanceType.HibernationSupported,
			Hypervisor:                    string(instanceType.Hypervisor),
			InferenceAcceleratorInfo:      inferenceAcceleratorInfo,
			InstanceStorageInfo:           instanceStorageInfo,
			InstanceStorageSupported:      *instanceType.InstanceStorageSupported,
			InstanceType:                  string(instanceType.InstanceType),
			MemoryInfo:                    memoryInfo,
		}
	}
	return listInstanceTypeResponse, nil
}
