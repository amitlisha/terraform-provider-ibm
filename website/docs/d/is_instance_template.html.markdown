---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: instance_template"
description: |-
  Retrives a specified IBM VPC instance template.
---

# ibm_is_instance_template
Retrieve information of an existing IBM VPC instance template. For more information, about VPC instance templates, see [creating an instance template](https://cloud.ibm.com/docs/vpc?topic=vpc-create-instance-template).

## Example usage
In the following example, you can get information of an instance template of VPC Generation-2 infrastructure by either name or identifier.

```terraform	
data "ibm_is_instance_template" "instancetemplate" {
	name = "test-instance-template"	   
}

```
```terraform	
data "ibm_is_instance_template" "instancetemplate" {
	identifier = "xxxxx-xxxxx-xxxxxx-xxxxx"	   
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `identifier` - (Optional, String) The id of the instance template.
- `name` - (Optional, String) The name of the instance template.



## Attribute reference
You can access the following attribute references after your data source is created. 

- `boot_volume` - (List) A nested block describes the boot volume configuration for the template.

	Nested scheme for `boot_volume`:
	- `delete_volume_on_instance_delete` - (String) You can configure to delete the boot volume based on instance deletion.
	- `encryption` - (String) The encryption key CRN such as HPCS, Key Protect, etc., is provided to encrypt the boot volume attached.
	- `iops` - (String) The IOPS for the boot volume.
	- `name` - (String) The name of the boot volume.
	- `profile` - (String) The profile for the boot volume configuration.
	- `size` - (String) The boot volume size to configure in giga bytes.
- `crn` - (String) The CRN of the instance template.
- `href` - (String) The URL of the instance template.
- `id` - (String) The ID of the instance template.
- `image` - (String) The ID of the image to create the template.
- `keys` - (String) List of SSH key IDs used to allow log in user to the instances.
- `name` - (String) The name of the instance template.
- `network_interfaces` - (List) A nested block describes the network interfaces for the template.

	Nested scheme for `network_interfaces`:
	- `name` - (String) The name of the interface.
	- `primary_ipv4_address` - (String) The IPv4 address assigned to the network interface.
	- `subnet` - (String) The VPC subnet to assign to the interface.
	- `security_groups` - (String) List of security groups of  the subnet.
	
- `placement_target` - (List) The placement restrictions to use for the virtual server instance.
  Nested scheme for `placement_target`:
    - `crn` - (String) The unique identifier for this placement target.
    - `href` - (String) The CRN for this placement target.
    - `id` - (String) The URL for this placement target.
		
- `profile` - (String) The number of instances created in the instance group.
- `primary_network_interfaces` - (List) A nested block describes the primary network interface for the template.

	Nested scheme for `primary_network_interfaces`:
	- `name` - (String) The name of the interface.
	- `primary_ipv4_address` - (String) The IPv4 address assigned to the primary network interface.
	- `subnet` - (String) The VPC subnet to assign to the interface.
	- `security_groups` - (String) List of security groups of the subnet.
- `resource_group` - (String) The resource group ID.	
- `total_volume_bandwidth` - (Integer) The amount of bandwidth (in megabits per second) allocated exclusively to instance storage volumes
- `user_data` -  (String) The user data provided for the instance.
- `volume_attachments` - (List) A nested block describes the storage volume configuration for the template.

	Nested scheme for `volume_attachments`:
	- `delete_volume_on_instance_delete` - (Bool) You can configure to delete the storage volume to delete based on instance deletion.
	- `name` - (String) The name of the boot volume.
	- `volume` - (String) The storage volume ID created in VPC.
	- `volume_prototype` - (List) A nested block describing prototype for the volume.

		Nested scheme for `volume_prototype`:
		- `capacity` - (String) The capacity of the volume in gigabytes. The specified minimum and maximum capacity values for creating or updating volumes can expand in the future.
		- `encryption_key` - (String) The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.
		- `iops` - (String) The maximum input/output operations per second (IOPS) for the volume.
		- `profile` - (String) The global unique name for the volume profile to use for the volume.
- `vpc` - (String) The VPC ID that the instance templates needs to be created.
- `zone` - (String) The name of the zone.
