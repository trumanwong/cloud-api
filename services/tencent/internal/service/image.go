package service

import (
	"context"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	v1 "tencent/api/instance/v1"
)

func (service *InstanceService) ListImages(ctx context.Context, request *v1.ListImagesRequest) (*v1.ListImagesResponse, error) {
	describeImagesRequest := cvm.NewDescribeImagesRequest()
	describeImagesRequest.Offset = request.Offset
	describeImagesRequest.Limit = request.Limit
	response, err := service.ic.ListImages(ctx, request.SecretId, request.SecretKey, request.Region, describeImagesRequest)
	if err != nil {
		return nil, err
	}

	listImageResponse := &v1.ListImagesResponse{
		Images:     make([]*v1.ListImagesResponse_Image, len(response.Response.ImageSet)),
		TotalCount: response.Response.TotalCount,
		RequestId:  response.Response.RequestId,
	}

	for i, image := range response.Response.ImageSet {
		listImageResponse.Images[i] = &v1.ListImagesResponse_Image{
			ImageId:          image.ImageId,
			OsName:           image.OsName,
			ImageType:        image.ImageType,
			CreatedTime:      image.CreatedTime,
			ImageName:        image.ImageName,
			ImageDescription: image.ImageDescription,
			ImageSize:        image.ImageSize,
			Architecture:     image.Architecture,
			ImageState:       image.ImageState,
			Platform:         image.Platform,
			ImageCreator:     image.ImageCreator,
			ImageSource:      image.ImageSource,
		}
	}
	return listImageResponse, nil
}
