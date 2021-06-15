package global

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Setting struct {
	BaseDir      string       `json:"base_dir"`
	Repositories []Repository `json:"repositories"`
	Includes     []Include    `json:"include"`
}

type Repository struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
	Type string `json:"type"`
}

type Include struct{}

func LoadSetting(filepath string) (Setting, error) {
	s := &Setting{}
	home, err := os.UserHomeDir()
	if err != nil {
		return Setting{}, err
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
	return Setting{}, errors.New("failed to load configuration")
}

func fillDefault(s Setting) (Setting, error) {
	if s.BaseDir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return Setting{}, err
		}
		s.BaseDir = fmt.Sprintf("%s/.snip", home)
	}
	return s, nil
}
