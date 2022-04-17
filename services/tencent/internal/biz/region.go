package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

type RegionResponse interface {
	ListAll(ctx context.Context, secretId, secretKey string, request *cvm.DescribeRegionsRequest) (*cvm.DescribeRegionsResponse, error)
}

// RegionUseCase is a Region UseCase.
type RegionUseCase struct {
	repo RegionResponse
	log  *log.Helper
}

// NewRegionUseCase new a Region UseCase.
func NewRegionUseCase(repo RegionResponse, logger log.Logger) *RegionUseCase {
	return &RegionUseCase{repo: repo, log: log.NewHelper(logger)}
}

// ListAll List All Regions
func (uc *RegionUseCase) ListAll(ctx context.Context, secretId, secretKey string, request *cvm.DescribeRegionsRequest) (*cvm.DescribeRegionsResponse, error) {
	return uc.repo.ListAll(ctx, secretId, secretKey, request)
}
