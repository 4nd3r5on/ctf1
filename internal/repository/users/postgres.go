package users

import (
	"context"

	"github.com/4nd3r5on/ctf1/internal/domain/entities"
	"github.com/4nd3r5on/ctf1/pkg/password_utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/4nd3rs0n/username"
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

func (ur *postgresUsersRepository) GetUsers(ctx context.Context) ([]entities.User, error) {
	rows, err := ur.pool.Query(ctx, ""+
		"SELECT (id, name, email, username_base, username_id, created_at) "+
		"FROM public.users")
	if err != nil {
		return []entities.User{}, err
	}

	var users []entities.User
	for rows.Next() {
		var user entities.User
		var usernameBase string
		var usernameID int
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.EMail,
			&usernameBase, &usernameID)
		if err != nil {
			return []entities.User{}, err
		}
		user.Username, err = username.NewUsername(usernameBase, usernameID)
		if err != nil {
			return []entities.User{}, err
		}
		users = append(users, user)
	}
	rows.Close()

	return users, nil
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
