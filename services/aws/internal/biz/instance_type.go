package biz

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/go-kratos/kratos/v2/log"
)

type InstanceTypeResponse interface {
	ListAll(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.DescribeInstanceTypesInput) (*ec2.DescribeInstanceTypesOutput, error)
}

// InstanceTypeUseCase is a InstanceType UseCase.
type InstanceTypeUseCase struct {
	repo InstanceTypeResponse
	log  *log.Helper
}

// NewInstanceTypeUseCase new a InstanceType UseCase.
func NewInstanceTypeUseCase(repo InstanceTypeResponse, logger log.Logger) *InstanceTypeUseCase {
	return &InstanceTypeUseCase{repo: repo, log: log.NewHelper(logger)}
}

// ListAll List All Regions
func (uc *InstanceTypeUseCase) ListAll(ctx context.Context, accessKeyId, secretAccessKey, region string, request *ec2.DescribeInstanceTypesInput) (*ec2.DescribeInstanceTypesOutput, error) {
	return uc.repo.ListAll(ctx, accessKeyId, secretAccessKey, region, request)
}
