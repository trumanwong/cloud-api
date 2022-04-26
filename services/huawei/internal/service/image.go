package service

import (
	"context"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ims/v2/model"
	v1 "huawei/api/instance/v1"
)

func (service *InstanceService) ListImages(ctx context.Context, request *v1.ListImagesRequest) (*v1.ListImagesResponse, error) {
	response, err := service.ic.ListImages(ctx, request.AccessKey, request.SecretKey, request.RegionId, &model.GlanceListImagesRequest{
		Limit:  &request.Limit,
		Marker: &request.Marker,
	})
	if err != nil {
		return nil, err
	}

	images := make([]*v1.ListImagesResponse_Image, 0)
	if response.Images != nil {
		images = make([]*v1.ListImagesResponse_Image, len(*response.Images))
	}
	listImageResponse := &v1.ListImagesResponse{
		Images: images,
	}
	for i, image := range *response.Images {
		imageType, _ := image.Imagetype.MarshalJSON()
		visibility, _ := image.Visibility.MarshalJSON()
		status, _ := image.Status.MarshalJSON()
		osBit, _ := image.OsBit.MarshalJSON()
		osType, _ := image.OsType.MarshalJSON()
		platform, _ := image.Platform.MarshalJSON()
		var description, maxRam, osVersion, activeAt string
		if image.Description != nil {
			description = *image.Description
		}
		var size int64
		if image.Size != nil {
			size = *image.Size
		}
		if image.MaxRam != nil {
			maxRam = *image.MaxRam
		}
		if image.OsVersion != nil {
			osVersion = *image.OsVersion
		}
		if image.ActiveAt != nil {
			activeAt = *image.ActiveAt
		}
		listImageResponse.Images[i] = &v1.ListImagesResponse_Image{
			ImageId:          image.Id,
			ImageName:        image.Name,
			ImageDescription: description,
			ImageSize:        size,
			ImageType:        string(imageType),
			Visibility:       string(visibility),
			Status:           string(status),
			MinRam:           image.MinRam,
			MaxRam:           maxRam,
			MinDisk:          image.MinDisk,
			OsBit:            string(osBit),
			Platform:         string(platform),
			OsType:           string(osType),
			OsVersion:        osVersion,
			CreatedAt:        image.CreatedAt,
			UpdatedAt:        image.UpdatedAt,
			ActiveAt:         activeAt,
		}
		listImageResponse.Marker = image.Id
	}
	if response.Next == nil {
		listImageResponse.Marker = ""
	}
	return listImageResponse, nil
}
