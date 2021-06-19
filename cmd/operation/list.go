package operation

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/takutakahashi/snip/pkg/cfg"
	"github.com/takutakahashi/snip/pkg/global"
	"gopkg.in/yaml.v2"
)

type ListOutput struct {
	Path        string
	Name        string
	Description string
}

func List(s global.SnipConfig) ([]ListOutput, error) {
	ret := []ListOutput{}
	for _, repo := range s.Repositories {
		path := fmt.Sprintf("%s/%s", s.Setting.BaseDir, repo.Name)
		out, err := listWithPath(path)
		if err != nil {
			return nil, err
		}
		for _, o := range out {
			o.Name = fmt.Sprintf("%s/%s", repo.Name, o.Name)
			ret = append(ret, o)
		}
	}
	return ret, nil
}

type s struct {
	Snippets []cfg.Snippet `json:"snippets"`
}

func listWithPath(path string) ([]ListOutput, error) {
	snippets := s{}
	ret := []ListOutput{}
	buf, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", path, ".root.snip.yaml"))
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(buf, &snippets); err != nil {
		return nil, err
	}
	if len(snippets.Snippets) == 0 {
		return nil, errors.New("no snippet found in configuration")
	}
	for _, snippet := range snippets.Snippets {
		lo := ListOutput{Name: snippet.Name}
		ret = append(ret, lo)
	}
	return ret, nil
}
