package interfaces

import (
	"context"

	"github.com/pixair-technologies/pixland-accounts/server/pkg/permittable"
	"github.com/pixair-technologies/pixland-accounts/server/pkg/role"
	"github.com/pixair-technologies/pixland-accounts/server/pkg/user"
)

type UpdatePermittableParam struct {
	UserID  user.ID
	RoleIDs []role.ID
}

type Permittable interface {
	GetUsersWithRoles(context.Context) (user.List, map[user.ID]role.List, error)
	UpdatePermittable(context.Context, UpdatePermittableParam) (*permittable.Permittable, error)
}



