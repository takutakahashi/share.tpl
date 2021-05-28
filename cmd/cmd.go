package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/takutakahashi/share.tpl/cmd/operation"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "life.go",
		Commands: []*cli.Command{
			{
				Name: "export",
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:  "set",
						Usage: "set variables. multiple value",
					},
				},
				Action: func(c *cli.Context) error {
					sets := c.StringSlice("set")
					path := c.Args().First()
					_ = sets
					data := map[string]string{}
					for _, s := range sets {
						sp := strings.Split(s, "=")
						if len(sp) != 2 {
							return errors.New("invalid args")
						}
						data[sp[0]] = sp[1]
					}
					out, err := operation.Export(operation.ExportOpt{
						Path: path,
						Type: "snippet",
						Data: data,
					})
					if err != nil {
						return err
					}
					fmt.Println(string(out.Files["stdout"]))
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
