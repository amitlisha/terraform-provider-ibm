---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_instance_network_interfaces"
description: |-
  Get information about NetworkInterfaceCollection

---

# ibm_is_instance_network_interfaces

Provides a read-only data source for NetworkInterfaceCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_subnet" "example" {
  name            = "example-subnet"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-1"
  ipv4_cidr_block = "10.240.0.0/24"
}

resource "ibm_is_ssh_key" "example" {
  name       = "example-ssh"
  public_key = file("~/.ssh/id_rsa.pub")
}

resource "ibm_is_instance" "example" {
  name    = "example-instance"
  image   = "a7a0626c-f97e-4180-afbe-0331ec62f32a"
  profile = "bc1-2x8"

  primary_network_interface {
    subnet = ibm_is_subnet.example.id
  }

  network_interfaces {
    name   = "eth1"
    subnet = ibm_is_subnet.example.id
  }

  vpc  = ibm_is_vpc.example.id
  zone = "us-south-1"
  keys = [ibm_is_ssh_key.example.id]
}

data "ibm_is_instance_network_interfaces" "example" {
	instance_name = ibm_is_instance.example.name
}
```

## Argument Reference

The following arguments are supported:

- `instance_name` - (Required, string) The name of instance.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- `network_interfaces` - (List) Collection of network interfaces. Nested `network_interfaces` blocks have the following structure:
	- `allow_ip_spoofing` - (Boolean) Indicates whether source IP spoofing is allowed on this interface. If false, source IP spoofing is prevented on this interface. If true, source IP spoofing is allowed on this interface.
	- `created_at` - (String) The date and time that the network interface was created.
	- `floating_ips` - (List) The floating IPs associated with this network interface. Nested `floating_ips` blocks have the following structure:
		- `address` - (String) The globally unique IP address.
		- `crn` - (String) The CRN for this floating IP.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this floating IP.
		- `id` - (String) The unique identifier for this floating IP.
		- `name` - (String) The unique user-defined name for this floating IP.
	- `href` - (String) The URL for this network interface.
	- `id` - (String) The unique identifier for this network interface.
	- `name` - (String) The user-defined name for this network interface.
	- `port_speed` - (Integer) The network interface port speed in Mbps.
	- `primary_ipv4_address` - (String) The primary IPv4 address.
	- `resource_type` - (String) The resource type.
	- `security_groups` - (List) Collection of security groups. Nested `security_groups` blocks have the following structure:
		- `crn` - (String) The security group's CRN.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The security group's canonical URL.
		- `id` - (String) The unique identifier for this security group.
		- `name` - (String) The user-defined name for this security group. Names must be unique within the VPC the security group resides in.
	- `status` - (String) The status of the network interface.
	- `subnet` - (List) The associated subnet. Nested `subnet` blocks have the following structure:
		- `crn` - (String) The CRN for this subnet.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this subnet.
		- `id` - (String) The unique identifier for this subnet.
		- `name` - (String) The user-defined name for this subnet.
	- `type` - (String) The type of this network interface as it relates to an instance.
