// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

package fortios

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccFortiOSSystemAutomationDestination_basic(t *testing.T) {
	rname := acctest.RandString(8)
	log.Printf("TestAccFortiOSSystemAutomationDestination_basic %s", rname)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiOSSystemAutomationDestinationConfig(rname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiOSSystemAutomationDestinationExists("fortios_system_automationdestination.trname"),
					resource.TestCheckResourceAttr("fortios_system_automationdestination.trname", "ha_group_id", "0"),
					resource.TestCheckResourceAttr("fortios_system_automationdestination.trname", "name", rname),
					resource.TestCheckResourceAttr("fortios_system_automationdestination.trname", "type", "fortigate"),
				),
			},
		},
	})
}

func testAccCheckFortiOSSystemAutomationDestinationExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found SystemAutomationDestination: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SystemAutomationDestination is set")
		}

		c := testAccProvider.Meta().(*FortiClient).Client

		i := rs.Primary.ID
		o, err := c.ReadSystemAutomationDestination(i)

		if err != nil {
			return fmt.Errorf("Error reading SystemAutomationDestination: %s", err)
		}

		if o == nil {
			return fmt.Errorf("Error creating SystemAutomationDestination: %s", n)
		}

		return nil
	}
}

func testAccCheckSystemAutomationDestinationDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*FortiClient).Client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fortios_system_automationdestination" {
			continue
		}

		i := rs.Primary.ID
		o, err := c.ReadSystemAutomationDestination(i)

		if err == nil {
			if o != nil {
				return fmt.Errorf("Error SystemAutomationDestination %s still exists", rs.Primary.ID)
			}
		}

		return nil
	}

	return nil
}

func testAccFortiOSSystemAutomationDestinationConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_system_automationdestination" "trname" {
  ha_group_id = 0
  name        = "%[1]s"
  type        = "fortigate"
}
`, name)
}
