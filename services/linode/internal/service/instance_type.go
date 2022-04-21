package service

import (
	"context"
	"github.com/linode/linodego"
	v1 "linode/api/instance/v1"
)

func (service *InstanceService) ListInstanceTypes(ctx context.Context, request *v1.ListInstanceTypesRequest) (*v1.ListInstanceTypesResponse, error) {
	options := &linodego.ListOptions{
		PageSize: int(request.PageSize),
		PageOptions: &linodego.PageOptions{
			Page: int(request.Page),
		},
	}

	result, err := service.itc.ListInstanceTypes(ctx, request.AccessToken, options)
	if err != nil {
		return nil, err
	}

	listInstanceTypesResponse := &v1.ListInstanceTypesResponse{
		InstanceTypes: make([]*v1.ListInstanceTypesResponse_InstanceType, len(result)),
		Page:          int32(options.Page),
		Pages:         int32(options.Pages),
		Results:       int32(options.Results),
	}
	for i, v := range result {
		var addons *v1.ListInstanceTypesResponse_InstanceType_Addon
		if v.Addons != nil && v.Addons.Backups != nil && v.Addons.Backups.Price != nil {
			addons = &v1.ListInstanceTypesResponse_InstanceType_Addon{
				Backups: &v1.ListInstanceTypesResponse_InstanceType_Addon_Backup{
					Price: &v1.ListInstanceTypesResponse_Price{
						Hourly:  v.Addons.Backups.Price.Hourly,
						Monthly: v.Addons.Backups.Price.Monthly,
					},
				},
			}
		}
		var price *v1.ListInstanceTypesResponse_Price
		if v.Price != nil {
			price = &v1.ListInstanceTypesResponse_Price{
				Hourly:  v.Price.Hourly,
				Monthly: v.Price.Monthly,
			}
		}
		listInstanceTypesResponse.InstanceTypes[i] = &v1.ListInstanceTypesResponse_InstanceType{
			Addons:     addons,
			Class:      string(v.Class),
			Disk:       int32(v.Disk),
			Gpus:       0,
			Id:         v.ID,
			Label:      v.Label,
			Memory:     int32(v.Memory),
			NetworkOut: int32(v.NetworkOut),
			Price:      price,
			Transfer:   int32(v.Transfer),
			Vcpus:      int32(v.VCPUs),
		}
	}
	return listInstanceTypesResponse, nil
}
