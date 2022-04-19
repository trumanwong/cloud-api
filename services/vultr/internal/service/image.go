package service

import (
	"context"
	"github.com/vultr/govultr/v2"
	v1 "vultr/api/instance/v1"
)

func (service *InstanceService) ListImages(ctx context.Context, request *v1.ListImagesRequest) (*v1.ListImagesResponse, error) {
	images, meta, err := service.ic.ListImages(ctx, request.AccessToken, &govultr.ListOptions{
		PerPage: int(request.PerPage),
		Cursor:  request.Cursor,
	})
	if err != nil {
		return nil, err
	}

	var link *v1.Meta_Link
	if meta.Links != nil {
		link = &v1.Meta_Link{
			Next: meta.Links.Next,
			Prev: meta.Links.Prev,
		}
	}
	listImageResponse := &v1.ListImagesResponse{
		Images: make([]*v1.ListImagesResponse_Image, len(images)),
		Meta: &v1.Meta{
			Total: int32(meta.Total),
			Link:  link,
		},
	}
	for i, image := range images {
		listImageResponse.Images[i] = &v1.ListImagesResponse_Image{
			Id:     int32(image.ID),
			Name:   image.Name,
			Arch:   image.Arch,
			Family: image.Family,
		}
	}
	return listImageResponse, nil
}
