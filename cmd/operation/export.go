package operation

import (
	"errors"
	"io/ioutil"

	"github.com/takutakahashi/share.tpl/pkg/cfg"
	"github.com/takutakahashi/share.tpl/pkg/parse"
)

type ExportOpt struct {
	Path string
	Type string
	Data map[string]string
}

type ExportOut struct {
	Files map[string][]byte
}

func Export(opt ExportOpt) (ExportOut, error) {
	switch opt.Type {
	case "snippet":
		return exportSnippet(opt)
	case "files":
		return exportFiles(opt)
	}
	return ExportOut{}, errors.New("unknown type")
}
func exportSnippet(opt ExportOpt) (ExportOut, error) {
	ret := ExportOut{
		Files: map[string][]byte{},
	}
	in, err := ioutil.ReadFile(opt.Path)
	if err != nil {
		return ExportOut{}, err
	}
	conf, err := cfg.ParsePath(opt.Path)
	if err != nil {
		return ExportOut{}, err
	}
	out, err := parse.Execute(conf, in, opt.Data)
	if err != nil {
		return ExportOut{}, err
	}
	ret.Files["stdout"] = out
	return ret, nil
}

func exportFiles(opt ExportOpt) (ExportOut, error) {
	ret := ExportOut{}
	return ret, nil
}
