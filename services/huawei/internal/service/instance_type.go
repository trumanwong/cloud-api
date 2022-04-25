package service

import (
	"context"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	v1 "huawei/api/instance/v1"
)

func (service *InstanceService) ListInstanceTypes(ctx context.Context, request *v1.ListInstanceTypesRequest) (*v1.ListInstanceTypesResponse, error) {
	response, err := service.itc.ListInstanceTypes(ctx, request.AccessKey, request.SecretKey, request.RegionId, request.ProjectId, &model.ListFlavorsRequest{})
	if err != nil {
		return nil, err
	}

	instanceTypes := make([]*v1.ListInstanceTypesResponse_InstanceType, 0)
	if response.Flavors != nil {
		instanceTypes = make([]*v1.ListInstanceTypesResponse_InstanceType, len(*response.Flavors))
		for i, instanceType := range *response.Flavors {
			var osExtraSpec *v1.ListInstanceTypesResponse_InstanceType_OsExtraSpec
			if instanceType.OsExtraSpecs != nil {
				osExtraSpec = &v1.ListInstanceTypesResponse_InstanceType_OsExtraSpec{
					EcsPerformanceType:            *instanceType.OsExtraSpecs.Ecsperformancetype,
					HwNumaNodes:                   *instanceType.OsExtraSpecs.HwnumaNodes,
					ResourceType:                  *instanceType.OsExtraSpecs.ResourceType,
					HPetSupport:                   *instanceType.OsExtraSpecs.HpetSupport,
					InstanceVNicType:              *instanceType.OsExtraSpecs.InstanceVnictype,
					InstanceVNicInstanceBandwidth: *instanceType.OsExtraSpecs.InstanceVnicinstanceBandwidth,
					InstanceVnicMaxCount:          *instanceType.OsExtraSpecs.InstanceVnicmaxCount,
					QuotaLocalDisk:                *instanceType.OsExtraSpecs.QuotalocalDisk,
					QuotaNvmeSsd:                  *instanceType.OsExtraSpecs.QuotanvmeSsd,
					ExtraSpecIoPersistentGrant:    *instanceType.OsExtraSpecs.ExtraSpeciopersistentGrant,
					EcsGeneration:                 *instanceType.OsExtraSpecs.Ecsgeneration,
					EcsVirtualizationEnvTypes:     *instanceType.OsExtraSpecs.EcsvirtualizationEnvTypes,
					PciPassthroughEnableGpu:       *instanceType.OsExtraSpecs.PciPassthroughenableGpu,
					PciPassthroughGpuSpecs:        *instanceType.OsExtraSpecs.PciPassthroughgpuSpecs,
					PciPassthroughAlias:           *instanceType.OsExtraSpecs.PciPassthroughalias,
					CondOperationStatus:           *instanceType.OsExtraSpecs.Condoperationstatus,
					CondOperationAz:               *instanceType.OsExtraSpecs.Condoperationaz,
					QuotaMaxRate:                  *instanceType.OsExtraSpecs.QuotamaxRate,
					QuotaMinRate:                  *instanceType.OsExtraSpecs.QuotaminRate,
					QuotaMaxPps:                   *instanceType.OsExtraSpecs.QuotamaxPps,
					CondOperationCharge:           *instanceType.OsExtraSpecs.Condoperationcharge,
					CondOperationChargeStop:       *instanceType.OsExtraSpecs.Condoperationchargestop,
					CondSpotOperationAz:           *instanceType.OsExtraSpecs.Condspotoperationaz,
					CondOperationRoles:            *instanceType.OsExtraSpecs.Condoperationroles,
					CondSpotOperationStatus:       *instanceType.OsExtraSpecs.Condspotoperationstatus,
					CondNetwork:                   *instanceType.OsExtraSpecs.Condnetwork,
					CondStorage:                   *instanceType.OsExtraSpecs.Condstorage,
					CondComputeLiveResizable:      *instanceType.OsExtraSpecs.CondcomputeliveResizable,
					CondCompute:                   *instanceType.OsExtraSpecs.Condcompute,
					InfoGpuName:                   *instanceType.OsExtraSpecs.Infogpuname,
					InfoCpuName:                   *instanceType.OsExtraSpecs.Infocpuname,
					QuotaGpu:                      *instanceType.OsExtraSpecs.Quotagpu,
					EcsInstanceArchitecture:       *instanceType.OsExtraSpecs.EcsinstanceArchitecture,
				}
			}
			var serverAttachableQuantity *v1.ListInstanceTypesResponse_InstanceType_ServerAttachableQuantity
			if instanceType.AttachableQuantity != nil {
				serverAttachableQuantity = &v1.ListInstanceTypesResponse_InstanceType_ServerAttachableQuantity{
					FreeScsi: instanceType.AttachableQuantity.FreeScsi,
					FreeBlk:  instanceType.AttachableQuantity.FreeBlk,
					FreeDisk: instanceType.AttachableQuantity.FreeDisk,
					FreeNic:  instanceType.AttachableQuantity.FreeNic,
				}
			}
			instanceTypes[i] = &v1.ListInstanceTypesResponse_InstanceType{
				Id:                       instanceType.Id,
				Name:                     instanceType.Name,
				VCpus:                    instanceType.Vcpus,
				MemorySize:               instanceType.Ram,
				OsExtraSpec:              osExtraSpec,
				ServerAttachableQuantity: serverAttachableQuantity,
			}
		}
	}
	listInstanceTypesResponse := &v1.ListInstanceTypesResponse{
		InstanceTypes: instanceTypes,
	}
	return listInstanceTypesResponse, nil
}
