package user_api

import (
	"context"
	"log/slog"
	"net/http"

	reg_api "github.com/4nd3r5on/ctf1/internal/api/user/reg"
	cfg "github.com/4nd3r5on/ctf1/internal/config"
	mailRepo "github.com/4nd3r5on/ctf1/internal/repository/mail_verification"
	userRepo "github.com/4nd3r5on/ctf1/internal/repository/users"
	"github.com/go-chi/chi/v5"
)

type UserApiOpts struct {
	SMTPConfig cfg.SMTPConfig
	URepo      userRepo.UsersRepository
	MRepo      mailRepo.MailVerificationRepository
	Logger     *slog.Logger

	LogoURL string
	// This URL should be the path where this code is called,
	// routes before this code, in format http(s)://example.com/api/v3
	ThisURL string
}

func NewUserAPI(ctx context.Context, r chi.Router, opts UserApiOpts) {
	var userRoute = "/user"
	r.Route(userRoute, func(r chi.Router) {
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				getUsers(ctx, w, r, opts.URepo)
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		})

		reg_api.NewRegistrationAPI(ctx, r, reg_api.RegistrationApiOpts{
			SMTPConfig: opts.SMTPConfig,
			URepo:      opts.URepo,
			MRepo:      opts.MRepo,
			Logger:     opts.Logger,

			LogoURL: opts.LogoURL,
			ThisURL: opts.ThisURL + userRoute,
		})
	})
}
