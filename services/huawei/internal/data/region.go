package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/region"
	ecsregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
	imsregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ims/v2/region"
	vpcregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/region"
	"huawei/internal/biz"
)

type regionResponse struct {
	data *Data
	log  *log.Helper
}

var staticImageRegions = map[string]*region.Region{
	"af-south-1":     imsregion.AF_SOUTH_1,
	"cn-north-4":     imsregion.CN_NORTH_4,
	"cn-north-1":     imsregion.CN_NORTH_1,
	"cn-east-2":      imsregion.CN_EAST_2,
	"cn-east-3":      imsregion.CN_EAST_3,
	"cn-south-1":     imsregion.CN_SOUTH_1,
	"cn-southwest-2": imsregion.CN_SOUTHWEST_2,
	"ap-southeast-2": imsregion.AP_SOUTHEAST_2,
	"ap-southeast-1": imsregion.AP_SOUTHEAST_1,
	"ap-southeast-3": imsregion.AP_SOUTHEAST_3,
	"cn-north-2":     imsregion.CN_NORTH_2,
	"cn-south-2":     imsregion.CN_SOUTH_2,
}

var staticEcsRegions = map[string]*region.Region{
	"cn-north-1":     ecsregion.CN_NORTH_1,
	"cn-north-4":     ecsregion.CN_NORTH_4,
	"cn-south-1":     ecsregion.CN_SOUTH_1,
	"cn-east-2":      ecsregion.CN_EAST_2,
	"cn-east-3":      ecsregion.CN_EAST_3,
	"cn-southwest-2": ecsregion.CN_SOUTHWEST_2,
	"ap-southeast-1": ecsregion.AP_SOUTHEAST_1,
	"ap-southeast-2": ecsregion.AP_SOUTHEAST_2,
	"ap-southeast-3": ecsregion.AP_SOUTHEAST_3,
	"af-south-1":     ecsregion.AF_SOUTH_1,
	"sa-brazil-1":    ecsregion.SA_BRAZIL_1,
	"la-north-2":     ecsregion.LA_NORTH_2,
	"cn-south-4":     ecsregion.CN_SOUTH_4,
	"na-mexico-1":    ecsregion.NA_MEXICO_1,
	"la-south-2":     ecsregion.LA_SOUTH_2,
}

var staticVpcRegions = map[string]*region.Region{
	"af-south-1":     vpcregion.AF_SOUTH_1,
	"cn-north-4":     vpcregion.CN_NORTH_4,
	"cn-north-1":     vpcregion.CN_NORTH_1,
	"cn-east-2":      vpcregion.CN_EAST_2,
	"cn-east-3":      vpcregion.CN_EAST_3,
	"cn-south-1":     vpcregion.CN_SOUTH_1,
	"cn-southwest-2": vpcregion.CN_SOUTHWEST_2,
	"ap-southeast-2": vpcregion.AP_SOUTHEAST_2,
	"cn-north-9":     vpcregion.CN_NORTH_9,
	"ap-southeast-1": vpcregion.AP_SOUTHEAST_1,
	"ap-southeast-3": vpcregion.AP_SOUTHEAST_3,
	"sa-brazil-1":    vpcregion.SA_BRAZIL_1,
	"la-north-2":     vpcregion.LA_NORTH_2,
}

// NewRegionRepo .
func NewRegionRepo(data *Data, logger log.Logger) biz.RegionResponse {
	return &regionResponse{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *regionResponse) ListRegions(ctx context.Context, regionType string) *biz.ListRegionResponse {
	regionFields := staticEcsRegions
	if regionType == "image" {
		regionFields = staticImageRegions
	} else if regionType == "vpc" {
		regionFields = staticVpcRegions
	}
	regions := make([]*region.Region, len(regionFields))
	i := 0
	for _, v := range staticImageRegions {
		regions[i] = &region.Region{
			Id:       v.Id,
			Endpoint: v.Endpoint,
		}
		i++
	}
	return &biz.ListRegionResponse{Regions: regions}
}
