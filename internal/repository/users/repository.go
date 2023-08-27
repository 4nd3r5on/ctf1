package users

import (
	"context"

	"github.com/4nd3r5on/ctf1/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateUserOpts struct {
	Name         string
	EMail        string
	UsernameBase string
	Password     string
}

type UsersRepository interface {
	CreateUser(ctx context.Context, opts CreateUserOpts) (UserID uuid.UUID, err error)
	// GetUsers returns list of an existing users
	GetUsers(ctx context.Context) ([]entities.User, error)
	// TestUsername returns available Username ID
	TestUsername(ctx context.Context, username string) (id int, err error)
}

func NewUsersRepository(pool *pgxpool.Pool) UsersRepository {
	return &postgresUsersRepository{
		pool: pool,
	}
}
