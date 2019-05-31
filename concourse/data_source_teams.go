package concourse

import (
	"fmt"
	"time"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceTeams() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTeamsRead,

		Schema: map[string]*schema.Schema{
			"names": {
				Type: schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceTeamsRead(d *schema.ResourceData, m interface{}) error {
	client := *m.(*CombinedConfig).client
	teams, err := client.ListTeams()
	
	if err != nil {
		return fmt.Errorf("err: %s", err)
	}

	d.SetId(time.Now().UTC().String())
	
	names := []string{}
	for _, t := range teams {
		name := t.Name
		names = append(names, name)
	}

	d.Set("names", names)

	return nil
}
