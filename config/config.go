package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Config struct {
	Panel struct {
		URL string `validate:"url" yaml:"url"`
		Key string `yaml:"key"`
		ID  string `yaml:"id"`
	} `yaml:"panel"`
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return !os.IsNotExist(err)
	}

	return true
}

func Get() *Config {
	// etc/psecurity/config.yaml

	dir := "/etc/psecurity"
	if !exists(dir) {
		return nil
	}

	path := filepath.Join(dir, "config.yaml")

	if !exists(path) {
		return nil
	}

	buf, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	var cfg *Config
	if err = yaml.Unmarshal(buf, &cfg); err != nil {
		return nil
	}

	return cfg
}

func Create() error {
	dir := "/etc/psecurity"
	if !exists(dir) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	path := filepath.Join(dir, "config.yaml")

	if exists(path) {
		return nil
	}

	fd, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fd.Close()

	cfg := Config{}

	enc := yaml.NewEncoder(fd)
	enc.SetIndent(2)
	enc.Encode(cfg)
	enc.Close()

	return nil
}
