package git

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGitRepo_Clone(t *testing.T) {
	type fields struct {
		dirpath    string
		uri        string
		revision   string
		credential *Credential
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{

		{
			name: "ok",
			fields: fields{
				dirpath:    "fill by test",
				uri:        "https://github.com/takutakahashi/snippets.git",
				revision:   "main",
				credential: &Credential{},
			},
		},
	}
	for _, tt := range tests {
		dir, _ := ioutil.TempDir("../../misc/", "")
		t.Run(tt.name, func(t *testing.T) {
			g := GitRepo{
				dirpath:    dir,
				uri:        tt.fields.uri,
				revision:   tt.fields.revision,
				credential: tt.fields.credential,
			}
			if err := g.Clone(); (err != nil) != tt.wantErr {
				t.Log(dir)
				t.Errorf("GitRepo.Clone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		os.RemoveAll(dir)
	}
}

func TestGitRepo_Pull(t *testing.T) {
	type fields struct {
		dirpath    string
		uri        string
		revision   string
		credential *Credential
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				dirpath:    "fill by test",
				uri:        "https://github.com/takutakahashi/snippets.git",
				revision:   "main",
				credential: &Credential{},
			},
		},
	}
	for _, tt := range tests {
		dir, _ := ioutil.TempDir("../../misc/", "")
		t.Run(tt.name, func(t *testing.T) {
			g := GitRepo{
				dirpath:    dir,
				uri:        tt.fields.uri,
				revision:   tt.fields.revision,
				credential: tt.fields.credential,
			}
			g.Clone()
			if err := g.Pull(); (err != nil) != tt.wantErr {
				t.Errorf("GitRepo.Pull() error = %v, wantErr %v", err, tt.wantErr)
			}
			os.RemoveAll(dir)
		})
	}
}

func TestGitRepo_Fetch(t *testing.T) {
	type fields struct {
		dirpath    string
		uri        string
		revision   string
		credential *Credential
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				dirpath:    "fill by test",
				uri:        "https://github.com/takutakahashi/snippets.git",
				revision:   "main",
				credential: &Credential{},
			},
		},
	}
	for _, tt := range tests {
		dir, _ := ioutil.TempDir("../../misc/", "")
		t.Run(tt.name, func(t *testing.T) {
			g := GitRepo{
				dirpath:    dir,
				uri:        tt.fields.uri,
				revision:   tt.fields.revision,
				credential: tt.fields.credential,
			}
			g.Clone()
			if err := g.Fetch(); (err != nil) != tt.wantErr {
				t.Errorf("GitRepo.Fetch() error = %v, wantErr %v", err, tt.wantErr)
			}
			os.RemoveAll(dir)
		})
	}
}
