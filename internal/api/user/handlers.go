package user_api

import (
	"context"
	"net/http"

	userRepo "github.com/4nd3r5on/ctf1/internal/repository/users"
)

func getUsers(ctx context.Context, w http.ResponseWriter, r *http.Request, ur userRepo.UsersRepository) {

}
