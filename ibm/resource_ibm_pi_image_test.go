// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
)

func TestAccIBMPIImagebasic(t *testing.T) {

	name := fmt.Sprintf("tf-pi-image-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPIImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIImageConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIImageExists("ibm_pi_image.power_image"),
					resource.TestCheckResourceAttr(
						"ibm_pi_image.power_image", "pi_image_name", name),
				),
			},
		},
	})
}

func testAccCheckIBMPIImageDestroy(s *terraform.State) error {

	sess, err := testAccProvider.Meta().(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_image" {
			continue
		}
		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		powerinstanceid := parts[0]
		imageC := st.NewIBMPIImageClient(sess, powerinstanceid)
		_, err = imageC.Get(parts[1], powerinstanceid)
		if err == nil {
			return fmt.Errorf("PI Image still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}
func testAccCheckIBMPIImageExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, err := testAccProvider.Meta().(ClientSession).IBMPISession()
		if err != nil {
			return err
		}
		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		powerinstanceid := parts[0]
		client := st.NewIBMPIImageClient(sess, powerinstanceid)

		_, err = client.Get(parts[1], powerinstanceid)
		if err != nil {
			return err
		}

		return nil

	}
}

func testAccCheckIBMPIImageConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_image" "power_image" {
		pi_image_name       = "%s"
		pi_image_id         = "cfc02954-8f6f-4e6b-96ae-40b24c90bd54"
		pi_cloud_instance_id = "%s"
	  }
	`, name, pi_cloud_instance_id)
}

func TestAccIBMPIImageCOSPublicImport(t *testing.T) {
	imageRes := "ibm_pi_image.cos_image"
	name := fmt.Sprintf("tf-pi-image-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPIImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIImageCOSPublicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIImageExists(imageRes),
					resource.TestCheckResourceAttr(imageRes, "pi_image_name", name),
					resource.TestCheckResourceAttrSet(imageRes, "image_id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIImageCOSPublicConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_image" "cos_image" {
		pi_image_name       = "%[1]s"
		pi_cloud_instance_id = "%[2]s"
		pi_image_bucket_name = "%[3]s"
		pi_image_bucket_access = "public"
		pi_image_bucket_region = "us-south"
		pi_image_bucket_file_name = "%[4]s"
		pi_image_storage_type = "tier1"
	}
	`, name, pi_cloud_instance_id, pi_image_bucket_name, pi_image_bucket_file_name)
}
