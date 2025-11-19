package workspace

import (
	"github.com/alpha344/pixland-accounts/server/pkg/gqlclient/gqlmodel"
)

type findByUserQuery struct {
	FindByUser []gqlmodel.Workspace `graphql:"findByUser(userId: $userId)"`
}



