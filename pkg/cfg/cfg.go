package cfg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Values      []Value `json:"values"`
}

type Value struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Default     string `json:"default"`
}

func ParsePath(filePath string) (Config, error) {
	cfg := Config{}
	if _, err := ioutil.ReadDir(filePath); err != nil {
		filePath = path.Dir(filePath)
	}
	f, err := ioutil.ReadFile(fmt.Sprintf("%s/.tnp.yaml", filePath))
	if err != nil {
		return cfg, err
	}
	if err := yaml.Unmarshal(f, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil

}

func Parse() (Config, error) {
	var cfg *Config
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	paths := []string{
		"%s/.share.yaml",
		"$s/.share/config.yaml",
		"%s/.local/share/config.yaml",
	}
	for _, path := range paths {
		cfg = &Config{}
		f, err := ioutil.ReadFile(fmt.Sprintf(path, home))
		if err != nil {
			cfg = nil
			continue
		}
		if err := yaml.Unmarshal(f, cfg); err != nil {
			cfg = nil
			continue
		}
	}
	if cfg == nil {
		return Config{}, errors.New("failed to load configuration")
	}
	return *cfg, nil
}
