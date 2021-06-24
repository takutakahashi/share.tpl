package git

import (
	"fmt"
	"os/exec"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
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
	if g.credential != nil {
		return g.cloneWithCredential()
	}
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
	if g.credential != nil {
		return g.pullWithCredential()
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
	if g.credential != nil {
		return g.checkoutWithCredential()
	}
	return g.cmd("checkout", []string{g.revision})
}

func (g GitRepo) Fetch() error {
	if g.credential != nil {
		return g.fetchWithCredential()
	}
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

func (g GitRepo) cloneWithCredential() error {

	_, err := git.PlainClone(g.dirpath, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: g.credential.Username,
			Password: g.credential.Password,
		},
		URL:           g.uri,
		ReferenceName: plumbing.NewBranchReferenceName(g.revision),
	})
	return err
}
func (g GitRepo) fetchWithCredential() error {
	msg := "failed to fetch"
	r, err := git.PlainOpen(g.dirpath)
	if err != nil {
		return errors.Wrap(err, msg)
	}
	err = r.Fetch(&git.FetchOptions{
		RemoteName: "origin",
		Auth: &http.BasicAuth{
			Username: g.credential.Username,
			Password: g.credential.Password,
		},
	})
	if err == git.NoErrAlreadyUpToDate {
		return nil
	}
	return errors.Wrap(err, msg)
}

func (g GitRepo) checkoutWithCredential() error {
	msg := "failed to checkout"
	auth := &http.BasicAuth{
		Username: g.credential.Username,
		Password: g.credential.Password,
	}
	r, err := git.PlainOpen(g.dirpath)
	if err != nil {
		return errors.Wrap(err, msg)
	}
	w, err := r.Worktree()
	if err != nil {
		return errors.Wrap(err, msg)
	}
	w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(g.revision),
		Create: true,
		Force:  true,
		Keep:   false,
	})
	err = w.Pull(&git.PullOptions{
		RemoteName:    "origin",
		ReferenceName: plumbing.NewBranchReferenceName(g.revision),
		Auth:          auth,
	})
	if err == git.NoErrAlreadyUpToDate {
		return nil
	}
	return errors.Wrap(err, msg)
}

func (g GitRepo) pullWithCredential() error {
	msg := "failed to pull"
	auth := &http.BasicAuth{
		Username: g.credential.Username,
		Password: g.credential.Password,
	}
	r, err := git.PlainOpen(g.dirpath)
	if err != nil {
		return errors.Wrap(err, msg)
	}
	w, err := r.Worktree()
	if err != nil {
		return errors.Wrap(err, msg)
	}
	err = w.Pull(&git.PullOptions{
		RemoteName:    "origin",
		ReferenceName: plumbing.NewBranchReferenceName(g.revision),
		Auth:          auth,
	})
	if err == git.NoErrAlreadyUpToDate {
		return nil
	}
	return errors.Wrap(err, msg)
}
