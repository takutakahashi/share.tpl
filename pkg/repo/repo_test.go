package repo

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestRepo_Update(t *testing.T) {
	os.Setenv("GITHUB_USERNAME", "takutakahashi")
	token, _ := ioutil.ReadFile("../../.ignore/token")
	os.Setenv("GITHUB_TOKEN", string(token))

	type fields struct {
		BaseDir    string
		Credential *Credential
		Name       string
		Type       string
		URI        string
		Revision   string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				BaseDir:    "fill by test",
				Credential: nil,
				Name:       "snippets",
				Type:       "git",
				URI:        "https://github.com/takutakahashi/snippets.git",
				Revision:   "main",
			},
		},
		{
			name: "with credential",
			fields: fields{
				BaseDir: "fill by test",
				Credential: &Credential{
					Username: Secret{
						EnvName: "GITHUB_USERNAME",
					},
					Password: Secret{
						EnvName: "GITHUB_TOKEN",
					},
				},
				Name:     "snippets",
				Type:     "git",
				URI:      "https://github.com/takutakahashi/private-snippets.git",
				Revision: "main",
			},
		},
	}
	for _, tt := range tests {
		dir, _ := ioutil.TempDir("../../misc", "base-")
		tt.fields.BaseDir = dir
		t.Run(tt.name, func(t *testing.T) {
			r := Repo{
				BaseDir:    tt.fields.BaseDir,
				Credential: tt.fields.Credential,
				Name:       tt.fields.Name,
				Type:       tt.fields.Type,
				URI:        tt.fields.URI,
				Revision:   tt.fields.Revision,
			}
			if err := r.Update(); (err != nil) != tt.wantErr {
				t.Errorf("Repo.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		os.RemoveAll(dir)
	}
}
