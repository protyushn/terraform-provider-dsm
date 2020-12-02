// **********
// Terraform Provider - SDKMS: main provider program
// **********
//       - Author:    fyoo at fortanix dot com
//       - Version:   0.1
//       - Date:      27/11/2020
// **********

package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"terraform-provider-sdkms/sdkms"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return sdkms.Provider()
		},
	})
}