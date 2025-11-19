package workspace

import (
	"github.com/pixair-technologies/pixland-accounts/server/pkg/gqlclient/gqlmodel"
)

type findByUserQuery struct {
	FindByUser []gqlmodel.Workspace `graphql:"findByUser(userId: $userId)"`
}



