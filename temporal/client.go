package temporal

import (
	"go.temporal.io/sdk/client"
	"workflow-service/config"
)

func NewClient(cfg *config.Config) (client.Client, error) {
	return client.Dial(client.Options{
		HostPort: cfg.TemporalHost,
	})
}
