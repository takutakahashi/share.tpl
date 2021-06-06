package operation

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"

	"github.com/takutakahashi/share.tpl/pkg/cfg"
)

func Show(path string) (string, error) {

	cfg, err := cfg.ParsePath(path)
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
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	dsc := `
  Description: {{ .Description }}
  Embedded values: 
    {{ .EmbedValues }}
  content:

{{ .F }}
  `

	tmpl, err := template.New("file.txt").Parse(dsc)
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
