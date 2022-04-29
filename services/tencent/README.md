# Tencent cloud api

## Quick Start

### Request a Security Credential

See the [Console Cloud API](https://console.cloud.tencent.com/capi) Keys page for key information.

1. `SecretId`: used to identify the identity of the API caller, which can be simply analogized to the username.
2. `SecretKey`: used to verify the identity of the API caller, which can be simply analogized to a password.
3. Users must keep the security credentials strictly to avoid leakage, otherwise the property safety will be endangered. If it has been leaked, please disable the security credential immediately.

For more information about credential, see [Signature Method](https://cloud.tencent.com/document/api/213/30654).

## List all regions

List all tencent cloud available regions.

### Try it out

```
curl --location --request GET 'http://127.0.0.1:8001/api/regions?secretId={{yourSecretId}}&secretKey={{yourSecretKey}}'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "Response": {
                "RegionSet": [
                    {
                        "Region": "ap-guangzhou",
                        "RegionName": "华南地区(广州)",
                        "RegionState": "AVAILABLE"
                    },
                    ...
                    {
                        "Region": "na-toronto",
                        "RegionName": "北美地区(多伦多)",
                        "RegionState": "AVAILABLE"
                    }
                ],
                "RequestId": "6ef60bec-0242-43af-bb20-270359fb54a7",
                "TotalCount": 19
            }
        }
    }
}
```

For more information about regions, see [DescribeRegions](https://cloud.tencent.com/document/api/213/15708).

## List all images

List all tencent cloud available images.

### Try it out

```shell
curl --location --request GET 'http://127.0.0.1:8001/api/images?secretId={{yourSecretId}}&secretKey={{yourSecretKey}}&region={{listAllRegionsApiFindRegion}}'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "Response": {
                "ImageSet": [
                    {
                        "Architecture": "arm",
                        "ImageDescription": "TencentOS Server 3.1 for ARM64 (TK4)",
                        "ImageId": "img-9a0eg8mt",
                        "ImageName": "TencentOS Server 3.1 for ARM64 (TK4)",
                        "ImageSize": 50,
                        "ImageSource": "OFFICIAL",
                        "ImageState": "NORMAL",
                        "ImageType": "PUBLIC_IMAGE",
                        "IsSupportCloudinit": true,
                        "OsName": "TencentOS Server 3.1 for ARM64 (TK4)",
                        "Platform": "TencentOS"
                    },
                    ...
                    {
                        "Architecture": "x86_64",
                        "ImageDescription": "CentOS 7.4 64位",
                        "ImageId": "img-8toqc6s3",
                        "ImageName": "CentOS 7.4 64位",
                        "ImageSize": 50,
                        "ImageSource": "OFFICIAL",
                        "ImageState": "NORMAL",
                        "ImageType": "PUBLIC_IMAGE",
                        "IsSupportCloudinit": true,
                        "OsName": "CentOS 7.4 64位",
                        "Platform": "CentOS"
                    }
                ],
                "RequestId": "6ef60bec-0242-43af-bb20-270359fb54a7",
                "TotalCount": 62
            }
        }
    }
}
```

For more information about images, see [DescribeImages](https://cloud.tencent.com/document/api/213/15715).

## List all instance types

List all tencent cloud available instance types.

### Try it out

```shell
curl --location --request GET 'http://127.0.0.1:8001/api/instance-types?secretId={{yourSecretId}}&secretKey={{yourSecretKey}}&region={{listAllRegionsApiFindRegion}}''

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "Response": {
                "InstanceTypeConfigSet": [
                    {
                        "CPU": 1,
                        "FPGA": 0,
                        "GPU": 0,
                        "InstanceFamily": "SA2",
                        "InstanceType": "SA2.SMALL1",
                        "Memory": 1,
                        "Zone": "ap-guangzhou-5"
                    },
                    ...
                    
                    {
                        "CPU": 232,
                        "FPGA": 0,
                        "GPU": 0,
                        "InstanceFamily": "SA3",
                        "InstanceType": "SA3.58XLARGE940",
                        "Memory": 940,
                        "Zone": "ap-guangzhou-5"
                    }
                ],
                "RequestId": "6ef60bec-0242-43af-bb20-270359fb54a7"
            }
        }
    }
}
```

For more information about instance types, see [DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)

## CreateInstances

Used to create one or more instances of the specified configuration.

### Try it out

```
curl --location --request POST 'http://127.0.0.1:8001/api/instances' \
--header 'Content-Type: application/json' \
--data-raw '{
    "secretId": "{{yourSecretId}}",
    "secretKey": "{{yourSecretKey}}",
    "region": "ap-guangzhou",
    "imageId": "img-9qabwvbn",
    "instanceType": "I1.MEDIUM4",
    "instanceName": "test",
    "systemDisk": {
        "disk_size": 80
    },
    "instanceCount": 2,
    "dryRun": true
}'

The response should like
{
	"Response": {
		"InstanceIdSet": [
			"ins-0s7wsh5x",
			"ins-03lw8hok"
		],
		"RequestId": "3c14def19-cfes-470e-b241-90787u6jf5uj"
	}
}
```

For more information about create instances, see [RunInstances](https://cloud.tencent.com/document/api/213/15730).

## ListInstances

Query the details of one or more instances.

### Try it out

```shell
curl --location --request GET 'http://127.0.0.1:8001/api/instances?secretId={{yourSecretId}}&secretKey={{yourSecretKey}}&region=ap-shanghai&offset=0&limit=20&instanceName={{yourInstanceName}}'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
        	"Response": {
                "InstanceSet": [
                    {
                        "CPU": 1,
                        "CamRoleName": "xx",
                        "CreatedTime": "2020-09-22T00:00:00+00:00",
                        "DataDisks": [
                            {
                                "CdcId": "xx",
                                "DeleteWithInstance": true,
                                "DiskId": "xx",
                                "DiskSize": 0,
                                "DiskType": "xx",
                                "Encrypt": true,
                                "KmsKeyId": "xx",
                                "SnapshotId": "xx",
                                "ThroughputPerformance": 0
                            }
                        ],
                        "DedicatedClusterId": "xx",
                        "DisasterRecoverGroupId": "xx",
                        "DisasterRecoverGroupIds": [
                            "xx"
                        ],
                        "ExpiredTime": "2020-09-22T00:00:00+00:00",
                        "HpcClusterId": "xx",
                        "IPv6Addresses": [
                            "xx"
                        ],
                        "ImageId": "xx",
                        "InstanceChargeType": "xx",
                        "InstanceId": "xx",
                        "InstanceName": "xx",
                        "InstanceState": "xx",
                        "InstanceType": "xx",
                        "InternetAccessible": {
                            "BandwidthPackageId": "xx",
                            "InternetChargeType": "xx",
                            "InternetMaxBandwidthOut": 1,
                            "PublicIpAssigned": true
                        },
                        "IsolatedSource": "xx",
                        "LatestOperation": "xx",
                        "LatestOperationRequestId": "xx",
                        "LatestOperationState": "xx",
                        "LoginSettings": {
                            "KeepImageLogin": "xx",
                            "KeyIds": [
                                "xx"
                            ],
                            "Password": "xx"
                        },
                        "MaintainExpireTime": "xx",
                        "Memory": 1,
                        "OsName": "xx",
                        "Placement": {
                            "HostId": "xx",
                            "HostIds": [
                                "xx"
                            ],
                            "HostIps": [
                                "xx"
                            ],
                            "ProjectId": 1174660,
                            "Zone": "xx"
                        },
                        "PrivateIpAddresses": [
                            "172.16.32.78"
                        ],
                        "PublicIpAddresses": [
                            "123.207.11.190"
                        ],
                        "RdmaIpAddresses": [
                            "xx"
                        ],
                        "RenewFlag": "xx",
                        "RestrictState": "xx",
                        "SecurityGroupIds": [
                            "sg-p1ezv4wz"
                        ],
                        "StopChargingMode": "xx",
                        "SystemDisk": {
                            "CdcId": "xx",
                            "DiskId": "xx",
                            "DiskSize": 50,
                            "DiskType": "xx"
                        },
                        "Tags": [
                            {
                                "Key": "xx",
                                "Value": "xx"
                            }
                        ],
                        "Uuid": "xx",
                        "VirtualPrivateCloud": {
                            "AsVpcGateway": false,
                            "Ipv6AddressCount": 1,
                            "PrivateIpAddresses": [
                                "xx"
                            ],
                            "SubnetId": "xx",
                            "VpcId": "xx"
                        }
                    }
                ],
                "RequestId": "xx",
                "TotalCount": 2
            }
        }
    }
}
```

For more information about list instances, see [DesctibeInstances](https://cloud.tencent.com/document/api/213/15728).

## StartInstances

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

For more information about start instances, see [StartInstances](https://cloud.tencent.com/document/api/213/15735).

## StopInstances

Used to close one or more instances.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8001/api/instances/stop' \
--header 'Content-Type: application/json' \
--data-raw '{
    "secretId": "",
    "secretKey": "",
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

For more information about stop instances, see [StopInstances](https://cloud.tencent.com/document/api/213/15743).

## RebootInstances

This api is used to resboot the instance.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8001/api/instances/reboot' \
--header 'Content-Type: application/json' \
--data-raw '{
    "secretId": "",
    "secretKey": "",
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

For more information about reboot instances, see [RebootInstances](https://cloud.tencent.com/document/api/213/15742).

## TerminateInstances

This api is used to terminate the instance.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8001/api/instances/terminate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "secretId": "",
    "secretKey": "",
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

For more information about terminate instances, see [TerminateInstances](https://cloud.tencent.com/document/api/213/15723).