package biz

import (
	"context"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"

	"github.com/go-kratos/kratos/v2/log"
)

// InstanceResponse is a Instance repo.
type InstanceResponse interface {
	CreateInstances(context.Context, string, string, string, string, *model.CreateServersRequest) (*model.CreateServersResponse, error)
	ListInstances(context.Context, string, string, string, string, *model.ListServersDetailsRequest) (*model.ListServersDetailsResponse, error)
	StartInstances(context.Context, string, string, string, string, *model.BatchStartServersRequest) (*model.BatchStartServersResponse, error)
	StopInstances(context.Context, string, string, string, string, *model.BatchStopServersRequest) (*model.BatchStopServersResponse, error)
	RebootInstances(context.Context, string, string, string, string, *model.BatchRebootServersRequest) (*model.BatchRebootServersResponse, error)
	DeleteInstances(context.Context, string, string, string, string, *model.DeleteServersRequest) (*model.DeleteServersResponse, error)
	ShowJob(context.Context, string, string, string, string, *model.ShowJobRequest) (*model.ShowJobResponse, error)
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

// CreateInstances creates Instances, and returns the new Instances.
func (uc *InstanceUseCase) CreateInstances(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.CreateServersRequest) (*model.CreateServersResponse, error) {
	return uc.repo.CreateInstances(ctx, accessKey, secretKey, regionId, projectId, request)
}

func (uc *InstanceUseCase) ListInstances(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.ListServersDetailsRequest) (*model.ListServersDetailsResponse, error) {
	return uc.repo.ListInstances(ctx, accessKey, secretKey, regionId, projectId, request)
}

func (uc *InstanceUseCase) StartInstances(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.BatchStartServersRequest) (*model.BatchStartServersResponse, error) {
	return uc.repo.StartInstances(ctx, accessKey, secretKey, regionId, projectId, request)
}

func (uc *InstanceUseCase) StopInstances(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.BatchStopServersRequest) (*model.BatchStopServersResponse, error) {
	return uc.repo.StopInstances(ctx, accessKey, secretKey, regionId, projectId, request)
}

func (uc *InstanceUseCase) RebootInstances(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.BatchRebootServersRequest) (*model.BatchRebootServersResponse, error) {
	return uc.repo.RebootInstances(ctx, accessKey, secretKey, regionId, projectId, request)
}

func (uc *InstanceUseCase) DeleteInstances(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.DeleteServersRequest) (*model.DeleteServersResponse, error) {
	return uc.repo.DeleteInstances(ctx, accessKey, secretKey, regionId, projectId, request)
}

func (uc *InstanceUseCase) ShowJob(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	return uc.repo.ShowJob(ctx, accessKey, secretKey, regionId, projectId, request)
}
