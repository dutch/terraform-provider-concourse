package concourse

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourcePipeline() *schema.Resource {
	return &schema.Resource{
		Create: resourcePipelineCreate,
		Read: resourcePipelineRead,
		Update: resourcePipelineUpdate,
		Delete: resourcePipelineDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourcePipelineCreate(d *schema.ResourceData, m interface{}) error {
	return resourcePipelineRead(d, m)
}

func resourcePipelineRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePipelineUpdate(d *schema.ResourceData, m interface{}) error {
	return resourcePipelineRead(d, m)
}

func resourcePipelineDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
