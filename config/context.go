package config

import "errors"

var (
	ErrContextNotFound      = errors.New("context not found")
	ErrCurrentContextNotSet = errors.New("current context not set")
	ErrContextsNotDefined   = errors.New("contexts not defined")
)

// Ctx represents a single context defined in the config file.
type Ctx struct {
	Name        string `yaml:"name"`
	AccessToken string `yaml:"accessToken"`
}
