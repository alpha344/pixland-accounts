package interfaces

import (
	"context"

	"github.com/alpha344/pixland-accounts/server/pkg/permittable"
	"github.com/alpha344/pixland-accounts/server/pkg/role"
	"github.com/alpha344/pixland-accounts/server/pkg/user"
)

type UpdatePermittableParam struct {
	UserID  user.ID
	RoleIDs []role.ID
}

type Permittable interface {
	GetUsersWithRoles(context.Context) (user.List, map[user.ID]role.List, error)
	UpdatePermittable(context.Context, UpdatePermittableParam) (*permittable.Permittable, error)
}



