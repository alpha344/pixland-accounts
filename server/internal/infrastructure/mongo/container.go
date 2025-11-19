package mongo

import (
	"context"

	"github.com/alpha344/pixland-accounts/server/internal/usecase/repo"
	"github.com/alpha344/pixlandx/mongox"
	"go.mongodb.org/mongo-driver/mongo"
)

func New(ctx context.Context, db *mongo.Database, useTransaction, needCompat bool, users []repo.User) (*repo.Container, error) {
	client := mongox.NewClientWithDatabase(db)
	if useTransaction {
		client = client.WithTransaction()
	}

	var ws repo.Workspace
	if needCompat {
		ws = NewWorkspaceCompat(client)
	} else {
		ws = NewWorkspace(client)
	}

	lock, err := NewLock(db.Collection("locks"))
	if err != nil {
		return nil, err
	}

	c := &repo.Container{
		User:        NewUser(client),
		Workspace:   ws,
		Role:        NewRole(client),
		Permittable: NewPermittable(client),
		Transaction: client.Transaction(),
		Users:       users,
		Config:      NewConfig(db.Collection("config"), lock),
	}

	return c, nil
}



