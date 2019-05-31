package concourse

import (
	"fmt"
	"github.com/concourse/concourse/fly/rc"
	"github.com/concourse/concourse/go-concourse/concourse"
)

type Config struct {
	Target string
}

type CombinedConfig struct {
	client *concourse.Client
}

func (c *Config) Client() (*CombinedConfig, error) {
	fly := rc.TargetName(c.Target)

	target, err := rc.LoadTarget(fly, false)

	if err != nil {
		return nil, fmt.Errorf("err: %s", err)
	}

	client := target.Client()

	return &CombinedConfig{
		client: &client,
	}, nil
}
