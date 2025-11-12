package temporal

import (
	"crypto/tls"
	"go.temporal.io/sdk/client"
	"workflow-service/config"
)

func NewClient(cfg *config.Config) (client.Client, error) {
	// dial to temporal server
	return client.Dial(client.Options{
		HostPort:          cfg.TemporalHost,
		Namespace:         cfg.Namespace,
		ConnectionOptions: client.ConnectionOptions{TLS: &tls.Config{}},
		Credentials:       client.NewAPIKeyStaticCredentials(cfg.APIKey),
	})
}
