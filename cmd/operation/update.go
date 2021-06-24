package operation

import (
	"github.com/pkg/errors"
	"github.com/takutakahashi/snip/pkg/repo"
)

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
