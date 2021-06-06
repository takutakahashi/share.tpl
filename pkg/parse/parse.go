package parse

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/takutakahashi/share.tpl/pkg/cfg"
)

func Execute(conf cfg.Config, in []byte, data map[string]string) ([]byte, error) {
	data, err := fill(conf, data)
	fmt.Println(conf)
	if err != nil {
		return nil, err
	}
	return execute(in, data)
}

func execute(in []byte, data map[string]string) ([]byte, error) {
	tmpl, err := template.New("file.txt").Funcs(sprig.TxtFuncMap()).Parse(string(in))
	if err != nil {
		return nil, err
	}
	result := bytes.Buffer{}
	if err := tmpl.Execute(&result, data); err != nil {
		return nil, err
	}
	return result.Bytes(), nil
}

func fill(conf cfg.Config, data map[string]string) (map[string]string, error) {
	for _, v := range conf.Values {
		if _, ok := data[v.Name]; !ok && v.Default != "" {
			data[v.Name] = v.Default
		}
		if _, ok := data[v.Name]; !ok {
			return nil, errors.New(fmt.Sprintf("value %s is not found.", v.Name))
		}
	}
	return data, nil
}
