package memory

import (
	"github.com/alpha344/pixland-accounts/server/internal/usecase/repo"
	"github.com/alpha344/pixlandx/usecasex"
)

func New() *repo.Container {
	return &repo.Container{
		User:        NewUser(),
		Workspace:   NewWorkspace(),
		Role:        NewRole(),
		Permittable: NewPermittable(),
		Transaction: &usecasex.NopTransaction{},
		Config:      NewConfig(),
	}
}



