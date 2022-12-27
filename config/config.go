package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	Z "github.com/rwxrob/bonzai/z"
	"gopkg.in/yaml.v3"
)

var (
	ErrMissingClientID     = errors.New("missing client ID")
	ErrMissingClientSecret = errors.New("missing client secret")
	ErrMissingServer       = errors.New("missing server")
)

// C holds the config parameters defined in the config file.
type C struct {
	ClientID     string `yaml:"clientID"`
	ClientSecret string `yaml:"clientSecret"`
	Server       string `yaml:"server"`
	Current      string `yaml:"currentContext"`
	Contexts     []Ctx  `yaml:"contexts"`
}

// New reads the config file and instantiates a new config object.
func New() (*C, error) {
	err := softInit()
	if err != nil {
		return nil, fmt.Errorf(
			"failed to initialize config: %s", err.Error(),
		)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf(
			"error reading %s: %s", path, err.Error(),
		)
	}

	c := new(C)
	err = c.unmarshal(data)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to unmarshal file to yaml: %s", err.Error(),
		)
	}

	return c, nil
}

// Validate validates the loaded configuration against a set of rules.
func (c *C) Validate() error {
	if c.ClientID == "" {
		return ErrMissingClientID
	}

	if c.ClientSecret == "" {
		return ErrMissingClientSecret
	}

	if c.Server == "" {
		return ErrMissingServer
	}

	return nil
}

// GetContexts returns a list of contexts.
func (c C) GetContexts() ([]Ctx, error) {
	if c.Contexts == nil || len(c.Contexts) == 0 {
		return nil, ErrContextsNotDefined
	}
	return c.Contexts, nil
}

// SetContexts sets the contexts. Note that you need to invoke the
// Save method in order to commit the changes.
func (c *C) SetContexts(ctxs []Ctx) {
	c.Contexts = ctxs
}

// InsertContext adds a new context to the list. Note that you need to
// invoke the Save method in order to commit the changes.
func (c *C) InsertContext(ctx Ctx) {
	contexts, err := c.GetContexts()
	if err != nil {
		contexts = make([]Ctx, 0)
	}

	contexts = append(contexts, ctx)
	c.SetContexts(contexts)
}

// GetContext returns the current context.
func (c C) GetContext() (*Ctx, error) {
	if c.Current == "" {
		return nil, ErrCurrentContextNotSet
	}

	contexts, err := c.GetContexts()
	if err != nil {
		return nil, err
	}

	for _, ctx := range contexts {
		if c.Current == ctx.Name {
			return &ctx, nil
		}
	}

	return nil, ErrContextNotFound
}

// SetContext sets the current context. Note that you need to invoke
// the Save method in order to commit the changes.
func (c *C) SetContext(name string) error {
	if !c.ContextExists(name) {
		return ErrContextNotFound
	}
	c.Current = name
	return nil
}

// RemoveContext deletes a context from the list. Note that you need
// to invoke the Save method in order to commit the changes.
func (c *C) RemoveContext(name string) error {
	contexts, err := c.GetContexts()
	if err != nil {
		return err
	}

	for i, ctx := range contexts {
		if ctx.Name == name {
			contexts = append(contexts[:i], contexts[i+1:]...)
			c.SetContexts(contexts)
			return nil
		}
	}

	return ErrContextNotFound
}

// Save commits the config to disk.
func (c C) Save() error {
	data, err := c.marshal()
	if err != nil {
		return fmt.Errorf("failed to marshal config: %s", err.Error())
	}

	return os.WriteFile(path, data, 0600)
}

// ContextExists checks if a given context exists in the list or not.
func (c C) ContextExists(name string) bool {
	contexts, err := c.GetContexts()
	if err != nil {
		return false
	}

	for _, ctx := range contexts {
		if ctx.Name == name {
			return true
		}
	}

	return false
}

// unmarshal unmarshals the raw yaml config into config.C struct.
func (c *C) unmarshal(data []byte) error {
	return yaml.Unmarshal(data, c)
}

// marshal marshals the config.C struct to yaml.
func (c *C) marshal() ([]byte, error) {
	return yaml.Marshal(c)
}

// path to the config file
var path string

// softInit creates the config file if it does not exists.
func softInit() error {
	err := Z.Conf.SoftInit()
	if err != nil {
		return err
	}

	cfg, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	path = filepath.Join(cfg, "mastodctl", "config.yaml")

	return nil
}
