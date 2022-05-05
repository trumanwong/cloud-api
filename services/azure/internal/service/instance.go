package service

import (
	v1 "azure/api/instance/v1"
	"azure/pkg/util/convert"
	"context"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-12-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
)

// CreateInstances Create Instances
func (service *InstanceService) CreateInstances(ctx context.Context, request *v1.CreateInstancesRequest) (*v1.CreateInstancesResponse, error) {
	virtualMachine := compute.VirtualMachine{}
	if request.HardwareProfile != nil {
		virtualMachine.HardwareProfile = &compute.HardwareProfile{
			VMSize: compute.VirtualMachineSizeTypes(request.HardwareProfile.VmSize),
		}
	}
	if request.StorageProfile != nil {
		var imageReference *compute.ImageReference
		var osDisk *compute.OSDisk
		var imageUri, vhdUri *compute.VirtualHardDisk
		var dataDisks []compute.DataDisk
		if request.StorageProfile.ImageReference != nil {
			imageReference = &compute.ImageReference{
				Publisher:               request.StorageProfile.ImageReference.Publisher,
				Offer:                   request.StorageProfile.ImageReference.Offer,
				Sku:                     request.StorageProfile.ImageReference.Sku,
				Version:                 request.StorageProfile.ImageReference.Version,
				ExactVersion:            request.StorageProfile.ImageReference.ExactVersion,
				SharedGalleryImageID:    request.StorageProfile.ImageReference.SharedGalleryImageId,
				CommunityGalleryImageID: request.StorageProfile.ImageReference.CommunityGalleryImageId,
				ID:                      request.StorageProfile.ImageReference.Id,
			}
		}
		if request.StorageProfile.OsDisk.ImageUri != nil {
			imageUri = &compute.VirtualHardDisk{
				URI: request.StorageProfile.OsDisk.ImageUri,
			}
		}
		if request.StorageProfile.OsDisk.VhdUri != nil {
			vhdUri = &compute.VirtualHardDisk{
				URI: request.StorageProfile.OsDisk.VhdUri,
			}
		}
		if request.StorageProfile.OsDisk != nil {
			osDisk = &compute.OSDisk{
				OsType:     compute.OperatingSystemTypes(request.StorageProfile.OsDisk.OsType),
				Name:       request.StorageProfile.OsDisk.Name,
				DiskSizeGB: request.StorageProfile.OsDisk.DiskSizeGb,
				Image:      imageUri,
				Vhd:        vhdUri,
			}
		}
		if request.StorageProfile.DataDisks != nil {
			dataDisks = make([]compute.DataDisk, len(request.StorageProfile.DataDisks))
			for i, v := range request.StorageProfile.DataDisks {
				dataDisks[i] = compute.DataDisk{
					Lun:        v.Lun,
					Name:       v.Name,
					DiskSizeGB: v.DiskSizeGb,
				}
			}
		}
		virtualMachine.StorageProfile = &compute.StorageProfile{
			ImageReference: imageReference,
			OsDisk:         osDisk,
			DataDisks:      &dataDisks,
		}
	}
	if request.OsProfile != nil {
		virtualMachine.OsProfile = &compute.OSProfile{
			ComputerName:  to.StringPtr(request.VmName),
			AdminUsername: request.OsProfile.AdminUsername,
			AdminPassword: request.OsProfile.AdminPassword,
		}
	}
	result, err := service.uc.CreateInstances(ctx, request.TenantId, request.ClientId, request.ClientSecret, request.SubscriptionId, request.ResourceGroupName, request.VmName, virtualMachine)
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
	result, err := service.uc.ListInstances(ctx, request.TenantId, request.ClientId, request.ClientSecret, request.SubscriptionId, request.ResourceGroupName, request.Filter)
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
	result, err := service.uc.StartInstances(ctx, request.TenantId, request.ClientId, request.ClientSecret, request.SubscriptionId, request.ResourceGroupName, request.VmName)
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
	result, err := service.uc.StopInstances(ctx, request.TenantId, request.ClientId, request.ClientSecret, request.SubscriptionId, request.ResourceGroupName, request.VmName, request.SkipShutdown)
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
	result, err := service.uc.RebootInstances(ctx, request.TenantId, request.ClientId, request.ClientSecret, request.SubscriptionId, request.ResourceGroupName, request.VmName)
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
	result, err := service.uc.DeleteInstances(ctx, request.TenantId, request.ClientId, request.ClientSecret, request.SubscriptionId, request.ResourceGroupName, request.VmName, request.ForceDeletion)
	if err != nil {
		return nil, err
	}

	data, err := convert.CastToAny(result)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteInstancesResponse{
		Data: data,
	}, nil
}
