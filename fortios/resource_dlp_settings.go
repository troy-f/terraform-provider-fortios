// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Designate logical storage for DLP fingerprint database.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDlpSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpSettingsUpdate,
		Read:   resourceDlpSettingsRead,
		Update: resourceDlpSettingsUpdate,
		Delete: resourceDlpSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"storage_device": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"size": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(16, 4294967295),
				Optional:     true,
				Computed:     true,
			},
			"db_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cache_mem_percent": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 15),
				Optional:     true,
				Computed:     true,
			},
			"chunk_size": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(100, 100000),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceDlpSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectDlpSettings(d)
	if err != nil {
		return fmt.Errorf("Error updating DlpSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateDlpSettings(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating DlpSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("DlpSettings")
	}

	return resourceDlpSettingsRead(d, m)
}

func resourceDlpSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteDlpSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting DlpSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadDlpSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error reading DlpSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpSettings(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DlpSettings resource from API: %v", err)
	}
	return nil
}

func flattenDlpSettingsStorageDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSettingsSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSettingsDbMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSettingsCacheMemPercent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpSettingsChunkSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDlpSettings(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("storage_device", flattenDlpSettingsStorageDevice(o["storage-device"], d, "storage_device")); err != nil {
		if !fortiAPIPatch(o["storage-device"]) {
			return fmt.Errorf("Error reading storage_device: %v", err)
		}
	}

	if err = d.Set("size", flattenDlpSettingsSize(o["size"], d, "size")); err != nil {
		if !fortiAPIPatch(o["size"]) {
			return fmt.Errorf("Error reading size: %v", err)
		}
	}

	if err = d.Set("db_mode", flattenDlpSettingsDbMode(o["db-mode"], d, "db_mode")); err != nil {
		if !fortiAPIPatch(o["db-mode"]) {
			return fmt.Errorf("Error reading db_mode: %v", err)
		}
	}

	if err = d.Set("cache_mem_percent", flattenDlpSettingsCacheMemPercent(o["cache-mem-percent"], d, "cache_mem_percent")); err != nil {
		if !fortiAPIPatch(o["cache-mem-percent"]) {
			return fmt.Errorf("Error reading cache_mem_percent: %v", err)
		}
	}

	if err = d.Set("chunk_size", flattenDlpSettingsChunkSize(o["chunk-size"], d, "chunk_size")); err != nil {
		if !fortiAPIPatch(o["chunk-size"]) {
			return fmt.Errorf("Error reading chunk_size: %v", err)
		}
	}

	return nil
}

func flattenDlpSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDlpSettingsStorageDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsDbMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsCacheMemPercent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpSettingsChunkSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDlpSettings(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("storage_device"); ok {
		t, err := expandDlpSettingsStorageDevice(d, v, "storage_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["storage-device"] = t
		}
	}

	if v, ok := d.GetOk("size"); ok {
		t, err := expandDlpSettingsSize(d, v, "size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["size"] = t
		}
	}

	if v, ok := d.GetOk("db_mode"); ok {
		t, err := expandDlpSettingsDbMode(d, v, "db_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["db-mode"] = t
		}
	}

	if v, ok := d.GetOk("cache_mem_percent"); ok {
		t, err := expandDlpSettingsCacheMemPercent(d, v, "cache_mem_percent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-mem-percent"] = t
		}
	}

	if v, ok := d.GetOk("chunk_size"); ok {
		t, err := expandDlpSettingsChunkSize(d, v, "chunk_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chunk-size"] = t
		}
	}

	return &obj, nil
}
