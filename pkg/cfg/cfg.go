package cfg

import (
	"fmt"
	"io/ioutil"
	"path"

	"gopkg.in/yaml.v3"
)

type Snippet struct {
	Name string `json:"name"`
}

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
	f, err := ioutil.ReadFile(fmt.Sprintf("%s/.snip.yaml", filePath))
	if err != nil {
		return cfg, err
	}
	if err := yaml.Unmarshal(f, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil

}
