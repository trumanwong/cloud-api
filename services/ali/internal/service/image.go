package service

import (
	v1 "ali/api/instance/v1"
	"context"
)

func (service *InstanceService) ListImage(ctx context.Context, request *v1.ListImageRequest) (*v1.ListImageResponse, error) {
	images, err := service.ic.ListImage(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, request.RegionId)
	if err != nil {
		return nil, err
	}

	listImageResponse := &v1.ListImageResponse{
		Images: make([]*v1.ListImageResponse_Image, len(images)),
	}
	for i, image := range images {
		listImageResponse.Images[i] = &v1.ListImageResponse_Image{
			ImageId:          *image.ImageId,
			ImageName:        *image.ImageName,
			ImageDescription: *image.Description,
			OsName:           *image.OSName,
			OsNameEn:         *image.OSNameEn,
			Architecture:     *image.Architecture,
			OsType:           *image.OSType,
		}
	}
	return listImageResponse, nil
}
