package service

import (
	v1 "aws/api/instance/v1"
	"context"
)

func (service *InstanceService) ListImage(ctx context.Context, request *v1.ListImageRequest) (*v1.ListImageResponse, error) {
	result, err := service.ic.ListImage(ctx, request.AccessKeyId, request.SecretKey, request.Region)
	if err != nil {
		return nil, err
	}

	listImageResponse := &v1.ListImageResponse{
		Images: make([]*v1.ListImageResponse_Image, len(result.Images)),
	}
	for i, image := range result.Images {
		listImageResponse.Images[i] = &v1.ListImageResponse_Image{
			Architecture:    string(image.Architecture),
			BootMode:        string(image.BootMode),
			CreationDate:    *image.CreationDate,
			DeprecationTime: *image.DeprecationTime,
			Description:     *image.Description,
			EnaSupport:      *image.EnaSupport,
			Hypervisor:      string(image.Hypervisor),
			ImageId:         *image.ImageId,
			ImageLocation:   *image.ImageLocation,
			ImageOwnerAlias: *image.ImageOwnerAlias,
			ImageOwnerId:    *image.OwnerId,
			ImageState:      string(image.State),
			ImageType:       string(image.ImageType),
			IsPublic:        *image.Public,
			KernelId:        *image.KernelId,
			Name:            *image.Name,
			Platform:        string(image.Platform),
			PlatformDetail:  *image.PlatformDetails,
		}
	}
	return listImageResponse, nil
}
