package repo

import (
	"fmt"
	"io/ioutil"

	"github.com/takutakahashi/snip/pkg/git"
)

type Repo struct {
	BaseDir    string     `json:"basedir"`
	Credential Credential `json:"credential"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	URI        string     `json:"uri"`
	Revision   string     `json:"revision"`
}

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r Repo) Update() error {
	repoDir := fmt.Sprintf("%s/%s", r.BaseDir, r.Name)
	if _, err := ioutil.ReadDir(repoDir); err != nil {
		if err := r.init(repoDir); err != nil {
			return err
		}
	}
	return r.update(repoDir)
}

func (r Repo) init(repoDir string) error {
	if r.Type == "git" {
		var cred *git.Credential = nil
		if r.Credential.Username != "" {
			cred = &git.Credential{
				Username: r.Credential.Username,
				Password: r.Credential.Password,
			}
		}
		g := git.New(repoDir, r.URI, r.Revision, cred)
		if err := g.Clone(); err != nil {
			return err
		}
	}
	return nil
}

func (r Repo) update(repoDir string) error {
	if r.Type == "git" {
		cred := &git.Credential{
			Username: r.Credential.Username,
			Password: r.Credential.Password,
		}
		g := git.New(repoDir, r.URI, r.Revision, cred)
		if err := g.Checkout(); err != nil {
			return err
		}
		return g.Pull()
	}
	return nil
}
