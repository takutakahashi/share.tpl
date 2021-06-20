package global

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type SnipConfig struct {
	Setting      Setting      `json:"setting"`
	Repositories []Repository `json:"repositories"`
	Includes     []Include    `json:"include"`
}

type Setting struct {
	BaseDir string `json:"basedir"`
}

type Repository struct {
	Name     string `json:"name"`
	URI      string `json:"uri"`
	Type     string `json:"type"`
	Revision string `json:"revision"`
}

type Include struct{}

func LoadSetting(filepath string) (SnipConfig, error) {
	s := &SnipConfig{}
	home, err := os.UserHomeDir()
	if err != nil {
		return SnipConfig{}, err
	}
	paths := []string{
		"%s/.snip.yaml",
		"%s/.snip/config.yaml",
		"%s/.local/snip/config.yaml",
	}
	if filepath != "" {
		paths = []string{filepath}
	}
	for _, path := range paths {
		f, err := ioutil.ReadFile(fmt.Sprintf(path, home))
		if err != nil {
			f, err = ioutil.ReadFile(path)
			if err != nil {
				s = nil
				continue
			}
		}
		if err := yaml.Unmarshal(f, s); err != nil {
			s = nil
			continue
		}
		if s != nil {
			return fillDefault(*s)
		}
	}
	return SnipConfig{}, errors.New("failed to load configuration")
}

func fillDefault(s SnipConfig) (SnipConfig, error) {
	if s.Setting.BaseDir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return SnipConfig{}, err
		}
		s.Setting.BaseDir = fmt.Sprintf("%s/.snip", home)
	}
	return s, nil
}
