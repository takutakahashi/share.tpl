package operation

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/takutakahashi/snip/pkg/cfg"
	"github.com/takutakahashi/snip/pkg/global"
	"github.com/urfave/cli/v2"
)

func CommandShow() *cli.Command {
	return &cli.Command{
		Name:        "show",
		Description: "show templates",
		Action:      DoShow,
	}
}

func DoShow(c *cli.Context) error {
	path := c.Args().First()
	s, err := global.LoadSetting(c.String("config"))
	if err != nil {
		return err
	}
	op, err := New(s)
	if err != nil {
		return err
	}
	out, err := op.Show(path)
	if err != nil {
		return err
	}
	fmt.Println(out)
	return nil

}

func (op Operation) Show(path string) (string, error) {
	p, err := op.extractPath(path)
	if err != nil {
		return "", err
	}
	cfg, err := cfg.ParsePath(p)
	if err != nil {
		return "", err
	}
	ev := ""
	for _, v := range cfg.Values {
		if v.Default == "" {
			v.Default = "None"
		}
		if v.Description == "" {
			v.Description = "no Description"
		}
		ev += fmt.Sprintf("%s ... %s, default: %s\n", v.Name, v.Description, v.Default)
	}
	var f []byte
	if snippet(path) {
		f, err = ioutil.ReadFile(fmt.Sprintf("%s/snippet", path))
	}
	if err != nil {
		return "", err
	}
	dsc := `
  Description: {{ .Description }}
  Embedded values: 
{{ .EmbedValues | indent 4 }}
{{- if .F -}}
  content: |
{{ .F | indent 4 }}
# end-of-content
{{- end -}}
  `

	tmpl, err := template.New("file.txt").Funcs(sprig.TxtFuncMap()).Parse(dsc)
	if err != nil {
		return "", err
	}
	result := bytes.Buffer{}
	if err := tmpl.Execute(&result, struct {
		Description string
		EmbedValues string
		F           string
	}{
		Description: cfg.Description,
		EmbedValues: ev,
		F:           string(f),
	}); err != nil {
		return "", err
	}
	return result.String(), nil
}
