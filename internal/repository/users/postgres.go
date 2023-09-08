package user_repo

import (
	"context"

	crypto_utils "github.com/4nd3r5on/ctf1/pkg/crypto"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/4nd3rs0n/username"
)

type postgresUsersRepository struct {
	pool *pgxpool.Pool
}

func (ur *postgresUsersRepository) CreateUser(ctx context.Context, opts CreateUserOpts) (uuid.UUID, error) {
	var userID uuid.UUID

	passwordHash, salt, err := crypto_utils.HashPassword(opts.Password)
	if err != nil {
		return uuid.Nil, err
	}

	err = ur.pool.QueryRow(ctx, ""+
		"INSERT INTO public.users "+
		"(name, email, password_hash, salt, username_base, username_id) "+
		"VALUES ($1, $2, $3, $4, $5, "+
		"COALESCE((SELECT MAX(username_id) FROM users WHERE username_base = $5), 0) "+
		"RETURNING id;",
		opts.Name, opts.EMail, passwordHash, salt, opts.UsernameBase).Scan(&userID)
	return userID, err
}

func (ur *postgresUsersRepository) RemoveUserByID(ctx context.Context, id uuid.UUID) error {
	_, err := ur.pool.Exec(ctx, "DELETE FROM public.users WHERE id = $1;", id)
	return err
}

func (ur *postgresUsersRepository) RemoveUserByUsername(ctx context.Context, uname username.Username) error {
	_, err := ur.pool.Exec(ctx, ""+
		"DELETE FROM public.users WHERE username_base=$1 AND username_id=$2;",
		uname.Base(), uname.ID())
	return err
}

func (ur *postgresUsersRepository) TestUsername(ctx context.Context, username string) (int, error) {
	var id int
	err := ur.pool.QueryRow(ctx,
		"COALESCE("+
			"(SELECT MAX(username_id) "+
			"FROM users "+
			"WHERE username_base = $1), 0);", username).Scan(&id)
	return id, err
}
