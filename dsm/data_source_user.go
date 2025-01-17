package dsm

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUserRead,
		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_email": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	req, err := m.(*api_client).APICallList("GET", "sys/v1/users")
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "[DSM SDK] Unable to call DSM provider API client",
			Detail:   fmt.Sprintf("[E]: API: GET sys/v1/users: %v", err),
		})
		return diags
	}

	user_id := ""
	for _, data := range req {
		if data.(map[string]interface{})["user_email"].(string) == d.Get("user_email").(string) {
			user_id = data.(map[string]interface{})["user_id"].(string)
			if err := d.Set("user_email", d.Get("user_email").(string)); err != nil {
				return diag.FromErr(err)
			}
			if err := d.Set("user_id", user_id); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	d.SetId(user_id)
	return nil
}
