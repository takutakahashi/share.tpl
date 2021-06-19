package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/takutakahashi/snip/cmd/operation"
	"github.com/takutakahashi/snip/pkg/global"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "snip",
		Action: func(c *cli.Context) error {
			sets := c.StringSlice("set")
			output := c.String("output")
			path := c.Args().First()
			data := map[string]string{}
			s, err := global.LoadSetting(c.String("config"))
			if err != nil {
				return err
			}
			op, err := operation.New(s)
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
			out, err := op.Export(operation.ExportOpt{
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
			return operation.Write(out.Files)
		},

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
		Commands: []*cli.Command{
			{
				Name:        "list",
				Description: "list templates",
				Action: func(c *cli.Context) error {
					s, err := global.LoadSetting(c.String("config"))
					if err != nil {
						return err
					}
					op, err := operation.New(s)
					if err != nil {
						return err
					}
					out, err := op.List()
					if err != nil {
						return err
					}
					return operation.PrintList(out)
				},
			},
			{
				Name:        "show",
				Description: "show templates",
				Action: func(c *cli.Context) error {
					path := c.Args().First()
					s, err := global.LoadSetting(c.String("config"))
					if err != nil {
						return err
					}
					op, err := operation.New(s)
					if err != nil {
						return err
					}
					out, err := op.Show(path)
					if err != nil {
						return err
					}
					fmt.Println(out)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
