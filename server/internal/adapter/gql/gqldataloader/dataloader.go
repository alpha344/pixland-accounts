package gqldataloader

//go:generate go run github.com/vektah/dataloaden UserLoader github.com/pixair-technologies/pixland-accounts/server/internal/adapter/gql/gqlmodel.ID *github.com/pixair-technologies/pixland-accounts/server/internal/adapter/gql/gqlmodel.User
//go:generate go run github.com/vektah/dataloaden WorkspaceLoader github.com/pixair-technologies/pixland-accounts/server/internal/adapter/gql/gqlmodel.ID *github.com/pixair-technologies/pixland-accounts/server/internal/adapter/gql/gqlmodel.Workspace



