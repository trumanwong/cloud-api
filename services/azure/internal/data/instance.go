package data

import (
	"azure/internal/biz"
	"context"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-12-01/compute"
	"github.com/go-kratos/kratos/v2/log"
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

func (r *instanceResponse) CreateInstances(ctx context.Context, tenantID, clientID, clientSecret, subscriptionId, resourceGroupName, vmName string, request compute.VirtualMachine) (*compute.VirtualMachinesCreateOrUpdateFuture, error) {
	client, err := getVMClient(tenantID, clientID, clientSecret, subscriptionId)
	if err != nil {
		return nil, err
	}
	result, err := client.CreateOrUpdate(ctx, resourceGroupName, vmName, request)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *instanceResponse) ListInstances(ctx context.Context, tenantID, clientID, clientSecret, subscriptionId, resourceGroupName, filter string) (*compute.VirtualMachineListResultPage, error) {
	client, err := getVMClient(tenantID, clientID, clientSecret, subscriptionId)
	if err != nil {
		return nil, err
	}
	result, err := client.List(ctx, resourceGroupName, filter)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *instanceResponse) StartInstances(ctx context.Context, tenantID, clientID, clientSecret, subscriptionId, resourceGroupName, vmName string) (*compute.VirtualMachinesStartFuture, error) {
	client, err := getVMClient(tenantID, clientID, clientSecret, subscriptionId)
	if err != nil {
		return nil, err
	}
	result, err := client.Start(ctx, resourceGroupName, vmName)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *instanceResponse) StopInstances(ctx context.Context, tenantID, clientID, clientSecret, subscriptionId, resourceGroupName, vmName string, skipShutDown *bool) (*compute.VirtualMachinesPowerOffFuture, error) {
	client, err := getVMClient(tenantID, clientID, clientSecret, subscriptionId)
	if err != nil {
		return nil, err
	}
	result, err := client.PowerOff(ctx, resourceGroupName, vmName, skipShutDown)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *instanceResponse) RebootInstances(ctx context.Context, tenantID, clientID, clientSecret, subscriptionId, resourceGroupName, vmName string) (*compute.VirtualMachinesRestartFuture, error) {
	client, err := getVMClient(tenantID, clientID, clientSecret, subscriptionId)
	if err != nil {
		return nil, err
	}
	result, err := client.Restart(ctx, resourceGroupName, vmName)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *instanceResponse) DeleteInstances(ctx context.Context, tenantID, clientID, clientSecret, subscriptionId, resourceGroupName, vmName string, forceDeletion *bool) (*compute.VirtualMachinesDeleteFuture, error) {
	client, err := getVMClient(tenantID, clientID, clientSecret, subscriptionId)
	if err != nil {
		return nil, err
	}
	result, err := client.Delete(ctx, resourceGroupName, vmName, forceDeletion)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
