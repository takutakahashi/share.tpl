package global

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Setting struct {
	Repositories []Repository `json:"repositories"`
}

type Repository struct {
	Path string `json:"path"`
}

func LoadSetting() (Setting, error) {
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
	for _, path := range paths {
		f, err := ioutil.ReadFile(fmt.Sprintf(path, home))
		if err != nil {
			s = nil
			continue
		}
		if err := yaml.Unmarshal(f, s); err != nil {
			s = nil
			continue
		}
		if s != nil {
			return *s, nil
		}
	}
	return Setting{}, errors.New("failed to load configuration")
}
