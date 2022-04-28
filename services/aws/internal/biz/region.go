package biz

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/go-kratos/kratos/v2/log"
)

type RegionResponse interface {
	ListRegions(context.Context, string, string, *ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error)
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
func (uc *RegionUseCase) ListAll(ctx context.Context, accessKeyId, secretAccessKey string, request *ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error) {
	return uc.repo.ListRegions(ctx, accessKeyId, secretAccessKey, request)
}
