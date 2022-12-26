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

	ErrNoContextsFound      = errors.New("no contexts found")
	ErrCurrentContextNotSet = errors.New("current context not set")
)

type Ctx struct {
	Name        string `yaml:"name"`
	AccessToken string `yaml:"accessToken"`
}

type C struct {
	ClientID     string `yaml:"clientID"`
	ClientSecret string `yaml:"clientSecret"`
	Server       string `yaml:"server"`
	Current      string `yaml:"currentContext"`
	Contexts     []Ctx  `yaml:"contexts"`
}

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

func (c C) GetContexts() []Ctx {
	return c.Contexts
}

func (c C) CurrentContext() (*Ctx, error) {
	if c.Current == "" {
		return nil, ErrCurrentContextNotSet
	}

	contexts := c.GetContexts()
	if contexts == nil {
		return nil, ErrNoContextsFound
	}

	for _, ctx := range c.GetContexts() {
		if c.Current == ctx.Name {
			return &ctx, nil
		}
	}

	return nil, ErrNoContextsFound
}

func (c *C) Save() error {
	data, err := c.marshal()
	if err != nil {
		return fmt.Errorf("failed to marshal config: %s", err.Error())
	}

	return os.WriteFile(path, data, 0600)
}

func (c *C) unmarshal(data []byte) error {
	return yaml.Unmarshal(data, c)
}

func (c *C) marshal() ([]byte, error) {
	return yaml.Marshal(c)
}

var path string

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
