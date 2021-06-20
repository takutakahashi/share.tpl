package repo

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestRepo_Update(t *testing.T) {
	type fields struct {
		BaseDir    string
		Credential Credential
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
				Credential: Credential{},
				Name:       "snippets",
				Type:       "git",
				URI:        "https://github.com/takutakahashi/snippets.git",
				Revision:   "main",
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
