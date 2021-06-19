package operation

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/takutakahashi/snip/pkg/global"
	"github.com/takutakahashi/snip/pkg/parse"
)

type Operation struct {
	g global.SnipConfig
}

func New(g global.SnipConfig) (Operation, error) {
	return Operation{
		g: g,
	}, nil
}

func Write(data map[string]parse.File) error {
	for filepath, buf := range data {
		if filepath == "stdout" {
			fmt.Println(string(buf.Data))
			continue
		}
		if err := os.MkdirAll(path.Dir(filepath), 0775); err != nil {
			return err
		}
		if err := ioutil.WriteFile(filepath, buf.Data, buf.Perm); err != nil {
			return err
		}
	}
	return nil
}

func PrintList(out []ListOutput) error {
	for _, o := range out {
		fmt.Println(o.Name)
	}
	return nil
}

func (op Operation) extractPath(path string) (string, error) {
	pathWithBase := fmt.Sprintf("%s/%s", op.g.Setting.BaseDir, path)
	if _, err := ioutil.ReadDir(pathWithBase); err == nil {
		return pathWithBase, nil
	}
	if _, err := ioutil.ReadDir(path); err == nil {
		return path, nil
	}
	return "", errors.New("The snippet with specified path is not found")
}
