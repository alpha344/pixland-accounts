package gql

import (
	"context"

	"github.com/alpha344/pixland-accounts/server/internal/adapter/gql/gqlmodel"
	"github.com/alpha344/pixland-accounts/server/internal/usecase/interfaces"
	"github.com/alpha344/pixlandx/rerror"
)

func (r *queryResolver) CheckPermission(ctx context.Context, input gqlmodel.CheckPermissionInput) (*gqlmodel.CheckPermissionPayload, error) {
	u := getUser(ctx)
	if u == nil {
		return nil, rerror.ErrNotFound
	}

	res, err := usecases(ctx).Cerbos.CheckPermission(ctx, u.ID(), interfaces.CheckPermissionParam{
		Service:  input.Service,
		Resource: input.Resource,
		Action:   input.Action,
	})
	if err != nil {
		return nil, err
	}

	return &gqlmodel.CheckPermissionPayload{
		Allowed: res.Allowed,
	}, nil
}



