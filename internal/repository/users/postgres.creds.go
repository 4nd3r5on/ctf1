package user_repo

import (
	"context"

	crypto_utils "github.com/4nd3r5on/ctf1/pkg/crypto"
	"github.com/google/uuid"
	"gitlab.com/4nd3rs0n/username"
)

func (ur *postgresUsersRepository) CheckCredsUsername(ctx context.Context, uname username.Username, password string) (bool, error) {
	var passwordHash, salt []byte
	err := ur.pool.QueryRow(ctx, "SELECT password_hash, salt "+
		"FROM users WHERE username_base=$1 AND username_id=$2",
		uname.Base(), uname.ID()).Scan(passwordHash, salt)
	if err != nil {
		return false, err
	}
	valid := crypto_utils.VerifyPassword(password, salt, passwordHash)

	return valid, nil
}

func (ur *postgresUsersRepository) CheckCredsID(ctx context.Context, userID uuid.UUID, password string) (bool, error) {
	var passwordHash, salt []byte
	err := ur.pool.QueryRow(ctx, "SELECT password_hash, salt "+
		"FROM users WHERE id=$1", userID).Scan(passwordHash, salt)
	if err != nil {
		return false, err
	}
	valid := crypto_utils.VerifyPassword(password, salt, passwordHash)

	return valid, nil
}

func (ur *postgresUsersRepository) CheckCredsMail(ctx context.Context, mail string, password string) (bool, error) {
	var passwordHash, salt []byte
	err := ur.pool.QueryRow(ctx, "SELECT TOP 1 password_hash, salt "+
		"FROM users WHERE email=$1", mail).Scan(passwordHash, salt)
	if err != nil {
		return false, err
	}

	valid := crypto_utils.VerifyPassword(password, salt, passwordHash)

	return valid, nil
}
