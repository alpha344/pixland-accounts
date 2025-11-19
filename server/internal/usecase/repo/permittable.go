package repo

import (
	"context"

	"github.com/pixair-technologies/pixland-accounts/server/pkg/id"
	"github.com/pixair-technologies/pixland-accounts/server/pkg/permittable"
	"github.com/pixair-technologies/pixland-accounts/server/pkg/user"
)

//go:generate mockgen -source=./permittable.go -destination=./mock_repo/mock_permittable.go -package mock_repo
type Permittable interface {
	FindByUserID(context.Context, user.ID) (*permittable.Permittable, error)
	FindByUserIDs(context.Context, user.IDList) (permittable.List, error)
	FindByRoleID(context.Context, id.RoleID) (permittable.List, error)
	Save(context.Context, permittable.Permittable) error
}



