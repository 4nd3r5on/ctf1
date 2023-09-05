package api

import (
	"context"
	"log/slog"

	user_api "github.com/4nd3r5on/ctf1/internal/api/user"
	"github.com/4nd3r5on/ctf1/internal/config"
	mailRepo "github.com/4nd3r5on/ctf1/internal/repository/mail_verification"
	userRepo "github.com/4nd3r5on/ctf1/internal/repository/users"
	"github.com/4nd3r5on/ctf1/pkg/chi_utils"
	chi "github.com/go-chi/chi/v5"
)

type ApiOpts struct {
	TlsEnabled bool
	Addr       string
	Port       int
	ApiPrefix  string
	ThisURL    string

	SMTPConfig config.SMTPConfig
	URepo      userRepo.UsersRepository
	MRepo      mailRepo.MailVerificationRepository
	Logger     *slog.Logger
}

func NewAPI(ctx context.Context, opts ApiOpts) *chi.Mux {
	r := chi.NewMux()
	r.Use(chi_utils.EnableCors)

	r.Route(opts.ApiPrefix, func(r chi.Router) {
		user_api.NewUserAPI(ctx, r, user_api.UserApiOpts{
			SMTPConfig: opts.SMTPConfig,
			URepo:      opts.URepo,
			MRepo:      opts.MRepo,
			Logger:     opts.Logger,

			LogoURL: opts.ThisURL + logoRoute,
			ThisURL: opts.ThisURL,
		})
		static(r, opts.Logger)
	})
	return r
}
