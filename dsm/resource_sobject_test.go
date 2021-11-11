// Code generated by apic. DO NOT EDIT.

package dsm

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

var (
	resourceSobject_createConfig = `resource "dsm_group" "example_group" {
  		name = "example_group"
	}

	resource "dsm_sobject" "example_sobject" {
		name     = "example_sobject"
		group_id = "${dsm_group.example_group.group_id}"
		key_size = 256
		key_ops = [
			"ENCRYPT",
			"DECRYPT",
	  	    "WRAPKEY",
	 	    "UNWRAPKEY",
			"DERIVEKEY",
			"MACGENERATE",
			"MACVERIFY",
			"APPMANAGEABLE"
		]
		obj_type = "AES"
	}`

	resourceSobject_updateConfig = `

    resource "dsm_sobject" "example_sobject" {
		name     = "example_sobject_updated"
	  	group_id = "${dsm_group.example_group.group_id}"
	 	key_size = 256
	  	key_ops = [
			"ENCRYPT",
		  	"DECRYPT",
			"WRAPKEY",
		   	"UNWRAPKEY",
		  	"DERIVEKEY",
		  	"MACGENERATE",
		  	"MACVERIFY",
		  	"APPMANAGEABLE"
	  	]
	  	obj_type = "AES"
  	}`
)

func TestAccResourceSobject(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDestroySobject,
		Steps: []resource.TestStep{
			{
				Config: resourceSobject_createConfig,
			},
		},
	})
}

func testAccCheckDestroySobject(s *terraform.State) (err error) {
	return err
}
