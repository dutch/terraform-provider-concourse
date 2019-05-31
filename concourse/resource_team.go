package concourse

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/concourse/concourse/atc"
)

func resourceTeam() *schema.Resource {
	return &schema.Resource{
		Create: resourceTeamCreate,
		Read: resourceTeamRead,
		Update: resourceTeamUpdate,
		Delete: resourceTeamDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceTeamCreate(d *schema.ResourceData, m interface{}) error {
	return resourceTeamUpdate(d, m)
}

func resourceTeamRead(d *schema.ResourceData, m interface{}) error {
	client := *m.(*CombinedConfig).client
	name := d.Get("name").(string)

	proxy := client.Team(name)
	team, _, err := proxy.Team(name)
	if err != nil {
		return fmt.Errorf("read error: %s", err)
	}

	d.Set("name", team.Name)

	return nil
}

func resourceTeamUpdate(d *schema.ResourceData, m interface{}) error {
	client := *m.(*CombinedConfig).client
	name := d.Get("name").(string)

	proxy := client.Team(name)
	_, _, _, err := proxy.CreateOrUpdate(atc.Team{
		Name: name,
		Auth: map[string]map[string][]string{
			"groups": make(map[string][]string, 0),
			"users": make(map[string][]string, 0),
		},
	})
	if err != nil {
		return fmt.Errorf("update error: %s", err)
	}

	d.SetId(name)

	return resourceTeamRead(d, m)
}

func resourceTeamDelete(d *schema.ResourceData, m interface{}) error {
	client := *m.(*CombinedConfig).client
	name := d.Get("name").(string)

	proxy := client.Team(name)
	err := proxy.DestroyTeam(name)
	if err != nil {
		return fmt.Errorf("delete error: %s", err)
	}

	d.SetId("")

	return nil
}
