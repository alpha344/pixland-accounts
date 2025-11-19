# PixLand Accounts - AI Coding Guidelines

## Architecture Overview
This is a Domain-Driven Design (DDD) microservice for centralized user, workspace, role, and permission management in PixLand's ecosystem. It exposes GraphQL APIs for CRUD operations and permission evaluations using Cerbos.

- **Domain Layer** (`pkg/`): Core business entities (User, Workspace, Role) with builder patterns and validation.
- **Application Layer** (`internal/usecase/`): Use cases orchestrated via interactors, with repository interfaces.
- **Infrastructure Layer** (`internal/infrastructure/`): MongoDB repos, Cerbos/Auth0 integrations.
- **Adapter Layer** (`internal/adapter/`): GraphQL resolvers generated from `schemas/*.graphql`.

Key flows: Other services query this via GraphQL for user/workspace data; permissions checked via Cerbos gRPC calls.

## Development Workflows
- **Local Dev**: `make dev` (Air hot-reload) after `make dev-install`.
- **Full Stack**: `make run` (Docker Compose with Cerbos + Accounts).
- **Testing**: `make test` (set `PIXLAND_ACCOUNTS_DB=mongodb://localhost:27017` for Mongo).
- **GraphQL Gen**: `make gql` after schema changes in `schemas/`.
- **Migrations**: `make run-migration` (auto-run on startup).

## Code Organization Patterns
- **Alphabetical Ordering**: Struct fields, functions, imports, and GraphQL types ordered alphabetically (e.g., `pkg/user/user.go`).
- **Builder Pattern**: Construct domain objects fluently: `user.New().ID(id.NewUserID()).Name("John").Build()`.
- **Context Handling**: Extract user/operator in adapters (`adapter.User(ctx)`), pass explicitly to use cases.
- **ID Types**: Use typed IDs like `id.NewWorkspaceID()` for type safety.
- **Error Handling**: Domain errors via `rerror.NewE(i18n.T("message"))`; internationalized.
- **GraphQL Resolvers**: Access use cases via `usecases(ctx).User.CreateUser(ctx, operator, input)`.

## Integration Points
- **Cerbos Permissions**: Evaluate via `CheckPermission` in `internal/usecase/interactor/cerbos.go` (builds Principal from user roles, queries Cerbos).
- **Auth0 JWT**: Auth middleware in `internal/app/app.go`; tokens validated via `appx.AuthMiddleware`.
- **MongoDB**: Repos in `internal/infrastructure/mongo/`; migrations in `migration/` (timestamped files).
- **Docker Network**: Attaches to `pixland` network for inter-service GraphQL calls.

## Testing Strategy
- **Domain Layer (`pkg/`)**: High coverage (80%+), test business logic, invariants.
- **Use Case Layer (`internal/usecase/`)**: High coverage, test workflows, use in-memory repos.
- **Adapter Layer (`internal/adapter/`)**: Test GraphQL flows, auth.
- Use `internal/infrastructure/memory` for unit tests; testcontainers for E2E.

## Common Patterns
- GraphQL resolver: `usecases(ctx).User.Method()` and `getOperator(ctx)`.
- Domain construction: `user.New().ID(id).Name(name).Build()`.
- Use case execution: `Run0`, `Run1`, etc. for transactions.
- Permission checking: `Usecase().WithReadableWorkspaces()`.
- Error handling: `rerror.NewE(i18n.T("message"))`.
- Repository: Interface segregation per domain.
- Model conversion: `gqlmodel.ToUser(domainUser)`.
- Context management: `AttachUser()`, `AttachOperator()`.
- Testing: `memory.New()` for in-memory.
- ID generation: `id.NewUserID()`.
- Authorization: Delegate to Cerbos via GraphQL.

## Examples
- **User Creation**: `user.New().NewID().Name(name).Email(email).PasswordPlainText(pwd).Build()`.
- **Permission Check**: `cerbos.CheckPermission(ctx, userID, interfaces.CheckPermissionParam{Service: "dashboard", Resource: "user", Action: "read"})`.
- **GraphQL Query**: Schemas in `schemas/user.graphql`; resolvers in `internal/adapter/gql/`.

Reference: `server/CLAUDE.md` for detailed DDD patterns and `README.md` for setup.