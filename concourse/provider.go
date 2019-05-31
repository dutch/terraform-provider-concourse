package concourse

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"target": {
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("CONCOURSE_TARGET", nil),
				Description: "The fly target for the Concourse instance.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"concourse_teams": dataSourceTeams(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"concourse_team": resourceTeam(),
			"concourse_pipeline": resourcePipeline(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Target: d.Get("target").(string),
	}

	return config.Client()
}
