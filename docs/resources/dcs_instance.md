---
subcategory: "Distributed Cache Service"
---

# huaweicloud\_dcs\_instance

Manages a DCS instance in the huaweicloud DCS Service.
This is an alternative to `huaweicloud_dcs_instance_v1`

## Example Usage

### DCS instance for Redis 3.0

```hcl
data "huaweicloud_dcs_az" "az_1" {
  code = "cn-north-1a"
}

resource "huaweicloud_networking_secgroup" "secgroup_1" {
  name        = "secgroup_1"
  description = "secgroup_1"
}
resource "huaweicloud_vpc" "vpc_1" {
  name = "terraform_provider_vpc1"
  cidr = "192.168.0.0/16"
}
resource "huaweicloud_vpc_subnet" "subnet_1" {
  name       = "huaweicloud_subnet"
  cidr       = "192.168.0.0/16"
  gateway_ip = "192.168.0.1"
  vpc_id     = huaweicloud_vpc.vpc_1.id
}

resource "huaweicloud_dcs_instance" "instance_1" {
  name              = "test_dcs_instance"
  engine            = "Redis"
  engine_version    = "3.0"
  password          = "Huawei_test"
  capacity          = 2
  vpc_id            = huaweicloud_vpc.vpc_1.id
  subnet_id         = huaweicloud_vpc_subnet.subnet_1.id
  security_group_id = huaweicloud_networking_secgroup.secgroup_1.id
  available_zones   = [data.huaweicloud_dcs_az.az_1.id]
  product_id        = "dcs.master_standby-h"
  save_days         = 1
  backup_type       = "manual"
  begin_at          = "00:00-01:00"
  period_type       = "weekly"
  backup_at         = [1]
}
```

### DCS instance for Redis 5.0

```hcl
resource "huaweicloud_dcs_instance" "instance_1" {
  name              = "test_dcs_instance"
  engine            = "Redis"
  engine_version    = "5.0"
  password          = "Huawei_test"
  capacity          = 2
  vpc_id            = huaweicloud_vpc.vpc_1.id
  subnet_id         = huaweicloud_vpc_subnet.subnet_1.id
  available_zones   = [data.huaweicloud_dcs_az.az_1.id]
  product_id        = "redis.ha.au1.large.r2.2-h"
  save_days         = 1
  backup_type       = "manual"
  begin_at          = "00:00-01:00"
  period_type       = "weekly"
  backup_at         = [1]

  whitelists {
    group_name = "test-group1"
    ip_address = ["192.168.10.100", "192.168.0.0/24"]
  }
  whitelists {
    group_name = "test-group2"
    ip_address = ["172.16.10.100", "172.16.0.0/24"]
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the DCS instance resource. If omitted, the provider-level region will be used. Changing this creates a new DCS instance resource.

* `name` - (Required, String) Indicates the name of an instance. It starts with English characters 
    and can only be composed of English letters, numbers, underscores and underscores. 
    When creating a single instance, the name is a string of 4 to 64 bits in length. 
    When creating instances in batches, the length of the name is a string of 4 to 56 characters, 
    and the format of the instance name is "custom name-n", where n starts from 000 and increases in sequence.
    For example, if you create two instances in batches and the custom name is dcs_demo, 
    the names of the two instances are dcs_demo-000 and dcs_demo-001.

* `description` - (Optional, String) Indicates the description of an instance. It is a character
    string containing not more than 1024 characters.

* `engine` - (Required, String, ForceNew) Indicates a cache engine. Options: Redis and Memcached. Changing this
    creates a new instance.

* `engine_version` - (Optional, String, ForceNew) Indicates the version of a message engine.When the cache engine is Redis, 
    the value is 3.0, 4.0 or 5.0. 
    Changing this creates a new instance.

* `capacity` - (Required, Float, ForceNew) Indicates the Cache capacity. Unit: GB.
    Redis3.0: Stand-alone and active/standby type instance values: 2, 4, 8, 16, 32, 64. 
    Proxy cluster instance specifications support 64, 128, 256, 512, and 1024.

    Redis4.0 and Redis5.0: Stand-alone and active/standby type instance 
    values: 0.125, 0.25, 0.5, 1, 2, 4, 8, 16, 32, 64. Cluster instance specifications 
    support 24, 32, 48, 64, 96, 128, 192, 256, 384, 512, 768, 1024.

    Memcached: Stand-alone and active/standby type instance values: 2, 4, 8, 16, 32, 64.
    Changing this creates a new instance.

* `access_user` - (Optional, String, ForceNew) Username used for accessing a DCS instance after password
    authentication. A username starts with a letter, consists of 1 to 64 characters,
    and supports only letters, digits, and hyphens (-).
    - When the cache engine is Memcached, this parameter is optional.
    - When the cache engine is Redis, this parameter does not need to be set.
    Changing this creates a new instance.

* `password` - (Optional, String, ForceNew) Password of a DCS instance.
    The password of a DCS Redis instance must meet the following complexity requirements:
    - Enter a string of 8 to 32 bits in length.
    - The new password cannot be the same as the old password.
    - Must contain three combinations of the following four characters: Lower case letters,
        uppercase letter, digital, Special characters include (`~!@#$%^&*()-_=+|[{}]:'",<.>/?).
    Changing this creates a new instance.

* `vpc_id` - (Required, String, ForceNew) Specifies the id of the VPC.
    Changing this creates a new instance.

* `subnet_id` - (Required, String, ForceNew) Specifies the id of the subnet.
    Changing this creates a new instance.

* `security_group_id` - (Optional, String) Specifies the id of the security group which the instance belongs to.
    This parameter is mandatory for Memcached and Redis 3.0 versions.

* `whitelists` - (Optional, List) Specifies the IP addresses which can access the instance.
    This parameter is valid for Redis 4.0 and 5.0 versions. The structure is described below.

* `whitelist_enable` - (Optional, Bool) Enable or disable the IP addresse whitelists. Default to true.
    If the whitelist is disabled, all IP addresses connected to the VPC can access the instance.

* `available_zones` - (Required, List, ForceNew) IDs of the AZs where cache nodes reside.
    If you are creating active/standby, Proxy cluster, and Cluster cluster instances to support 
    cross-zone deployment, you can specify the standby zone for the standby node. When specifying 
    availability zones for nodes, separate them with commas.
    Changing this creates a new instance.

* `product_id` - (Required, String, ForceNew) Product ID or Names used to differentiate DCS instance types.
    Changing this creates a new instance.

* `maintain_begin` - (Optional, String) Indicates the time at which a maintenance time window starts.
    Format: HH:mm:ss.
    The start time and end time of a maintenance time window must indicate the time segment of
	a supported maintenance time window. For details, see section Querying Maintenance Time Windows.
    The start time must be set to 22:00, 02:00, 06:00, 10:00, 14:00, or 18:00.
    Parameters maintain_begin and maintain_end must be set in pairs. If parameter maintain_begin
	is left blank, parameter maintain_end is also blank. In this case, the system automatically
	allocates the default start time 02:00.

* `maintain_end` - (Optional, String) Indicates the time at which a maintenance time window ends.
    Format: HH:mm:ss.
    The start time and end time of a maintenance time window must indicate the time segment of
	a supported maintenance time window. For details, see section Querying Maintenance Time Windows.
    The end time is four hours later than the start time. For example, if the start time is 22:00,
	the end time is 02:00.
    Parameters maintain_begin and maintain_end must be set in pairs. If parameter maintain_end is left
	blank, parameter maintain_begin is also blank. In this case, the system automatically allocates
	the default end time 06:00.

* `save_days` - (Required, Int, ForceNew) Retention time. Unit: day. Range: 1–7. Changing this creates a new instance.

* `backup_type` - (Required, String, ForceNew) Backup type. Options:
    auto: automatic backup.
    manual: manual backup.
    Changing this creates a new instance.

* `begin_at` - (Required, String, ForceNew) Time at which backup starts. "00:00-01:00" indicates that backup
    starts at 00:00:00. Changing this creates a new instance.

* `period_type` - (Required, String, ForceNew) Interval at which backup is performed. Currently, only weekly
    backup is supported. Changing this creates a new instance.

* `backup_at` - (Required, List, ForceNew) Day in a week on which backup starts. Range: 1–7. Where: 1
    indicates Monday; 7 indicates Sunday. Changing this creates a new instance.

* `enterprise_project_id` - (Optional, String, ForceNew) The enterprise project id of the dcs instance. Changing this creates a new instance.

* `tags` - (Optional, Map) The key/value pairs to associate with the dcs instance.

The `whitelists` block supports:

* `group_name` - (Required, String) Specifies the name of IP address group.

* `ip_address` - (Required, List) Specifies the list of IP address or CIDR which can be whitelisted for an instance.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.
* `vpc_name` - Indicates the name of a vpc.
* `subnet_name` - Indicates the name of a subnet.
* `security_group_name` - Indicates the name of a security group.
* `order_id` - An order ID is generated only in the monthly or yearly billing mode.
    In other billing modes, no value is returned for this parameter.
* `resource_spec_code` - Resource specifications.
    dcs.single_node: indicates a DCS instance in single-node mode.
    dcs.master_standby: indicates a DCS instance in master/standby mode.
    dcs.cluster: indicates a DCS instance in cluster mode.
* `used_memory` - Size of the used memory. Unit: MB.
* `internal_version` - Internal DCS version.
* `max_memory` - Overall memory size. Unit: MB.
* `user_id` - Indicates a user ID.
* `ip` - Cache node's IP address in tenant's VPC.
* `port` - Port of the cache node.
