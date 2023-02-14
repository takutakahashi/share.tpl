package operation

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/takutakahashi/snip/pkg/cfg"
	"github.com/takutakahashi/snip/pkg/global"
	"github.com/takutakahashi/snip/pkg/parse"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

type ExportOpt struct {
	Path          string            `yaml:"path"`
	OutputDirPath string            `yaml:"output"`
	Sets          map[string]string `yaml:"sets"`
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
			&cli.StringFlag{
				Name:  "from",
				Usage: "set config file path",
			},
		},
		Action: DoExport,
	}
}

func DoExport(c *cli.Context) error {
	setsConfigPath := c.String("from")
	logrus.Info(setsConfigPath)
	s, err := global.LoadSetting(c.String("config"))
	if err != nil {
		return err
	}
	var o ExportOpt
	if setsConfigPath != "" {
		o, err = parseOpt(setsConfigPath)
		if err != nil {
			return err
		}
	} else {
		sets := c.StringSlice("set")
		output := c.String("output")
		path := c.Args().First()
		data := map[string]string{}
		for _, s := range sets {
			sp := strings.Split(s, "=")
			if len(sp) != 2 {
				return errors.New("invalid args")
			}
			data[sp[0]] = sp[1]

		}
		o = ExportOpt{
			Path:          path,
			OutputDirPath: output,
			Sets:          data,
		}

	}
	logrus.Info(o)
	op, err := New(s)
	if err != nil {
		return err
	}
	out, err := op.Export(o)
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
	out, err := parse.Execute(conf, in, opt.Sets)
	if err != nil {
		return ExportOut{}, err
	}
	ret.Files["stdout"] = parse.File{Data: out}
	return ret, nil
}

func exportDir(opt ExportOpt) (ExportOut, error) {
	conf, err := cfg.ParsePath(fmt.Sprintf("%s/.snip.yaml", opt.Path))
	if err != nil {
		return ExportOut{}, err
	}
	if opt.OutputDirPath == "" {
		if conf.Output == "" {
			return ExportOut{}, errors.New("output path is not found")
		}
		opt.OutputDirPath = conf.Output
	}
	files, err := parse.ExecuteFiles(conf, opt.Path, opt.OutputDirPath, opt.Sets)
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

func parseOpt(p string) (ExportOpt, error) {
	ret := ExportOpt{}
	f, err := os.ReadFile(p)
	if err != nil {
		return ret, err
	}
	err = yaml.Unmarshal(f, &ret)
	return ret, err
}
