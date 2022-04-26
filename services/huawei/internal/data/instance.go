package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	"huawei/internal/biz"
)

type instanceResponse struct {
	data *Data
	log  *log.Helper
}

// NewInstanceRepo .
func NewInstanceRepo(data *Data, logger log.Logger) biz.InstanceResponse {
	return &instanceResponse{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *instanceResponse) CreateInstances(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.CreateServersRequest) (*model.CreateServersResponse, error) {
	client := getEcsClient(accessKey, secretKey, regionId, projectId)
	result, err := client.CreateServers(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *instanceResponse) ListInstances(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.ListServersDetailsRequest) (*model.ListServersDetailsResponse, error) {
	client := getEcsClient(accessKey, secretKey, regionId, projectId)
	result, err := client.ListServersDetails(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *instanceResponse) StartInstances(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.BatchStartServersRequest) (*model.BatchStartServersResponse, error) {
	client := getEcsClient(accessKey, secretKey, regionId, projectId)
	result, err := client.BatchStartServers(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *instanceResponse) StopInstances(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.BatchStopServersRequest) (*model.BatchStopServersResponse, error) {
	client := getEcsClient(accessKey, secretKey, regionId, projectId)
	result, err := client.BatchStopServers(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *instanceResponse) RebootInstances(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.BatchRebootServersRequest) (*model.BatchRebootServersResponse, error) {
	client := getEcsClient(accessKey, secretKey, regionId, projectId)
	result, err := client.BatchRebootServers(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *instanceResponse) DeleteInstances(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.DeleteServersRequest) (*model.DeleteServersResponse, error) {
	client := getEcsClient(accessKey, secretKey, regionId, projectId)
	result, err := client.DeleteServers(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *instanceResponse) ShowJob(ctx context.Context, accessKey, secretKey, regionId, projectId string, request *model.ShowJobRequest) (*model.ShowJobResponse, error) {
	client := getEcsClient(accessKey, secretKey, regionId, projectId)
	result, err := client.ShowJob(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}
