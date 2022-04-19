package service

import (
	v1 "ali/api/instance/v1"
	"context"
)

func (service *InstanceService) ListImages(ctx context.Context, request *v1.ListImagesRequest) (*v1.ListImagesResponse, error) {
	images, err := service.ic.ListImage(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, request.RegionId)
	if err != nil {
		return nil, err
	}

	listImageResponse := &v1.ListImagesResponse{
		Images: make([]*v1.ListImagesResponse_Image, len(images)),
	}
	for i, image := range images {
		listImageResponse.Images[i] = &v1.ListImagesResponse_Image{
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
