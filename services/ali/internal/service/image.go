package service

import (
	v1 "ali/api/instance/v1"
	"ali/internal/biz"
	"context"
)

func (service *InstanceService) ListImage(ctx context.Context, request *v1.ListImageRequest) (*v1.ListImageResponse, error) {
	images, err := service.ic.ListImage(ctx, &biz.ListImageRequest{
		AccessKeyId:     request.AccessKeyId,
		AccessKeySecret: request.AccessKeySecret,
		RegionId:        request.RegionId,
	})
	if err != nil {
		return nil, err
	}

	listImageResponse := &v1.ListImageResponse{
		Images: make([]*v1.ListImageResponse_Image, len(images)),
	}
	for i, image := range images {
		listImageResponse.Images[i] = &v1.ListImageResponse_Image{
			ImageName:    image.ImageName,
			OsName:       image.OsName,
			OsNameEn:     image.OsNameEn,
			Architecture: image.Architecture,
			OsType:       image.OsType,
		}
	}
	return listImageResponse, nil
}
