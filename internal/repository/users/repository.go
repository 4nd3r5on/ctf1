package user_repo

import (
	"context"

	"github.com/4nd3r5on/ctf1/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/4nd3rs0n/username"
)

type CreateUserOpts struct {
	Name         string
	EMail        string
	UsernameBase string
	Password     string
}

type UsersRepository interface {
	CreateUser(ctx context.Context, opts CreateUserOpts) (UserID uuid.UUID, err error)

	GetUsers(ctx context.Context) ([]entities.User, error)
	GetPublicUsers(ctx context.Context) ([]entities.PublicUserInfo, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (entities.User, error)
	GetUserByUsername(ctx context.Context, username username.Username) (entities.User, error)

	CheckCredsUsername(ctx context.Context, username username.Username, password string) (valid bool, err error)
	CheckCredsID(ctx context.Context, userID uuid.UUID, password string) (valid bool, err error)
	CheckCredsMail(ctx context.Context, mail string, password string) (valid bool, err error)

	RemoveUserByID(ctx context.Context, id uuid.UUID) error
	RemoveUserByUsername(ctx context.Context, uname username.Username) error

	// TestUsername returns available Username ID
	TestUsername(ctx context.Context, username string) (id int, err error)
}

func NewUsersRepository(pool *pgxpool.Pool) UsersRepository {
	return &postgresUsersRepository{
		pool: pool,
	}
}
