package repo

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/takutakahashi/snip/pkg/git"
)

type Repo struct {
	BaseDir    string      `json:"basedir"`
	Credential *Credential `json:"credential"`
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	URI        string      `json:"uri"`
	Revision   string      `json:"revision"`
}

type Credential struct {
	Username Secret `json:"username"`
	Password Secret `json:"password"`
}

type Secret struct {
	EnvName string `json:"envName"`
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
		if r.Revision == "" {
			return errors.New("git revision is not defined")
		}
		var cred *git.Credential = nil
		if r.Credential != nil {
			cred = &git.Credential{
				Username: os.Getenv(r.Credential.Username.EnvName),
				Password: os.Getenv(r.Credential.Password.EnvName),
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
		var cred *git.Credential = nil
		if r.Credential != nil {
			cred = &git.Credential{
				Username: os.Getenv(r.Credential.Username.EnvName),
				Password: os.Getenv(r.Credential.Password.EnvName),
			}
		}
		g := git.New(repoDir, r.URI, r.Revision, cred)
		if err := g.Checkout(); err != nil {
			return errors.Wrap(err, "failed to checkout repository")
		}
		return errors.Wrap(g.Pull(), "failed to pull")
	}
	return nil
}
