package operation

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/takutakahashi/share.tpl/pkg/cfg"
)

func Show(path string) (string, error) {
	if snippet(path) {
		path = fmt.Sprintf("%s/snippet", path)
	}
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
{{ .EmbedValues | indent 4 }}
  content: |
{{ .F | indent 4 }}
# end-of-content
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
