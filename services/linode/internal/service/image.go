package service

import (
	"context"
	"github.com/linode/linodego"
	v1 "linode/api/instance/v1"
)

func (service *InstanceService) ListImages(ctx context.Context, request *v1.ListImagesRequest) (*v1.ListImagesResponse, error) {
	options := &linodego.ListOptions{
		PageSize: int(request.PageSize),
		PageOptions: &linodego.PageOptions{
			Page: int(request.Page),
		},
	}
	images, err := service.ic.ListImages(ctx, request.AccessToken, options)
	if err != nil {
		return nil, err
	}

	listImageResponse := &v1.ListImagesResponse{
		Images:  make([]*v1.ListImagesResponse_Image, len(images)),
		Page:    int32(options.Page),
		Pages:   int32(options.Pages),
		Results: int32(options.Results),
	}
	for i, image := range images {
		listImageResponse.Images[i] = &v1.ListImagesResponse_Image{
			Created:     image.Created.Format("2006-01-02 15:04:05"),
			CreatedBy:   image.CreatedBy,
			Deprecated:  image.Deprecated,
			Description: image.Description,
			Expiry:      image.Expiry.Format("2006-01-02 15:04:05"),
			Id:          image.ID,
			IsPublic:    image.IsPublic,
			Label:       image.Label,
			Size:        int32(image.Size),
			Status:      string(image.Status),
			Type:        image.Type,
			Vendor:      image.Vendor,
		}
	}
	return listImageResponse, nil
}
