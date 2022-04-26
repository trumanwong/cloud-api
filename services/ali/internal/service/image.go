package service

import (
	v1 "ali/api/instance/v1"
	"context"
)

func (service *InstanceService) ListImages(ctx context.Context, request *v1.ListImagesRequest) (*v1.ListImagesResponse, error) {
	result, err := service.ic.ListImages(ctx, request.AccessKeyId, request.AccessKeySecret, request.Endpoint, request.RegionId)
	if err != nil {
		return nil, err
	}

	var images []*v1.ListImagesResponse_Image
	if result.Body.Images != nil {
		images = make([]*v1.ListImagesResponse_Image, len(result.Body.Images.Image))
		for i, v := range result.Body.Images.Image {
			images[i] = &v1.ListImagesResponse_Image{
				CreationTime:            *v.CreationTime,
				Status:                  *v.Status,
				ImageFamily:             *v.ImageFamily,
				Progress:                *v.Progress,
				IsCopied:                *v.IsCopied,
				IsSupportIoOptimized:    *v.IsSupportIoOptimized,
				ImageOwnerAlias:         *v.ImageOwnerAlias,
				IsSupportCloudInit:      *v.IsSupportCloudinit,
				ImageVersion:            *v.ImageVersion,
				Usage:                   *v.Usage,
				IsSelfShared:            *v.IsSelfShared,
				Description:             *v.Description,
				Size:                    *v.Size,
				ResourceGroupId:         *v.ResourceGroupId,
				Platform:                *v.Platform,
				OsNameEn:                *v.OSNameEn,
				ImageName:               *v.ImageName,
				OsName:                  *v.OSName,
				ImageId:                 *v.ImageId,
				OsType:                  *v.OSType,
				IsSubscribed:            *v.IsSubscribed,
				ProductCode:             *v.ProductCode,
				Architecture:            *v.Architecture,
				IsPublic:                *v.IsPublic,
				ImageOwnerId:            *v.ImageOwnerId,
				LoginAsNonRootSupported: *v.LoginAsNonRootSupported,
				SupplierName:            *v.SupplierName,
			}
		}
	}
	return &v1.ListImagesResponse{
		PageSize:   *result.Body.PageSize,
		PageNumber: *result.Body.PageNumber,
		RequestId:  *result.Body.RequestId,
		TotalCount: *result.Body.TotalCount,
		RegionId:   *result.Body.RegionId,
		Images:     images,
	}, nil
}
