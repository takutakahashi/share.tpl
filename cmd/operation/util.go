package operation

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func Write(data map[string][]byte) error {
	for filepath, buf := range data {
		if filepath == "stdout" {
			fmt.Println(string(buf))
			continue
		}
		if err := os.MkdirAll(path.Dir(filepath), 0775); err != nil {
			return err
		}
		if err := ioutil.WriteFile(filepath, buf, 0775); err != nil {
			return err
		}
	}
	return nil
}
