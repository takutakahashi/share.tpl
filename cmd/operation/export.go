package operation

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/takutakahashi/snip/pkg/cfg"
	"github.com/takutakahashi/snip/pkg/global"
	"github.com/takutakahashi/snip/pkg/parse"
	"github.com/urfave/cli/v2"
)

type ExportOpt struct {
	Path          string
	OutputDirPath string
	Data          map[string]string
}

type ExportOut struct {
	Files map[string]parse.File
}

func CommandExport() *cli.Command {
	return &cli.Command{
		Name:        "export",
		Description: "export templates",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:  "set",
				Usage: "set variables. multiple value",
			},
			&cli.StringFlag{
				Name:  "output",
				Usage: "output dir path",
			},
			&cli.StringFlag{
				Name:  "config",
				Usage: "config path",
			},
		},
		Action: DoExport,
	}
}

func DoExport(c *cli.Context) error {

	sets := c.StringSlice("set")
	output := c.String("output")
	path := c.Args().First()
	data := map[string]string{}
	s, err := global.LoadSetting(c.String("config"))
	if err != nil {
		return err
	}
	op, err := New(s)
	if err != nil {
		return err
	}
	for _, s := range sets {
		sp := strings.Split(s, "=")
		if len(sp) != 2 {
			return errors.New("invalid args")
		}
		data[sp[0]] = sp[1]
	}
	out, err := op.Export(ExportOpt{
		Path:          path,
		OutputDirPath: output,
		Data:          data,
	})
	if err != nil {
		return err
	}
	if os.Getenv("DEBUG") != "" {
		fmt.Println(out.Files)
	}
	return Write(out.Files)
}

func (op Operation) Export(opt ExportOpt) (ExportOut, error) {
	path, err := op.extractPath(opt.Path)
	if err != nil {
		return ExportOut{}, err
	}
	opt.Path = path
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
		if info.Name() == ".snip.yaml" {
			conf = true
		}
	}
	return snippet && conf
}

func exportFile(opt ExportOpt) (ExportOut, error) {
	opt.Path = fmt.Sprintf("%s/snippet", opt.Path)
	ret := ExportOut{
		Files: map[string]parse.File{},
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
	ret.Files["stdout"] = parse.File{Data: out}
	return ret, nil
}

func exportDir(opt ExportOpt) (ExportOut, error) {
	if opt.OutputDirPath == "" {
		return ExportOut{}, errors.New("output path is not found")
	}
	conf, err := cfg.ParsePath(fmt.Sprintf("%s/.snip.yaml", opt.Path))
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
