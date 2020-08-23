// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure WAN optimization peers.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWanoptPeer() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptPeerCreate,
		Read:   resourceWanoptPeerRead,
		Update: resourceWanoptPeerUpdate,
		Delete: resourceWanoptPeerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"peer_host_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWanoptPeerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWanoptPeer(d)
	if err != nil {
		return fmt.Errorf("Error creating WanoptPeer resource while getting object: %v", err)
	}

	o, err := c.CreateWanoptPeer(obj)

	if err != nil {
		return fmt.Errorf("Error creating WanoptPeer resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WanoptPeer")
	}

	return resourceWanoptPeerRead(d, m)
}

func resourceWanoptPeerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWanoptPeer(d)
	if err != nil {
		return fmt.Errorf("Error updating WanoptPeer resource while getting object: %v", err)
	}

	o, err := c.UpdateWanoptPeer(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WanoptPeer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WanoptPeer")
	}

	return resourceWanoptPeerRead(d, m)
}

func resourceWanoptPeerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWanoptPeer(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WanoptPeer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptPeerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWanoptPeer(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WanoptPeer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptPeer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WanoptPeer resource from API: %v", err)
	}
	return nil
}

func flattenWanoptPeerPeerHostId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptPeerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWanoptPeer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("peer_host_id", flattenWanoptPeerPeerHostId(o["peer-host-id"], d, "peer_host_id")); err != nil {
		if !fortiAPIPatch(o["peer-host-id"]) {
			return fmt.Errorf("Error reading peer_host_id: %v", err)
		}
	}

	if err = d.Set("ip", flattenWanoptPeerIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	return nil
}

func flattenWanoptPeerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWanoptPeerPeerHostId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptPeerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptPeer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("peer_host_id"); ok {
		t, err := expandWanoptPeerPeerHostId(d, v, "peer_host_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-host-id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandWanoptPeerIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	return &obj, nil
}
