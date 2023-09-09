package operation

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	_ "embed"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

//go:embed src/config.yaml
var templ string

func CommandInit() *cli.Command {
	return &cli.Command{
		Name:        "init",
		Description: "init config",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "path",
				Usage: "new dir path",
			},
		},
		Action: DoInit,
	}
}

func DoInit(c *cli.Context) error {
	s := c.String("path")
	if s == "" {
		s = fmt.Sprintf("%s/.config/snip/config.yaml", os.Getenv("HOME"))
	}
	if err := os.MkdirAll(filepath.Dir(s), 0755); err != nil {
		return err
	}
	if b, err := os.ReadFile(s); err == nil && len(b) != 0 {
		return fmt.Errorf("config file already exists")
	}
	f, err := os.Create(s)
	if err != nil {
		return err
	}
	defer f.Close()
	tpl, err := template.New("config.yaml").Parse(templ)
	if err != nil {
		return err
	}
	if err := tpl.Execute(f,
		struct{ BaseDir string }{BaseDir: fmt.Sprintf("%s/.config/snip", os.Getenv("HOME"))}); err != nil {
		return err
	}
	logrus.Infof("Created config.yaml to %s", s)
	return nil
}
