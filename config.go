package elasticsearch

import "github.com/Selly-Modules/natsio"

type Config struct {
	ApiKey string
	Nats   natsio.Config
}
