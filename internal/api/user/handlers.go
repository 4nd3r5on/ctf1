package user_api

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/4nd3r5on/ctf1/internal/domain/entities"
	users_logic "github.com/4nd3r5on/ctf1/internal/domain/logic/users"
	userRepo "github.com/4nd3r5on/ctf1/internal/repository/users"
	"gitlab.com/4nd3rs0n/errorsx/domain_errors"
)

type getUsersResp struct {
	Users []entities.PublicUserInfo `json:"users"`
}

func getUsers(ctx context.Context, w http.ResponseWriter, r *http.Request,
	ur userRepo.UsersRepository, logger *slog.Logger) {

	list, errx := users_logic.GetUsersList(ctx, ur)
	if errx != nil {
		domain_errors.HttpHandleDomainErr(errx, logger, w, r)
		return
	}

	resp, err := json.Marshal(getUsersResp{Users: list})
	if err != nil {
		msg := "Failed to marshal response body"
		logger.Error(msg, slog.String("error", err.Error()))
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
