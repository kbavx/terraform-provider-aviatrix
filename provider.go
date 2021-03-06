package aviatrix

import (
	"os"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a schema.Provider for Aviatrix.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"controller_ip": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: envDefaultFunc("AVIATRIX_CONTROLLER_IP"),
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: envDefaultFunc("AVIATRIX_USERNAME"),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: envDefaultFunc("AVIATRIX_PASSWORD"),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"aviatrix_account": resourceAccount(),
			"aviatrix_admin_email" : resourceAdminEmail(),
			"aviatrix_customer_id" : resourceCustomerID(),
			"aviatrix_gateway": resourceAviatrixGateway(),
			"aviatrix_tunnel" : resourceTunnel(),
			"aviatrix_transpeer" : resourceTranspeer(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"aviatrix_caller_identity": dataSourceAviatrixCallerIdentity(),
			"aviatrix_account": dataSourceAviatrixAccount(),
			"aviatrix_gateway": dataSourceAviatrixGateway(),
		},
		ConfigureFunc: aviatrixConfigure,
	}
}

func envDefaultFunc(k string) schema.SchemaDefaultFunc {
	return func() (interface{}, error) {
		if v := os.Getenv(k); v != "" {
			return v, nil
		}

		return nil, nil
	}
}

func envDefaultFuncAllowMissing(k string) schema.SchemaDefaultFunc {
	return func() (interface{}, error) {
		v := os.Getenv(k)
		return v, nil
	}
}

func aviatrixConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		ControllerIP: d.Get("controller_ip").(string),
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
	}
	return config.Client()
}
