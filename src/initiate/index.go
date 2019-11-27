package initiate

import (
	router "projects/browsertalk/src/apis/rest"
	"projects/browsertalk/src/config"
)

// Start ...
func Start() {
	config.Init()

	cfg := config.Get()

	router.Build(cfg)
}
