package biz

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/go-kratos/kratos/v2/log"
)

type RegionResponse interface {
	ListAll(ctx context.Context, accessKeyId, secretAccessKey, token string, request *ec2.DescribeRegionsInput) ([]types.Region, error)
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
func (uc *RegionUseCase) ListAll(ctx context.Context, accessKeyId, secretAccessKey, token string, request *ec2.DescribeRegionsInput) ([]types.Region, error) {
	return uc.repo.ListAll(ctx, accessKeyId, secretAccessKey, token, request)
}
