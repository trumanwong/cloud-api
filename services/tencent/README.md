# Tencent cloud api

## Get Started

### Request a Security Credential

See the [Console Cloud API](https://console.cloud.tencent.com/capi) Keys page for key information.

1. `SecretId`: used to identify the identity of the API caller, which can be simply analogized to the username.
2. `SecretKey`: used to verify the identity of the API caller, which can be simply analogized to a password.
3. Users must keep the security credentials strictly to avoid leakage, otherwise the property safety will be endangered. If it has been leaked, please disable the security credential immediately.

For more details, please refer to [Signature Method](https://cloud.tencent.com/document/api/213/30654).

## List all regions

List all tencent cloud available regions.

```shell
curl --location --request GET 'http://127.0.0.1:8001/api/regions?secretId={{yourSecretId}}&secretKey={{yourSecretKey}}'
```

examples:

```
{
    "totalCount": "19",
    "regions": [
        {
            "region": "ap-guangzhou",
            "regionName": "华南地区(广州)",
            "regionState": "AVAILABLE"
        },
        ...
        {
            "region": "na-toronto",
            "regionName": "北美地区(多伦多)",
            "regionState": "AVAILABLE"
        }
    ],
    "requestId": "6ef60bec-0242-43af-bb20-270359fb54a7"
}
```

For more details, please refer to [DescribeRegions](https://cloud.tencent.com/document/api/213/15708).

## List all images

List all tencent cloud available images.

```shell
curl --location --request GET 'http://127.0.0.1:8001/api/images?secretId={{yourSecretId}}&secretKey={{yourSecretKey}}&region={{listAllRegionsApiFindRegion}}'
```

examples:

```
{
    "images": [
        {
            "imageId": "img-9a0eg8mt",
            "osName": "TencentOS Server 3.1 for ARM64 (TK4)",
            "imageType": "PUBLIC_IMAGE",
            "imageName": "TencentOS Server 3.1 for ARM64 (TK4)",
            "imageDescription": "TencentOS Server 3.1 for ARM64 (TK4)",
            "imageSize": "50",
            "architecture": "arm",
            "imageState": "NORMAL",
            "platform": "TencentOS",
            "imageSource": "OFFICIAL"
        },
        ...
        {
            "imageId": "img-8toqc6s3",
            "osName": "CentOS 7.4 64位",
            "imageType": "PUBLIC_IMAGE",
            "imageName": "CentOS 7.4 64位",
            "imageDescription": "CentOS 7.4 64位",
            "imageSize": "50",
            "architecture": "x86_64",
            "imageState": "NORMAL",
            "platform": "CentOS",
            "imageSource": "OFFICIAL"
        }
    ],
    "totalCount": "62",
    "requestId": "6ef60bec-0242-43af-bb20-270359fb54a7"
}
```

For more details, please refer to [DescribeImages](https://cloud.tencent.com/document/api/213/15715).

## List all instance types

List all tencent cloud available instance types.

```shell
curl --location --request GET 'http://127.0.0.1:8001/api/instance-types?secretId={{yourSecretId}}&secretKey={{yourSecretKey}}&region={{listAllRegionsApiFindRegion}}''
```

examples:

```
{
    "instanceTypes": [
        {
            "zone": "ap-guangzhou-5",
            "instanceType": "SA2.SMALL1",
            "instanceFamily": "SA2",
            "gpu": "0",
            "cpu": "1",
            "memory": "1",
            "fpga": "0"
        },
        ...
        {
            "zone": "ap-guangzhou-5",
            "instanceType": "SA3.58XLARGE940",
            "instanceFamily": "SA3",
            "gpu": "0",
            "cpu": "232",
            "memory": "940",
            "fpga": "0"
        }
    ],
    "requestId": "6ef60bec-0242-43af-bb20-270359fb54a7"
}
```

For more details, please refer to [DescribeInstanceTypeConfigs](https://cloud.tencent.com/document/api/213/15749)

## CreateInstance

