package biz

import (
	"context"
	"github.com/vultr/govultr/v2"

	"github.com/go-kratos/kratos/v2/log"
)

type ListInstancesResponse struct {
	Instances []govultr.Instance `json:"instances"`
	Meta      *govultr.Meta      `json:"meta"`
}

// InstanceResponse is a Instance repo.
type InstanceResponse interface {
	CreateInstance(context.Context, string, *govultr.InstanceCreateReq) (*govultr.Instance, error)
	ListInstances(context.Context, string, *govultr.ListOptions) (*ListInstancesResponse, error)
	StartInstance(context.Context, string, string) error
	StopInstance(context.Context, string, string) error
	RebootInstance(context.Context, string, string) error
	DeleteInstance(context.Context, string, string) error
}

// InstanceUseCase is a Instance UseCase.
type InstanceUseCase struct {
	repo InstanceResponse
	log  *log.Helper
}

// NewInstanceUseCase new a Instance UseCase.
func NewInstanceUseCase(repo InstanceResponse, logger log.Logger) *InstanceUseCase {
	return &InstanceUseCase{repo: repo, log: log.NewHelper(logger)}
}

// CreateInstance create Instance, and returns the new Instance.
func (uc *InstanceUseCase) CreateInstance(ctx context.Context, accessToken string, request *govultr.InstanceCreateReq) (*govultr.Instance, error) {
	return uc.repo.CreateInstance(ctx, accessToken, request)
}

func (uc *InstanceUseCase) ListInstances(ctx context.Context, accessToken string, request *govultr.ListOptions) (*ListInstancesResponse, error) {
	return uc.repo.ListInstances(ctx, accessToken, request)
}

func (uc *InstanceUseCase) StartInstance(ctx context.Context, accessToken, instanceId string) error {
	return uc.repo.StartInstance(ctx, accessToken, instanceId)
}

func (uc *InstanceUseCase) StopInstance(ctx context.Context, accessToken, instanceId string) error {
	return uc.repo.StopInstance(ctx, accessToken, instanceId)
}

func (uc *InstanceUseCase) RebootInstance(ctx context.Context, accessToken, instanceId string) error {
	return uc.repo.RebootInstance(ctx, accessToken, instanceId)
}

func (uc *InstanceUseCase) DeleteInstance(ctx context.Context, accessToken, instanceId string) error {
	return uc.repo.DeleteInstance(ctx, accessToken, instanceId)
}
