package gql

import (
	"context"

	"github.com/pixair-technologies/pixland-accounts/server/internal/adapter/gql/gqlmodel"
	"github.com/pixair-technologies/pixland-accounts/server/internal/usecase/interfaces"
	"github.com/pixair-technologies/pixland-accounts/server/pkg/id"
)

func (r *mutationResolver) UpdatePermittable(ctx context.Context, input gqlmodel.UpdatePermittableInput) (*gqlmodel.UpdatePermittablePayload, error) {
	userId, err := gqlmodel.ToID[id.User](input.UserID)
	if err != nil {
		return nil, err
	}

	roleIds, err := gqlmodel.ToIDs[id.Role](input.RoleIds)
	if err != nil {
		return nil, err
	}

	permittable, err := usecases(ctx).Permittable.UpdatePermittable(ctx, interfaces.UpdatePermittableParam{
		UserID:  userId,
		RoleIDs: roleIds,
	})
	if err != nil {
		return nil, err
	}

	return &gqlmodel.UpdatePermittablePayload{
		Permittable: gqlmodel.ToPermittable(permittable),
	}, nil
}



