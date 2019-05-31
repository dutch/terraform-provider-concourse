package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/dutch/terraform-provider-concourse/concourse"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: concourse.Provider})
}
