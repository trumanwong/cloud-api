package service

import (
	v1 "ali/api/instance/v1"
	"context"
)

func (service *InstanceService) ListInstanceTypes(ctx context.Context, request *v1.ListInstanceTypesRequest) (*v1.ListInstanceTypesResponse, error) {
	result, err := service.itc.ListInstanceTypes(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint)
	if err != nil {
		return nil, err
	}

	var instanceTypes []*v1.ListInstanceTypesResponse_InstanceType
	if result.Body.InstanceTypes != nil {
		instanceTypes = make([]*v1.ListInstanceTypesResponse_InstanceType, len(result.Body.InstanceTypes.InstanceType))
		for i, v := range result.Body.InstanceTypes.InstanceType {
			instanceTypes[i] = &v1.ListInstanceTypesResponse_InstanceType{
				LocalStorageCategory:        *v.LocalStorageCategory,
				PrimaryEniQueueNumber:       *v.PrimaryEniQueueNumber,
				MemorySize:                  *v.MemorySize,
				LocalStorageCapacity:        *v.LocalStorageCapacity,
				InstanceFamilyLevel:         *v.InstanceFamilyLevel,
				InstancePpsRx:               *v.InstancePpsRx,
				EniIpv6AddressQuantity:      *v.EniIpv6AddressQuantity,
				MaximumQueueNumberPerEni:    *v.MaximumQueueNumberPerEni,
				InstanceTypeId:              *v.InstanceTypeId,
				InstanceBandwidthRx:         *v.InstanceBandwidthRx,
				SecondaryEniQueueNumber:     *v.SecondaryEniQueueNumber,
				GpuSpec:                     *v.GPUSpec,
				QueuePairNumber:             *v.QueuePairNumber,
				EriQuatity:                  *v.EriQuantity,
				GpuAmount:                   *v.GPUAmount,
				TotalEniQueueQuantity:       *v.TotalEniQueueQuantity,
				NvmeSupport:                 *v.NvmeSupport,
				DiskQuantity:                *v.DiskQuantity,
				InitialCredit:               *v.InitialCredit,
				LocalStorageAmount:          *v.LocalStorageAmount,
				BaselineCredit:              *v.BaselineCredit,
				InstancePpsTx:               *v.InstancePpsTx,
				EniPrivateIpAddressQuantity: *v.EniPrivateIpAddressQuantity,
				CpuCoreCount:                *v.CpuCoreCount,
				InstanceTypeFamily:          *v.InstanceTypeFamily,
				EniQuantity:                 *v.EniQuantity,
			}
		}
	}
	return &v1.ListInstanceTypesResponse{
		RequestId:     *result.Body.RequestId,
		InstanceTypes: instanceTypes,
		NextToken:     *result.Body.NextToken,
	}, nil
}
