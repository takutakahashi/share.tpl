package git

import (
	"fmt"
	"os/exec"

	"github.com/pkg/errors"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GitRepo struct {
	dirpath    string
	uri        string
	revision   string
	credential *Credential
}

func New(dirpath, uri, branch string, cred *Credential) GitRepo {
	return GitRepo{
		dirpath:    dirpath,
		uri:        uri,
		revision:   branch,
		credential: cred,
	}
}

func (g GitRepo) Clone() error {
	args := []string{}
	if g.revision != "" {
		args = append(args, "-b", g.revision)
	}
	args = append(args, g.uri, g.dirpath)
	return g.cmd("clone", args)
}

func (g GitRepo) Pull() error {
	if g.revision == "" {
		return nil
	}
	return g.cmd("pull", []string{"origin", g.revision})
}

func (g GitRepo) Checkout() error {
	if g.revision == "" {
		return nil
	}
	if err := g.Fetch(); err != nil {
		return err
	}
	return g.cmd("checkout", []string{g.revision})
}

func (g GitRepo) Fetch() error {
	return g.cmd("fetch", []string{})
}

func (g GitRepo) cmd(action string, args []string) error {
	if action != "clone" {
		base := []string{"--git-dir", fmt.Sprintf("%s/.git", g.dirpath), "--work-tree", g.dirpath, action}
		base = append(base, args...)
		args = base
	} else {
		base := []string{action}
		base = append(base, args...)
		args = base
	}
	_, err := exec.Command("git", args...).Output()
	return errors.Wrapf(err, "failed to execute %s", action)
}
