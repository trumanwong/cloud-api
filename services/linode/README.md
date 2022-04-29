# Linode cloud api

## Quick Start

### Request a Security Credential

Before you begin, you need to sign up for an Linode account and retrieve your [Personal Access Token](https://cloud.linode.com/profile/tokens) or the [Create Personal Access Token](https://www.linode.com/docs/api/profile/#personal-access-token-create) endpoint.

## List all regions

Lists the Regions available for Linode services. Not all services are guaranteed to be available in all Regions.

### Try it out

```
curl --location -g --request GET 'http://127.0.0.1:8006/api/regions?accessToken={{yourAccessToken}}'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "Regions": [
                {
                    "capabilities": [
                        "Linodes",
                        "NodeBalancers",
                        "Block Storage",
                        "GPU Linodes",
                        "Kubernetes",
                        "Cloud Firewall",
                        "Vlans",
                        "Block Storage Migrations",
                        "Managed Databases"
                    ],
                    "country": "in",
                    "id": "ap-west",
                    "resolvers": {
                        "ipv4": "172.105.34.5,172.105.35.5,172.105.36.5,172.105.37.5,172.105.38.5,172.105.39.5,172.105.40.5,172.105.41.5,172.105.42.5,172.105.43.5",
                        "ipv6": "2400:8904::f03c:91ff:fea5:659,2400:8904::f03c:91ff:fea5:9282,2400:8904::f03c:91ff:fea5:b9b3,2400:8904::f03c:91ff:fea5:925a,2400:8904::f03c:91ff:fea5:22cb,2400:8904::f03c:91ff:fea5:227a,2400:8904::f03c:91ff:fea5:924c,2400:8904::f03c:91ff:fea5:f7e2,2400:8904::f03c:91ff:fea5:2205,2400:8904::f03c:91ff:fea5:9207"
                    },
                    "status": "ok"
                },
                ...
                {
                    "capabilities": [
                        "Linodes",
                        "NodeBalancers",
                        "Block Storage",
                        "Kubernetes",
                        "Cloud Firewall",
                        "Block Storage Migrations",
                        "Managed Databases"
                    ],
                    "country": "jp",
                    "id": "ap-northeast",
                    "resolvers": {
                        "ipv4": "139.162.66.5,139.162.67.5,139.162.68.5,139.162.69.5,139.162.70.5,139.162.71.5,139.162.72.5,139.162.73.5,139.162.74.5,139.162.75.5",
                        "ipv6": "2400:8902::3,2400:8902::6,2400:8902::c,2400:8902::4,2400:8902::2,2400:8902::8,2400:8902::7,2400:8902::5,2400:8902::b,2400:8902::9"
                    },
                    "status": "ok"
                }
            ]
        }
    }
}
```

For more information about regions, see [List Regions](https://www.linode.com/docs/api/regions/#regions-list).

## List all images

Returns a paginated list of Images.

### Try it out

```shell
curl --location -g --request GET 'http://127.0.0.1:8006/api/images?accessToken={{yourAccessToken}}'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "images": [
                {
                    "created_by": "linode",
                    "deprecated": false,
                    "description": "",
                    "id": "linode/almalinux8",
                    "is_public": true,
                    "label": "AlmaLinux 8",
                    "size": 1700,
                    "status": "available",
                    "type": "manual",
                    "vendor": "AlmaLinux"
                },
                ...
                {
                    "created_by": "linode",
                    "deprecated": true,
                    "description": "",
                    "id": "linode/ubuntu21.04",
                    "is_public": true,
                    "label": "Ubuntu 21.04",
                    "size": 3500,
                    "status": "available",
                    "type": "manual",
                    "vendor": "Ubuntu"
                }
            ]
        }
    }
}
```

For more information about images, see [List Images](https://www.linode.com/docs/api/images/#images-list).

## List all plans

Returns collection of Linode Types, including pricing and specifications for each Type.

### Try it out

```shell
curl --location -g --request GET 'http://127.0.0.1:8006/api/instance-types?accessToken={{yourAccessToken}}'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "InstanceTypes": [
                {
                    "addons": {
                        "backups": {
                            "price": {
                                "hourly": 0.003,
                                "monthly": 2
                            }
                        }
                    },
                    "class": "nanode",
                    "disk": 25600,
                    "id": "g6-nanode-1",
                    "label": "Nanode 1GB",
                    "memory": 1024,
                    "network_out": 1000,
                    "price": {
                        "hourly": 0.0075,
                        "monthly": 5
                    },
                    "transfer": 1000,
                    "vcpus": 1
                },
                ...
                {
                    "addons": {
                        "backups": {
                            "price": {
                                "hourly": 0.24,
                                "monthly": 160
                            }
                        }
                    },
                    "class": "gpu",
                    "disk": 2621440,
                    "id": "g1-gpu-rtx6000-4",
                    "label": "Dedicated 128GB + RTX6000 GPU x4",
                    "memory": 131072,
                    "network_out": 10000,
                    "price": {
                        "hourly": 6,
                        "monthly": 4000
                    },
                    "transfer": 20000,
                    "vcpus": 24
                }
            ]
        }
    }
}
```

For more information about plans, see [List Linode Types](https://www.linode.com/docs/api/linode-types/#types-list)

## Create Instance

Creates a Linode Instance on your Account.

### Try it out

```
curl --location --request POST 'http://127.0.0.1:8005/api/instance' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessToken": "{{yourAccessToken}}",
    "region": "ap-northeast",
    "instanceType": "g1-gpu-rtx6000-4",
    "imageId": "linode/ubuntu21.04",
    "label": "test",
    "password": "test"
}'
```

For more information about create instances, see [Create Instance](https://www.linode.com/docs/api/linode-instances/#linode-create).

## List Instances

Returns a paginated list of Linodes you have permission to view.

### Try it out

```shell
curl --location -g --request GET 'http://127.0.0.1:8006/api/instances?accessToken={{yourAccessToken}}'
```

For more information about list instances, see [List Instances](https://www.linode.com/docs/api/linode-instances/#linodes-list).

## Start Instance

Boots a Linode you have permission to modify.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8006/api/instance/start' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessToken": "{{yourAccessToken}}",
    "instanceId": "xxxxx"
}'
```

For more information about start instance, see [Boot Instance](https://www.linode.com/docs/api/linode-instances/#linode-boot).

## Halt Instance

Shuts down a Linode you have permission to modify.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8006/api/instances/stop' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessToken": "{{yourAccessToken}}",
    "instanceId": "xxxxx"
}'
```

For more information about stop instance, see [Linode Shut Down](https://www.linode.com/docs/api/linode-instances/#linode-shut-down).

## Reboot Instance

Reboots a Linode you have permission to modify.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8006/api/instance/reboot' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessToken": "{{yourAccessToken}}",
    "instanceId": "xxxxx"
}'
```

For more information about reboot instance, see [Linode Reboot](https://www.linode.com/docs/api/linode-instances/#linode-reboot).

## Delete Instance

Deletes a Linode you have permission to `read_write`.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8006/api/instance/terminate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessToken": "{{yourAccessToken}}",
    "instanceId": "xxxxx"
}'
```

For more information about delete instance, see [Linode Delete](https://www.linode.com/docs/api/linode-instances/#linode-delete).