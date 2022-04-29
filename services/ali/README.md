# Ali cloud api

## Quick Start

### Request a Security Credential

Before you begin, you need to sign up for an Alibaba Cloud account and retrieve your [Credentials](https://usercenter.console.aliyun.com/#/manage/ak).

## List all regions

Query the Alibaba Cloud regions you can use.

### Try it out

```
curl --location --request GET 'http://127.0.0.1:8000/api/regions?accessKeyId={{yourAccessKeyId}}&accessKeySecret={{yourAccessKeySecret}}'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "body": {
                "Regions": {
                    "Region": [
                        {
                            "LocalName": "华北1（青岛）",
                            "RegionEndpoint": "ecs.cn-qingdao.aliyuncs.com",
                            "RegionId": "cn-qingdao"
                        },
                    	...
                    	{
                            "LocalName": "华东5（南京-本地地域）",
                            "RegionEndpoint": "ecs.cn-nanjing.aliyuncs.com",
                            "RegionId": "cn-nanjing"
                        }
                    ]
                },
                "RequestId": "xxxxxx"
            },
            "headers": {
                "access-control-allow-headers": "X-Requested-With, X-Sequence, _aop_secret, _aop_signature, x-acs-action, x-acs-version, x-acs-date, Content-Type",
                "access-control-allow-methods": "POST, GET, OPTIONS, PUT, DELETE",
                "access-control-allow-origin": "*",
                "access-control-max-age": "172800",
                "connection": "keep-alive",
                "content-type": "application/json;charset=utf-8",
                "date": "Fri, 29 Apr 2022 07:02:12 GMT",
                "vary": "Accept-Encoding",
                "x-acs-request-id": "xxxxxx",
                "x-acs-trace-id": "xxxxxx"
            }
        }
    }
}
```

For more information about regions, see [DescribeRegions](https://help.aliyun.com/document_detail/25609.html?spm=api-workbench.API Document.0.0.72781e0f2JLh3n).

## List all images

Query the mirror resources you can use.

### Try it out

```shell
curl --location --request GET 'http://127.0.0.1:8000/api/images?accessKeyId={{yourAccessKeyId}}&accessKeySecret={{yourAccessKeySecret}}&endpoint=ecs.cn-beijing.aliyuncs.com&regionId=cn-beijing'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "body": {
                "Images": {
                    "Image": [
                        {
                            "Architecture": "x86_64",
                            "CreationTime": "2022-04-18T08:44:32Z",
                            "Description": "",
                            "DiskDeviceMappings": {},
                            "ImageFamily": "",
                            "ImageId": "aliyun_2_1903_x64_20G_alibase_20220408.vhd",
                            "ImageName": "aliyun_2_1903_x64_20G_alibase_20220408.vhd",
                            "ImageOwnerAlias": "system",
                            "ImageVersion": "",
                            "IsCopied": false,
                            "IsPublic": true,
                            "IsSelfShared": "",
                            "IsSubscribed": false,
                            "IsSupportCloudinit": true,
                            "IsSupportIoOptimized": true,
                            "LoginAsNonRootSupported": false,
                            "OSName": "Alibaba Cloud Linux  2.1903 LTS 64位",
                            "OSNameEn": "Alibaba Cloud Linux 2.1903 LTS 64 bit",
                            "OSType": "linux",
                            "Platform": "Aliyun",
                            "ProductCode": "",
                            "Progress": "100%",
                            "ResourceGroupId": "",
                            "Size": 20,
                            "Status": "Available",
                            "Tags": {},
                            "Usage": "instance"
                        },
                        ...
                        {
                            "Architecture": "x86_64",
                            "CreationTime": "2022-04-02T03:05:34Z",
                            "Description": "",
                            "DiskDeviceMappings": {},
                            "ImageFamily": "",
                            "ImageId": "debian_10_12_x64_20G_alibase_20220328.vhd",
                            "ImageName": "debian_10_12_x64_20G_alibase_20220328.vhd",
                            "ImageOwnerAlias": "system",
                            "ImageVersion": "",
                            "IsCopied": false,
                            "IsPublic": true,
                            "IsSelfShared": "",
                            "IsSubscribed": false,
                            "IsSupportCloudinit": true,
                            "IsSupportIoOptimized": true,
                            "LoginAsNonRootSupported": false,
                            "OSName": "Debian  10.12 64位",
                            "OSNameEn": "Debian  10.12 64 bit",
                            "OSType": "linux",
                            "Platform": "Debian",
                            "ProductCode": "",
                            "Progress": "100%",
                            "ResourceGroupId": "",
                            "Size": 20,
                            "Status": "Available",
                            "Tags": {},
                            "Usage": "instance"
                        }
                    ]
                },
                "PageNumber": 1,
                "PageSize": 10,
                "RegionId": "cn-beijing",
                "RequestId": "xxxxx",
                "TotalCount": 136
            },
            "headers": {
                "access-control-allow-headers": "X-Requested-With, X-Sequence, _aop_secret, _aop_signature, x-acs-action, x-acs-version, x-acs-date, Content-Type",
                "access-control-allow-methods": "POST, GET, OPTIONS, PUT, DELETE",
                "access-control-allow-origin": "*",
                "access-control-max-age": "172800",
                "connection": "keep-alive",
                "content-type": "application/json;charset=utf-8",
                "date": "Fri, 29 Apr 2022 07:09:12 GMT",
                "vary": "Accept-Encoding",
                "x-acs-request-id": "xxxxx",
                "x-acs-trace-id": "xxxxx"
            }
        }
    }
}
```

For more information about images, see [DescribeImages](https://help.aliyun.com/document_detail/25534.html?spm=api-workbench.API%20Document.0.0.72781e0f2JLh3n).

## List all instance types

Query information about all instance specifications provided by ECS, or query information about a specified instance specification.

### Try it out

```shell
curl --location --request GET 'http://127.0.0.1:8000/api/instance-types?accessKeyId={{yourAccessKeyId}}&accessKeySecret={{yourAccessKeySecret}}&endpoint=ecs.cn-beijing.aliyuncs.com'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "body": {
                "InstanceTypes": {
                    "InstanceType": [
                        {
                            "CpuCoreCount": 1,
                            "DiskQuantity": 17,
                            "EniIpv6AddressQuantity": 0,
                            "EniPrivateIpAddressQuantity": 1,
                            "EniQuantity": 1,
                            "EniTotalQuantity": 1,
                            "EniTrunkSupported": false,
                            "EriQuantity": 0,
                            "GPUAmount": 0,
                            "GPUSpec": "",
                            "InstanceFamilyLevel": "EntryLevel",
                            "InstanceTypeFamily": "ecs.t1",
                            "InstanceTypeId": "ecs.t1.xsmall",
                            "LocalStorageCategory": "",
                            "MemorySize": 0.5,
                            "NvmeSupport": "unsupported",
                            "PrimaryEniQueueNumber": 1,
                            "SecondaryEniQueueNumber": 1,
                            "TotalEniQueueQuantity": 1
                        },
                        ...
                        {
                            "CpuCoreCount": 64,
                            "DiskQuantity": 17,
                            "EniIpv6AddressQuantity": 1,
                            "EniPrivateIpAddressQuantity": 6,
                            "EniQuantity": 2,
                            "EniTotalQuantity": 2,
                            "EniTrunkSupported": false,
                            "EriQuantity": 0,
                            "GPUAmount": 0,
                            "GPUSpec": "",
                            "InstanceBandwidthRx": 1228800,
                            "InstanceBandwidthTx": 1228800,
                            "InstanceFamilyLevel": "EntryLevel",
                            "InstancePpsRx": 600000,
                            "InstancePpsTx": 600000,
                            "InstanceTypeFamily": "ecs.s6",
                            "InstanceTypeId": "ecs.s6-c1m2.16xlarge",
                            "LocalStorageCategory": "",
                            "MemorySize": 128,
                            "NvmeSupport": "unsupported",
                            "PrimaryEniQueueNumber": 2,
                            "SecondaryEniQueueNumber": 1,
                            "TotalEniQueueQuantity": 3
                        }
                    ]
                },
                "NextToken": "",
                "RequestId": "4E47123B-014C-5D2C-B00B-C8A3B671DA16"
            },
            "headers": {
                "access-control-allow-headers": "X-Requested-With, X-Sequence, _aop_secret, _aop_signature, x-acs-action, x-acs-version, x-acs-date, Content-Type",
                "access-control-allow-methods": "POST, GET, OPTIONS, PUT, DELETE",
                "access-control-allow-origin": "*",
                "access-control-max-age": "172800",
                "connection": "keep-alive",
                "content-type": "application/json;charset=utf-8",
                "date": "Fri, 29 Apr 2022 07:11:27 GMT",
                "vary": "Accept-Encoding",
                "x-acs-request-id": "xxxxx",
                "x-acs-trace-id": "xxxxx"
            }
        }
    }
}
```

For more information about instance types, see [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html?spm=api-workbench.API%20Document.0.0.72781e0f2JLh3n)

## CreateInstances

Create one or more pay-as-you-go or annual and monthly ECS instances.

### Try it out

```
curl --location --request POST 'http://127.0.0.1:8000/api/instances' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessKeyId": "{{yourAccessKeyId}}",
    "accessKeySecret": "{{yourAccessKeySecret}}",
    "endpoint": "ecs.cn-beijing.aliyuncs.com",
    "regionId": "cn-beijing",
    "imageId": "debian_10_12_x64_20G_alibase_20220328.vhd",
    "instanceType": "ecs.t1.xsmall",
    "systemDiskSize": "50",
    "instanceName": "test",
    "uniqueSuffix": true,
    "amount": 10,
    "password": "test",
    "dry_run": true
}'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "body": {
                "RequestId": "473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E",
                "OrderId": "1234567890",
                "TradePrice": 0.165,
                "InstanceIdSets": [
                "[\"i-bp67acfmxazb4pd2****\", \"i-bp1i43l28m7u48p1****\", \"i-bp12yqg7jdyxl11f****\"]"
                ]
            },
            "headers": {
                "access-control-allow-headers": "X-Requested-With, X-Sequence, _aop_secret, _aop_signature, x-acs-action, x-acs-version, x-acs-date, Content-Type",
                "access-control-allow-methods": "POST, GET, OPTIONS, PUT, DELETE",
                "access-control-allow-origin": "*",
                "access-control-max-age": "172800",
                "connection": "keep-alive",
                "content-type": "application/json;charset=utf-8",
                "date": "Fri, 29 Apr 2022 07:11:27 GMT",
                "vary": "Accept-Encoding",
                "x-acs-request-id": "xxxxx",
                "x-acs-trace-id": "xxxxx"
            }
        }
    }
}
```

For more information about create instances, see [RunInstances](https://next.api.aliyun.com/document/Ecs/2014-05-26/RunInstances).

## ListInstances

Query the details of one or more instances.

### Try it out

```shell
curl --location --request GET 'http://127.0.0.1:8000/api/instances?{{yourAccessKeyId}}&accessKeySecret={{yourAccessKeySecret}}&endpoint=ecs.cn-beijing.aliyuncs.com&region_id=cn-beijing'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "body": {
                "Instances": [
                	{
                      "CreationTime": "2017-12-10T04:04Z",
                      "SerialNumber": "51d1353b-22bf-4567-a176-8b3e12e4****",
                      "Status": "Running",
                      "DeploymentSetId": "ds-bp67acfmxazb4p****",
                      "KeyPairName": "testKeyPairName",
                      "SaleCycle": "month",
                      "SpotStrategy": "NoSpot",
                      "DeviceAvailable": true,
                      "LocalStorageCapacity": 1000,
                      "Description": "testDescription",
                      "SpotDuration": 1,
                      "InstanceNetworkType": "vpc",
                      "InstanceName": "InstanceNameTest",
                      "OSNameEn": "CentOS  7.4 64 bit",
                      "HpcClusterId": "hpc-bp67acfmxazb4p****",
                      "SpotPriceLimit": 0.98,
                      "Memory": 16384,
                      "OSName": "CentOS  7.4 64 位",
                      "DeploymentSetGroupNo": 1,
                      "ImageId": "m-bp67acfmxazb4p****",
                      "VlanId": "10",
                      "ClusterId": "c-bp67acfmxazb4p****",
                      "GPUSpec": "NVIDIA V100",
                      "AutoReleaseTime": "2017-12-10T04:04Z",
                      "DeletionProtection": true,
                      "StoppedMode": "KeepCharging",
                      "GPUAmount": 4,
                      "HostName": "testHostName",
                      "InstanceId": "i-bp67acfmxazb4p****",
                      "InternetMaxBandwidthOut": 5,
                      "InternetMaxBandwidthIn": 50,
                      "InstanceType": "ecs.g5.large",
                      "InstanceChargeType": "PostPaid",
                      "RegionId": "cn-hangzhou",
                      "IoOptimized": true,
                      "StartTime": "2017-12-10T04:04Z",
                      "Cpu": 8,
                      "LocalStorageAmount": 2,
                      "ExpiredTime": "2017-12-10T04:04Z",
                      "ResourceGroupId": "rg-bp67acfmxazb4p****",
                      "InternetChargeType": "PayByTraffic",
                      "ZoneId": "cn-hangzhou-g",
                      "Recyclable": true,
                      "ISP": "null",
                      "CreditSpecification": "Standard",
                      "InstanceTypeFamily": "ecs.g5",
                      "OSType": "linux",
                      "NetworkInterfaces": [
                        {
                          "Type": "Primary",
                          "MacAddress": "00:16:3e:32:b4:**",
                          "PrimaryIpAddress": "172.17.**.***",
                          "NetworkInterfaceId": "eni-2zeh9atclduxvf1z****",
                          "PrivateIpSets": [
                            {
                              "PrivateIpAddress": "172.17.**.**",
                              "Primary": true
                            }
                          ],
                          "Ipv6Sets": [
                            {
                              "Ipv6Address": "2408:4321:180:1701:94c7:bc38:3bfa:***"
                            }
                          ]
                        }
                      ],
                      "OperationLocks": [
                        {
                          "LockMsg": "The specified instance is locked due to financial reason.",
                          "LockReason": "Recycling"
                        }
                      ],
                      "Tags": [
                        {
                          "TagValue": "TestValue",
                          "TagKey": "TestKey"
                        }
                      ],
                      "RdmaIpAddress": [
                        "10.10.10.102"
                      ],
                      "SecurityGroupIds": [
                        "sg-bp67acfmxazb4p****"
                      ],
                      "PublicIpAddress": [
                        "121.40.**.**"
                      ],
                      "InnerIpAddress": [
                        "10.170.**.**"
                      ],
                      "VpcAttributes": {
                        "VpcId": "vpc-2zeuphj08tt7q3brd****",
                        "NatIpAddress": "172.17.**.**",
                        "VSwitchId": "vsw-2zeh0r1pabwtg6wcs****",
                        "PrivateIpAddress": [
                          "172.17.**.**"
                        ]
                      },
                      "EipAddress": {
                        "IsSupportUnassociate": true,
                        "InternetChargeType": "PayByTraffic",
                        "IpAddress": "42.112.**.**",
                        "Bandwidth": 5,
                        "AllocationId": "eip-2ze88m67qx5z****"
                      },
                      "HibernationOptions": {
                        "Configured": true
                      },
                      "DedicatedHostAttribute": {
                        "DedicatedHostId": "dh-bp67acfmxazb4p****",
                        "DedicatedHostName": "testDedicatedHostName",
                        "DedicatedHostClusterId": "dc-bp67acfmxazb4h****"
                      },
                      "EcsCapacityReservationAttr": {
                        "CapacityReservationPreference": "cr-bp67acfmxazb4p****",
                        "CapacityReservationId": "cr-bp67acfmxazb4p****"
                      },
                      "DedicatedInstanceAttribute": {
                        "Affinity": "default",
                        "Tenancy": "default"
                      },
                      "CpuOptions": {
                        "Numa": "2",
                        "CoreCount": 2,
                        "ThreadsPerCore": 4
                      },
                      "MetadataOptions": {
                        "HttpEndpoint": "enabled",
                        "HttpPutResponseHopLimit": 0,
                        "HttpTokens": "optional"
                      },
                      "ImageOptions": {
                        "LoginAsNonRoot": true
                      }
                    }
                ],
                "NextToken": "",
                "PageNumber": 1,
                "PageSize": 10,
                "RequestId": "xxxxxx",
                "TotalCount": 1
            },
            "headers": {
                "access-control-allow-headers": "X-Requested-With, X-Sequence, _aop_secret, _aop_signature, x-acs-action, x-acs-version, x-acs-date, Content-Type",
                "access-control-allow-methods": "POST, GET, OPTIONS, PUT, DELETE",
                "access-control-allow-origin": "*",
                "access-control-max-age": "172800",
                "connection": "keep-alive",
                "content-length": "139",
                "content-type": "application/json;charset=utf-8",
                "date": "Fri, 29 Apr 2022 07:14:37 GMT",
                "x-acs-request-id": "xxxxxx",
                "x-acs-trace-id": "xxxxxx"
            }
        }
    }
}
```

For more information about list instances, see [DesctibeInstances](https://next.api.aliyun.com/document/Ecs/2014-05-26/DescribeInstances).

## StartInstances

Launch one or more instances that are in a stopped state.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8000/api/instances/start' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessKeyId": "{{yourAccessKeyId}}",
    "accessKeySecret": "{{yourAccessKeySecret}}",
    "endpoint": "ecs.cn-beijing.aliyuncs.com",
    "regionId": "cn-beijing",
    "instanceIds": ["i-bp67acfmxazb4p****"]
}'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "body": {
                "RequestId": "473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E",
                "InstanceResponses": [
                    {
                    "Code": "200",
                    "Message": "success",
                    "InstanceId": "i-bp67acfmxazb4p****",
                    "CurrentStatus": "Starting",
                    "PreviousStatus": "Stopped"
                    }
                ]
            },
            "headers": {
                "access-control-allow-headers": "X-Requested-With, X-Sequence, _aop_secret, _aop_signature, x-acs-action, x-acs-version, x-acs-date, Content-Type",
                "access-control-allow-methods": "POST, GET, OPTIONS, PUT, DELETE",
                "access-control-allow-origin": "*",
                "access-control-max-age": "172800",
                "connection": "keep-alive",
                "content-type": "application/json;charset=utf-8",
                "date": "Fri, 29 Apr 2022 07:11:27 GMT",
                "vary": "Accept-Encoding",
                "x-acs-request-id": "xxxxx",
                "x-acs-trace-id": "xxxxx"
            }
        }
    }
}
```

For more information about start instances, see [StartInstances](https://next.api.aliyun.com/document/Ecs/2014-05-26/StartInstances).

## StopInstances

After success, the ECS instance first enters the Stopping state, and finally enters the Stopped state.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8000/api/instances/stop' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessKeyId": "{{yourAccessKeyId}}",
    "accessKeySecret": "{{yourAccessKeySecret}}",
    "endpoint": "ecs.cn-beijing.aliyuncs.com",
    "regionId": "cn-beijing",
    "instanceIds": ["i-bp67acfmxazb4p****"]
}'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "body": {
                "RequestId": "1C488B66-B819-4D14-8711-C4EAAA13AC01",
                "InstanceResponses": [
                    {
                    "Code": "200",
                    "Message": "success",
                    "InstanceId": "i-bp67acfmxazb4p****",
                    "CurrentStatus": "Stopping",
                    "PreviousStatus": "Running"
                    }
                ]
            },
            "headers": {
                "access-control-allow-headers": "X-Requested-With, X-Sequence, _aop_secret, _aop_signature, x-acs-action, x-acs-version, x-acs-date, Content-Type",
                "access-control-allow-methods": "POST, GET, OPTIONS, PUT, DELETE",
                "access-control-allow-origin": "*",
                "access-control-max-age": "172800",
                "connection": "keep-alive",
                "content-type": "application/json;charset=utf-8",
                "date": "Fri, 29 Apr 2022 07:11:27 GMT",
                "vary": "Accept-Encoding",
                "x-acs-request-id": "xxxxx",
                "x-acs-trace-id": "xxxxx"
            }
        }
    }
}
```

For more information about stop instances, see [StopInstances](https://next.api.aliyun.com/document/Ecs/2014-05-26/StopInstances).

## RebootInstances

The ECS instance first enters the Stopping state and finally enters the Running state.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8000/api/instances/reboot' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessKeyId": "{{yourAccessKeyId}}",
    "accessKeySecret": "{{yourAccessKeySecret}}",
    "endpoint": "ecs.cn-beijing.aliyuncs.com",
    "regionId": "cn-beijing",
    "instanceIds": ["i-bp67acfmxazb4p****"]
}'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "body": {
                "RequestId": "473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E",
                "InstanceResponses": [
                    {
                        "Code": "200",
                        "Message": "success",
                        "InstanceId": "i-bp1g6zv0ce8oghu7****",
                        "CurrentStatus": "Stopping",
                        "PreviousStatus": "Running"
                    }
                ]
            },
            "headers": {
                "access-control-allow-headers": "X-Requested-With, X-Sequence, _aop_secret, _aop_signature, x-acs-action, x-acs-version, x-acs-date, Content-Type",
                "access-control-allow-methods": "POST, GET, OPTIONS, PUT, DELETE",
                "access-control-allow-origin": "*",
                "access-control-max-age": "172800",
                "connection": "keep-alive",
                "content-type": "application/json;charset=utf-8",
                "date": "Fri, 29 Apr 2022 07:11:27 GMT",
                "vary": "Accept-Encoding",
                "x-acs-request-id": "xxxxx",
                "x-acs-trace-id": "xxxxx"
            }
        }
    }
}
```

For more information about reboot instances, see [RebootInstances](https://next.api.aliyun.com/document/Ecs/2014-05-26/RebootInstances).

## TerminateInstances

Release one or more pay-as-you-go instances or expired annual and monthly subscription instances

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8000/api/instances/terminate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessKeyId": "{{yourAccessKeyId}}",
    "accessKeySecret": "{{yourAccessKeySecret}}",
    "endpoint": "ecs.cn-beijing.aliyuncs.com",
    "regionId": "cn-beijing",
    "instanceIds": ["i-bp67acfmxazb4p****"]
}'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "body": {
                "RequestId": "7B7813C6-57BF-41XX-B12B-F172F65A6046"
            },
            "headers": {
                "access-control-allow-headers": "X-Requested-With, X-Sequence, _aop_secret, _aop_signature, x-acs-action, x-acs-version, x-acs-date, Content-Type",
                "access-control-allow-methods": "POST, GET, OPTIONS, PUT, DELETE",
                "access-control-allow-origin": "*",
                "access-control-max-age": "172800",
                "connection": "keep-alive",
                "content-type": "application/json;charset=utf-8",
                "date": "Fri, 29 Apr 2022 07:11:27 GMT",
                "vary": "Accept-Encoding",
                "x-acs-request-id": "xxxxx",
                "x-acs-trace-id": "xxxxx"
            }
        }
    }
}
```

For more information about terminate instances, see [DeleteInstances](https://next.api.aliyun.com/document/Ecs/2014-05-26/DeleteInstances).