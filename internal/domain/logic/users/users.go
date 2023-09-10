package users_logic

import (
	"context"
	"net/http"

	"github.com/4nd3r5on/ctf1/internal/domain/entities"
	userRepo "github.com/4nd3r5on/ctf1/internal/repository/users"
	"gitlab.com/4nd3rs0n/errorsx/domain_errors"
)

func GetUsersList(
	ctx context.Context,
	ur userRepo.UsersRepository,
) ([]entities.PublicUserInfo, domain_errors.DomainErr) {
	users, err := ur.GetPublicUsers(ctx)
	if err != nil {
		return []entities.PublicUserInfo{},
			domain_errors.NewDomainErr(err.Error(), domain_errors.DomainErrInfo{
				Log:        true,
				HttpCode:   http.StatusInternalServerError,
				ApiMessage: "Failed to get list of users",
			})
	}
	return users, nil
}
