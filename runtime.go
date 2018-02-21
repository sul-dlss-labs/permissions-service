package app

import (
	"github.com/sul-dlss-labs/permissions-service/config"
)

// NewRuntime creates a new application level runtime that
// encapsulates the shared services for this application
func NewRuntime(config *config.Config) (*Runtime, error) {
	return &Runtime{
		config: config,
	}, nil
}

// Runtime encapsulates the shared services for this application
type Runtime struct {
	config *config.Config
}

// Config returns the config for this application
func (r *Runtime) Config() *config.Config {
	return r.config
}
