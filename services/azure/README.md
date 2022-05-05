# Azure api

## Quick Start

### Request a Security Credential

For more information about credential, see [Register your client application with Azure AD](https://docs.microsoft.com/en-us/rest/api/azure/#register-your-client-application-with-azure-ad).

## List all regions

Gets all available geo-locations.

### Try it out

```
curl --location --request GET 'http://127.0.0.1:8004/api/regions?tenantId={{yourTenantId}}&clientSecret={{yourClientSecret}}&&subscriptionId={{yourSubscriptionId}}'
```

For more information about regions, see [Subscriptions - List Locations](https://docs.microsoft.com/en-us/rest/api/resources/subscriptions/list-locations).

## List all images

Gets the list of Images in the subscription.

### Try it out

```shell
curl --location --request GET 'http://127.0.0.1:8004/api/images?tenantId={{yourTenantId}}&clientSecret={{yourClientSecret}}&&subscriptionId={{yourSubscriptionId}}'
```

For more information about images, see [Images - List](https://docs.microsoft.com/en-us/rest/api/compute/images/list).

## List all instance types

Gets the list of Microsoft.Compute SKUs available for your Subscription.

### Try it out

```shell
curl --location --request GET 'http://127.0.0.1:8004/api/instance-types?tenantId={{yourTenantId}}&clientSecret={{yourClientSecret}}&&subscriptionId={{yourSubscriptionId}}'
```

For more information about instance types, see [Resource Skus - List](https://docs.microsoft.com/en-us/rest/api/compute/resource-skus/list)

## CreateInstances

The operation to create or update a virtual machine.

### Try it out

```
curl --location --request POST 'http://127.0.0.1:8004/api/instances' \
--header 'Content-Type: application/json' \
--data-raw '{
    "tenantId": "{{yourTenantId}}",
    "clientId": "{{yourClientId}}",
    "clientSecret": "{{yourClientSecret}}",
    "subscriptionId": "{{yourSubscriptionId}}",
    "resourceGroupName": "{{yourResourceGroupName}}",
    "vmName": "{{yourVmName}}",
    "hardwareProfile": {
    	"vmSize": "Standard_D1_v2"
    },
    "storageProfile": {
    	"osDisk": {
    		"name": "test",
    		"imageUri": "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/{existing-generalized-os-image-blob-name}.vhd",
    		"vhdUri": "http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk.vhd",
    		"osType": "Windows"
    	}
    },
    "osProfile": {
    	"adminUsername": "{{your-username}}",
    	"adminPassword": "{{your-password}}"
    }
}'
```

For more information about create instances, see [Virtual Machines - Create Or Update](https://docs.microsoft.com/en-us/rest/api/compute/virtual-machines/create-or-update).

## ListInstances

Lists all of the virtual machines in the specified resource group.

### Try it out

```shell
curl --location --request GET 'http://127.0.0.1:8004/api/instances?tenantId={{yourTenantId}}&clientSecret={{yourClientSecret}}&&subscriptionId={{yourSubscriptionId}}&resourceGroupName={{yourResourceGroupName}}'
```

For more information about list instances, see [Virtual Machines - List](https://docs.microsoft.com/en-us/rest/api/compute/virtual-machines/list).

## StartInstances

The operation to start a virtual machine..

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8004/api/instances/start' \
--header 'Content-Type: application/json' \
--data-raw '{
    "tenantId": "{{yourTenantId}}",
    "clientId": "{{yourClientId}}",
    "clientSecret": "{{yourClientSecret}}",
    "subscriptionId": "{{yourSubscriptionId}}",
    "resourceGroupName": "{{yourResourceGroupName}}",
    "vmName": "{{yourVmName}}"
}'
```

For more information about start instances, see [Virtual Machines - Start](https://docs.microsoft.com/en-us/rest/api/compute/virtual-machines/start).

## StopInstances

The operation to power off (stop) a virtual machine.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8004/api/instances/stop' \
--header 'Content-Type: application/json' \
--data-raw '{
    "tenantId": "{{yourTenantId}}",
    "clientId": "{{yourClientId}}",
    "clientSecret": "{{yourClientSecret}}",
    "subscriptionId": "{{yourSubscriptionId}}",
    "resourceGroupName": "{{yourResourceGroupName}}",
    "vmName": "{{yourVmName}}"
}'
```

For more information about stop instances, see [Virtual Machines - Power Off](https://docs.microsoft.com/en-us/rest/api/compute/virtual-machines/power-off).

## RebootInstances

The operation to restart a virtual machine.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8004/api/instances/reboot' \
--header 'Content-Type: application/json' \
--data-raw '{
    "tenantId": "{{yourTenantId}}",
    "clientId": "{{yourClientId}}",
    "clientSecret": "{{yourClientSecret}}",
    "subscriptionId": "{{yourSubscriptionId}}",
    "resourceGroupName": "{{yourResourceGroupName}}",
    "vmName": "{{yourVmName}}"
}'
```

For more information about stop instances, see [Virtual Machines - Restart](https://docs.microsoft.com/en-us/rest/api/compute/virtual-machines/restart).

## TerminateInstances

The operation to delete a virtual machine.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8004/api/instances/terminate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "tenantId": "{{yourTenantId}}",
    "clientId": "{{yourClientId}}",
    "clientSecret": "{{yourClientSecret}}",
    "subscriptionId": "{{yourSubscriptionId}}",
    "resourceGroupName": "{{yourResourceGroupName}}",
    "vmName": "{{yourVmName}}"
}'
```

For more information about delete instances, see [Virtual Machines - Delete](https://docs.microsoft.com/en-us/rest/api/compute/virtual-machines/delete).