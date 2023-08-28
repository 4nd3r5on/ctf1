package users

import (
	"context"

	"github.com/4nd3r5on/ctf1/internal/domain/entities"
	"gitlab.com/4nd3rs0n/username"
)

func (ur *postgresUsersRepository) GetUsers(ctx context.Context) ([]entities.User, error) {
	rows, err := ur.pool.Query(ctx, ""+
		"SELECT id, name, email, username_base, username_id, created_at "+
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
			&usernameBase, &usernameID,
			&user.CreatedAt)
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

// This function is maybe bad by the reasons of the architecture and code clarity,
// code is almost duplicate, but it's required for optimization.
// It's expensive in Golang to move data from []entities.Users type to []entities.PublicUserInfo,
// request from DB only data that we need and put it to the right type from the begining is way faster.
func (ur *postgresUsersRepository) GetPublicUsers(ctx context.Context) ([]entities.PublicUserInfo, error) {
	rows, err := ur.pool.Query(ctx, ""+
		"SELECT id, name, username_base, username_id, created_at "+
		"FROM public.users")
	if err != nil {
		return []entities.PublicUserInfo{}, err
	}

	var users []entities.PublicUserInfo
	for rows.Next() {
		var user entities.PublicUserInfo
		var usernameBase string
		var usernameID int
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&usernameBase, &usernameID,
			&user.CreatedAt)
		if err != nil {
			return []entities.PublicUserInfo{}, err
		}
		user.Username, err = username.NewUsername(usernameBase, usernameID)
		if err != nil {
			return []entities.PublicUserInfo{}, err
		}
		users = append(users, user)
	}
	rows.Close()

	return users, nil
}
