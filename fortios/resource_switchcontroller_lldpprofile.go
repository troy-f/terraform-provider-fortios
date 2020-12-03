// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch LLDP profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerLldpProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerLldpProfileCreate,
		Read:   resourceSwitchControllerLldpProfileRead,
		Update: resourceSwitchControllerLldpProfileUpdate,
		Delete: resourceSwitchControllerLldpProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required:     true,
				ForceNew:     true,
			},
			"med_tlvs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"n8021_tlvs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"n8023_tlvs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_isl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_isl_hello_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),
				Optional:     true,
				Computed:     true,
			},
			"auto_isl_receive_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 90),
				Optional:     true,
				Computed:     true,
			},
			"auto_isl_port_group": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 9),
				Optional:     true,
				Computed:     true,
			},
			"med_network_policy": &schema.Schema{
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
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4094),
							Optional:     true,
							Computed:     true,
						},
						"priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"dscp": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"custom_tlvs": &schema.Schema{
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
						"oui": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"subtype": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"information_string": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceSwitchControllerLldpProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerLldpProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerLldpProfile resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerLldpProfile(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerLldpProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerLldpProfile")
	}

	return resourceSwitchControllerLldpProfileRead(d, m)
}

func resourceSwitchControllerLldpProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerLldpProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLldpProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerLldpProfile(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLldpProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerLldpProfile")
	}

	return resourceSwitchControllerLldpProfileRead(d, m)
}

func resourceSwitchControllerLldpProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchControllerLldpProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerLldpProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerLldpProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchControllerLldpProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLldpProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerLldpProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLldpProfile resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerLldpProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedTlvs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfile8021Tlvs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfile8023Tlvs(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslHelloTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslReceiveTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslPortGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedNetworkPolicy(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenSwitchControllerLldpProfileMedNetworkPolicyName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = flattenSwitchControllerLldpProfileMedNetworkPolicyStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := i["vlan"]; ok {
			tmp["vlan"] = flattenSwitchControllerLldpProfileMedNetworkPolicyVlan(i["vlan"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			tmp["priority"] = flattenSwitchControllerLldpProfileMedNetworkPolicyPriority(i["priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp"
		if _, ok := i["dscp"]; ok {
			tmp["dscp"] = flattenSwitchControllerLldpProfileMedNetworkPolicyDscp(i["dscp"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSwitchControllerLldpProfileMedNetworkPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedNetworkPolicyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedNetworkPolicyVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedNetworkPolicyPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedNetworkPolicyDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileCustomTlvs(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenSwitchControllerLldpProfileCustomTlvsName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui"
		if _, ok := i["oui"]; ok {
			tmp["oui"] = flattenSwitchControllerLldpProfileCustomTlvsOui(i["oui"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subtype"
		if _, ok := i["subtype"]; ok {
			tmp["subtype"] = flattenSwitchControllerLldpProfileCustomTlvsSubtype(i["subtype"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "information_string"
		if _, ok := i["information-string"]; ok {
			tmp["information_string"] = flattenSwitchControllerLldpProfileCustomTlvsInformationString(i["information-string"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSwitchControllerLldpProfileCustomTlvsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileCustomTlvsOui(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileCustomTlvsSubtype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileCustomTlvsInformationString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerLldpProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerLldpProfileName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("med_tlvs", flattenSwitchControllerLldpProfileMedTlvs(o["med-tlvs"], d, "med_tlvs")); err != nil {
		if !fortiAPIPatch(o["med-tlvs"]) {
			return fmt.Errorf("Error reading med_tlvs: %v", err)
		}
	}

	if err = d.Set("n8021_tlvs", flattenSwitchControllerLldpProfile8021Tlvs(o["802.1-tlvs"], d, "n8021_tlvs")); err != nil {
		if !fortiAPIPatch(o["802.1-tlvs"]) {
			return fmt.Errorf("Error reading n8021_tlvs: %v", err)
		}
	}

	if err = d.Set("n8023_tlvs", flattenSwitchControllerLldpProfile8023Tlvs(o["802.3-tlvs"], d, "n8023_tlvs")); err != nil {
		if !fortiAPIPatch(o["802.3-tlvs"]) {
			return fmt.Errorf("Error reading n8023_tlvs: %v", err)
		}
	}

	if err = d.Set("auto_isl", flattenSwitchControllerLldpProfileAutoIsl(o["auto-isl"], d, "auto_isl")); err != nil {
		if !fortiAPIPatch(o["auto-isl"]) {
			return fmt.Errorf("Error reading auto_isl: %v", err)
		}
	}

	if err = d.Set("auto_isl_hello_timer", flattenSwitchControllerLldpProfileAutoIslHelloTimer(o["auto-isl-hello-timer"], d, "auto_isl_hello_timer")); err != nil {
		if !fortiAPIPatch(o["auto-isl-hello-timer"]) {
			return fmt.Errorf("Error reading auto_isl_hello_timer: %v", err)
		}
	}

	if err = d.Set("auto_isl_receive_timeout", flattenSwitchControllerLldpProfileAutoIslReceiveTimeout(o["auto-isl-receive-timeout"], d, "auto_isl_receive_timeout")); err != nil {
		if !fortiAPIPatch(o["auto-isl-receive-timeout"]) {
			return fmt.Errorf("Error reading auto_isl_receive_timeout: %v", err)
		}
	}

	if err = d.Set("auto_isl_port_group", flattenSwitchControllerLldpProfileAutoIslPortGroup(o["auto-isl-port-group"], d, "auto_isl_port_group")); err != nil {
		if !fortiAPIPatch(o["auto-isl-port-group"]) {
			return fmt.Errorf("Error reading auto_isl_port_group: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("med_network_policy", flattenSwitchControllerLldpProfileMedNetworkPolicy(o["med-network-policy"], d, "med_network_policy")); err != nil {
			if !fortiAPIPatch(o["med-network-policy"]) {
				return fmt.Errorf("Error reading med_network_policy: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("med_network_policy"); ok {
			if err = d.Set("med_network_policy", flattenSwitchControllerLldpProfileMedNetworkPolicy(o["med-network-policy"], d, "med_network_policy")); err != nil {
				if !fortiAPIPatch(o["med-network-policy"]) {
					return fmt.Errorf("Error reading med_network_policy: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("custom_tlvs", flattenSwitchControllerLldpProfileCustomTlvs(o["custom-tlvs"], d, "custom_tlvs")); err != nil {
			if !fortiAPIPatch(o["custom-tlvs"]) {
				return fmt.Errorf("Error reading custom_tlvs: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_tlvs"); ok {
			if err = d.Set("custom_tlvs", flattenSwitchControllerLldpProfileCustomTlvs(o["custom-tlvs"], d, "custom_tlvs")); err != nil {
				if !fortiAPIPatch(o["custom-tlvs"]) {
					return fmt.Errorf("Error reading custom_tlvs: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchControllerLldpProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerLldpProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedTlvs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfile8021Tlvs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfile8023Tlvs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslHelloTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslReceiveTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslPortGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSwitchControllerLldpProfileMedNetworkPolicyName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandSwitchControllerLldpProfileMedNetworkPolicyStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vlan"], _ = expandSwitchControllerLldpProfileMedNetworkPolicyVlan(d, i["vlan"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandSwitchControllerLldpProfileMedNetworkPolicyPriority(d, i["priority"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp"], _ = expandSwitchControllerLldpProfileMedNetworkPolicyDscp(d, i["dscp"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicyVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicyPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicyDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileCustomTlvs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSwitchControllerLldpProfileCustomTlvsName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["oui"], _ = expandSwitchControllerLldpProfileCustomTlvsOui(d, i["oui"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subtype"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["subtype"], _ = expandSwitchControllerLldpProfileCustomTlvsSubtype(d, i["subtype"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "information_string"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["information-string"], _ = expandSwitchControllerLldpProfileCustomTlvsInformationString(d, i["information_string"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerLldpProfileCustomTlvsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileCustomTlvsOui(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileCustomTlvsSubtype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileCustomTlvsInformationString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerLldpProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerLldpProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("med_tlvs"); ok {
		t, err := expandSwitchControllerLldpProfileMedTlvs(d, v, "med_tlvs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["med-tlvs"] = t
		}
	}

	if v, ok := d.GetOk("n8021_tlvs"); ok {
		t, err := expandSwitchControllerLldpProfile8021Tlvs(d, v, "n8021_tlvs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["802.1-tlvs"] = t
		}
	}

	if v, ok := d.GetOk("n8023_tlvs"); ok {
		t, err := expandSwitchControllerLldpProfile8023Tlvs(d, v, "n8023_tlvs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["802.3-tlvs"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl"); ok {
		t, err := expandSwitchControllerLldpProfileAutoIsl(d, v, "auto_isl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_hello_timer"); ok {
		t, err := expandSwitchControllerLldpProfileAutoIslHelloTimer(d, v, "auto_isl_hello_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-hello-timer"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_receive_timeout"); ok {
		t, err := expandSwitchControllerLldpProfileAutoIslReceiveTimeout(d, v, "auto_isl_receive_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-receive-timeout"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_port_group"); ok {
		t, err := expandSwitchControllerLldpProfileAutoIslPortGroup(d, v, "auto_isl_port_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-port-group"] = t
		}
	}

	if v, ok := d.GetOk("med_network_policy"); ok {
		t, err := expandSwitchControllerLldpProfileMedNetworkPolicy(d, v, "med_network_policy")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["med-network-policy"] = t
		}
	}

	if v, ok := d.GetOk("custom_tlvs"); ok {
		t, err := expandSwitchControllerLldpProfileCustomTlvs(d, v, "custom_tlvs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-tlvs"] = t
		}
	}

	return &obj, nil
}
