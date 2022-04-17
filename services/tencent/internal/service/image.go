package service

import (
	"context"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	v1 "tencent/api/instance/v1"
)

func (service *InstanceService) ListImage(ctx context.Context, request *v1.ListImageRequest) (*v1.ListImageResponse, error) {
	response, err := service.ic.ListAll(ctx, request.SecretId, request.SecretKey, request.Region, cvm.NewDescribeImagesRequest())
	if err != nil {
		return nil, err
	}

	listImageResponse := &v1.ListImageResponse{
		Images:     make([]*v1.ListImageResponse_Image, len(response.Response.ImageSet)),
		TotalCount: *response.Response.TotalCount,
		RequestId:  *response.Response.RequestId,
	}

	for i, image := range response.Response.ImageSet {
		listImageResponse.Images[i] = &v1.ListImageResponse_Image{
			ImageId:          *image.ImageId,
			ImageName:        *image.ImageName,
			ImageDescription: *image.ImageDescription,
			OsName:           *image.OsName,
			Architecture:     *image.Architecture,
			Platform:         *image.Platform,
		}
	}
	return listImageResponse, nil
}
