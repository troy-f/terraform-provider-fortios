---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_h2qposuprovidernai"
description: |-
  Configure online sign up (OSU) provider NAI list.
---

# fortios_wirelesscontrollerhotspot20_h2qposuprovidernai
Configure online sign up (OSU) provider NAI list. Applies to FortiOS Version `>= 7.0.2`.

## Argument Reference

The following arguments are supported:

* `name` - OSU provider NAI ID.
* `nai_list` - OSU NAI list. The structure of `nai_list` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `nai_list` block supports:

* `name` - OSU NAI ID.
* `osu_nai` - OSU NAI.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 H2QpOsuProviderNai can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontrollerhotspot20_h2qposuprovidernai.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontrollerhotspot20_h2qposuprovidernai.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
