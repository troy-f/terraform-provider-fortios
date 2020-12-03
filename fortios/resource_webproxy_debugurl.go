// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure debug URL addresses.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWebProxyDebugUrl() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyDebugUrlCreate,
		Read:   resourceWebProxyDebugUrlRead,
		Update: resourceWebProxyDebugUrlUpdate,
		Delete: resourceWebProxyDebugUrlDelete,

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
			"url_pattern": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Required:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"exact": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWebProxyDebugUrlCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWebProxyDebugUrl(d)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyDebugUrl resource while getting object: %v", err)
	}

	o, err := c.CreateWebProxyDebugUrl(obj)

	if err != nil {
		return fmt.Errorf("Error creating WebProxyDebugUrl resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebProxyDebugUrl")
	}

	return resourceWebProxyDebugUrlRead(d, m)
}

func resourceWebProxyDebugUrlUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWebProxyDebugUrl(d)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyDebugUrl resource while getting object: %v", err)
	}

	o, err := c.UpdateWebProxyDebugUrl(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyDebugUrl resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebProxyDebugUrl")
	}

	return resourceWebProxyDebugUrlRead(d, m)
}

func resourceWebProxyDebugUrlDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWebProxyDebugUrl(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyDebugUrl resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyDebugUrlRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWebProxyDebugUrl(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyDebugUrl resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyDebugUrl(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyDebugUrl resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyDebugUrlName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyDebugUrlUrlPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyDebugUrlStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyDebugUrlExact(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebProxyDebugUrl(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWebProxyDebugUrlName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("url_pattern", flattenWebProxyDebugUrlUrlPattern(o["url-pattern"], d, "url_pattern")); err != nil {
		if !fortiAPIPatch(o["url-pattern"]) {
			return fmt.Errorf("Error reading url_pattern: %v", err)
		}
	}

	if err = d.Set("status", flattenWebProxyDebugUrlStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("exact", flattenWebProxyDebugUrlExact(o["exact"], d, "exact")); err != nil {
		if !fortiAPIPatch(o["exact"]) {
			return fmt.Errorf("Error reading exact: %v", err)
		}
	}

	return nil
}

func flattenWebProxyDebugUrlFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebProxyDebugUrlName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyDebugUrlUrlPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyDebugUrlStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyDebugUrlExact(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyDebugUrl(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWebProxyDebugUrlName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("url_pattern"); ok {
		t, err := expandWebProxyDebugUrlUrlPattern(d, v, "url_pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-pattern"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandWebProxyDebugUrlStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("exact"); ok {
		t, err := expandWebProxyDebugUrlExact(d, v, "exact")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exact"] = t
		}
	}

	return &obj, nil
}
