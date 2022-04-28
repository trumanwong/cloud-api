package biz

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-12-01/compute"
	"github.com/go-kratos/kratos/v2/log"
)

// InstanceResponse is a Instance repo.
type InstanceResponse interface {
	CreateInstances(context.Context, string, string, string, string, string, string, compute.VirtualMachine) (*compute.VirtualMachinesCreateOrUpdateFuture, error)
	ListInstances(context.Context, string, string, string, string, string, string) (*compute.VirtualMachineListResultPage, error)
	StartInstances(context.Context, string, string, string, string, string, string) (*compute.VirtualMachinesStartFuture, error)
	StopInstances(context.Context, string, string, string, string, string, string, *bool) (*compute.VirtualMachinesPowerOffFuture, error)
	RebootInstances(context.Context, string, string, string, string, string, string) (*compute.VirtualMachinesRestartFuture, error)
	DeleteInstances(context.Context, string, string, string, string, string, string, *bool) (*compute.VirtualMachinesDeleteFuture, error)
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
func (uc *InstanceUseCase) CreateInstances(ctx context.Context, tenantID, clientID, clientSecret, subscriptionId, resourceGroupName, vmName string, request compute.VirtualMachine) (*compute.VirtualMachinesCreateOrUpdateFuture, error) {
	return uc.repo.CreateInstances(ctx, tenantID, clientID, clientSecret, subscriptionId, resourceGroupName, vmName, request)
}

func (uc *InstanceUseCase) ListInstances(ctx context.Context, tenantID, clientID, clientSecret, subscriptionId, resourceGroupName, filter string) (*compute.VirtualMachineListResultPage, error) {
	return uc.repo.ListInstances(ctx, tenantID, clientID, clientSecret, subscriptionId, resourceGroupName, filter)
}

func (uc *InstanceUseCase) StartInstances(ctx context.Context, tenantID, clientID, clientSecret, subscriptionId, resourceGroupName, vmName string) (*compute.VirtualMachinesStartFuture, error) {
	return uc.repo.StartInstances(ctx, tenantID, clientID, clientSecret, subscriptionId, resourceGroupName, vmName)
}

func (uc *InstanceUseCase) StopInstances(ctx context.Context, tenantID, clientID, clientSecret, subscriptionId, resourceGroupName, vmName string, skipShutDown *bool) (*compute.VirtualMachinesPowerOffFuture, error) {
	return uc.repo.StopInstances(ctx, tenantID, clientID, clientSecret, subscriptionId, resourceGroupName, vmName, skipShutDown)
}

func (uc *InstanceUseCase) RebootInstances(ctx context.Context, tenantID, clientID, clientSecret, subscriptionId, resourceGroupName, vmName string) (*compute.VirtualMachinesRestartFuture, error) {
	return uc.repo.RebootInstances(ctx, tenantID, clientID, clientSecret, subscriptionId, resourceGroupName, vmName)
}

func (uc *InstanceUseCase) DeleteInstances(ctx context.Context, tenantID, clientID, clientSecret, subscriptionId, resourceGroupName, vmName string, forceDeletion *bool) (*compute.VirtualMachinesDeleteFuture, error) {
	return uc.repo.DeleteInstances(ctx, tenantID, clientID, clientSecret, subscriptionId, resourceGroupName, vmName, forceDeletion)
}
