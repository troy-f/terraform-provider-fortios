---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_multicastflow"
description: |-
  Configure multicast-flow.
---

# fortios_router_multicastflow
Configure multicast-flow.

## Example Usage

```hcl
resource "fortios_router_multicastflow" "trname" {
  name = "1"

  flows {
    group_addr  = "224.252.0.0"
    source_addr = "224.112.0.0"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `comments` - Comment.
* `flows` - Multicast-flow entries. The structure of `flows` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `flows` block supports:

* `id` - Flow ID.
* `group_addr` - Multicast group IP address.
* `source_addr` - Multicast source IP address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router MulticastFlow can be imported using any of these accepted formats:
```
$ terraform import fortios_router_multicastflow.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_router_multicastflow.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
