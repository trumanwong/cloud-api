# AWS api

## Quick Start

### To get your access key ID and secret access key.[ ](https://aws.github.io/aws-sdk-go-v2/docs/getting-started/#to-get-your-access-key-id-and-secret-access-key)

1. Open the [IAM console](https://console.aws.amazon.com/iam/home)
2. On the navigation menu, choose **Users**.
3. Choose your IAM user name (not the check box).
4. Open the **Security credentials** tab, and then choose **Create access key**.
5. To see the new access key, choose __Show__. Your credentials resemble the following:
    - Access key ID: `AKIAIOSFODNN7EXAMPLE`
    - Secret access key: `wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY`
6. To download the key pair, choose **Download .csv file**. Store the keys in a secure location.

For more information about credential, see [Get your AWS access keys](https://aws.github.io/aws-sdk-go-v2/docs/getting-started/#get-your-aws-access-keys).

## List all regions

Describes the Regions that are enabled for your account, or all Regions.

### Try it out

```
curl --location --request GET 'http://127.0.0.1:8003/api/regions?accessKeyId={{yourAccessKeyId}}&secretKey={{yourSecretKey}}'
```

For more information about regions, see [DescribeRegions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeRegions.html).

## List all images

Describes the specified images (AMIs, AKIs, and ARIs) available to you or all of the images available to you.

### Try it out

```shell
curl --location --request GET 'http://127.0.0.1:8003/api/images?accessKeyId={{yourAccessKeyId}}&secretKey={{yourSecretKey}}'
```

For more information about images, see [DescribeImages](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeImages.html).

## List all instance types

Describes the details of the instance types that are offered in a location.

### Try it out

```shell
curl --location --request GET 'http://127.0.0.1:8003/api/instance-types?accessKeyId={{yourAccessKeyId}}&secretKey={{yourSecretKey}}&region={{selectRegion}}'
```

For more information about instance types, see [DescribeInstanceTypes](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstanceTypes.html).

## CreateInstances

Launches the specified number of instances using an AMI for which you have permissions.

### Try it out

```
curl --location --request POST 'http://127.0.0.1:8003/api/instances' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessKeyId": "{{yourAccessKeyId}}",
    "secretKey": "{{yourSecretKey}}",
    "region": "{{selectRegion}}",
    "imageId": "{{selectImageId}}",
    "instanceType": "{{selectInstanceType}}",
    "minCount": 1,
    "maxCount": 3,
    "dryRun": true
}'
```

For more information about create instances, see [RunInstances](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RunInstances.html).

## ListInstances

Describes the specified instances or all instances.

### Try it out

```shell
curl --location --request GET 'http://127.0.0.1:8003/api/instances?accessKeyId={{yourAccessKeyId}}&secretKey={{yourSecretKey}}&region={{selectRegion}}'
```

For more information about list instances, see [DescribeInstances](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstances.html).

## StartInstances

Starts an Amazon EBS-backed instance that you've previously stopped.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8003/api/instances/start' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessKeyId": "{{yourAccessKeyId}}",
    "secretKey": "{{yourSecretKey}}",
    "region": "{{selectRegion}}",
    "instanceIds": ["i-1234567890abcdef0"],
    "dryRun": true
}'
```

For more information about start instances, see [StartInstances](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_StartInstances.html).

## StopInstances

Stops an Amazon EBS-backed instance.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8003/api/instances/stop' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessKeyId": "{{yourAccessKeyId}}",
    "secretKey": "{{yourSecretKey}}",
    "region": "{{selectRegion}}",
    "instanceIds": ["i-1234567890abcdef0"],
    "dryRun": true
}'
```

For more information about stop instances, see [StopInstances](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_StopInstances.html).

## RebootInstances

Requests a reboot of the specified instances.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8003/api/instances/reboot' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessKeyId": "{{yourAccessKeyId}}",
    "secretKey": "{{yourSecretKey}}",
    "region": "{{selectRegion}}",
    "instanceIds": ["i-1234567890abcdef0"],
    "dryRun": true
}'
```

For more information about stop instances, see [RebootInstances](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RebootInstances.html).

## TerminateInstances

Shuts down the specified instances.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8003/api/instances/terminate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessKeyId": "{{yourAccessKeyId}}",
    "secretKey": "{{yourSecretKey}}",
    "region": "{{selectRegion}}",
    "instanceIds": ["i-1234567890abcdef0"],
    "dryRun": true
}'
```

For more information about stop instances, see [TerminateInstances](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TerminateInstances.html).