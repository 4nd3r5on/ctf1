package users

import (
	"context"

	"github.com/4nd3r5on/ctf1/pkg/password_utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresUsersRepository struct {
	pool *pgxpool.Pool
}

func (ur *postgresUsersRepository) CreateUser(ctx context.Context, opts CreateUserOpts) (uuid.UUID, error) {
	var userID uuid.UUID

	passwordHash, salt, err := password_utils.HashPassword(opts.Password)
	if err != nil {
		return uuid.Nil, err
	}

	err = ur.pool.QueryRow(ctx, ""+
		"INSERT INTO public.users "+
		"(name, email, password_hash, salt, username_base, username_id) "+
		"VALUES ($1, $2, $3, $4, $5, "+
		"COALESCE(SELECT MAX(username_id) FROM users WHERE username_base = $5)) "+
		"RETURNING id;",
		opts.Name, opts.EMail, passwordHash, salt, opts.UsernameBase).Scan(&userID)
	return userID, err
}

func (ur *postgresUsersRepository) TestUsername(ctx context.Context, username string) (int, error) {
	var id int
	err := ur.pool.QueryRow(ctx,
		"COALESCE("+
			"(SELECT MAX(username_id) "+
			"FROM user_data "+
			"WHERE username = $1), 0)", username).Scan(&id)
	return id, err
}
