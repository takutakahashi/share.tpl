package operation

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func CommandNew() *cli.Command {
	return &cli.Command{
		Name:        "new",
		Description: "create new snippet",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "path",
				Usage: "new dir path",
			},
		},
		Action: DoNew,
	}
}

//go:embed src/snip.yaml
var snipyaml string

func DoNew(c *cli.Context) error {
	path := c.String("path")
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}
	f, err := os.Create(fmt.Sprintf("%s/.snip.yaml", path))
	if err != nil {
		return err
	}
	logrus.Info(snipyaml)
	defer f.Close()
	_, err = f.WriteString(snipyaml)
	return err
}
