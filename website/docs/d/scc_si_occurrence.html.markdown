---
layout: "ibm"
subcategory: "Security and Compliance Center"
page_title: "IBM : ibm_scc_si_occurence"
description: |-
  Get information about scc_si_occurence
---

# ibm_scc_si_occurence

Provides a read-only data source for scc_si_occurence. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_si_occurence" "scc_si_occurence" {
	occurrence_id = "occurrence_id"
	provider_id = "provider_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `account_id` - (Optional, String) Account ID is optional, if not provided value will be inferred from the token retrieved from the IBM Cloud API key.
* `occurrence_id` - (Required, Forces new resource, String) Second part of occurrence `name`: providers/{provider_id}/occurrences/{occurrence_id}.
* `provider_id` - (Required, Forces new resource, String) Part of the parent. This field contains the provider ID. For example: providers/{provider_id}.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the scc_si_occurence.
* `context` - (Optional, List) 
Nested scheme for **context**:
	* `component_name` - (Optional, String) The name of the component the occurrence applies to.
	* `environment_name` - (Optional, String) The name of the environment the occurrence applies to.
	* `region` - (Optional, String) The IBM Cloud region.
	* `resource_crn` - (Optional, String) The resource CRN (e.g. certificate CRN, image CRN).
	* `resource_id` - (Optional, String) The resource ID, in case the CRN is not available.
	* `resource_name` - (Optional, String) The user-friendly resource name.
	* `resource_type` - (Optional, String) The resource type name (e.g. Pod, Cluster, Certificate, Image).
	* `service_crn` - (Optional, String) The service CRN (e.g. CertMgr Instance CRN).
	* `service_name` - (Optional, String) The service name (e.g. CertMgr).
	* `toolchain_id` - (Optional, String) The id of the toolchain the occurrence applies to.

* `create_time` - (Optional, String) Output only. The time this `Occurrence` was created.

* `finding` - (Optional, List) Finding provides details about a finding occurrence.
Nested scheme for **finding**:
	* `certainty` - (Optional, String) Note provider-assigned confidence on the validity of an occurrence- LOW&#58; Low Certainty- MEDIUM&#58; Medium Certainty- HIGH&#58; High Certainty.
	  * Constraints: Allowable values are: `LOW`, `MEDIUM`, `HIGH`.
	* `data_transferred` - (Optional, List) It provides details about data transferred between clients and servers.
	Nested scheme for **data_transferred**:
		* `client_bytes` - (Optional, Integer) The number of client bytes transferred.
		* `client_packets` - (Optional, Integer) The number of client packets transferred.
		* `server_bytes` - (Optional, Integer) The number of server bytes transferred.
		* `server_packets` - (Optional, Integer) The number of server packets transferred.
	* `network_connection` - (Optional, List) It provides details about a network connection.
	Nested scheme for **network_connection**:
		* `client` - (Optional, List) It provides details about a socket address.
		Nested scheme for **client**:
			* `address` - (Required, String) The IP address of this socket address.
			* `port` - (Optional, Integer) The port number of this socket address.
		* `direction` - (Optional, String) The direction of this network connection.
		* `protocol` - (Optional, String) The protocol of this network connection.
		* `server` - (Optional, List) It provides details about a socket address.
		Nested scheme for **server**:
			* `address` - (Required, String) The IP address of this socket address.
			* `port` - (Optional, Integer) The port number of this socket address.
	* `next_steps` - (Optional, List) Remediation steps for the issues reported in this finding. They override the note's next steps.
	Nested scheme for **next_steps**:
		* `title` - (Optional, String) Title of this next step.
		* `url` - (Optional, String) The URL associated to this next steps.
	* `severity` - (Optional, String) Note provider-assigned severity/impact ranking- LOW&#58; Low Impact- MEDIUM&#58; Medium Impact- HIGH&#58; High Impact- CRITICAL&#58; Critical Impact.
	  * Constraints: Allowable values are: `LOW`, `MEDIUM`, `HIGH`, `CRITICAL`.

* `id` - (Required, String) The id of the occurrence.

* `kind` - (Required, String) The type of note. Use this field to filter notes and occurences by kind. - FINDING&#58; The note and occurrence represent a finding. - KPI&#58; The note and occurrence represent a KPI value. - CARD&#58; The note represents a card showing findings and related metric values. - CARD_CONFIGURED&#58; The note represents a card configured for a user account. - SECTION&#58; The note represents a section in a dashboard.
  * Constraints: Allowable values are: `FINDING`, `KPI`, `CARD`, `CARD_CONFIGURED`, `SECTION`.

* `kpi` - (Optional, List) Kpi provides details about a KPI occurrence.
Nested scheme for **kpi**:
	* `total` - (Optional, Float) The total value of this KPI.
	* `value` - (Required, Float) The value of this KPI.

* `note_name` - (Required, String) An analysis note associated with this image, in the form "{account_id}/providers/{provider_id}/notes/{note_id}" This field can be used as a filter in list requests.

* `remediation` - (Optional, String) A description of actions that can be taken to remedy the `Note`.

* `resource_url` - (Optional, String) The unique URL of the resource, image or the container, for which the `Occurrence` applies. For example, https://gcr.io/provider/image@sha256:foo. This field can be used as a filter in list requests.

* `update_time` - (Optional, String) Output only. The time this `Occurrence` was last updated.

