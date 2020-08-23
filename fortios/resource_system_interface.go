// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure interfaces.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemInterfaceCreate,
		Read:   resourceSystemInterfaceRead,
		Update: resourceSystemInterfaceUpdate,
		Delete: resourceSystemInterfaceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"vdom": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Required:     true,
			},
			"vrf": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"cli_conn_status": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"fortilink": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"distance": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"priority": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"dhcp_relay_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_relay_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_relay_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_relay_agent_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gwdetect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ping_serv_status": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"detectserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"detectprotocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_priority": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 50),
				Optional:     true,
				Computed:     true,
			},
			"fail_detect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fail_detect_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fail_alert_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fail_action_on_extender": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fail_alert_interfaces": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dhcp_client_identifier": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 48),
				Optional:     true,
				Computed:     true,
			},
			"dhcp_renew_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 604800),
				Optional:     true,
				Computed:     true,
			},
			"ipunnumbered": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
			"pppoe_unnumbered_negotiate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
			},
			"idle_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32767),
				Optional:     true,
				Computed:     true,
			},
			"detected_peer_mtu": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"disc_retry_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"padt_retry_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"service_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"ac_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"lcp_echo_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32767),
				Optional:     true,
				Computed:     true,
			},
			"lcp_max_echo_fails": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32767),
				Optional:     true,
				Computed:     true,
			},
			"defaultgw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pptp_client": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pptp_user": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
			"pptp_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
			},
			"pptp_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pptp_auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pptp_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"arpforward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ndiscforward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"broadcast_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bfd_desired_min_tx": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100000),
				Optional:     true,
				Computed:     true,
			},
			"bfd_detect_mult": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 50),
				Optional:     true,
				Computed:     true,
			},
			"bfd_required_min_rx": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100000),
				Optional:     true,
				Computed:     true,
			},
			"l2forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmp_send_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmp_accept_redirect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlanforward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stpforward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stpforward_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ips_sniffer_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ident_accept": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipmac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"macaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"substitute_dst_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"speed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netbios_forward": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wins_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dedicated_to": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trust_ip_1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trust_ip_2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trust_ip_3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trust_ip6_1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trust_ip6_2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trust_ip6_3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mtu_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mtu": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"wccp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netflow_sampler": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sflow_sampler": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"drop_overlapped_fragment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"drop_fragment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scan_botnet_connections": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sample_rate": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 99999),
				Optional:     true,
				Computed:     true,
			},
			"polling_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"sample_direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"explicit_web_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"explicit_ftp_proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"proxy_captive_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_mss": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"inbandwidth": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),
				Optional:     true,
				Computed:     true,
			},
			"outbandwidth": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),
				Optional:     true,
				Computed:     true,
			},
			"egress_shaping_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"disconnect_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10000),
				Optional:     true,
				Computed:     true,
			},
			"spillover_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),
				Optional:     true,
				Computed:     true,
			},
			"ingress_spillover_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16776000),
				Optional:     true,
				Computed:     true,
			},
			"weight": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"external": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlanid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 4094),
				Optional:     true,
				Computed:     true,
			},
			"forward_domain": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2147483647),
				Optional:     true,
				Computed:     true,
			},
			"remote_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"member": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"lacp_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lacp_ha_slave": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lacp_speed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"min_links": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 32),
				Optional:     true,
				Computed:     true,
			},
			"min_links_down": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"link_up_delay": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(50, 3600000),
				Optional:     true,
				Computed:     true,
			},
			"priority_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aggregate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"redundant_interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"managed_device": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"devindex": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"vindex": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"switch": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"alias": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 25),
				Optional:     true,
				Computed:     true,
			},
			"security_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"captive_portal": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"security_mac_auth_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_external_web": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"security_external_logout": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"replacemsg_override_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"security_redirect_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"security_exempt_list": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"security_groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"device_identification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_user_identification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_identification_active_scan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_access_list": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"device_netscan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lldp_reception": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lldp_transmission": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lldp_network_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"fortiheartbeat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"broadcast_forticlient_discovery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"endpoint_compliance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"estimated_upstream_bandwidth": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"estimated_downstream_bandwidth": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"vrrp_virtual_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrrp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vrid": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),
							Optional:     true,
							Computed:     true,
						},
						"version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vrgrp": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"vrip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),
							Optional:     true,
							Computed:     true,
						},
						"adv_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),
							Optional:     true,
							Computed:     true,
						},
						"start_time": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),
							Optional:     true,
							Computed:     true,
						},
						"preempt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"accept_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vrdst": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vrdst_priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 254),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"proxy_arp": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4294967295),
										Optional:     true,
										Computed:     true,
									},
									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmp_index": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"secondary_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secondaryip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional:     true,
							Computed:     true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"allowaccess": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gwdetect": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ping_serv_status": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"detectserver": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"detectprotocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ha_priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 50),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"preserve_session_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_auth_extension_device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_discover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortilink_stacking": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortilink_split_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internal": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"fortilink_backup_link": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"switch_controller_access_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_traffic_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"switch_controller_igmp_snooping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_dhcp_snooping": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_dhcp_snooping_verify_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_dhcp_snooping_option82": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_arp_inspection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_learning_limit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),
				Optional:     true,
				Computed:     true,
			},
			"color": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"tagging": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"category": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 64),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
					},
				},
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip6_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nd_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nd_cert": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"nd_security_level": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"nd_timestamp_delta": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 3600),
							Optional:     true,
							Computed:     true,
						},
						"nd_timestamp_fuzz": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 60),
							Optional:     true,
							Computed:     true,
						},
						"nd_cga_modifier": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_dns_server_override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_extra_addr": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"ip6_allowaccess": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_send_adv": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_manage_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_other_flag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_max_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(4, 1800),
							Optional:     true,
							Computed:     true,
						},
						"ip6_min_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(3, 1350),
							Optional:     true,
							Computed:     true,
						},
						"ip6_link_mtu": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1280, 16000),
							Optional:     true,
							Computed:     true,
						},
						"ip6_reachable_time": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 3600000),
							Optional:     true,
							Computed:     true,
						},
						"ip6_retrans_time": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional:     true,
							Computed:     true,
						},
						"ip6_default_life": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 9000),
							Optional:     true,
							Computed:     true,
						},
						"ip6_hop_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"autoconf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_upstream_interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"ip6_subnet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_prefix_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"autonomous_flag": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"onlink_flag": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"valid_life_time": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4294967295),
										Optional:     true,
										Computed:     true,
									},
									"preferred_life_time": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4294967295),
										Optional:     true,
										Computed:     true,
									},
									"rdnss": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dnssl": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"domain": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 79),
													Optional:     true,
													Computed:     true,
												},
											},
										},
									},
								},
							},
						},
						"ip6_delegated_prefix_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix_id": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4294967295),
										Optional:     true,
										Computed:     true,
									},
									"upstream_interface": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),
										Optional:     true,
										Computed:     true,
									},
									"autonomous_flag": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"onlink_flag": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"subnet": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"rdnss_service": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"rdnss": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"dhcp6_relay_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dhcp6_relay_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dhcp6_relay_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dhcp6_client_options": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dhcp6_prefix_delegation": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dhcp6_information_request": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dhcp6_prefix_hint": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dhcp6_prefix_hint_plt": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional:     true,
							Computed:     true,
						},
						"dhcp6_prefix_hint_vlt": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional:     true,
							Computed:     true,
						},
						"vrrp_virtual_mac6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vrip6_link_local": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vrrp6": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vrid": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 255),
										Optional:     true,
										Computed:     true,
									},
									"vrgrp": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),
										Optional:     true,
										Computed:     true,
									},
									"vrip6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"priority": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 255),
										Optional:     true,
										Computed:     true,
									},
									"adv_interval": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 255),
										Optional:     true,
										Computed:     true,
									},
									"start_time": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 255),
										Optional:     true,
										Computed:     true,
									},
									"preempt": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"accept_mode": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vrdst6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			}},
	}
}

func resourceSystemInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemInterface(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemInterface resource while getting object: %v", err)
	}

	var o map[string]interface{}
	v, _ := d.GetOk("type")
	if v == "physical" {
		o, err = c.UpdateSystemInterface(obj, (*obj)["name"].(string))
	} else {
		o, err = c.CreateSystemInterface(obj)
	}

	if err != nil {
		return fmt.Errorf("Error creating SystemInterface resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemInterface")
	}

	return resourceSystemInterfaceRead(d, m)
}

func resourceSystemInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemInterface(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterface resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemInterface(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemInterface")
	}

	return resourceSystemInterfaceRead(d, m)
}

func resourceSystemInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	v, _ := d.GetOk("type")
	if v != "physical" {
		err := c.DeleteSystemInterface(mkey)
		if err != nil {
			return fmt.Errorf("Error deleting SystemInterface resource: %v", err)
		}
	}

	d.SetId("")

	return nil
}

func resourceSystemInterfaceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemInterface(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemInterface resource from API: %v", err)
	}
	return nil
}

func flattenSystemInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceCliConnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceFortilink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfacePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDhcpRelayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDhcpRelayIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDhcpRelayType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDhcpRelayAgentOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceManagementIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceGwdetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfacePingServStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDetectserver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDetectprotocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceFailDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceFailDetectOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceFailAlertMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceFailActionOnExtender(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceFailAlertInterfaces(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenSystemInterfaceFailAlertInterfacesName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemInterfaceFailAlertInterfacesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDhcpClientIdentifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDhcpRenewTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpunnumbered(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfacePppoeUnnumberedNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfacePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDetectedPeerMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDiscRetryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfacePadtRetryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceAcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceLcpEchoInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceLcpMaxEchoFails(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDefaultgw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDnsServerOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfacePptpClient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfacePptpUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfacePptpPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfacePptpServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfacePptpAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfacePptpTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceArpforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceNdiscforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceBroadcastForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceBfdDesiredMinTx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceBfdDetectMult(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceBfdRequiredMinRx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceL2Forward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIcmpSendRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIcmpAcceptRedirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVlanforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceStpforward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceStpforwardMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpsSnifferMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIdentAccept(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpmac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSubst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceMacaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSubstituteDstMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSpeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceNetbiosForward(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceWinsIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDedicatedTo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceTrustIp1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemInterfaceTrustIp2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemInterfaceTrustIp3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemInterfaceTrustIp61(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceTrustIp62(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceTrustIp63(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceMtuOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceWccp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceNetflowSampler(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSflowSampler(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDropOverlappedFragment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDropFragment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceScanBotnetConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSrcCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSampleRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfacePollingInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSampleDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceExplicitWebProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceExplicitFtpProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceProxyCaptivePortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceTcpMss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceInbandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceOutbandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceEgressShapingProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDisconnectThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIngressSpilloverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceExternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceForwardDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceRemoteIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceMember(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := i["interface-name"]; ok {
			tmp["interface_name"] = flattenSystemInterfaceMemberInterfaceName(i["interface-name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemInterfaceMemberInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceLacpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceLacpHaSlave(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceLacpSpeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceMinLinks(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceMinLinksDown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceLinkUpDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfacePriorityOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceAggregate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceRedundantInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceManagedDevice(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenSystemInterfaceManagedDeviceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemInterfaceManagedDeviceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDevindex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVindex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSwitch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceAlias(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSecurityMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceCaptivePortal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSecurityMacAuthBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSecurityExternalWeb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSecurityExternalLogout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceReplacemsgOverrideGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSecurityRedirectUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSecurityExemptList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSecurityGroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenSystemInterfaceSecurityGroupsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemInterfaceSecurityGroupsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDeviceIdentification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDeviceUserIdentification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDeviceIdentificationActiveScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDeviceAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceDeviceNetscan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceLldpReception(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceLldpTransmission(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceLldpNetworkPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceFortiheartbeat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceBroadcastForticlientDiscovery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceEndpointCompliance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceEstimatedUpstreamBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceEstimatedDownstreamBandwidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpVirtualMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := i["vrid"]; ok {
			tmp["vrid"] = flattenSystemInterfaceVrrpVrid(i["vrid"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := i["version"]; ok {
			tmp["version"] = flattenSystemInterfaceVrrpVersion(i["version"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := i["vrgrp"]; ok {
			tmp["vrgrp"] = flattenSystemInterfaceVrrpVrgrp(i["vrgrp"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip"
		if _, ok := i["vrip"]; ok {
			tmp["vrip"] = flattenSystemInterfaceVrrpVrip(i["vrip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			tmp["priority"] = flattenSystemInterfaceVrrpPriority(i["priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := i["adv-interval"]; ok {
			tmp["adv_interval"] = flattenSystemInterfaceVrrpAdvInterval(i["adv-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := i["start-time"]; ok {
			tmp["start_time"] = flattenSystemInterfaceVrrpStartTime(i["start-time"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := i["preempt"]; ok {
			tmp["preempt"] = flattenSystemInterfaceVrrpPreempt(i["preempt"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_mode"
		if _, ok := i["accept-mode"]; ok {
			tmp["accept_mode"] = flattenSystemInterfaceVrrpAcceptMode(i["accept-mode"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst"
		if _, ok := i["vrdst"]; ok {
			tmp["vrdst"] = flattenSystemInterfaceVrrpVrdst(i["vrdst"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst_priority"
		if _, ok := i["vrdst-priority"]; ok {
			tmp["vrdst_priority"] = flattenSystemInterfaceVrrpVrdstPriority(i["vrdst-priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = flattenSystemInterfaceVrrpStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proxy_arp"
		if _, ok := i["proxy-arp"]; ok {
			tmp["proxy_arp"] = flattenSystemInterfaceVrrpProxyArp(i["proxy-arp"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemInterfaceVrrpVrid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpVrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpVrip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpAdvInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpStartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpPreempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpAcceptMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpVrdst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpVrdstPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpProxyArp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenSystemInterfaceVrrpProxyArpId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = flattenSystemInterfaceVrrpProxyArpIp(i["ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemInterfaceVrrpProxyArpId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceVrrpProxyArpIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceRole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSnmpIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSecondaryIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSecondaryip(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenSystemInterfaceSecondaryipId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = flattenSystemInterfaceSecondaryipIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowaccess"
		if _, ok := i["allowaccess"]; ok {
			tmp["allowaccess"] = flattenSystemInterfaceSecondaryipAllowaccess(i["allowaccess"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gwdetect"
		if _, ok := i["gwdetect"]; ok {
			tmp["gwdetect"] = flattenSystemInterfaceSecondaryipGwdetect(i["gwdetect"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ping_serv_status"
		if _, ok := i["ping-serv-status"]; ok {
			tmp["ping_serv_status"] = flattenSystemInterfaceSecondaryipPingServStatus(i["ping-serv-status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectserver"
		if _, ok := i["detectserver"]; ok {
			tmp["detectserver"] = flattenSystemInterfaceSecondaryipDetectserver(i["detectserver"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectprotocol"
		if _, ok := i["detectprotocol"]; ok {
			tmp["detectprotocol"] = flattenSystemInterfaceSecondaryipDetectprotocol(i["detectprotocol"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := i["ha-priority"]; ok {
			tmp["ha_priority"] = flattenSystemInterfaceSecondaryipHaPriority(i["ha-priority"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemInterfaceSecondaryipId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSecondaryipIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSecondaryipAllowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSecondaryipGwdetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSecondaryipPingServStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSecondaryipDetectserver(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSecondaryipDetectprotocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSecondaryipHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfacePreserveSessionRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceAutoAuthExtensionDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceApDiscover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceFortilinkStacking(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceFortilinkSplitInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceInternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceFortilinkBackupLink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSwitchControllerAccessVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSwitchControllerTrafficPolicy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSwitchControllerIgmpSnooping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSwitchControllerDhcpSnooping(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSwitchControllerDhcpSnoopingVerifyMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSwitchControllerDhcpSnoopingOption82(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSwitchControllerArpInspection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceSwitchControllerLearningLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceTagging(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenSystemInterfaceTaggingName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			tmp["category"] = flattenSystemInterfaceTaggingCategory(i["category"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := i["tags"]; ok {
			tmp["tags"] = flattenSystemInterfaceTaggingTags(i["tags"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemInterfaceTaggingName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceTaggingCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceTaggingTags(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenSystemInterfaceTaggingTagsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemInterfaceTaggingTagsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ip6_mode"
	if _, ok := i["ip6-mode"]; ok {
		result["ip6_mode"] = flattenSystemInterfaceIpv6Ip6Mode(i["ip6-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_mode"
	if _, ok := i["nd-mode"]; ok {
		result["nd_mode"] = flattenSystemInterfaceIpv6NdMode(i["nd-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_cert"
	if _, ok := i["nd-cert"]; ok {
		result["nd_cert"] = flattenSystemInterfaceIpv6NdCert(i["nd-cert"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_security_level"
	if _, ok := i["nd-security-level"]; ok {
		result["nd_security_level"] = flattenSystemInterfaceIpv6NdSecurityLevel(i["nd-security-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_timestamp_delta"
	if _, ok := i["nd-timestamp-delta"]; ok {
		result["nd_timestamp_delta"] = flattenSystemInterfaceIpv6NdTimestampDelta(i["nd-timestamp-delta"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_timestamp_fuzz"
	if _, ok := i["nd-timestamp-fuzz"]; ok {
		result["nd_timestamp_fuzz"] = flattenSystemInterfaceIpv6NdTimestampFuzz(i["nd-timestamp-fuzz"], d, pre_append)
	}

	pre_append = pre + ".0." + "nd_cga_modifier"
	if _, ok := i["nd-cga-modifier"]; ok {
		result["nd_cga_modifier"] = flattenSystemInterfaceIpv6NdCgaModifier(i["nd-cga-modifier"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_dns_server_override"
	if _, ok := i["ip6-dns-server-override"]; ok {
		result["ip6_dns_server_override"] = flattenSystemInterfaceIpv6Ip6DnsServerOverride(i["ip6-dns-server-override"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_address"
	if _, ok := i["ip6-address"]; ok {
		result["ip6_address"] = flattenSystemInterfaceIpv6Ip6Address(i["ip6-address"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_extra_addr"
	if _, ok := i["ip6-extra-addr"]; ok {
		result["ip6_extra_addr"] = flattenSystemInterfaceIpv6Ip6ExtraAddr(i["ip6-extra-addr"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_allowaccess"
	if _, ok := i["ip6-allowaccess"]; ok {
		result["ip6_allowaccess"] = flattenSystemInterfaceIpv6Ip6Allowaccess(i["ip6-allowaccess"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_send_adv"
	if _, ok := i["ip6-send-adv"]; ok {
		result["ip6_send_adv"] = flattenSystemInterfaceIpv6Ip6SendAdv(i["ip6-send-adv"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_manage_flag"
	if _, ok := i["ip6-manage-flag"]; ok {
		result["ip6_manage_flag"] = flattenSystemInterfaceIpv6Ip6ManageFlag(i["ip6-manage-flag"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_other_flag"
	if _, ok := i["ip6-other-flag"]; ok {
		result["ip6_other_flag"] = flattenSystemInterfaceIpv6Ip6OtherFlag(i["ip6-other-flag"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_max_interval"
	if _, ok := i["ip6-max-interval"]; ok {
		result["ip6_max_interval"] = flattenSystemInterfaceIpv6Ip6MaxInterval(i["ip6-max-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_min_interval"
	if _, ok := i["ip6-min-interval"]; ok {
		result["ip6_min_interval"] = flattenSystemInterfaceIpv6Ip6MinInterval(i["ip6-min-interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_link_mtu"
	if _, ok := i["ip6-link-mtu"]; ok {
		result["ip6_link_mtu"] = flattenSystemInterfaceIpv6Ip6LinkMtu(i["ip6-link-mtu"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_reachable_time"
	if _, ok := i["ip6-reachable-time"]; ok {
		result["ip6_reachable_time"] = flattenSystemInterfaceIpv6Ip6ReachableTime(i["ip6-reachable-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_retrans_time"
	if _, ok := i["ip6-retrans-time"]; ok {
		result["ip6_retrans_time"] = flattenSystemInterfaceIpv6Ip6RetransTime(i["ip6-retrans-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_default_life"
	if _, ok := i["ip6-default-life"]; ok {
		result["ip6_default_life"] = flattenSystemInterfaceIpv6Ip6DefaultLife(i["ip6-default-life"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_hop_limit"
	if _, ok := i["ip6-hop-limit"]; ok {
		result["ip6_hop_limit"] = flattenSystemInterfaceIpv6Ip6HopLimit(i["ip6-hop-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "autoconf"
	if _, ok := i["autoconf"]; ok {
		result["autoconf"] = flattenSystemInterfaceIpv6Autoconf(i["autoconf"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_upstream_interface"
	if _, ok := i["ip6-upstream-interface"]; ok {
		result["ip6_upstream_interface"] = flattenSystemInterfaceIpv6Ip6UpstreamInterface(i["ip6-upstream-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_subnet"
	if _, ok := i["ip6-subnet"]; ok {
		result["ip6_subnet"] = flattenSystemInterfaceIpv6Ip6Subnet(i["ip6-subnet"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_prefix_list"
	if _, ok := i["ip6-prefix-list"]; ok {
		result["ip6_prefix_list"] = flattenSystemInterfaceIpv6Ip6PrefixList(i["ip6-prefix-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "ip6_delegated_prefix_list"
	if _, ok := i["ip6-delegated-prefix-list"]; ok {
		result["ip6_delegated_prefix_list"] = flattenSystemInterfaceIpv6Ip6DelegatedPrefixList(i["ip6-delegated-prefix-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_service"
	if _, ok := i["dhcp6-relay-service"]; ok {
		result["dhcp6_relay_service"] = flattenSystemInterfaceIpv6Dhcp6RelayService(i["dhcp6-relay-service"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_type"
	if _, ok := i["dhcp6-relay-type"]; ok {
		result["dhcp6_relay_type"] = flattenSystemInterfaceIpv6Dhcp6RelayType(i["dhcp6-relay-type"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_relay_ip"
	if _, ok := i["dhcp6-relay-ip"]; ok {
		result["dhcp6_relay_ip"] = flattenSystemInterfaceIpv6Dhcp6RelayIp(i["dhcp6-relay-ip"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_client_options"
	if _, ok := i["dhcp6-client-options"]; ok {
		result["dhcp6_client_options"] = flattenSystemInterfaceIpv6Dhcp6ClientOptions(i["dhcp6-client-options"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_delegation"
	if _, ok := i["dhcp6-prefix-delegation"]; ok {
		result["dhcp6_prefix_delegation"] = flattenSystemInterfaceIpv6Dhcp6PrefixDelegation(i["dhcp6-prefix-delegation"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_information_request"
	if _, ok := i["dhcp6-information-request"]; ok {
		result["dhcp6_information_request"] = flattenSystemInterfaceIpv6Dhcp6InformationRequest(i["dhcp6-information-request"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_hint"
	if _, ok := i["dhcp6-prefix-hint"]; ok {
		result["dhcp6_prefix_hint"] = flattenSystemInterfaceIpv6Dhcp6PrefixHint(i["dhcp6-prefix-hint"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_hint_plt"
	if _, ok := i["dhcp6-prefix-hint-plt"]; ok {
		result["dhcp6_prefix_hint_plt"] = flattenSystemInterfaceIpv6Dhcp6PrefixHintPlt(i["dhcp6-prefix-hint-plt"], d, pre_append)
	}

	pre_append = pre + ".0." + "dhcp6_prefix_hint_vlt"
	if _, ok := i["dhcp6-prefix-hint-vlt"]; ok {
		result["dhcp6_prefix_hint_vlt"] = flattenSystemInterfaceIpv6Dhcp6PrefixHintVlt(i["dhcp6-prefix-hint-vlt"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrrp_virtual_mac6"
	if _, ok := i["vrrp-virtual-mac6"]; ok {
		result["vrrp_virtual_mac6"] = flattenSystemInterfaceIpv6VrrpVirtualMac6(i["vrrp-virtual-mac6"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrip6_link_local"
	if _, ok := i["vrip6_link_local"]; ok {
		result["vrip6_link_local"] = flattenSystemInterfaceIpv6Vrip6_Link_Local(i["vrip6_link_local"], d, pre_append)
	}

	pre_append = pre + ".0." + "vrrp6"
	if _, ok := i["vrrp6"]; ok {
		result["vrrp6"] = flattenSystemInterfaceIpv6Vrrp6(i["vrrp6"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemInterfaceIpv6Ip6Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6NdMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6NdCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6NdSecurityLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6NdTimestampDelta(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6NdTimestampFuzz(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6NdCgaModifier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6DnsServerOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6ExtraAddr(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = flattenSystemInterfaceIpv6Ip6ExtraAddrPrefix(i["prefix"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemInterfaceIpv6Ip6ExtraAddrPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6Allowaccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6SendAdv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6ManageFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6OtherFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6MaxInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6MinInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6LinkMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6ReachableTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6RetransTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6DefaultLife(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6HopLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Autoconf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6UpstreamInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6Subnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6PrefixList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = flattenSystemInterfaceIpv6Ip6PrefixListPrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := i["autonomous-flag"]; ok {
			tmp["autonomous_flag"] = flattenSystemInterfaceIpv6Ip6PrefixListAutonomousFlag(i["autonomous-flag"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := i["onlink-flag"]; ok {
			tmp["onlink_flag"] = flattenSystemInterfaceIpv6Ip6PrefixListOnlinkFlag(i["onlink-flag"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valid_life_time"
		if _, ok := i["valid-life-time"]; ok {
			tmp["valid_life_time"] = flattenSystemInterfaceIpv6Ip6PrefixListValidLifeTime(i["valid-life-time"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_life_time"
		if _, ok := i["preferred-life-time"]; ok {
			tmp["preferred_life_time"] = flattenSystemInterfaceIpv6Ip6PrefixListPreferredLifeTime(i["preferred-life-time"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := i["rdnss"]; ok {
			tmp["rdnss"] = flattenSystemInterfaceIpv6Ip6PrefixListRdnss(i["rdnss"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dnssl"
		if _, ok := i["dnssl"]; ok {
			tmp["dnssl"] = flattenSystemInterfaceIpv6Ip6PrefixListDnssl(i["dnssl"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemInterfaceIpv6Ip6PrefixListPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6PrefixListAutonomousFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6PrefixListOnlinkFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6PrefixListValidLifeTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6PrefixListPreferredLifeTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6PrefixListRdnss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6PrefixListDnssl(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {
			tmp["domain"] = flattenSystemInterfaceIpv6Ip6PrefixListDnsslDomain(i["domain"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemInterfaceIpv6Ip6PrefixListDnsslDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6DelegatedPrefixList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_id"
		if _, ok := i["prefix-id"]; ok {
			tmp["prefix_id"] = flattenSystemInterfaceIpv6Ip6DelegatedPrefixListPrefixId(i["prefix-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "upstream_interface"
		if _, ok := i["upstream-interface"]; ok {
			tmp["upstream_interface"] = flattenSystemInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface(i["upstream-interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := i["autonomous-flag"]; ok {
			tmp["autonomous_flag"] = flattenSystemInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag(i["autonomous-flag"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := i["onlink-flag"]; ok {
			tmp["onlink_flag"] = flattenSystemInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag(i["onlink-flag"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := i["subnet"]; ok {
			tmp["subnet"] = flattenSystemInterfaceIpv6Ip6DelegatedPrefixListSubnet(i["subnet"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss_service"
		if _, ok := i["rdnss-service"]; ok {
			tmp["rdnss_service"] = flattenSystemInterfaceIpv6Ip6DelegatedPrefixListRdnssService(i["rdnss-service"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := i["rdnss"]; ok {
			tmp["rdnss"] = flattenSystemInterfaceIpv6Ip6DelegatedPrefixListRdnss(i["rdnss"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemInterfaceIpv6Ip6DelegatedPrefixListPrefixId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6DelegatedPrefixListSubnet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6DelegatedPrefixListRdnssService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Ip6DelegatedPrefixListRdnss(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6RelayService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6RelayType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6RelayIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6ClientOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6PrefixDelegation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6InformationRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6PrefixHint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6PrefixHintPlt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Dhcp6PrefixHintVlt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6VrrpVirtualMac6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrip6_Link_Local(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := i["vrid"]; ok {
			tmp["vrid"] = flattenSystemInterfaceIpv6Vrrp6Vrid(i["vrid"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := i["vrgrp"]; ok {
			tmp["vrgrp"] = flattenSystemInterfaceIpv6Vrrp6Vrgrp(i["vrgrp"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip6"
		if _, ok := i["vrip6"]; ok {
			tmp["vrip6"] = flattenSystemInterfaceIpv6Vrrp6Vrip6(i["vrip6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			tmp["priority"] = flattenSystemInterfaceIpv6Vrrp6Priority(i["priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := i["adv-interval"]; ok {
			tmp["adv_interval"] = flattenSystemInterfaceIpv6Vrrp6AdvInterval(i["adv-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := i["start-time"]; ok {
			tmp["start_time"] = flattenSystemInterfaceIpv6Vrrp6StartTime(i["start-time"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := i["preempt"]; ok {
			tmp["preempt"] = flattenSystemInterfaceIpv6Vrrp6Preempt(i["preempt"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_mode"
		if _, ok := i["accept-mode"]; ok {
			tmp["accept_mode"] = flattenSystemInterfaceIpv6Vrrp6AcceptMode(i["accept-mode"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst6"
		if _, ok := i["vrdst6"]; ok {
			tmp["vrdst6"] = flattenSystemInterfaceIpv6Vrrp6Vrdst6(i["vrdst6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = flattenSystemInterfaceIpv6Vrrp6Status(i["status"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemInterfaceIpv6Vrrp6Vrid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6Vrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6Vrip6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6Priority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6AdvInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6StartTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6Preempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6AcceptMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6Vrdst6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemInterfaceIpv6Vrrp6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemInterfaceName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemInterfaceVdom(o["vdom"], d, "vdom")); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("vrf", flattenSystemInterfaceVrf(o["vrf"], d, "vrf")); err != nil {
		if !fortiAPIPatch(o["vrf"]) {
			return fmt.Errorf("Error reading vrf: %v", err)
		}
	}

	if err = d.Set("cli_conn_status", flattenSystemInterfaceCliConnStatus(o["cli-conn-status"], d, "cli_conn_status")); err != nil {
		if !fortiAPIPatch(o["cli-conn-status"]) {
			return fmt.Errorf("Error reading cli_conn_status: %v", err)
		}
	}

	if err = d.Set("fortilink", flattenSystemInterfaceFortilink(o["fortilink"], d, "fortilink")); err != nil {
		if !fortiAPIPatch(o["fortilink"]) {
			return fmt.Errorf("Error reading fortilink: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemInterfaceMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("distance", flattenSystemInterfaceDistance(o["distance"], d, "distance")); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystemInterfacePriority(o["priority"], d, "priority")); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_service", flattenSystemInterfaceDhcpRelayService(o["dhcp-relay-service"], d, "dhcp_relay_service")); err != nil {
		if !fortiAPIPatch(o["dhcp-relay-service"]) {
			return fmt.Errorf("Error reading dhcp_relay_service: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_ip", flattenSystemInterfaceDhcpRelayIp(o["dhcp-relay-ip"], d, "dhcp_relay_ip")); err != nil {
		if !fortiAPIPatch(o["dhcp-relay-ip"]) {
			return fmt.Errorf("Error reading dhcp_relay_ip: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_type", flattenSystemInterfaceDhcpRelayType(o["dhcp-relay-type"], d, "dhcp_relay_type")); err != nil {
		if !fortiAPIPatch(o["dhcp-relay-type"]) {
			return fmt.Errorf("Error reading dhcp_relay_type: %v", err)
		}
	}

	if err = d.Set("dhcp_relay_agent_option", flattenSystemInterfaceDhcpRelayAgentOption(o["dhcp-relay-agent-option"], d, "dhcp_relay_agent_option")); err != nil {
		if !fortiAPIPatch(o["dhcp-relay-agent-option"]) {
			return fmt.Errorf("Error reading dhcp_relay_agent_option: %v", err)
		}
	}

	if err = d.Set("management_ip", flattenSystemInterfaceManagementIp(o["management-ip"], d, "management_ip")); err != nil {
		if !fortiAPIPatch(o["management-ip"]) {
			return fmt.Errorf("Error reading management_ip: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemInterfaceIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("allowaccess", flattenSystemInterfaceAllowaccess(o["allowaccess"], d, "allowaccess")); err != nil {
		if !fortiAPIPatch(o["allowaccess"]) {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("gwdetect", flattenSystemInterfaceGwdetect(o["gwdetect"], d, "gwdetect")); err != nil {
		if !fortiAPIPatch(o["gwdetect"]) {
			return fmt.Errorf("Error reading gwdetect: %v", err)
		}
	}

	if err = d.Set("ping_serv_status", flattenSystemInterfacePingServStatus(o["ping-serv-status"], d, "ping_serv_status")); err != nil {
		if !fortiAPIPatch(o["ping-serv-status"]) {
			return fmt.Errorf("Error reading ping_serv_status: %v", err)
		}
	}

	if err = d.Set("detectserver", flattenSystemInterfaceDetectserver(o["detectserver"], d, "detectserver")); err != nil {
		if !fortiAPIPatch(o["detectserver"]) {
			return fmt.Errorf("Error reading detectserver: %v", err)
		}
	}

	if err = d.Set("detectprotocol", flattenSystemInterfaceDetectprotocol(o["detectprotocol"], d, "detectprotocol")); err != nil {
		if !fortiAPIPatch(o["detectprotocol"]) {
			return fmt.Errorf("Error reading detectprotocol: %v", err)
		}
	}

	if err = d.Set("ha_priority", flattenSystemInterfaceHaPriority(o["ha-priority"], d, "ha_priority")); err != nil {
		if !fortiAPIPatch(o["ha-priority"]) {
			return fmt.Errorf("Error reading ha_priority: %v", err)
		}
	}

	if err = d.Set("fail_detect", flattenSystemInterfaceFailDetect(o["fail-detect"], d, "fail_detect")); err != nil {
		if !fortiAPIPatch(o["fail-detect"]) {
			return fmt.Errorf("Error reading fail_detect: %v", err)
		}
	}

	if err = d.Set("fail_detect_option", flattenSystemInterfaceFailDetectOption(o["fail-detect-option"], d, "fail_detect_option")); err != nil {
		if !fortiAPIPatch(o["fail-detect-option"]) {
			return fmt.Errorf("Error reading fail_detect_option: %v", err)
		}
	}

	if err = d.Set("fail_alert_method", flattenSystemInterfaceFailAlertMethod(o["fail-alert-method"], d, "fail_alert_method")); err != nil {
		if !fortiAPIPatch(o["fail-alert-method"]) {
			return fmt.Errorf("Error reading fail_alert_method: %v", err)
		}
	}

	if err = d.Set("fail_action_on_extender", flattenSystemInterfaceFailActionOnExtender(o["fail-action-on-extender"], d, "fail_action_on_extender")); err != nil {
		if !fortiAPIPatch(o["fail-action-on-extender"]) {
			return fmt.Errorf("Error reading fail_action_on_extender: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("fail_alert_interfaces", flattenSystemInterfaceFailAlertInterfaces(o["fail-alert-interfaces"], d, "fail_alert_interfaces")); err != nil {
			if !fortiAPIPatch(o["fail-alert-interfaces"]) {
				return fmt.Errorf("Error reading fail_alert_interfaces: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fail_alert_interfaces"); ok {
			if err = d.Set("fail_alert_interfaces", flattenSystemInterfaceFailAlertInterfaces(o["fail-alert-interfaces"], d, "fail_alert_interfaces")); err != nil {
				if !fortiAPIPatch(o["fail-alert-interfaces"]) {
					return fmt.Errorf("Error reading fail_alert_interfaces: %v", err)
				}
			}
		}
	}

	if err = d.Set("dhcp_client_identifier", flattenSystemInterfaceDhcpClientIdentifier(o["dhcp-client-identifier"], d, "dhcp_client_identifier")); err != nil {
		if !fortiAPIPatch(o["dhcp-client-identifier"]) {
			return fmt.Errorf("Error reading dhcp_client_identifier: %v", err)
		}
	}

	if err = d.Set("dhcp_renew_time", flattenSystemInterfaceDhcpRenewTime(o["dhcp-renew-time"], d, "dhcp_renew_time")); err != nil {
		if !fortiAPIPatch(o["dhcp-renew-time"]) {
			return fmt.Errorf("Error reading dhcp_renew_time: %v", err)
		}
	}

	if err = d.Set("ipunnumbered", flattenSystemInterfaceIpunnumbered(o["ipunnumbered"], d, "ipunnumbered")); err != nil {
		if !fortiAPIPatch(o["ipunnumbered"]) {
			return fmt.Errorf("Error reading ipunnumbered: %v", err)
		}
	}

	if err = d.Set("username", flattenSystemInterfaceUsername(o["username"], d, "username")); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("pppoe_unnumbered_negotiate", flattenSystemInterfacePppoeUnnumberedNegotiate(o["pppoe-unnumbered-negotiate"], d, "pppoe_unnumbered_negotiate")); err != nil {
		if !fortiAPIPatch(o["pppoe-unnumbered-negotiate"]) {
			return fmt.Errorf("Error reading pppoe_unnumbered_negotiate: %v", err)
		}
	}

	if err = d.Set("password", flattenSystemInterfacePassword(o["password"], d, "password")); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("idle_timeout", flattenSystemInterfaceIdleTimeout(o["idle-timeout"], d, "idle_timeout")); err != nil {
		if !fortiAPIPatch(o["idle-timeout"]) {
			return fmt.Errorf("Error reading idle_timeout: %v", err)
		}
	}

	if err = d.Set("detected_peer_mtu", flattenSystemInterfaceDetectedPeerMtu(o["detected-peer-mtu"], d, "detected_peer_mtu")); err != nil {
		if !fortiAPIPatch(o["detected-peer-mtu"]) {
			return fmt.Errorf("Error reading detected_peer_mtu: %v", err)
		}
	}

	if err = d.Set("disc_retry_timeout", flattenSystemInterfaceDiscRetryTimeout(o["disc-retry-timeout"], d, "disc_retry_timeout")); err != nil {
		if !fortiAPIPatch(o["disc-retry-timeout"]) {
			return fmt.Errorf("Error reading disc_retry_timeout: %v", err)
		}
	}

	if err = d.Set("padt_retry_timeout", flattenSystemInterfacePadtRetryTimeout(o["padt-retry-timeout"], d, "padt_retry_timeout")); err != nil {
		if !fortiAPIPatch(o["padt-retry-timeout"]) {
			return fmt.Errorf("Error reading padt_retry_timeout: %v", err)
		}
	}

	if err = d.Set("service_name", flattenSystemInterfaceServiceName(o["service-name"], d, "service_name")); err != nil {
		if !fortiAPIPatch(o["service-name"]) {
			return fmt.Errorf("Error reading service_name: %v", err)
		}
	}

	if err = d.Set("ac_name", flattenSystemInterfaceAcName(o["ac-name"], d, "ac_name")); err != nil {
		if !fortiAPIPatch(o["ac-name"]) {
			return fmt.Errorf("Error reading ac_name: %v", err)
		}
	}

	if err = d.Set("lcp_echo_interval", flattenSystemInterfaceLcpEchoInterval(o["lcp-echo-interval"], d, "lcp_echo_interval")); err != nil {
		if !fortiAPIPatch(o["lcp-echo-interval"]) {
			return fmt.Errorf("Error reading lcp_echo_interval: %v", err)
		}
	}

	if err = d.Set("lcp_max_echo_fails", flattenSystemInterfaceLcpMaxEchoFails(o["lcp-max-echo-fails"], d, "lcp_max_echo_fails")); err != nil {
		if !fortiAPIPatch(o["lcp-max-echo-fails"]) {
			return fmt.Errorf("Error reading lcp_max_echo_fails: %v", err)
		}
	}

	if err = d.Set("defaultgw", flattenSystemInterfaceDefaultgw(o["defaultgw"], d, "defaultgw")); err != nil {
		if !fortiAPIPatch(o["defaultgw"]) {
			return fmt.Errorf("Error reading defaultgw: %v", err)
		}
	}

	if err = d.Set("dns_server_override", flattenSystemInterfaceDnsServerOverride(o["dns-server-override"], d, "dns_server_override")); err != nil {
		if !fortiAPIPatch(o["dns-server-override"]) {
			return fmt.Errorf("Error reading dns_server_override: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenSystemInterfaceAuthType(o["auth-type"], d, "auth_type")); err != nil {
		if !fortiAPIPatch(o["auth-type"]) {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("pptp_client", flattenSystemInterfacePptpClient(o["pptp-client"], d, "pptp_client")); err != nil {
		if !fortiAPIPatch(o["pptp-client"]) {
			return fmt.Errorf("Error reading pptp_client: %v", err)
		}
	}

	if err = d.Set("pptp_user", flattenSystemInterfacePptpUser(o["pptp-user"], d, "pptp_user")); err != nil {
		if !fortiAPIPatch(o["pptp-user"]) {
			return fmt.Errorf("Error reading pptp_user: %v", err)
		}
	}

	if err = d.Set("pptp_password", flattenSystemInterfacePptpPassword(o["pptp-password"], d, "pptp_password")); err != nil {
		if !fortiAPIPatch(o["pptp-password"]) {
			return fmt.Errorf("Error reading pptp_password: %v", err)
		}
	}

	if err = d.Set("pptp_server_ip", flattenSystemInterfacePptpServerIp(o["pptp-server-ip"], d, "pptp_server_ip")); err != nil {
		if !fortiAPIPatch(o["pptp-server-ip"]) {
			return fmt.Errorf("Error reading pptp_server_ip: %v", err)
		}
	}

	if err = d.Set("pptp_auth_type", flattenSystemInterfacePptpAuthType(o["pptp-auth-type"], d, "pptp_auth_type")); err != nil {
		if !fortiAPIPatch(o["pptp-auth-type"]) {
			return fmt.Errorf("Error reading pptp_auth_type: %v", err)
		}
	}

	if err = d.Set("pptp_timeout", flattenSystemInterfacePptpTimeout(o["pptp-timeout"], d, "pptp_timeout")); err != nil {
		if !fortiAPIPatch(o["pptp-timeout"]) {
			return fmt.Errorf("Error reading pptp_timeout: %v", err)
		}
	}

	if err = d.Set("arpforward", flattenSystemInterfaceArpforward(o["arpforward"], d, "arpforward")); err != nil {
		if !fortiAPIPatch(o["arpforward"]) {
			return fmt.Errorf("Error reading arpforward: %v", err)
		}
	}

	if err = d.Set("ndiscforward", flattenSystemInterfaceNdiscforward(o["ndiscforward"], d, "ndiscforward")); err != nil {
		if !fortiAPIPatch(o["ndiscforward"]) {
			return fmt.Errorf("Error reading ndiscforward: %v", err)
		}
	}

	if err = d.Set("broadcast_forward", flattenSystemInterfaceBroadcastForward(o["broadcast-forward"], d, "broadcast_forward")); err != nil {
		if !fortiAPIPatch(o["broadcast-forward"]) {
			return fmt.Errorf("Error reading broadcast_forward: %v", err)
		}
	}

	if err = d.Set("bfd", flattenSystemInterfaceBfd(o["bfd"], d, "bfd")); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("bfd_desired_min_tx", flattenSystemInterfaceBfdDesiredMinTx(o["bfd-desired-min-tx"], d, "bfd_desired_min_tx")); err != nil {
		if !fortiAPIPatch(o["bfd-desired-min-tx"]) {
			return fmt.Errorf("Error reading bfd_desired_min_tx: %v", err)
		}
	}

	if err = d.Set("bfd_detect_mult", flattenSystemInterfaceBfdDetectMult(o["bfd-detect-mult"], d, "bfd_detect_mult")); err != nil {
		if !fortiAPIPatch(o["bfd-detect-mult"]) {
			return fmt.Errorf("Error reading bfd_detect_mult: %v", err)
		}
	}

	if err = d.Set("bfd_required_min_rx", flattenSystemInterfaceBfdRequiredMinRx(o["bfd-required-min-rx"], d, "bfd_required_min_rx")); err != nil {
		if !fortiAPIPatch(o["bfd-required-min-rx"]) {
			return fmt.Errorf("Error reading bfd_required_min_rx: %v", err)
		}
	}

	if err = d.Set("l2forward", flattenSystemInterfaceL2Forward(o["l2forward"], d, "l2forward")); err != nil {
		if !fortiAPIPatch(o["l2forward"]) {
			return fmt.Errorf("Error reading l2forward: %v", err)
		}
	}

	if err = d.Set("icmp_send_redirect", flattenSystemInterfaceIcmpSendRedirect(o["icmp-send-redirect"], d, "icmp_send_redirect")); err != nil {
		if !fortiAPIPatch(o["icmp-send-redirect"]) {
			return fmt.Errorf("Error reading icmp_send_redirect: %v", err)
		}
	}

	if err = d.Set("icmp_accept_redirect", flattenSystemInterfaceIcmpAcceptRedirect(o["icmp-accept-redirect"], d, "icmp_accept_redirect")); err != nil {
		if !fortiAPIPatch(o["icmp-accept-redirect"]) {
			return fmt.Errorf("Error reading icmp_accept_redirect: %v", err)
		}
	}

	if err = d.Set("vlanforward", flattenSystemInterfaceVlanforward(o["vlanforward"], d, "vlanforward")); err != nil {
		if !fortiAPIPatch(o["vlanforward"]) {
			return fmt.Errorf("Error reading vlanforward: %v", err)
		}
	}

	if err = d.Set("stpforward", flattenSystemInterfaceStpforward(o["stpforward"], d, "stpforward")); err != nil {
		if !fortiAPIPatch(o["stpforward"]) {
			return fmt.Errorf("Error reading stpforward: %v", err)
		}
	}

	if err = d.Set("stpforward_mode", flattenSystemInterfaceStpforwardMode(o["stpforward-mode"], d, "stpforward_mode")); err != nil {
		if !fortiAPIPatch(o["stpforward-mode"]) {
			return fmt.Errorf("Error reading stpforward_mode: %v", err)
		}
	}

	if err = d.Set("ips_sniffer_mode", flattenSystemInterfaceIpsSnifferMode(o["ips-sniffer-mode"], d, "ips_sniffer_mode")); err != nil {
		if !fortiAPIPatch(o["ips-sniffer-mode"]) {
			return fmt.Errorf("Error reading ips_sniffer_mode: %v", err)
		}
	}

	if err = d.Set("ident_accept", flattenSystemInterfaceIdentAccept(o["ident-accept"], d, "ident_accept")); err != nil {
		if !fortiAPIPatch(o["ident-accept"]) {
			return fmt.Errorf("Error reading ident_accept: %v", err)
		}
	}

	if err = d.Set("ipmac", flattenSystemInterfaceIpmac(o["ipmac"], d, "ipmac")); err != nil {
		if !fortiAPIPatch(o["ipmac"]) {
			return fmt.Errorf("Error reading ipmac: %v", err)
		}
	}

	if err = d.Set("subst", flattenSystemInterfaceSubst(o["subst"], d, "subst")); err != nil {
		if !fortiAPIPatch(o["subst"]) {
			return fmt.Errorf("Error reading subst: %v", err)
		}
	}

	if err = d.Set("macaddr", flattenSystemInterfaceMacaddr(o["macaddr"], d, "macaddr")); err != nil {
		if !fortiAPIPatch(o["macaddr"]) {
			return fmt.Errorf("Error reading macaddr: %v", err)
		}
	}

	if err = d.Set("substitute_dst_mac", flattenSystemInterfaceSubstituteDstMac(o["substitute-dst-mac"], d, "substitute_dst_mac")); err != nil {
		if !fortiAPIPatch(o["substitute-dst-mac"]) {
			return fmt.Errorf("Error reading substitute_dst_mac: %v", err)
		}
	}

	if err = d.Set("speed", flattenSystemInterfaceSpeed(o["speed"], d, "speed")); err != nil {
		if !fortiAPIPatch(o["speed"]) {
			return fmt.Errorf("Error reading speed: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemInterfaceStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("netbios_forward", flattenSystemInterfaceNetbiosForward(o["netbios-forward"], d, "netbios_forward")); err != nil {
		if !fortiAPIPatch(o["netbios-forward"]) {
			return fmt.Errorf("Error reading netbios_forward: %v", err)
		}
	}

	if err = d.Set("wins_ip", flattenSystemInterfaceWinsIp(o["wins-ip"], d, "wins_ip")); err != nil {
		if !fortiAPIPatch(o["wins-ip"]) {
			return fmt.Errorf("Error reading wins_ip: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemInterfaceType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("dedicated_to", flattenSystemInterfaceDedicatedTo(o["dedicated-to"], d, "dedicated_to")); err != nil {
		if !fortiAPIPatch(o["dedicated-to"]) {
			return fmt.Errorf("Error reading dedicated_to: %v", err)
		}
	}

	if err = d.Set("trust_ip_1", flattenSystemInterfaceTrustIp1(o["trust-ip-1"], d, "trust_ip_1")); err != nil {
		if !fortiAPIPatch(o["trust-ip-1"]) {
			return fmt.Errorf("Error reading trust_ip_1: %v", err)
		}
	}

	if err = d.Set("trust_ip_2", flattenSystemInterfaceTrustIp2(o["trust-ip-2"], d, "trust_ip_2")); err != nil {
		if !fortiAPIPatch(o["trust-ip-2"]) {
			return fmt.Errorf("Error reading trust_ip_2: %v", err)
		}
	}

	if err = d.Set("trust_ip_3", flattenSystemInterfaceTrustIp3(o["trust-ip-3"], d, "trust_ip_3")); err != nil {
		if !fortiAPIPatch(o["trust-ip-3"]) {
			return fmt.Errorf("Error reading trust_ip_3: %v", err)
		}
	}

	if err = d.Set("trust_ip6_1", flattenSystemInterfaceTrustIp61(o["trust-ip6-1"], d, "trust_ip6_1")); err != nil {
		if !fortiAPIPatch(o["trust-ip6-1"]) {
			return fmt.Errorf("Error reading trust_ip6_1: %v", err)
		}
	}

	if err = d.Set("trust_ip6_2", flattenSystemInterfaceTrustIp62(o["trust-ip6-2"], d, "trust_ip6_2")); err != nil {
		if !fortiAPIPatch(o["trust-ip6-2"]) {
			return fmt.Errorf("Error reading trust_ip6_2: %v", err)
		}
	}

	if err = d.Set("trust_ip6_3", flattenSystemInterfaceTrustIp63(o["trust-ip6-3"], d, "trust_ip6_3")); err != nil {
		if !fortiAPIPatch(o["trust-ip6-3"]) {
			return fmt.Errorf("Error reading trust_ip6_3: %v", err)
		}
	}

	if err = d.Set("mtu_override", flattenSystemInterfaceMtuOverride(o["mtu-override"], d, "mtu_override")); err != nil {
		if !fortiAPIPatch(o["mtu-override"]) {
			return fmt.Errorf("Error reading mtu_override: %v", err)
		}
	}

	if err = d.Set("mtu", flattenSystemInterfaceMtu(o["mtu"], d, "mtu")); err != nil {
		if !fortiAPIPatch(o["mtu"]) {
			return fmt.Errorf("Error reading mtu: %v", err)
		}
	}

	if err = d.Set("wccp", flattenSystemInterfaceWccp(o["wccp"], d, "wccp")); err != nil {
		if !fortiAPIPatch(o["wccp"]) {
			return fmt.Errorf("Error reading wccp: %v", err)
		}
	}

	if err = d.Set("netflow_sampler", flattenSystemInterfaceNetflowSampler(o["netflow-sampler"], d, "netflow_sampler")); err != nil {
		if !fortiAPIPatch(o["netflow-sampler"]) {
			return fmt.Errorf("Error reading netflow_sampler: %v", err)
		}
	}

	if err = d.Set("sflow_sampler", flattenSystemInterfaceSflowSampler(o["sflow-sampler"], d, "sflow_sampler")); err != nil {
		if !fortiAPIPatch(o["sflow-sampler"]) {
			return fmt.Errorf("Error reading sflow_sampler: %v", err)
		}
	}

	if err = d.Set("drop_overlapped_fragment", flattenSystemInterfaceDropOverlappedFragment(o["drop-overlapped-fragment"], d, "drop_overlapped_fragment")); err != nil {
		if !fortiAPIPatch(o["drop-overlapped-fragment"]) {
			return fmt.Errorf("Error reading drop_overlapped_fragment: %v", err)
		}
	}

	if err = d.Set("drop_fragment", flattenSystemInterfaceDropFragment(o["drop-fragment"], d, "drop_fragment")); err != nil {
		if !fortiAPIPatch(o["drop-fragment"]) {
			return fmt.Errorf("Error reading drop_fragment: %v", err)
		}
	}

	if err = d.Set("scan_botnet_connections", flattenSystemInterfaceScanBotnetConnections(o["scan-botnet-connections"], d, "scan_botnet_connections")); err != nil {
		if !fortiAPIPatch(o["scan-botnet-connections"]) {
			return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
		}
	}

	if err = d.Set("src_check", flattenSystemInterfaceSrcCheck(o["src-check"], d, "src_check")); err != nil {
		if !fortiAPIPatch(o["src-check"]) {
			return fmt.Errorf("Error reading src_check: %v", err)
		}
	}

	if err = d.Set("sample_rate", flattenSystemInterfaceSampleRate(o["sample-rate"], d, "sample_rate")); err != nil {
		if !fortiAPIPatch(o["sample-rate"]) {
			return fmt.Errorf("Error reading sample_rate: %v", err)
		}
	}

	if err = d.Set("polling_interval", flattenSystemInterfacePollingInterval(o["polling-interval"], d, "polling_interval")); err != nil {
		if !fortiAPIPatch(o["polling-interval"]) {
			return fmt.Errorf("Error reading polling_interval: %v", err)
		}
	}

	if err = d.Set("sample_direction", flattenSystemInterfaceSampleDirection(o["sample-direction"], d, "sample_direction")); err != nil {
		if !fortiAPIPatch(o["sample-direction"]) {
			return fmt.Errorf("Error reading sample_direction: %v", err)
		}
	}

	if err = d.Set("explicit_web_proxy", flattenSystemInterfaceExplicitWebProxy(o["explicit-web-proxy"], d, "explicit_web_proxy")); err != nil {
		if !fortiAPIPatch(o["explicit-web-proxy"]) {
			return fmt.Errorf("Error reading explicit_web_proxy: %v", err)
		}
	}

	if err = d.Set("explicit_ftp_proxy", flattenSystemInterfaceExplicitFtpProxy(o["explicit-ftp-proxy"], d, "explicit_ftp_proxy")); err != nil {
		if !fortiAPIPatch(o["explicit-ftp-proxy"]) {
			return fmt.Errorf("Error reading explicit_ftp_proxy: %v", err)
		}
	}

	if err = d.Set("proxy_captive_portal", flattenSystemInterfaceProxyCaptivePortal(o["proxy-captive-portal"], d, "proxy_captive_portal")); err != nil {
		if !fortiAPIPatch(o["proxy-captive-portal"]) {
			return fmt.Errorf("Error reading proxy_captive_portal: %v", err)
		}
	}

	if err = d.Set("tcp_mss", flattenSystemInterfaceTcpMss(o["tcp-mss"], d, "tcp_mss")); err != nil {
		if !fortiAPIPatch(o["tcp-mss"]) {
			return fmt.Errorf("Error reading tcp_mss: %v", err)
		}
	}

	if err = d.Set("inbandwidth", flattenSystemInterfaceInbandwidth(o["inbandwidth"], d, "inbandwidth")); err != nil {
		if !fortiAPIPatch(o["inbandwidth"]) {
			return fmt.Errorf("Error reading inbandwidth: %v", err)
		}
	}

	if err = d.Set("outbandwidth", flattenSystemInterfaceOutbandwidth(o["outbandwidth"], d, "outbandwidth")); err != nil {
		if !fortiAPIPatch(o["outbandwidth"]) {
			return fmt.Errorf("Error reading outbandwidth: %v", err)
		}
	}

	if err = d.Set("egress_shaping_profile", flattenSystemInterfaceEgressShapingProfile(o["egress-shaping-profile"], d, "egress_shaping_profile")); err != nil {
		if !fortiAPIPatch(o["egress-shaping-profile"]) {
			return fmt.Errorf("Error reading egress_shaping_profile: %v", err)
		}
	}

	if err = d.Set("disconnect_threshold", flattenSystemInterfaceDisconnectThreshold(o["disconnect-threshold"], d, "disconnect_threshold")); err != nil {
		if !fortiAPIPatch(o["disconnect-threshold"]) {
			return fmt.Errorf("Error reading disconnect_threshold: %v", err)
		}
	}

	if err = d.Set("spillover_threshold", flattenSystemInterfaceSpilloverThreshold(o["spillover-threshold"], d, "spillover_threshold")); err != nil {
		if !fortiAPIPatch(o["spillover-threshold"]) {
			return fmt.Errorf("Error reading spillover_threshold: %v", err)
		}
	}

	if err = d.Set("ingress_spillover_threshold", flattenSystemInterfaceIngressSpilloverThreshold(o["ingress-spillover-threshold"], d, "ingress_spillover_threshold")); err != nil {
		if !fortiAPIPatch(o["ingress-spillover-threshold"]) {
			return fmt.Errorf("Error reading ingress_spillover_threshold: %v", err)
		}
	}

	if err = d.Set("weight", flattenSystemInterfaceWeight(o["weight"], d, "weight")); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemInterfaceInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("external", flattenSystemInterfaceExternal(o["external"], d, "external")); err != nil {
		if !fortiAPIPatch(o["external"]) {
			return fmt.Errorf("Error reading external: %v", err)
		}
	}

	if err = d.Set("vlanid", flattenSystemInterfaceVlanid(o["vlanid"], d, "vlanid")); err != nil {
		if !fortiAPIPatch(o["vlanid"]) {
			return fmt.Errorf("Error reading vlanid: %v", err)
		}
	}

	if err = d.Set("forward_domain", flattenSystemInterfaceForwardDomain(o["forward-domain"], d, "forward_domain")); err != nil {
		if !fortiAPIPatch(o["forward-domain"]) {
			return fmt.Errorf("Error reading forward_domain: %v", err)
		}
	}

	if err = d.Set("remote_ip", flattenSystemInterfaceRemoteIp(o["remote-ip"], d, "remote_ip")); err != nil {
		if !fortiAPIPatch(o["remote-ip"]) {
			return fmt.Errorf("Error reading remote_ip: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("member", flattenSystemInterfaceMember(o["member"], d, "member")); err != nil {
			if !fortiAPIPatch(o["member"]) {
				return fmt.Errorf("Error reading member: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("member"); ok {
			if err = d.Set("member", flattenSystemInterfaceMember(o["member"], d, "member")); err != nil {
				if !fortiAPIPatch(o["member"]) {
					return fmt.Errorf("Error reading member: %v", err)
				}
			}
		}
	}

	if err = d.Set("lacp_mode", flattenSystemInterfaceLacpMode(o["lacp-mode"], d, "lacp_mode")); err != nil {
		if !fortiAPIPatch(o["lacp-mode"]) {
			return fmt.Errorf("Error reading lacp_mode: %v", err)
		}
	}

	if err = d.Set("lacp_ha_slave", flattenSystemInterfaceLacpHaSlave(o["lacp-ha-slave"], d, "lacp_ha_slave")); err != nil {
		if !fortiAPIPatch(o["lacp-ha-slave"]) {
			return fmt.Errorf("Error reading lacp_ha_slave: %v", err)
		}
	}

	if err = d.Set("lacp_speed", flattenSystemInterfaceLacpSpeed(o["lacp-speed"], d, "lacp_speed")); err != nil {
		if !fortiAPIPatch(o["lacp-speed"]) {
			return fmt.Errorf("Error reading lacp_speed: %v", err)
		}
	}

	if err = d.Set("min_links", flattenSystemInterfaceMinLinks(o["min-links"], d, "min_links")); err != nil {
		if !fortiAPIPatch(o["min-links"]) {
			return fmt.Errorf("Error reading min_links: %v", err)
		}
	}

	if err = d.Set("min_links_down", flattenSystemInterfaceMinLinksDown(o["min-links-down"], d, "min_links_down")); err != nil {
		if !fortiAPIPatch(o["min-links-down"]) {
			return fmt.Errorf("Error reading min_links_down: %v", err)
		}
	}

	if err = d.Set("algorithm", flattenSystemInterfaceAlgorithm(o["algorithm"], d, "algorithm")); err != nil {
		if !fortiAPIPatch(o["algorithm"]) {
			return fmt.Errorf("Error reading algorithm: %v", err)
		}
	}

	if err = d.Set("link_up_delay", flattenSystemInterfaceLinkUpDelay(o["link-up-delay"], d, "link_up_delay")); err != nil {
		if !fortiAPIPatch(o["link-up-delay"]) {
			return fmt.Errorf("Error reading link_up_delay: %v", err)
		}
	}

	if err = d.Set("priority_override", flattenSystemInterfacePriorityOverride(o["priority-override"], d, "priority_override")); err != nil {
		if !fortiAPIPatch(o["priority-override"]) {
			return fmt.Errorf("Error reading priority_override: %v", err)
		}
	}

	if err = d.Set("aggregate", flattenSystemInterfaceAggregate(o["aggregate"], d, "aggregate")); err != nil {
		if !fortiAPIPatch(o["aggregate"]) {
			return fmt.Errorf("Error reading aggregate: %v", err)
		}
	}

	if err = d.Set("redundant_interface", flattenSystemInterfaceRedundantInterface(o["redundant-interface"], d, "redundant_interface")); err != nil {
		if !fortiAPIPatch(o["redundant-interface"]) {
			return fmt.Errorf("Error reading redundant_interface: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("managed_device", flattenSystemInterfaceManagedDevice(o["managed-device"], d, "managed_device")); err != nil {
			if !fortiAPIPatch(o["managed-device"]) {
				return fmt.Errorf("Error reading managed_device: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("managed_device"); ok {
			if err = d.Set("managed_device", flattenSystemInterfaceManagedDevice(o["managed-device"], d, "managed_device")); err != nil {
				if !fortiAPIPatch(o["managed-device"]) {
					return fmt.Errorf("Error reading managed_device: %v", err)
				}
			}
		}
	}

	if err = d.Set("devindex", flattenSystemInterfaceDevindex(o["devindex"], d, "devindex")); err != nil {
		if !fortiAPIPatch(o["devindex"]) {
			return fmt.Errorf("Error reading devindex: %v", err)
		}
	}

	if err = d.Set("vindex", flattenSystemInterfaceVindex(o["vindex"], d, "vindex")); err != nil {
		if !fortiAPIPatch(o["vindex"]) {
			return fmt.Errorf("Error reading vindex: %v", err)
		}
	}

	if err = d.Set("switch", flattenSystemInterfaceSwitch(o["switch"], d, "switch")); err != nil {
		if !fortiAPIPatch(o["switch"]) {
			return fmt.Errorf("Error reading switch: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemInterfaceDescription(o["description"], d, "description")); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("alias", flattenSystemInterfaceAlias(o["alias"], d, "alias")); err != nil {
		if !fortiAPIPatch(o["alias"]) {
			return fmt.Errorf("Error reading alias: %v", err)
		}
	}

	if err = d.Set("security_mode", flattenSystemInterfaceSecurityMode(o["security-mode"], d, "security_mode")); err != nil {
		if !fortiAPIPatch(o["security-mode"]) {
			return fmt.Errorf("Error reading security_mode: %v", err)
		}
	}

	if err = d.Set("captive_portal", flattenSystemInterfaceCaptivePortal(o["captive-portal"], d, "captive_portal")); err != nil {
		if !fortiAPIPatch(o["captive-portal"]) {
			return fmt.Errorf("Error reading captive_portal: %v", err)
		}
	}

	if err = d.Set("security_mac_auth_bypass", flattenSystemInterfaceSecurityMacAuthBypass(o["security-mac-auth-bypass"], d, "security_mac_auth_bypass")); err != nil {
		if !fortiAPIPatch(o["security-mac-auth-bypass"]) {
			return fmt.Errorf("Error reading security_mac_auth_bypass: %v", err)
		}
	}

	if err = d.Set("security_external_web", flattenSystemInterfaceSecurityExternalWeb(o["security-external-web"], d, "security_external_web")); err != nil {
		if !fortiAPIPatch(o["security-external-web"]) {
			return fmt.Errorf("Error reading security_external_web: %v", err)
		}
	}

	if err = d.Set("security_external_logout", flattenSystemInterfaceSecurityExternalLogout(o["security-external-logout"], d, "security_external_logout")); err != nil {
		if !fortiAPIPatch(o["security-external-logout"]) {
			return fmt.Errorf("Error reading security_external_logout: %v", err)
		}
	}

	if err = d.Set("replacemsg_override_group", flattenSystemInterfaceReplacemsgOverrideGroup(o["replacemsg-override-group"], d, "replacemsg_override_group")); err != nil {
		if !fortiAPIPatch(o["replacemsg-override-group"]) {
			return fmt.Errorf("Error reading replacemsg_override_group: %v", err)
		}
	}

	if err = d.Set("security_redirect_url", flattenSystemInterfaceSecurityRedirectUrl(o["security-redirect-url"], d, "security_redirect_url")); err != nil {
		if !fortiAPIPatch(o["security-redirect-url"]) {
			return fmt.Errorf("Error reading security_redirect_url: %v", err)
		}
	}

	if err = d.Set("security_exempt_list", flattenSystemInterfaceSecurityExemptList(o["security-exempt-list"], d, "security_exempt_list")); err != nil {
		if !fortiAPIPatch(o["security-exempt-list"]) {
			return fmt.Errorf("Error reading security_exempt_list: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("security_groups", flattenSystemInterfaceSecurityGroups(o["security-groups"], d, "security_groups")); err != nil {
			if !fortiAPIPatch(o["security-groups"]) {
				return fmt.Errorf("Error reading security_groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("security_groups"); ok {
			if err = d.Set("security_groups", flattenSystemInterfaceSecurityGroups(o["security-groups"], d, "security_groups")); err != nil {
				if !fortiAPIPatch(o["security-groups"]) {
					return fmt.Errorf("Error reading security_groups: %v", err)
				}
			}
		}
	}

	if err = d.Set("device_identification", flattenSystemInterfaceDeviceIdentification(o["device-identification"], d, "device_identification")); err != nil {
		if !fortiAPIPatch(o["device-identification"]) {
			return fmt.Errorf("Error reading device_identification: %v", err)
		}
	}

	if err = d.Set("device_user_identification", flattenSystemInterfaceDeviceUserIdentification(o["device-user-identification"], d, "device_user_identification")); err != nil {
		if !fortiAPIPatch(o["device-user-identification"]) {
			return fmt.Errorf("Error reading device_user_identification: %v", err)
		}
	}

	if err = d.Set("device_identification_active_scan", flattenSystemInterfaceDeviceIdentificationActiveScan(o["device-identification-active-scan"], d, "device_identification_active_scan")); err != nil {
		if !fortiAPIPatch(o["device-identification-active-scan"]) {
			return fmt.Errorf("Error reading device_identification_active_scan: %v", err)
		}
	}

	if err = d.Set("device_access_list", flattenSystemInterfaceDeviceAccessList(o["device-access-list"], d, "device_access_list")); err != nil {
		if !fortiAPIPatch(o["device-access-list"]) {
			return fmt.Errorf("Error reading device_access_list: %v", err)
		}
	}

	if err = d.Set("device_netscan", flattenSystemInterfaceDeviceNetscan(o["device-netscan"], d, "device_netscan")); err != nil {
		if !fortiAPIPatch(o["device-netscan"]) {
			return fmt.Errorf("Error reading device_netscan: %v", err)
		}
	}

	if err = d.Set("lldp_reception", flattenSystemInterfaceLldpReception(o["lldp-reception"], d, "lldp_reception")); err != nil {
		if !fortiAPIPatch(o["lldp-reception"]) {
			return fmt.Errorf("Error reading lldp_reception: %v", err)
		}
	}

	if err = d.Set("lldp_transmission", flattenSystemInterfaceLldpTransmission(o["lldp-transmission"], d, "lldp_transmission")); err != nil {
		if !fortiAPIPatch(o["lldp-transmission"]) {
			return fmt.Errorf("Error reading lldp_transmission: %v", err)
		}
	}

	if err = d.Set("lldp_network_policy", flattenSystemInterfaceLldpNetworkPolicy(o["lldp-network-policy"], d, "lldp_network_policy")); err != nil {
		if !fortiAPIPatch(o["lldp-network-policy"]) {
			return fmt.Errorf("Error reading lldp_network_policy: %v", err)
		}
	}

	if err = d.Set("fortiheartbeat", flattenSystemInterfaceFortiheartbeat(o["fortiheartbeat"], d, "fortiheartbeat")); err != nil {
		if !fortiAPIPatch(o["fortiheartbeat"]) {
			return fmt.Errorf("Error reading fortiheartbeat: %v", err)
		}
	}

	if err = d.Set("broadcast_forticlient_discovery", flattenSystemInterfaceBroadcastForticlientDiscovery(o["broadcast-forticlient-discovery"], d, "broadcast_forticlient_discovery")); err != nil {
		if !fortiAPIPatch(o["broadcast-forticlient-discovery"]) {
			return fmt.Errorf("Error reading broadcast_forticlient_discovery: %v", err)
		}
	}

	if err = d.Set("endpoint_compliance", flattenSystemInterfaceEndpointCompliance(o["endpoint-compliance"], d, "endpoint_compliance")); err != nil {
		if !fortiAPIPatch(o["endpoint-compliance"]) {
			return fmt.Errorf("Error reading endpoint_compliance: %v", err)
		}
	}

	if err = d.Set("estimated_upstream_bandwidth", flattenSystemInterfaceEstimatedUpstreamBandwidth(o["estimated-upstream-bandwidth"], d, "estimated_upstream_bandwidth")); err != nil {
		if !fortiAPIPatch(o["estimated-upstream-bandwidth"]) {
			return fmt.Errorf("Error reading estimated_upstream_bandwidth: %v", err)
		}
	}

	if err = d.Set("estimated_downstream_bandwidth", flattenSystemInterfaceEstimatedDownstreamBandwidth(o["estimated-downstream-bandwidth"], d, "estimated_downstream_bandwidth")); err != nil {
		if !fortiAPIPatch(o["estimated-downstream-bandwidth"]) {
			return fmt.Errorf("Error reading estimated_downstream_bandwidth: %v", err)
		}
	}

	if err = d.Set("vrrp_virtual_mac", flattenSystemInterfaceVrrpVirtualMac(o["vrrp-virtual-mac"], d, "vrrp_virtual_mac")); err != nil {
		if !fortiAPIPatch(o["vrrp-virtual-mac"]) {
			return fmt.Errorf("Error reading vrrp_virtual_mac: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("vrrp", flattenSystemInterfaceVrrp(o["vrrp"], d, "vrrp")); err != nil {
			if !fortiAPIPatch(o["vrrp"]) {
				return fmt.Errorf("Error reading vrrp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vrrp"); ok {
			if err = d.Set("vrrp", flattenSystemInterfaceVrrp(o["vrrp"], d, "vrrp")); err != nil {
				if !fortiAPIPatch(o["vrrp"]) {
					return fmt.Errorf("Error reading vrrp: %v", err)
				}
			}
		}
	}

	if err = d.Set("role", flattenSystemInterfaceRole(o["role"], d, "role")); err != nil {
		if !fortiAPIPatch(o["role"]) {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("snmp_index", flattenSystemInterfaceSnmpIndex(o["snmp-index"], d, "snmp_index")); err != nil {
		if !fortiAPIPatch(o["snmp-index"]) {
			return fmt.Errorf("Error reading snmp_index: %v", err)
		}
	}

	if err = d.Set("secondary_ip", flattenSystemInterfaceSecondaryIp(o["secondary-IP"], d, "secondary_ip")); err != nil {
		if !fortiAPIPatch(o["secondary-IP"]) {
			return fmt.Errorf("Error reading secondary_ip: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("secondaryip", flattenSystemInterfaceSecondaryip(o["secondaryip"], d, "secondaryip")); err != nil {
			if !fortiAPIPatch(o["secondaryip"]) {
				return fmt.Errorf("Error reading secondaryip: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("secondaryip"); ok {
			if err = d.Set("secondaryip", flattenSystemInterfaceSecondaryip(o["secondaryip"], d, "secondaryip")); err != nil {
				if !fortiAPIPatch(o["secondaryip"]) {
					return fmt.Errorf("Error reading secondaryip: %v", err)
				}
			}
		}
	}

	if err = d.Set("preserve_session_route", flattenSystemInterfacePreserveSessionRoute(o["preserve-session-route"], d, "preserve_session_route")); err != nil {
		if !fortiAPIPatch(o["preserve-session-route"]) {
			return fmt.Errorf("Error reading preserve_session_route: %v", err)
		}
	}

	if err = d.Set("auto_auth_extension_device", flattenSystemInterfaceAutoAuthExtensionDevice(o["auto-auth-extension-device"], d, "auto_auth_extension_device")); err != nil {
		if !fortiAPIPatch(o["auto-auth-extension-device"]) {
			return fmt.Errorf("Error reading auto_auth_extension_device: %v", err)
		}
	}

	if err = d.Set("ap_discover", flattenSystemInterfaceApDiscover(o["ap-discover"], d, "ap_discover")); err != nil {
		if !fortiAPIPatch(o["ap-discover"]) {
			return fmt.Errorf("Error reading ap_discover: %v", err)
		}
	}

	if err = d.Set("fortilink_stacking", flattenSystemInterfaceFortilinkStacking(o["fortilink-stacking"], d, "fortilink_stacking")); err != nil {
		if !fortiAPIPatch(o["fortilink-stacking"]) {
			return fmt.Errorf("Error reading fortilink_stacking: %v", err)
		}
	}

	if err = d.Set("fortilink_split_interface", flattenSystemInterfaceFortilinkSplitInterface(o["fortilink-split-interface"], d, "fortilink_split_interface")); err != nil {
		if !fortiAPIPatch(o["fortilink-split-interface"]) {
			return fmt.Errorf("Error reading fortilink_split_interface: %v", err)
		}
	}

	if err = d.Set("internal", flattenSystemInterfaceInternal(o["internal"], d, "internal")); err != nil {
		if !fortiAPIPatch(o["internal"]) {
			return fmt.Errorf("Error reading internal: %v", err)
		}
	}

	if err = d.Set("fortilink_backup_link", flattenSystemInterfaceFortilinkBackupLink(o["fortilink-backup-link"], d, "fortilink_backup_link")); err != nil {
		if !fortiAPIPatch(o["fortilink-backup-link"]) {
			return fmt.Errorf("Error reading fortilink_backup_link: %v", err)
		}
	}

	if err = d.Set("switch_controller_access_vlan", flattenSystemInterfaceSwitchControllerAccessVlan(o["switch-controller-access-vlan"], d, "switch_controller_access_vlan")); err != nil {
		if !fortiAPIPatch(o["switch-controller-access-vlan"]) {
			return fmt.Errorf("Error reading switch_controller_access_vlan: %v", err)
		}
	}

	if err = d.Set("switch_controller_traffic_policy", flattenSystemInterfaceSwitchControllerTrafficPolicy(o["switch-controller-traffic-policy"], d, "switch_controller_traffic_policy")); err != nil {
		if !fortiAPIPatch(o["switch-controller-traffic-policy"]) {
			return fmt.Errorf("Error reading switch_controller_traffic_policy: %v", err)
		}
	}

	if err = d.Set("switch_controller_igmp_snooping", flattenSystemInterfaceSwitchControllerIgmpSnooping(o["switch-controller-igmp-snooping"], d, "switch_controller_igmp_snooping")); err != nil {
		if !fortiAPIPatch(o["switch-controller-igmp-snooping"]) {
			return fmt.Errorf("Error reading switch_controller_igmp_snooping: %v", err)
		}
	}

	if err = d.Set("switch_controller_dhcp_snooping", flattenSystemInterfaceSwitchControllerDhcpSnooping(o["switch-controller-dhcp-snooping"], d, "switch_controller_dhcp_snooping")); err != nil {
		if !fortiAPIPatch(o["switch-controller-dhcp-snooping"]) {
			return fmt.Errorf("Error reading switch_controller_dhcp_snooping: %v", err)
		}
	}

	if err = d.Set("switch_controller_dhcp_snooping_verify_mac", flattenSystemInterfaceSwitchControllerDhcpSnoopingVerifyMac(o["switch-controller-dhcp-snooping-verify-mac"], d, "switch_controller_dhcp_snooping_verify_mac")); err != nil {
		if !fortiAPIPatch(o["switch-controller-dhcp-snooping-verify-mac"]) {
			return fmt.Errorf("Error reading switch_controller_dhcp_snooping_verify_mac: %v", err)
		}
	}

	if err = d.Set("switch_controller_dhcp_snooping_option82", flattenSystemInterfaceSwitchControllerDhcpSnoopingOption82(o["switch-controller-dhcp-snooping-option82"], d, "switch_controller_dhcp_snooping_option82")); err != nil {
		if !fortiAPIPatch(o["switch-controller-dhcp-snooping-option82"]) {
			return fmt.Errorf("Error reading switch_controller_dhcp_snooping_option82: %v", err)
		}
	}

	if err = d.Set("switch_controller_arp_inspection", flattenSystemInterfaceSwitchControllerArpInspection(o["switch-controller-arp-inspection"], d, "switch_controller_arp_inspection")); err != nil {
		if !fortiAPIPatch(o["switch-controller-arp-inspection"]) {
			return fmt.Errorf("Error reading switch_controller_arp_inspection: %v", err)
		}
	}

	if err = d.Set("switch_controller_learning_limit", flattenSystemInterfaceSwitchControllerLearningLimit(o["switch-controller-learning-limit"], d, "switch_controller_learning_limit")); err != nil {
		if !fortiAPIPatch(o["switch-controller-learning-limit"]) {
			return fmt.Errorf("Error reading switch_controller_learning_limit: %v", err)
		}
	}

	if err = d.Set("color", flattenSystemInterfaceColor(o["color"], d, "color")); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("tagging", flattenSystemInterfaceTagging(o["tagging"], d, "tagging")); err != nil {
			if !fortiAPIPatch(o["tagging"]) {
				return fmt.Errorf("Error reading tagging: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tagging"); ok {
			if err = d.Set("tagging", flattenSystemInterfaceTagging(o["tagging"], d, "tagging")); err != nil {
				if !fortiAPIPatch(o["tagging"]) {
					return fmt.Errorf("Error reading tagging: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ipv6", flattenSystemInterfaceIpv6(o["ipv6"], d, "ipv6")); err != nil {
			if !fortiAPIPatch(o["ipv6"]) {
				return fmt.Errorf("Error reading ipv6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ipv6"); ok {
			if err = d.Set("ipv6", flattenSystemInterfaceIpv6(o["ipv6"], d, "ipv6")); err != nil {
				if !fortiAPIPatch(o["ipv6"]) {
					return fmt.Errorf("Error reading ipv6: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceCliConnStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceFortilink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfacePriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDhcpRelayService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDhcpRelayIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDhcpRelayType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDhcpRelayAgentOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceManagementIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceGwdetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfacePingServStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDetectserver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDetectprotocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceHaPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceFailDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceFailDetectOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceFailAlertMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceFailActionOnExtender(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceFailAlertInterfaces(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandSystemInterfaceFailAlertInterfacesName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceFailAlertInterfacesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDhcpClientIdentifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDhcpRenewTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpunnumbered(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfacePppoeUnnumberedNegotiate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfacePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDetectedPeerMtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDiscRetryTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfacePadtRetryTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceAcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceLcpEchoInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceLcpMaxEchoFails(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDefaultgw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDnsServerOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfacePptpClient(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfacePptpUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfacePptpPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfacePptpServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfacePptpAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfacePptpTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceArpforward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceNdiscforward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceBroadcastForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceBfd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceBfdDesiredMinTx(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceBfdDetectMult(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceBfdRequiredMinRx(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceL2Forward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIcmpSendRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIcmpAcceptRedirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVlanforward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceStpforward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceStpforwardMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpsSnifferMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIdentAccept(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpmac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSubst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceMacaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSubstituteDstMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSpeed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceNetbiosForward(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceWinsIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDedicatedTo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceTrustIp1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceTrustIp2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceTrustIp3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceTrustIp61(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceTrustIp62(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceTrustIp63(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceMtuOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceMtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceWccp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceNetflowSampler(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSflowSampler(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDropOverlappedFragment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDropFragment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceScanBotnetConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSrcCheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSampleRate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfacePollingInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSampleDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceExplicitWebProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceExplicitFtpProxy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceProxyCaptivePortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceTcpMss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceInbandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceOutbandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceEgressShapingProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDisconnectThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIngressSpilloverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceExternal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVlanid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceForwardDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceRemoteIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface-name"], _ = expandSystemInterfaceMemberInterfaceName(d, i["interface_name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceMemberInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceLacpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceLacpHaSlave(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceLacpSpeed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceMinLinks(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceMinLinksDown(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceLinkUpDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfacePriorityOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceAggregate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceRedundantInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceManagedDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandSystemInterfaceManagedDeviceName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceManagedDeviceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDevindex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVindex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceAlias(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSecurityMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceCaptivePortal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSecurityMacAuthBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSecurityExternalWeb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSecurityExternalLogout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceReplacemsgOverrideGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSecurityRedirectUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSecurityExemptList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSecurityGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandSystemInterfaceSecurityGroupsName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceSecurityGroupsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDeviceIdentification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDeviceUserIdentification(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDeviceIdentificationActiveScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDeviceAccessList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceDeviceNetscan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceLldpReception(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceLldpTransmission(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceLldpNetworkPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceFortiheartbeat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceBroadcastForticlientDiscovery(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceEndpointCompliance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceEstimatedUpstreamBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceEstimatedDownstreamBandwidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpVirtualMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrid"], _ = expandSystemInterfaceVrrpVrid(d, i["vrid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["version"], _ = expandSystemInterfaceVrrpVersion(d, i["version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrgrp"], _ = expandSystemInterfaceVrrpVrgrp(d, i["vrgrp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrip"], _ = expandSystemInterfaceVrrpVrip(d, i["vrip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandSystemInterfaceVrrpPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["adv-interval"], _ = expandSystemInterfaceVrrpAdvInterval(d, i["adv_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-time"], _ = expandSystemInterfaceVrrpStartTime(d, i["start_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["preempt"], _ = expandSystemInterfaceVrrpPreempt(d, i["preempt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["accept-mode"], _ = expandSystemInterfaceVrrpAcceptMode(d, i["accept_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrdst"], _ = expandSystemInterfaceVrrpVrdst(d, i["vrdst"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst_priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrdst-priority"], _ = expandSystemInterfaceVrrpVrdstPriority(d, i["vrdst_priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandSystemInterfaceVrrpStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "proxy_arp"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["proxy-arp"], _ = expandSystemInterfaceVrrpProxyArp(d, i["proxy_arp"], pre_append)
		} else {
			tmp["proxy-arp"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceVrrpVrid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpVrgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpVrip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpAdvInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpStartTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpPreempt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpAcceptMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpVrdst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpVrdstPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpProxyArp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandSystemInterfaceVrrpProxyArpId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandSystemInterfaceVrrpProxyArpIp(d, i["ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceVrrpProxyArpId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceVrrpProxyArpIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceRole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSnmpIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSecondaryIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSecondaryip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandSystemInterfaceSecondaryipId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandSystemInterfaceSecondaryipIp(d, i["ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowaccess"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowaccess"], _ = expandSystemInterfaceSecondaryipAllowaccess(d, i["allowaccess"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gwdetect"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["gwdetect"], _ = expandSystemInterfaceSecondaryipGwdetect(d, i["gwdetect"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ping_serv_status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ping-serv-status"], _ = expandSystemInterfaceSecondaryipPingServStatus(d, i["ping_serv_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectserver"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["detectserver"], _ = expandSystemInterfaceSecondaryipDetectserver(d, i["detectserver"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detectprotocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["detectprotocol"], _ = expandSystemInterfaceSecondaryipDetectprotocol(d, i["detectprotocol"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ha-priority"], _ = expandSystemInterfaceSecondaryipHaPriority(d, i["ha_priority"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceSecondaryipId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSecondaryipIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSecondaryipAllowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSecondaryipGwdetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSecondaryipPingServStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSecondaryipDetectserver(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSecondaryipDetectprotocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSecondaryipHaPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfacePreserveSessionRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceAutoAuthExtensionDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceApDiscover(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceFortilinkStacking(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceFortilinkSplitInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceInternal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceFortilinkBackupLink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSwitchControllerAccessVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSwitchControllerTrafficPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSwitchControllerIgmpSnooping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSwitchControllerDhcpSnooping(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSwitchControllerDhcpSnoopingVerifyMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSwitchControllerDhcpSnoopingOption82(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSwitchControllerArpInspection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceSwitchControllerLearningLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceTagging(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandSystemInterfaceTaggingName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["category"], _ = expandSystemInterfaceTaggingCategory(d, i["category"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tags"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tags"], _ = expandSystemInterfaceTaggingTags(d, i["tags"], pre_append)
		} else {
			tmp["tags"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceTaggingName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceTaggingCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceTaggingTags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandSystemInterfaceTaggingTagsName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceTaggingTagsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ip6_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-mode"], _ = expandSystemInterfaceIpv6Ip6Mode(d, i["ip6_mode"], pre_append)
	}
	pre_append = pre + ".0." + "nd_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["nd-mode"], _ = expandSystemInterfaceIpv6NdMode(d, i["nd_mode"], pre_append)
	}
	pre_append = pre + ".0." + "nd_cert"
	if _, ok := d.GetOk(pre_append); ok {
		result["nd-cert"], _ = expandSystemInterfaceIpv6NdCert(d, i["nd_cert"], pre_append)
	}
	pre_append = pre + ".0." + "nd_security_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["nd-security-level"], _ = expandSystemInterfaceIpv6NdSecurityLevel(d, i["nd_security_level"], pre_append)
	}
	pre_append = pre + ".0." + "nd_timestamp_delta"
	if _, ok := d.GetOk(pre_append); ok {
		result["nd-timestamp-delta"], _ = expandSystemInterfaceIpv6NdTimestampDelta(d, i["nd_timestamp_delta"], pre_append)
	}
	pre_append = pre + ".0." + "nd_timestamp_fuzz"
	if _, ok := d.GetOk(pre_append); ok {
		result["nd-timestamp-fuzz"], _ = expandSystemInterfaceIpv6NdTimestampFuzz(d, i["nd_timestamp_fuzz"], pre_append)
	}
	pre_append = pre + ".0." + "nd_cga_modifier"
	if _, ok := d.GetOk(pre_append); ok {
		result["nd-cga-modifier"], _ = expandSystemInterfaceIpv6NdCgaModifier(d, i["nd_cga_modifier"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_dns_server_override"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-dns-server-override"], _ = expandSystemInterfaceIpv6Ip6DnsServerOverride(d, i["ip6_dns_server_override"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_address"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-address"], _ = expandSystemInterfaceIpv6Ip6Address(d, i["ip6_address"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_extra_addr"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-extra-addr"], _ = expandSystemInterfaceIpv6Ip6ExtraAddr(d, i["ip6_extra_addr"], pre_append)
	} else {
		result["ip6-extra-addr"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ip6_allowaccess"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-allowaccess"], _ = expandSystemInterfaceIpv6Ip6Allowaccess(d, i["ip6_allowaccess"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_send_adv"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-send-adv"], _ = expandSystemInterfaceIpv6Ip6SendAdv(d, i["ip6_send_adv"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_manage_flag"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-manage-flag"], _ = expandSystemInterfaceIpv6Ip6ManageFlag(d, i["ip6_manage_flag"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_other_flag"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-other-flag"], _ = expandSystemInterfaceIpv6Ip6OtherFlag(d, i["ip6_other_flag"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_max_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-max-interval"], _ = expandSystemInterfaceIpv6Ip6MaxInterval(d, i["ip6_max_interval"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_min_interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-min-interval"], _ = expandSystemInterfaceIpv6Ip6MinInterval(d, i["ip6_min_interval"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_link_mtu"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-link-mtu"], _ = expandSystemInterfaceIpv6Ip6LinkMtu(d, i["ip6_link_mtu"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_reachable_time"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-reachable-time"], _ = expandSystemInterfaceIpv6Ip6ReachableTime(d, i["ip6_reachable_time"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_retrans_time"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-retrans-time"], _ = expandSystemInterfaceIpv6Ip6RetransTime(d, i["ip6_retrans_time"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_default_life"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-default-life"], _ = expandSystemInterfaceIpv6Ip6DefaultLife(d, i["ip6_default_life"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_hop_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-hop-limit"], _ = expandSystemInterfaceIpv6Ip6HopLimit(d, i["ip6_hop_limit"], pre_append)
	}
	pre_append = pre + ".0." + "autoconf"
	if _, ok := d.GetOk(pre_append); ok {
		result["autoconf"], _ = expandSystemInterfaceIpv6Autoconf(d, i["autoconf"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_upstream_interface"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-upstream-interface"], _ = expandSystemInterfaceIpv6Ip6UpstreamInterface(d, i["ip6_upstream_interface"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_subnet"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-subnet"], _ = expandSystemInterfaceIpv6Ip6Subnet(d, i["ip6_subnet"], pre_append)
	}
	pre_append = pre + ".0." + "ip6_prefix_list"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-prefix-list"], _ = expandSystemInterfaceIpv6Ip6PrefixList(d, i["ip6_prefix_list"], pre_append)
	} else {
		result["ip6-prefix-list"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "ip6_delegated_prefix_list"
	if _, ok := d.GetOk(pre_append); ok {
		result["ip6-delegated-prefix-list"], _ = expandSystemInterfaceIpv6Ip6DelegatedPrefixList(d, i["ip6_delegated_prefix_list"], pre_append)
	} else {
		result["ip6-delegated-prefix-list"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "dhcp6_relay_service"
	if _, ok := d.GetOk(pre_append); ok {
		result["dhcp6-relay-service"], _ = expandSystemInterfaceIpv6Dhcp6RelayService(d, i["dhcp6_relay_service"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_relay_type"
	if _, ok := d.GetOk(pre_append); ok {
		result["dhcp6-relay-type"], _ = expandSystemInterfaceIpv6Dhcp6RelayType(d, i["dhcp6_relay_type"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_relay_ip"
	if _, ok := d.GetOk(pre_append); ok {
		result["dhcp6-relay-ip"], _ = expandSystemInterfaceIpv6Dhcp6RelayIp(d, i["dhcp6_relay_ip"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_client_options"
	if _, ok := d.GetOk(pre_append); ok {
		result["dhcp6-client-options"], _ = expandSystemInterfaceIpv6Dhcp6ClientOptions(d, i["dhcp6_client_options"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_prefix_delegation"
	if _, ok := d.GetOk(pre_append); ok {
		result["dhcp6-prefix-delegation"], _ = expandSystemInterfaceIpv6Dhcp6PrefixDelegation(d, i["dhcp6_prefix_delegation"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_information_request"
	if _, ok := d.GetOk(pre_append); ok {
		result["dhcp6-information-request"], _ = expandSystemInterfaceIpv6Dhcp6InformationRequest(d, i["dhcp6_information_request"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_prefix_hint"
	if _, ok := d.GetOk(pre_append); ok {
		result["dhcp6-prefix-hint"], _ = expandSystemInterfaceIpv6Dhcp6PrefixHint(d, i["dhcp6_prefix_hint"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_prefix_hint_plt"
	if _, ok := d.GetOk(pre_append); ok {
		result["dhcp6-prefix-hint-plt"], _ = expandSystemInterfaceIpv6Dhcp6PrefixHintPlt(d, i["dhcp6_prefix_hint_plt"], pre_append)
	}
	pre_append = pre + ".0." + "dhcp6_prefix_hint_vlt"
	if _, ok := d.GetOk(pre_append); ok {
		result["dhcp6-prefix-hint-vlt"], _ = expandSystemInterfaceIpv6Dhcp6PrefixHintVlt(d, i["dhcp6_prefix_hint_vlt"], pre_append)
	}
	pre_append = pre + ".0." + "vrrp_virtual_mac6"
	if _, ok := d.GetOk(pre_append); ok {
		result["vrrp-virtual-mac6"], _ = expandSystemInterfaceIpv6VrrpVirtualMac6(d, i["vrrp_virtual_mac6"], pre_append)
	}
	pre_append = pre + ".0." + "vrip6_link_local"
	if _, ok := d.GetOk(pre_append); ok {
		result["vrip6_link_local"], _ = expandSystemInterfaceIpv6Vrip6_Link_Local(d, i["vrip6_link_local"], pre_append)
	}
	pre_append = pre + ".0." + "vrrp6"
	if _, ok := d.GetOk(pre_append); ok {
		result["vrrp6"], _ = expandSystemInterfaceIpv6Vrrp6(d, i["vrrp6"], pre_append)
	} else {
		result["vrrp6"] = make([]string, 0)
	}

	return result, nil
}

func expandSystemInterfaceIpv6Ip6Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6NdMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6NdCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6NdSecurityLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6NdTimestampDelta(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6NdTimestampFuzz(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6NdCgaModifier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6DnsServerOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6ExtraAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix"], _ = expandSystemInterfaceIpv6Ip6ExtraAddrPrefix(d, i["prefix"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceIpv6Ip6ExtraAddrPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6Allowaccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6SendAdv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6ManageFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6OtherFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6MaxInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6MinInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6LinkMtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6ReachableTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6RetransTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6DefaultLife(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6HopLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Autoconf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6UpstreamInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6Subnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6PrefixList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix"], _ = expandSystemInterfaceIpv6Ip6PrefixListPrefix(d, i["prefix"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["autonomous-flag"], _ = expandSystemInterfaceIpv6Ip6PrefixListAutonomousFlag(d, i["autonomous_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["onlink-flag"], _ = expandSystemInterfaceIpv6Ip6PrefixListOnlinkFlag(d, i["onlink_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "valid_life_time"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["valid-life-time"], _ = expandSystemInterfaceIpv6Ip6PrefixListValidLifeTime(d, i["valid_life_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preferred_life_time"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["preferred-life-time"], _ = expandSystemInterfaceIpv6Ip6PrefixListPreferredLifeTime(d, i["preferred_life_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rdnss"], _ = expandSystemInterfaceIpv6Ip6PrefixListRdnss(d, i["rdnss"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dnssl"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dnssl"], _ = expandSystemInterfaceIpv6Ip6PrefixListDnssl(d, i["dnssl"], pre_append)
		} else {
			tmp["dnssl"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceIpv6Ip6PrefixListPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6PrefixListAutonomousFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6PrefixListOnlinkFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6PrefixListValidLifeTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6PrefixListPreferredLifeTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6PrefixListRdnss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6PrefixListDnssl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["domain"], _ = expandSystemInterfaceIpv6Ip6PrefixListDnsslDomain(d, i["domain"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceIpv6Ip6PrefixListDnsslDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6DelegatedPrefixList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix-id"], _ = expandSystemInterfaceIpv6Ip6DelegatedPrefixListPrefixId(d, i["prefix_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "upstream_interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["upstream-interface"], _ = expandSystemInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface(d, i["upstream_interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "autonomous_flag"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["autonomous-flag"], _ = expandSystemInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag(d, i["autonomous_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "onlink_flag"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["onlink-flag"], _ = expandSystemInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag(d, i["onlink_flag"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["subnet"], _ = expandSystemInterfaceIpv6Ip6DelegatedPrefixListSubnet(d, i["subnet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss_service"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rdnss-service"], _ = expandSystemInterfaceIpv6Ip6DelegatedPrefixListRdnssService(d, i["rdnss_service"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rdnss"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rdnss"], _ = expandSystemInterfaceIpv6Ip6DelegatedPrefixListRdnss(d, i["rdnss"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceIpv6Ip6DelegatedPrefixListPrefixId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6DelegatedPrefixListUpstreamInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6DelegatedPrefixListAutonomousFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6DelegatedPrefixListOnlinkFlag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6DelegatedPrefixListSubnet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6DelegatedPrefixListRdnssService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Ip6DelegatedPrefixListRdnss(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6RelayService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6RelayType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6RelayIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6ClientOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6PrefixDelegation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6InformationRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6PrefixHint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6PrefixHintPlt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Dhcp6PrefixHintVlt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6VrrpVirtualMac6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrip6_Link_Local(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrid"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrid"], _ = expandSystemInterfaceIpv6Vrrp6Vrid(d, i["vrid"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrgrp"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrgrp"], _ = expandSystemInterfaceIpv6Vrrp6Vrgrp(d, i["vrgrp"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrip6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrip6"], _ = expandSystemInterfaceIpv6Vrrp6Vrip6(d, i["vrip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandSystemInterfaceIpv6Vrrp6Priority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "adv_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["adv-interval"], _ = expandSystemInterfaceIpv6Vrrp6AdvInterval(d, i["adv_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_time"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-time"], _ = expandSystemInterfaceIpv6Vrrp6StartTime(d, i["start_time"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preempt"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["preempt"], _ = expandSystemInterfaceIpv6Vrrp6Preempt(d, i["preempt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["accept-mode"], _ = expandSystemInterfaceIpv6Vrrp6AcceptMode(d, i["accept_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vrdst6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vrdst6"], _ = expandSystemInterfaceIpv6Vrrp6Vrdst6(d, i["vrdst6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandSystemInterfaceIpv6Vrrp6Status(d, i["status"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemInterfaceIpv6Vrrp6Vrid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6Vrgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6Vrip6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6Priority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6AdvInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6StartTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6Preempt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6AcceptMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6Vrdst6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemInterfaceIpv6Vrrp6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemInterface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemInterfaceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		t, err := expandSystemInterfaceVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("vrf"); ok {
		t, err := expandSystemInterfaceVrf(d, v, "vrf")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf"] = t
		}
	}

	if v, ok := d.GetOk("cli_conn_status"); ok {
		t, err := expandSystemInterfaceCliConnStatus(d, v, "cli_conn_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cli-conn-status"] = t
		}
	}

	if v, ok := d.GetOk("fortilink"); ok {
		t, err := expandSystemInterfaceFortilink(d, v, "fortilink")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandSystemInterfaceMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok {
		t, err := expandSystemInterfaceDistance(d, v, "distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok {
		t, err := expandSystemInterfacePriority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_relay_service"); ok {
		t, err := expandSystemInterfaceDhcpRelayService(d, v, "dhcp_relay_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-relay-service"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_relay_ip"); ok {
		t, err := expandSystemInterfaceDhcpRelayIp(d, v, "dhcp_relay_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-relay-ip"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_relay_type"); ok {
		t, err := expandSystemInterfaceDhcpRelayType(d, v, "dhcp_relay_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-relay-type"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_relay_agent_option"); ok {
		t, err := expandSystemInterfaceDhcpRelayAgentOption(d, v, "dhcp_relay_agent_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-relay-agent-option"] = t
		}
	}

	if v, ok := d.GetOk("management_ip"); ok {
		t, err := expandSystemInterfaceManagementIp(d, v, "management_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["management-ip"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandSystemInterfaceIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok {
		t, err := expandSystemInterfaceAllowaccess(d, v, "allowaccess")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("gwdetect"); ok {
		t, err := expandSystemInterfaceGwdetect(d, v, "gwdetect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gwdetect"] = t
		}
	}

	if v, ok := d.GetOk("ping_serv_status"); ok {
		t, err := expandSystemInterfacePingServStatus(d, v, "ping_serv_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ping-serv-status"] = t
		}
	}

	if v, ok := d.GetOk("detectserver"); ok {
		t, err := expandSystemInterfaceDetectserver(d, v, "detectserver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detectserver"] = t
		}
	}

	if v, ok := d.GetOk("detectprotocol"); ok {
		t, err := expandSystemInterfaceDetectprotocol(d, v, "detectprotocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detectprotocol"] = t
		}
	}

	if v, ok := d.GetOk("ha_priority"); ok {
		t, err := expandSystemInterfaceHaPriority(d, v, "ha_priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-priority"] = t
		}
	}

	if v, ok := d.GetOk("fail_detect"); ok {
		t, err := expandSystemInterfaceFailDetect(d, v, "fail_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-detect"] = t
		}
	}

	if v, ok := d.GetOk("fail_detect_option"); ok {
		t, err := expandSystemInterfaceFailDetectOption(d, v, "fail_detect_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-detect-option"] = t
		}
	}

	if v, ok := d.GetOk("fail_alert_method"); ok {
		t, err := expandSystemInterfaceFailAlertMethod(d, v, "fail_alert_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-alert-method"] = t
		}
	}

	if v, ok := d.GetOk("fail_action_on_extender"); ok {
		t, err := expandSystemInterfaceFailActionOnExtender(d, v, "fail_action_on_extender")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-action-on-extender"] = t
		}
	}

	if v, ok := d.GetOk("fail_alert_interfaces"); ok {
		t, err := expandSystemInterfaceFailAlertInterfaces(d, v, "fail_alert_interfaces")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-alert-interfaces"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_client_identifier"); ok {
		t, err := expandSystemInterfaceDhcpClientIdentifier(d, v, "dhcp_client_identifier")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-client-identifier"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_renew_time"); ok {
		t, err := expandSystemInterfaceDhcpRenewTime(d, v, "dhcp_renew_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-renew-time"] = t
		}
	}

	if v, ok := d.GetOk("ipunnumbered"); ok {
		t, err := expandSystemInterfaceIpunnumbered(d, v, "ipunnumbered")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipunnumbered"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok {
		t, err := expandSystemInterfaceUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("pppoe_unnumbered_negotiate"); ok {
		t, err := expandSystemInterfacePppoeUnnumberedNegotiate(d, v, "pppoe_unnumbered_negotiate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pppoe-unnumbered-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandSystemInterfacePassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("idle_timeout"); ok {
		t, err := expandSystemInterfaceIdleTimeout(d, v, "idle_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("detected_peer_mtu"); ok {
		t, err := expandSystemInterfaceDetectedPeerMtu(d, v, "detected_peer_mtu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["detected-peer-mtu"] = t
		}
	}

	if v, ok := d.GetOk("disc_retry_timeout"); ok {
		t, err := expandSystemInterfaceDiscRetryTimeout(d, v, "disc_retry_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disc-retry-timeout"] = t
		}
	}

	if v, ok := d.GetOk("padt_retry_timeout"); ok {
		t, err := expandSystemInterfacePadtRetryTimeout(d, v, "padt_retry_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["padt-retry-timeout"] = t
		}
	}

	if v, ok := d.GetOk("service_name"); ok {
		t, err := expandSystemInterfaceServiceName(d, v, "service_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-name"] = t
		}
	}

	if v, ok := d.GetOk("ac_name"); ok {
		t, err := expandSystemInterfaceAcName(d, v, "ac_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ac-name"] = t
		}
	}

	if v, ok := d.GetOk("lcp_echo_interval"); ok {
		t, err := expandSystemInterfaceLcpEchoInterval(d, v, "lcp_echo_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lcp-echo-interval"] = t
		}
	}

	if v, ok := d.GetOk("lcp_max_echo_fails"); ok {
		t, err := expandSystemInterfaceLcpMaxEchoFails(d, v, "lcp_max_echo_fails")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lcp-max-echo-fails"] = t
		}
	}

	if v, ok := d.GetOk("defaultgw"); ok {
		t, err := expandSystemInterfaceDefaultgw(d, v, "defaultgw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["defaultgw"] = t
		}
	}

	if v, ok := d.GetOk("dns_server_override"); ok {
		t, err := expandSystemInterfaceDnsServerOverride(d, v, "dns_server_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server-override"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok {
		t, err := expandSystemInterfaceAuthType(d, v, "auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("pptp_client"); ok {
		t, err := expandSystemInterfacePptpClient(d, v, "pptp_client")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pptp-client"] = t
		}
	}

	if v, ok := d.GetOk("pptp_user"); ok {
		t, err := expandSystemInterfacePptpUser(d, v, "pptp_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pptp-user"] = t
		}
	}

	if v, ok := d.GetOk("pptp_password"); ok {
		t, err := expandSystemInterfacePptpPassword(d, v, "pptp_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pptp-password"] = t
		}
	}

	if v, ok := d.GetOk("pptp_server_ip"); ok {
		t, err := expandSystemInterfacePptpServerIp(d, v, "pptp_server_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pptp-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("pptp_auth_type"); ok {
		t, err := expandSystemInterfacePptpAuthType(d, v, "pptp_auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pptp-auth-type"] = t
		}
	}

	if v, ok := d.GetOk("pptp_timeout"); ok {
		t, err := expandSystemInterfacePptpTimeout(d, v, "pptp_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pptp-timeout"] = t
		}
	}

	if v, ok := d.GetOk("arpforward"); ok {
		t, err := expandSystemInterfaceArpforward(d, v, "arpforward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arpforward"] = t
		}
	}

	if v, ok := d.GetOk("ndiscforward"); ok {
		t, err := expandSystemInterfaceNdiscforward(d, v, "ndiscforward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ndiscforward"] = t
		}
	}

	if v, ok := d.GetOk("broadcast_forward"); ok {
		t, err := expandSystemInterfaceBroadcastForward(d, v, "broadcast_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["broadcast-forward"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok {
		t, err := expandSystemInterfaceBfd(d, v, "bfd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("bfd_desired_min_tx"); ok {
		t, err := expandSystemInterfaceBfdDesiredMinTx(d, v, "bfd_desired_min_tx")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd-desired-min-tx"] = t
		}
	}

	if v, ok := d.GetOk("bfd_detect_mult"); ok {
		t, err := expandSystemInterfaceBfdDetectMult(d, v, "bfd_detect_mult")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd-detect-mult"] = t
		}
	}

	if v, ok := d.GetOk("bfd_required_min_rx"); ok {
		t, err := expandSystemInterfaceBfdRequiredMinRx(d, v, "bfd_required_min_rx")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd-required-min-rx"] = t
		}
	}

	if v, ok := d.GetOk("l2forward"); ok {
		t, err := expandSystemInterfaceL2Forward(d, v, "l2forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l2forward"] = t
		}
	}

	if v, ok := d.GetOk("icmp_send_redirect"); ok {
		t, err := expandSystemInterfaceIcmpSendRedirect(d, v, "icmp_send_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-send-redirect"] = t
		}
	}

	if v, ok := d.GetOk("icmp_accept_redirect"); ok {
		t, err := expandSystemInterfaceIcmpAcceptRedirect(d, v, "icmp_accept_redirect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmp-accept-redirect"] = t
		}
	}

	if v, ok := d.GetOk("vlanforward"); ok {
		t, err := expandSystemInterfaceVlanforward(d, v, "vlanforward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlanforward"] = t
		}
	}

	if v, ok := d.GetOk("stpforward"); ok {
		t, err := expandSystemInterfaceStpforward(d, v, "stpforward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stpforward"] = t
		}
	}

	if v, ok := d.GetOk("stpforward_mode"); ok {
		t, err := expandSystemInterfaceStpforwardMode(d, v, "stpforward_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["stpforward-mode"] = t
		}
	}

	if v, ok := d.GetOk("ips_sniffer_mode"); ok {
		t, err := expandSystemInterfaceIpsSnifferMode(d, v, "ips_sniffer_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sniffer-mode"] = t
		}
	}

	if v, ok := d.GetOk("ident_accept"); ok {
		t, err := expandSystemInterfaceIdentAccept(d, v, "ident_accept")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ident-accept"] = t
		}
	}

	if v, ok := d.GetOk("ipmac"); ok {
		t, err := expandSystemInterfaceIpmac(d, v, "ipmac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipmac"] = t
		}
	}

	if v, ok := d.GetOk("subst"); ok {
		t, err := expandSystemInterfaceSubst(d, v, "subst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subst"] = t
		}
	}

	if v, ok := d.GetOk("macaddr"); ok {
		t, err := expandSystemInterfaceMacaddr(d, v, "macaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["macaddr"] = t
		}
	}

	if v, ok := d.GetOk("substitute_dst_mac"); ok {
		t, err := expandSystemInterfaceSubstituteDstMac(d, v, "substitute_dst_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["substitute-dst-mac"] = t
		}
	}

	if v, ok := d.GetOk("speed"); ok {
		t, err := expandSystemInterfaceSpeed(d, v, "speed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["speed"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemInterfaceStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("netbios_forward"); ok {
		t, err := expandSystemInterfaceNetbiosForward(d, v, "netbios_forward")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netbios-forward"] = t
		}
	}

	if v, ok := d.GetOk("wins_ip"); ok {
		t, err := expandSystemInterfaceWinsIp(d, v, "wins_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-ip"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemInterfaceType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("dedicated_to"); ok {
		t, err := expandSystemInterfaceDedicatedTo(d, v, "dedicated_to")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dedicated-to"] = t
		}
	}

	if v, ok := d.GetOk("trust_ip_1"); ok {
		t, err := expandSystemInterfaceTrustIp1(d, v, "trust_ip_1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ip-1"] = t
		}
	}

	if v, ok := d.GetOk("trust_ip_2"); ok {
		t, err := expandSystemInterfaceTrustIp2(d, v, "trust_ip_2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ip-2"] = t
		}
	}

	if v, ok := d.GetOk("trust_ip_3"); ok {
		t, err := expandSystemInterfaceTrustIp3(d, v, "trust_ip_3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ip-3"] = t
		}
	}

	if v, ok := d.GetOk("trust_ip6_1"); ok {
		t, err := expandSystemInterfaceTrustIp61(d, v, "trust_ip6_1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ip6-1"] = t
		}
	}

	if v, ok := d.GetOk("trust_ip6_2"); ok {
		t, err := expandSystemInterfaceTrustIp62(d, v, "trust_ip6_2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ip6-2"] = t
		}
	}

	if v, ok := d.GetOk("trust_ip6_3"); ok {
		t, err := expandSystemInterfaceTrustIp63(d, v, "trust_ip6_3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trust-ip6-3"] = t
		}
	}

	if v, ok := d.GetOk("mtu_override"); ok {
		t, err := expandSystemInterfaceMtuOverride(d, v, "mtu_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtu-override"] = t
		}
	}

	if v, ok := d.GetOk("mtu"); ok {
		t, err := expandSystemInterfaceMtu(d, v, "mtu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtu"] = t
		}
	}

	if v, ok := d.GetOk("wccp"); ok {
		t, err := expandSystemInterfaceWccp(d, v, "wccp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wccp"] = t
		}
	}

	if v, ok := d.GetOk("netflow_sampler"); ok {
		t, err := expandSystemInterfaceNetflowSampler(d, v, "netflow_sampler")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netflow-sampler"] = t
		}
	}

	if v, ok := d.GetOk("sflow_sampler"); ok {
		t, err := expandSystemInterfaceSflowSampler(d, v, "sflow_sampler")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sflow-sampler"] = t
		}
	}

	if v, ok := d.GetOk("drop_overlapped_fragment"); ok {
		t, err := expandSystemInterfaceDropOverlappedFragment(d, v, "drop_overlapped_fragment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drop-overlapped-fragment"] = t
		}
	}

	if v, ok := d.GetOk("drop_fragment"); ok {
		t, err := expandSystemInterfaceDropFragment(d, v, "drop_fragment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drop-fragment"] = t
		}
	}

	if v, ok := d.GetOk("scan_botnet_connections"); ok {
		t, err := expandSystemInterfaceScanBotnetConnections(d, v, "scan_botnet_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-botnet-connections"] = t
		}
	}

	if v, ok := d.GetOk("src_check"); ok {
		t, err := expandSystemInterfaceSrcCheck(d, v, "src_check")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-check"] = t
		}
	}

	if v, ok := d.GetOk("sample_rate"); ok {
		t, err := expandSystemInterfaceSampleRate(d, v, "sample_rate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sample-rate"] = t
		}
	}

	if v, ok := d.GetOk("polling_interval"); ok {
		t, err := expandSystemInterfacePollingInterval(d, v, "polling_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["polling-interval"] = t
		}
	}

	if v, ok := d.GetOk("sample_direction"); ok {
		t, err := expandSystemInterfaceSampleDirection(d, v, "sample_direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sample-direction"] = t
		}
	}

	if v, ok := d.GetOk("explicit_web_proxy"); ok {
		t, err := expandSystemInterfaceExplicitWebProxy(d, v, "explicit_web_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["explicit-web-proxy"] = t
		}
	}

	if v, ok := d.GetOk("explicit_ftp_proxy"); ok {
		t, err := expandSystemInterfaceExplicitFtpProxy(d, v, "explicit_ftp_proxy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["explicit-ftp-proxy"] = t
		}
	}

	if v, ok := d.GetOk("proxy_captive_portal"); ok {
		t, err := expandSystemInterfaceProxyCaptivePortal(d, v, "proxy_captive_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-captive-portal"] = t
		}
	}

	if v, ok := d.GetOk("tcp_mss"); ok {
		t, err := expandSystemInterfaceTcpMss(d, v, "tcp_mss")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-mss"] = t
		}
	}

	if v, ok := d.GetOk("inbandwidth"); ok {
		t, err := expandSystemInterfaceInbandwidth(d, v, "inbandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inbandwidth"] = t
		}
	}

	if v, ok := d.GetOk("outbandwidth"); ok {
		t, err := expandSystemInterfaceOutbandwidth(d, v, "outbandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outbandwidth"] = t
		}
	}

	if v, ok := d.GetOk("egress_shaping_profile"); ok {
		t, err := expandSystemInterfaceEgressShapingProfile(d, v, "egress_shaping_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["egress-shaping-profile"] = t
		}
	}

	if v, ok := d.GetOk("disconnect_threshold"); ok {
		t, err := expandSystemInterfaceDisconnectThreshold(d, v, "disconnect_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disconnect-threshold"] = t
		}
	}

	if v, ok := d.GetOk("spillover_threshold"); ok {
		t, err := expandSystemInterfaceSpilloverThreshold(d, v, "spillover_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spillover-threshold"] = t
		}
	}

	if v, ok := d.GetOk("ingress_spillover_threshold"); ok {
		t, err := expandSystemInterfaceIngressSpilloverThreshold(d, v, "ingress_spillover_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ingress-spillover-threshold"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok {
		t, err := expandSystemInterfaceWeight(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandSystemInterfaceInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("external"); ok {
		t, err := expandSystemInterfaceExternal(d, v, "external")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external"] = t
		}
	}

	if v, ok := d.GetOk("vlanid"); ok {
		t, err := expandSystemInterfaceVlanid(d, v, "vlanid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlanid"] = t
		}
	}

	if v, ok := d.GetOk("forward_domain"); ok {
		t, err := expandSystemInterfaceForwardDomain(d, v, "forward_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-domain"] = t
		}
	}

	if v, ok := d.GetOk("remote_ip"); ok {
		t, err := expandSystemInterfaceRemoteIp(d, v, "remote_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-ip"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok {
		t, err := expandSystemInterfaceMember(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("lacp_mode"); ok {
		t, err := expandSystemInterfaceLacpMode(d, v, "lacp_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lacp-mode"] = t
		}
	}

	if v, ok := d.GetOk("lacp_ha_slave"); ok {
		t, err := expandSystemInterfaceLacpHaSlave(d, v, "lacp_ha_slave")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lacp-ha-slave"] = t
		}
	}

	if v, ok := d.GetOk("lacp_speed"); ok {
		t, err := expandSystemInterfaceLacpSpeed(d, v, "lacp_speed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lacp-speed"] = t
		}
	}

	if v, ok := d.GetOk("min_links"); ok {
		t, err := expandSystemInterfaceMinLinks(d, v, "min_links")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-links"] = t
		}
	}

	if v, ok := d.GetOk("min_links_down"); ok {
		t, err := expandSystemInterfaceMinLinksDown(d, v, "min_links_down")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["min-links-down"] = t
		}
	}

	if v, ok := d.GetOk("algorithm"); ok {
		t, err := expandSystemInterfaceAlgorithm(d, v, "algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["algorithm"] = t
		}
	}

	if v, ok := d.GetOk("link_up_delay"); ok {
		t, err := expandSystemInterfaceLinkUpDelay(d, v, "link_up_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-up-delay"] = t
		}
	}

	if v, ok := d.GetOk("priority_override"); ok {
		t, err := expandSystemInterfacePriorityOverride(d, v, "priority_override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-override"] = t
		}
	}

	if v, ok := d.GetOk("aggregate"); ok {
		t, err := expandSystemInterfaceAggregate(d, v, "aggregate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregate"] = t
		}
	}

	if v, ok := d.GetOk("redundant_interface"); ok {
		t, err := expandSystemInterfaceRedundantInterface(d, v, "redundant_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant-interface"] = t
		}
	}

	if v, ok := d.GetOk("managed_device"); ok {
		t, err := expandSystemInterfaceManagedDevice(d, v, "managed_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["managed-device"] = t
		}
	}

	if v, ok := d.GetOk("devindex"); ok {
		t, err := expandSystemInterfaceDevindex(d, v, "devindex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devindex"] = t
		}
	}

	if v, ok := d.GetOk("vindex"); ok {
		t, err := expandSystemInterfaceVindex(d, v, "vindex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vindex"] = t
		}
	}

	if v, ok := d.GetOk("switch"); ok {
		t, err := expandSystemInterfaceSwitch(d, v, "switch")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandSystemInterfaceDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("alias"); ok {
		t, err := expandSystemInterfaceAlias(d, v, "alias")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alias"] = t
		}
	}

	if v, ok := d.GetOk("security_mode"); ok {
		t, err := expandSystemInterfaceSecurityMode(d, v, "security_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-mode"] = t
		}
	}

	if v, ok := d.GetOk("captive_portal"); ok {
		t, err := expandSystemInterfaceCaptivePortal(d, v, "captive_portal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["captive-portal"] = t
		}
	}

	if v, ok := d.GetOk("security_mac_auth_bypass"); ok {
		t, err := expandSystemInterfaceSecurityMacAuthBypass(d, v, "security_mac_auth_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-mac-auth-bypass"] = t
		}
	}

	if v, ok := d.GetOk("security_external_web"); ok {
		t, err := expandSystemInterfaceSecurityExternalWeb(d, v, "security_external_web")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-external-web"] = t
		}
	}

	if v, ok := d.GetOk("security_external_logout"); ok {
		t, err := expandSystemInterfaceSecurityExternalLogout(d, v, "security_external_logout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-external-logout"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_override_group"); ok {
		t, err := expandSystemInterfaceReplacemsgOverrideGroup(d, v, "replacemsg_override_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-override-group"] = t
		}
	}

	if v, ok := d.GetOk("security_redirect_url"); ok {
		t, err := expandSystemInterfaceSecurityRedirectUrl(d, v, "security_redirect_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-redirect-url"] = t
		}
	}

	if v, ok := d.GetOk("security_exempt_list"); ok {
		t, err := expandSystemInterfaceSecurityExemptList(d, v, "security_exempt_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-exempt-list"] = t
		}
	}

	if v, ok := d.GetOk("security_groups"); ok {
		t, err := expandSystemInterfaceSecurityGroups(d, v, "security_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-groups"] = t
		}
	}

	if v, ok := d.GetOk("device_identification"); ok {
		t, err := expandSystemInterfaceDeviceIdentification(d, v, "device_identification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-identification"] = t
		}
	}

	if v, ok := d.GetOk("device_user_identification"); ok {
		t, err := expandSystemInterfaceDeviceUserIdentification(d, v, "device_user_identification")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-user-identification"] = t
		}
	}

	if v, ok := d.GetOk("device_identification_active_scan"); ok {
		t, err := expandSystemInterfaceDeviceIdentificationActiveScan(d, v, "device_identification_active_scan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-identification-active-scan"] = t
		}
	}

	if v, ok := d.GetOk("device_access_list"); ok {
		t, err := expandSystemInterfaceDeviceAccessList(d, v, "device_access_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-access-list"] = t
		}
	}

	if v, ok := d.GetOk("device_netscan"); ok {
		t, err := expandSystemInterfaceDeviceNetscan(d, v, "device_netscan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-netscan"] = t
		}
	}

	if v, ok := d.GetOk("lldp_reception"); ok {
		t, err := expandSystemInterfaceLldpReception(d, v, "lldp_reception")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-reception"] = t
		}
	}

	if v, ok := d.GetOk("lldp_transmission"); ok {
		t, err := expandSystemInterfaceLldpTransmission(d, v, "lldp_transmission")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-transmission"] = t
		}
	}

	if v, ok := d.GetOk("lldp_network_policy"); ok {
		t, err := expandSystemInterfaceLldpNetworkPolicy(d, v, "lldp_network_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-network-policy"] = t
		}
	}

	if v, ok := d.GetOk("fortiheartbeat"); ok {
		t, err := expandSystemInterfaceFortiheartbeat(d, v, "fortiheartbeat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiheartbeat"] = t
		}
	}

	if v, ok := d.GetOk("broadcast_forticlient_discovery"); ok {
		t, err := expandSystemInterfaceBroadcastForticlientDiscovery(d, v, "broadcast_forticlient_discovery")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["broadcast-forticlient-discovery"] = t
		}
	}

	if v, ok := d.GetOk("endpoint_compliance"); ok {
		t, err := expandSystemInterfaceEndpointCompliance(d, v, "endpoint_compliance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endpoint-compliance"] = t
		}
	}

	if v, ok := d.GetOk("estimated_upstream_bandwidth"); ok {
		t, err := expandSystemInterfaceEstimatedUpstreamBandwidth(d, v, "estimated_upstream_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["estimated-upstream-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("estimated_downstream_bandwidth"); ok {
		t, err := expandSystemInterfaceEstimatedDownstreamBandwidth(d, v, "estimated_downstream_bandwidth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["estimated-downstream-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("vrrp_virtual_mac"); ok {
		t, err := expandSystemInterfaceVrrpVirtualMac(d, v, "vrrp_virtual_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrrp-virtual-mac"] = t
		}
	}

	if v, ok := d.GetOk("vrrp"); ok {
		t, err := expandSystemInterfaceVrrp(d, v, "vrrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrrp"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok {
		t, err := expandSystemInterfaceRole(d, v, "role")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("snmp_index"); ok {
		t, err := expandSystemInterfaceSnmpIndex(d, v, "snmp_index")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["snmp-index"] = t
		}
	}

	if v, ok := d.GetOk("secondary_ip"); ok {
		t, err := expandSystemInterfaceSecondaryIp(d, v, "secondary_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-IP"] = t
		}
	}

	if v, ok := d.GetOk("secondaryip"); ok {
		t, err := expandSystemInterfaceSecondaryip(d, v, "secondaryip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondaryip"] = t
		}
	}

	if v, ok := d.GetOk("preserve_session_route"); ok {
		t, err := expandSystemInterfacePreserveSessionRoute(d, v, "preserve_session_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preserve-session-route"] = t
		}
	}

	if v, ok := d.GetOk("auto_auth_extension_device"); ok {
		t, err := expandSystemInterfaceAutoAuthExtensionDevice(d, v, "auto_auth_extension_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-auth-extension-device"] = t
		}
	}

	if v, ok := d.GetOk("ap_discover"); ok {
		t, err := expandSystemInterfaceApDiscover(d, v, "ap_discover")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-discover"] = t
		}
	}

	if v, ok := d.GetOk("fortilink_stacking"); ok {
		t, err := expandSystemInterfaceFortilinkStacking(d, v, "fortilink_stacking")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink-stacking"] = t
		}
	}

	if v, ok := d.GetOk("fortilink_split_interface"); ok {
		t, err := expandSystemInterfaceFortilinkSplitInterface(d, v, "fortilink_split_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink-split-interface"] = t
		}
	}

	if v, ok := d.GetOk("internal"); ok {
		t, err := expandSystemInterfaceInternal(d, v, "internal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal"] = t
		}
	}

	if v, ok := d.GetOk("fortilink_backup_link"); ok {
		t, err := expandSystemInterfaceFortilinkBackupLink(d, v, "fortilink_backup_link")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink-backup-link"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_access_vlan"); ok {
		t, err := expandSystemInterfaceSwitchControllerAccessVlan(d, v, "switch_controller_access_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-access-vlan"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_traffic_policy"); ok {
		t, err := expandSystemInterfaceSwitchControllerTrafficPolicy(d, v, "switch_controller_traffic_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-traffic-policy"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_igmp_snooping"); ok {
		t, err := expandSystemInterfaceSwitchControllerIgmpSnooping(d, v, "switch_controller_igmp_snooping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-igmp-snooping"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_dhcp_snooping"); ok {
		t, err := expandSystemInterfaceSwitchControllerDhcpSnooping(d, v, "switch_controller_dhcp_snooping")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-dhcp-snooping"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_dhcp_snooping_verify_mac"); ok {
		t, err := expandSystemInterfaceSwitchControllerDhcpSnoopingVerifyMac(d, v, "switch_controller_dhcp_snooping_verify_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-dhcp-snooping-verify-mac"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_dhcp_snooping_option82"); ok {
		t, err := expandSystemInterfaceSwitchControllerDhcpSnoopingOption82(d, v, "switch_controller_dhcp_snooping_option82")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-dhcp-snooping-option82"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_arp_inspection"); ok {
		t, err := expandSystemInterfaceSwitchControllerArpInspection(d, v, "switch_controller_arp_inspection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-arp-inspection"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_learning_limit"); ok {
		t, err := expandSystemInterfaceSwitchControllerLearningLimit(d, v, "switch_controller_learning_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-learning-limit"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok {
		t, err := expandSystemInterfaceColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("tagging"); ok {
		t, err := expandSystemInterfaceTagging(d, v, "tagging")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tagging"] = t
		}
	}

	if v, ok := d.GetOk("ipv6"); ok {
		t, err := expandSystemInterfaceIpv6(d, v, "ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6"] = t
		}
	}

	return &obj, nil
}
