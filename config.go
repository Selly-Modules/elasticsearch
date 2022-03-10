package elasticsearch

import "github.com/Selly-Modules/natsio"

// Config int client elasticsearch
type Config struct {
	ApiKey string
	Nats   natsio.Config
}
