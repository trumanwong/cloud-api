# Vultr cloud api

## Quick Start

### Request a Security Credential

Before you begin, you need to sign up for an Vultr account and retrieve your [customer portal](https://my.vultr.com/settings/#settingsapi).

## List all regions

List all Regions at Vultr.

### Try it out

```
curl --location -g --request GET 'http://127.0.0.1:8005/api/regions?accessToken={{yourAccessToken}}'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "meta": {
                "Links": {
                    "next": "",
                    "prev": ""
                },
                "total": 25
            },
            "regions": [
                {
                    "city": "Amsterdam",
                    "continent": "Europe",
                    "country": "NL",
                    "id": "ams",
                    "options": [
                        "ddos_protection",
                        "block_storage_storage_opt",
                        "block_storage_high_perf",
                        "load_balancers",
                        "kubernetes"
                    ]
                },
                ...
                {
                    "city": "Toronto",
                    "continent": "North America",
                    "country": "CA",
                    "id": "yto",
                    "options": [
                        "ddos_protection",
                        "block_storage_storage_opt",
                        "load_balancers",
                        "kubernetes"
                    ]
                }
            ]
        }
    }
}
```

For more information about regions, see [List Regions](https://www.vultr.com/api/#operation/list-regions).

## List all images

List the OS images available for installation at Vultr.

### Try it out

```shell
curl --location -g --request GET 'http://127.0.0.1:8005/api/images??accessToken={{yourAccessToken}}'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "images": [
                {
                    "arch": "x64",
                    "family": "windows",
                    "id": 124,
                    "name": "Windows 2012 R2 Standard x64"
                },
                ...
                {
                    "arch": "x64",
                    "family": "openbsd",
                    "id": 1797,
                    "name": "OpenBSD 7.1 x64"
                }
            ],
            "meta": {
                "Links": {
                    "next": "",
                    "prev": ""
                },
                "total": 34
            }
        }
    }
}
```

For more information about images, see [List OS](https://www.vultr.com/api/#tag/os).

## List all plans

Get a list of all VPS plans at Vultr.

### Try it out

```shell
curl --location --request GET 'http://127.0.0.1:8001/api/instance-types?secretId={{yourSecretId}}&secretKey={{yourSecretKey}}&region={{listAllRegionsApiFindRegion}}''

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "meta": {
                "Links": {
                    "next": "bmV4dF9fdm9jLW0tMzJjLTI1NmdiLTMyMDBzLWFtZA==",
                    "prev": ""
                },
                "total": 107
            },
            "plans": [
                {
                    "bandwidth": 1024,
                    "disk": 25,
                    "disk_count": 1,
                    "id": "vc2-1c-1gb",
                    "locations": [
                        "ewr",
                        "ord",
                        "dfw",
                        "sea",
                        "lax",
                        "atl",
                        "ams",
                        "lhr",
                        "fra",
                        "sjc",
                        "syd",
                        "yto",
                        "cdg",
                        "nrt",
                        "waw",
                        "mad",
                        "icn",
                        "mia",
                        "sgp",
                        "sto",
                        "mex",
                        "mel"
                    ],
                    "monthly_cost": 5,
                    "ram": 1024,
                    "type": "vc2",
                    "vcpu_count": 1
                },
                ...
                {
                    "bandwidth": 12288,
                    "disk": 3200,
                    "disk_count": 1,
                    "id": "voc-m-32c-256gb-3200s-amd",
                    "locations": [
                        "ewr",
                        "ord",
                        "dfw",
                        "sea",
                        "lax",
                        "atl",
                        "ams",
                        "lhr",
                        "fra",
                        "sjc",
                        "syd",
                        "yto",
                        "cdg",
                        "nrt",
                        "waw",
                        "icn",
                        "mia",
                        "sgp",
                        "sto",
                        "mel",
                        "hnl",
                        "bom"
                    ],
                    "monthly_cost": 1565,
                    "ram": 262144,
                    "type": "voc",
                    "vcpu_count": 32
                }
            ]
        }
    }
}
```

For more information about plans, see [List Plans](https://www.vultr.com/api/#operation/list-plans)

## Create Instance

Create a new VPS Instance in a region with the desired plan.

### Try it out

```
curl --location --request POST 'http://127.0.0.1:8005/api/instance' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessToken": "{{yourAccessToken}}",
    "region": "ams",
    "plan": "vc2-1c-1gb",
    "image_id": "124",
    "hostname": "test"
}'
```

For more information about create instances, see [Create Instance](https://www.vultr.com/api/#operation/create-instance).

## List Instances

Query the details of one or more instances.

### Try it out

```shell
curl --location -g --request GET 'http://127.0.0.1:8005/api/instances?accessToken={{yourAccessToken}}'

The response should like
{
    "data": {
        "@type": "type.googleapis.com/google.protobuf.Struct",
        "value": {
            "meta": {
                "Links": {
                    "next": "WxYzExampleNext",
                    "prev": ""
                },
                "total": 3
            },
            "instances": [
                {
                    "id": "cb676a46-66fd-4dfb-b839-443f2e6c0b60",
                    "os": "CentOS SELinux 8 x64",
                    "ram": 2048,
                    "disk": 55,
                    "main_ip": "192.0.2.123",
                    "vcpu_count": 1,
                    "region": "atl",
                    "plan": "vc2-6c-16gb",
                    "date_created": "2020-10-10T01:56:20+00:00",
                    "status": "active",
                    "allowed_bandwidth": 2000,
                    "netmask_v4": "255.255.252.0",
                    "gateway_v4": "192.0.2.1",
                    "power_status": "running",
                    "server_status": "ok",
                    "v6_network": "2001:0db8:1112:18fb::",
                    "v6_main_ip": "2001:0db8:1112:18fb:0200:00ff:fe00:0000",
                    "v6_network_size": 64,
                    "label": "Example Instance",
                    "internal_ip": "",
                    "kvm": "https://my.vultr.com/subs/vps/novnc/api.php?data=00example11223344",
                    "hostname": "my_hostname",
                    "os_id": 215,
                    "app_id": 0,
                    "image_id": "",
                    "firewall_group_id": "",
                    "features": [],
                    "tags": []
                    }
                ],
        }
    }
}
```

For more information about list instances, see [List Instances](https://www.vultr.com/api/#operation/list-instances).

## Start Instance

Start an Instance.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8005/api/instance/start' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessToken": "{{yourAccessToken}}",
    "instanceId": "xxxxx"
}'
```

For more information about start instance, see [Start Instance](https://www.vultr.com/api/#operation/start-instance).

## Halt Instance

Halt an Instance.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8005/api/instances/stop' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessToken": "{{yourAccessToken}}",
    "instanceId": "xxxxx"
}'
```

For more information about stop instance, see [Halt Instance](https://www.vultr.com/api/#operation/halt-instance).

## Reboot Instance

Reboot an Instance.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8005/api/instance/reboot' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessToken": "{{yourAccessToken}}",
    "instanceId": "xxxxx"
}'
```

For more information about reboot instance, see [Reboot Instance](https://www.vultr.com/api/#operation/reboot-instance).

## Delete Instance

Delete an Instance.

### Try it out

```
curl --location --request PUT 'http://127.0.0.1:8005/api/instance/terminate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "accessToken": "{{yourAccessToken}}",
    "instanceId": "xxxxx"
}'
```

For more information about delete instance, see [Delete Instance](https://www.vultr.com/api/#operation/delete-instance).