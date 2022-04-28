package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/linode/linodego"
)

// InstanceResponse is a Instance repo.
type InstanceResponse interface {
	CreateInstance(context.Context, string, linodego.InstanceCreateOptions) (*linodego.Instance, error)
	ListInstances(context.Context, string, *linodego.ListOptions) (*ListInstancesResponse, error)
	StartInstance(context.Context, string, int, int) error
	StopInstance(context.Context, string, int) error
	RebootInstance(context.Context, string, int, int) error
	DeleteInstance(context.Context, string, int) error
}

type ListInstancesResponse struct {
	Instances []linodego.Instance
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
func (uc *InstanceUseCase) CreateInstance(ctx context.Context, accessToken string, request linodego.InstanceCreateOptions) (*linodego.Instance, error) {
	return uc.repo.CreateInstance(ctx, accessToken, request)
}

func (uc *InstanceUseCase) ListInstances(ctx context.Context, accessToken string, request *linodego.ListOptions) (*ListInstancesResponse, error) {
	return uc.repo.ListInstances(ctx, accessToken, request)
}

func (uc *InstanceUseCase) StartInstance(ctx context.Context, accessToken string, instanceId, configId int) error {
	return uc.repo.StartInstance(ctx, accessToken, instanceId, configId)
}

func (uc *InstanceUseCase) StopInstance(ctx context.Context, accessToken string, instanceId int) error {
	return uc.repo.StopInstance(ctx, accessToken, instanceId)
}

func (uc *InstanceUseCase) RebootInstance(ctx context.Context, accessToken string, instanceId, configId int) error {
	return uc.repo.RebootInstance(ctx, accessToken, instanceId, configId)
}

func (uc *InstanceUseCase) DeleteInstance(ctx context.Context, accessToken string, instanceId int) error {
	return uc.repo.DeleteInstance(ctx, accessToken, instanceId)
}
