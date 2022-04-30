# Huawei cloud api

## Quick Start

### Request a Security Credential

Before you begin, you need to sign up for an Huawei cloud account and retrieve your [AK/SK](https://support.huaweicloud.com/qs-iam/iam_01_0031.html).

For more information about credential, see [Authentication](https://support.huaweicloud.com/devg-apisign/api-sign-provide-aksk.html).

## List all regions

List all huawei cloud ECS/IMS(IMAGES)/VPC available regions.

### Try it out

```
curl --location --request GET 'http://127.0.0.1:8002/api/regions?regionType=ecs'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "regions": [
                {
                    "Endpoint": "https://ecs.cn-north-4.myhuaweicloud.com",
                    "Id": "cn-north-4"
                },
                ...
                {
                    "Endpoint": "https://ecs.na-mexico-1.myhuaweicloud.com",
                    "Id": "na-mexico-1"
                }
            ]
        }
    }
}
```

For more information about regions, see [DescribeRegions](https://cloud.tencent.com/document/api/213/15708).

## List all images

List all huawei cloud available images.

### Try it out

```shell
curl --location --request GET 'http://127.0.0.1:8002/api/images?accessKey={{yourHuaweiAccessKey}}&secretKey={{yourHuaweiSecretKey}}&regionId=cn-south-1'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "first": "/v2/images?__isregistered=true",
            "images": [
                {
                    "__account_code": "",
                    "__data_origin": "image,cn-south-1,25604208-2d47-498b-a054-6d8d36dbd3b4",
                    "__image_size": "1848573952",
                    "__image_source_type": "uds",
                    "__imagetype": "market",
                    "__is_offshelved": "false",
                    "__isregistered": "true",
                    "__lazyloading": "True",
                    "__originalimagename": "25604208-2d47-498b-a054-6d8d36dbd3b4",
                    "__os_bit": "64",
                    "__os_feature_list": "{\"nic_hotplug\":\"true\", \"disk_hotplug\": \"true\", \"user_data\": \"true\", \"ssh_key\": \"true\", \"hostname_inject\": \"true\", \"onekey_resetpasswd\": \"true\"}",
                    "__os_type": "Linux",
                    "__os_version": "openEuler 20.03 64bit",
                    "__platform": "openEuler",
                    "__productcode": "0199bc7498734da8a78b8d66f213be00",
                    "__support_agent_list": "hss",
                    "__support_arm": "true",
                    "active_at": "2022-04-28T10:47:52Z",
                    "checksum": "99914b932bd37a50b983c5e7c90ae93b",
                    "container_format": "bare",
                    "created_at": "2022-04-28T10:47:20Z",
                    "disk_format": "zvhd2",
                    "file": "/v2/images/88faab67-9e31-48d9-8027-f8a025da94de/file",
                    "hw_firmware_type": "uefi",
                    "hw_vif_multiqueue_enabled": "true",
                    "id": "88faab67-9e31-48d9-8027-f8a025da94de",
                    "max_ram": "0",
                    "min_disk": 40,
                    "min_ram": 1024,
                    "name": "COPY_远程授权激活系统镜像文件",
                    "owner": "055ae22e000025062fe1c00d2c57a287",
                    "protected": true,
                    "schema": "/v2/schemas/image",
                    "self": "/v2/images/88faab67-9e31-48d9-8027-f8a025da94de",
                    "size": 2,
                    "status": "active",
                    "tags": [],
                    "updated_at": "2022-04-28T11:25:17Z",
                    "virtual_env_type": "FusionCompute",
                    "visibility": "public"
                },
                ...
                {
                    "__data_origin": "image,cn-east-2,1bbe7ee1-f280-49a3-b2ac-be9dce37a098",
                    "__description": "云时通标准套件镜像",
                    "__image_size": "9260425216",
                    "__image_source_type": "uds",
                    "__imagetype": "market",
                    "__is_offshelved": "false",
                    "__isregistered": "true",
                    "__lazyloading": "True",
                    "__originalimagename": "1bbe7ee1-f280-49a3-b2ac-be9dce37a098",
                    "__os_bit": "64",
                    "__os_feature_list": "{\"nic_hotplug\":\"true\", \"disk_hotplug\": \"true\", \"user_data\": \"true\", \"ssh_key\":\"true\", \"hostname_inject\": \"true\", \"onekey_resetpasswd\": \"true\"}",
                    "__os_type": "Linux",
                    "__os_version": "CentOS 7.9 64bit",
                    "__platform": "CentOS",
                    "__productcode": "3a11562842d393e4065db612ab1188ba",
                    "__support_agent_list": "hss,ces",
                    "__support_kvm": "true",
                    "active_at": "2022-04-01T04:26:35Z",
                    "checksum": "99914b932bd37a50b983c5e7c90ae93b",
                    "container_format": "bare",
                    "created_at": "2022-04-01T02:12:49Z",
                    "disk_format": "zvhd2",
                    "file": "/v2/images/81d3fbb5-9f44-47f3-9ce2-9fe43658f027/file",
                    "hw_vif_multiqueue_enabled": "true",
                    "id": "81d3fbb5-9f44-47f3-9ce2-9fe43658f027",
                    "max_ram": "0",
                    "min_disk": 40,
                    "min_ram": 1024,
                    "name": "cn-south-1_yst-buzi-std",
                    "owner": "055ae22e000025062fe1c00d2c57a287",
                    "protected": true,
                    "schema": "/v2/schemas/image",
                    "self": "/v2/images/81d3fbb5-9f44-47f3-9ce2-9fe43658f027",
                    "size": 2,
                    "status": "active",
                    "tags": [],
                    "updated_at": "2022-04-01T04:35:10Z",
                    "virtual_env_type": "FusionCompute",
                    "visibility": "public"
                }
            ],
            "next": "/v2/images?__isregistered=true&marker=81d3fbb5-9f44-47f3-9ce2-9fe43658f027",
            "schema": "/v2/schemas/images"
        }
    }
}
```

For more information about images, see [List Images](https://support.huaweicloud.com/api-ims/ims_03_0702.html).

## List all instance types

This API is used to query details about ECS flavors and extended flavor information.

### Try it out

```shell
curl --location --request GET 'http://127.0.0.1:8002/api/instance-types?accessKey=={{yourHuaweiAccessKey}}&secretKey={{yourHuaweiSecretKey}}&regionId=cn-south-1'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "flavors": [
                {
                    "OS-FLV-DISABLED:disabled": false,
                    "OS-FLV-EXT-DATA:ephemeral": 0,
                    "attachableQuantity": {
                        "free_blk": 24,
                        "free_disk": 24,
                        "free_nic": 8,
                        "free_scsi": 0
                    },
                    "disk": "0",
                    "id": "ac6.12xlarge.2",
                    "links": [
                        {
                            "href": "https://ecs.cn-south-1.myhuaweicloud.com/v1.0/ded721d4e0204e3880aa06dccddf65c0/flavors/ac6.12xlarge.2",
                            "rel": "self",
                            "type": ""
                        },
                        {
                            "href": "https://ecs.cn-south-1.myhuaweicloud.com/ded721d4e0204e3880aa06dccddf65c0/flavors/ac6.12xlarge.2",
                            "rel": "bookmark",
                            "type": ""
                        }
                    ],
                    "name": "ac6.12xlarge.2",
                    "os-flavor-access:is_public": true,
                    "os_extra_specs": {
                        "cond:compute": "autorecovery",
                        "cond:operation:az": "cn-south-1a(abandon),cn-south-2b(abandon),cn-south-1f(obt_sellout),cn-south-1e(abandon),cn-south-1(abandon),cn-south-1c(abandon)",
                        "cond:operation:roles": "op_gated_ecs_c6a",
                        "cond:operation:status": "abandon",
                        "ecs:generation": "c6",
                        "ecs:performancetype": "computingv3",
                        "ecs:virtualization_env_types": "CloudCompute",
                        "hw:numa_nodes": "1",
                        "info:cpu:name": "AMD EPYC 7742 2.25GHz",
                        "quota:max_pps": "2500000",
                        "quota:max_rate": "24000",
                        "quota:min_rate": "8000",
                        "resource_type": "IOoptimizedaC6_2"
                    },
                    "ram": 98304,
                    "rxtx_cap": "",
                    "rxtx_factor": 1,
                    "rxtx_quota": "",
                    "swap": "",
                    "vcpus": "48"
                },
                ...
                {
                    "OS-FLV-DISABLED:disabled": false,
                    "OS-FLV-EXT-DATA:ephemeral": 0,
                    "attachableQuantity": {
                        "free_blk": 24,
                        "free_disk": 60,
                        "free_nic": 2,
                        "free_scsi": 60
                    },
                    "disk": "0",
                    "id": "t6.xlarge.4",
                    "links": [
                        {
                            "href": "https://ecs.cn-south-1.myhuaweicloud.com/v1.0/ded721d4e0204e3880aa06dccddf65c0/flavors/t6.xlarge.4",
                            "rel": "self",
                            "type": ""
                        },
                        {
                            "href": "https://ecs.cn-south-1.myhuaweicloud.com/ded721d4e0204e3880aa06dccddf65c0/flavors/t6.xlarge.4",
                            "rel": "bookmark",
                            "type": ""
                        }
                    ],
                    "name": "t6.xlarge.4",
                    "os-flavor-access:is_public": true,
                    "os_extra_specs": {
                        "cond:compute": "autorecovery",
                        "cond:operation:az": "cn-south-1a(sellout),cn-south-2b(abandon),cn-south-1c(sellout)",
                        "cond:operation:status": "abandon",
                        "ecs:generation": "s6",
                        "ecs:performancetype": "entry",
                        "ecs:virtualization_env_types": "CloudCompute",
                        "hw:numa_nodes": "1",
                        "info:cpu:name": "Intel SkyLake 6161 2.2GHz",
                        "quota:max_pps": "200000",
                        "quota:max_rate": "1000",
                        "quota:min_rate": "200",
                        "resource_type": "IOoptimizedT6"
                    },
                    "ram": 16384,
                    "rxtx_cap": "",
                    "rxtx_factor": 1,
                    "rxtx_quota": "",
                    "swap": "",
                    "vcpus": "4"
                }
            ]
        }
    }
}
```

For more information about instance types, see [Flavor Management](https://support.huaweicloud.com/api-ecs/zh-cn_topic_0020212656.html)

## Create Instances

Used to create one or more instances of the specified configuration.

### Try it out

```
curl --location --request POST 'http://127.0.0.1:8002/api/instances' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessKey": "{{yourHuaweiAccessKey}}",
    "secretKey": "{{yourHuaweiSecretKey}}",
    "regionId": "cn-east-2",
    "imageId": "88faab67-9e31-48d9-8027-f8a025da94de",
    "name": "test",
    "instanceType": "ac6.12xlarge.2",
    "vpcId": "{{yourVpcId}}",
    "nics": [{
        "subnetId": "{{yourSubnetId}}"
    }],
    "rootVolume": {
        "volumeType": "SATA",
        "size": 60
    }
}'
```

For more information about create instances, see [Create Instances](https://support.huaweicloud.com/api-ecs/ecs_02_0101.html).

## List Instances

Query the details of one or more instances.

### Try it out

```shell
curl --location -g --request GET 'http://127.0.0.1:8002/api/instances?accessKey={{yourHuaweiAccessKey}}&secretKey={{yourHuaweiSecretKey}}&regionId=cn-south-1&project_id={{yourProjectId}}'
```

For more information about list instances, see [List Instances](https://support.huaweicloud.com/api-ecs/zh-cn_topic_0094148850.html).

## Start Instances

Used to start one or more instances.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8001/api/instances/start' \
--header 'Content-Type: application/json' \
--data-raw '{
    "secretId": "{{yourSecretId}}",
    "secretKey": "{{yourSecretKey}}",
    "region": "ap-guangzhou",
    "instance_ids": ["ins-r8hr2upy", "ins-5d8a23rs"]
}'

The response should like
{
  "Response": {
    "RequestId": "d39d6c09-44e9-4e80-8661-77b5ff3cbc15"
  }
}
```

For more information about start instances, see [StartInstances](https://support.huaweicloud.com/api-ecs/ecs_02_0301.html).

## StopInstances

Used to close one or more instances.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8002/api/instances/stop' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessKey": "{{yourHuaweiAccessKey}}",
    "secretKey": "{{yourHuaweiSecretKey}}",
    "regionId": "cn-east-2",
    "projectId": "{{yourProjectId}}",
    "servers": [{
        "id": "ins-r8hr2upy"
    }]
}'
```

For more information about stop instances, see [StopInstances](https://support.huaweicloud.com/api-ecs/ecs_02_0303.html).

## RebootInstances

This api is used to resboot the instance.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8002/api/instances/reboot' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessKey": "{{yourHuaweiAccessKey}}",
    "secretKey": "{{yourHuaweiSecretKey}}",
    "regionId": "cn-east-2",
    "projectId": "{{yourProjectId}}",
    "servers": [{
        "id": "ins-r8hr2upy"
    }]
}'
```

For more information about stop instances, see [RebootInstances](https://support.huaweicloud.com/api-ecs/ecs_02_0302.html).

## TerminateInstances

This api is used to terminate the instance.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8001/api/instances/terminate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessKey": "{{yourHuaweiAccessKey}}",
    "secretKey": "{{yourHuaweiSecretKey}}",
    "regionId": "cn-east-2",
    "projectId": "{{yourProjectId}}",
    "servers": [{
        "id": "ins-r8hr2upy"
    }]
}'
```

For more information about stop instances, see [TerminateInstances](https://support.huaweicloud.com/api-ecs/ecs_02_0103.html).