package operation

import (
	"errors"
	"strings"

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
	return op.Exec(path, data)
}

func (op Operation) Exec(path string, data map[string]string) error {
	return nil
}
