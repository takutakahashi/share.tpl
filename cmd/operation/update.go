package operation

import (
	"github.com/pkg/errors"
	"github.com/takutakahashi/snip/pkg/global"
	"github.com/takutakahashi/snip/pkg/repo"
	"github.com/urfave/cli/v2"
)

func CommandUpdate() *cli.Command {
	return &cli.Command{
		Name:        "update",
		Description: "update repositories",
		Action:      DoUpdate,
	}
}

func DoUpdate(c *cli.Context) error {
	s, err := global.LoadSetting(c.String("config"))
	if err != nil {
		return err
	}
	op, err := New(s)
	if err != nil {
		return err
	}
	return op.Update()

}

func (op Operation) Update() error {
	for _, r := range op.g.Repositories {
		repo := repo.Repo{
			BaseDir:    op.g.Setting.BaseDir,
			Name:       r.Name,
			Type:       r.Type,
			URI:        r.URI,
			Revision:   r.Revision,
			Credential: &r.Credential,
		}
		if err := repo.Update(); err != nil {
			return errors.Wrapf(err, "failed to update %s", r.Name)
		}
	}
	return nil
}
