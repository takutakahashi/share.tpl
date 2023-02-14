package operation

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/takutakahashi/snip/pkg/global"
	"github.com/urfave/cli/v2"
)

func CommandExec() *cli.Command {
	return &cli.Command{
		Name:        "exec",
		Description: "execute command from templates",
		Action:      DoExec,
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:  "set",
				Usage: "set variables. multiple value",
			},
			&cli.BoolFlag{
				Name:    "quiet",
				Aliases: []string{"q"},
				Usage:   "execute with all prompt `yes`",
			},
			&cli.BoolFlag{
				Name:    "dry-run",
				Aliases: []string{"d"},
				Usage:   "only output, not execute",
			},
		},
	}
}

func DoExec(c *cli.Context) error {
	sets := c.StringSlice("set")
	path := c.Args().First()
	dryRun := c.Bool("dry-run")
	quiet := c.Bool("quiet")
	data := map[string]string{}
	for _, s := range sets {
		sp := strings.Split(s, "=")
		if len(sp) != 2 {
			return errors.New("invalid args")
		}
		data[sp[0]] = sp[1]
	}
	s, err := global.LoadSetting(c.String("config"))
	if err != nil {
		return err
	}
	op, err := New(s)
	if err != nil {
		return err
	}
	return op.Exec(path, data, dryRun, quiet)
}

func (op Operation) Exec(path string, data map[string]string, dryRun, quiet bool) error {
	out, err := op.Export(ExportOpt{
		Path:          path,
		OutputDirPath: "",
		Sets:          data,
	})
	if err != nil {
		return err
	}
	if _, ok := out.Files["stdout"]; !ok {
		return errors.New("Executable command is not found. use stdout snippets")
	}
	f, err := ioutil.TempFile(op.g.Setting.BaseDir, "tmp-")
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	_, err = f.Write(out.Files["stdout"].Data)
	if err != nil {
		return err
	}
	if dryRun {
		logrus.Infof("bash %s", string(out.Files["stdout"].Data))
		return nil
	}
	if !quiet {
		logrus.Info("You are attempt to exec below script:")
		logrus.Infof("\n---\n%s\n---", string(out.Files["stdout"].Data))
		if !confirm() {
			return errors.New("execution was declined.")
		}
	}
	stdout, err := exec.Command("bash", f.Name()).CombinedOutput()
	fmt.Print(string(stdout))
	return err
}

func confirm() bool {
	var s string

	fmt.Printf("(y/N): ")
	_, err := fmt.Scan(&s)
	if err != nil {
		panic(err)
	}

	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {
		return true
	}
	return false
}
