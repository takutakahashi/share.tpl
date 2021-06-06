package operation

import (
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
	if _, err := ioutil.ReadFile(opt.Path); err == nil {
		return exportFile(opt)
	} else if _, err := ioutil.ReadDir(opt.Path); err == nil {
		return exportDir(opt)
	} else {
		return ExportOut{}, err
	}
}

func exportFile(opt ExportOpt) (ExportOut, error) {
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
