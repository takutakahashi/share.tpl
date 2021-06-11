package operation

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/takutakahashi/share.tpl/pkg/cfg"
	"github.com/takutakahashi/share.tpl/pkg/parse"
)

type ExportOpt struct {
	Path          string
	OutputDirPath string
	Data          map[string]string
}

type ExportOut struct {
	Files map[string][]byte
}

func Export(opt ExportOpt) (ExportOut, error) {
	if snippet(opt.Path) {
		return exportFile(opt)
	} else {
		return exportDir(opt)
	}
}

func snippet(target string) bool {
	fis, err := ioutil.ReadDir(target)
	if err != nil {
		return false
	}
	var snippet, conf bool = false, false
	for _, info := range fis {
		if info.Name() == "snippet" {
			snippet = true
		}
		if info.Name() == ".share.yaml" {
			conf = true
		}
	}
	return snippet && conf
}

func exportFile(opt ExportOpt) (ExportOut, error) {
	opt.Path = fmt.Sprintf("%s/snippet", opt.Path)
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

func exportDir(opt ExportOpt) (ExportOut, error) {
	if opt.OutputDirPath == "" {
		return ExportOut{}, errors.New("output path is not found")
	}
	conf, err := cfg.ParsePath(fmt.Sprintf("%s/.share.yaml", opt.Path))
	if err != nil {
		return ExportOut{}, err
	}
	files, err := parse.ExecuteFiles(conf, opt.Path, opt.OutputDirPath, opt.Data)
	if err != nil {
		return ExportOut{}, err
	}
	ret := ExportOut{
		Files: files,
	}
	return ret, nil
}

func exportFiles(opt ExportOpt) (ExportOut, error) {
	ret := ExportOut{}
	return ret, nil
}
