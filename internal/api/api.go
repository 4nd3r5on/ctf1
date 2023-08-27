package api

import (
	"context"

	user_api "github.com/4nd3r5on/ctf1/internal/api/user"
	"github.com/4nd3r5on/ctf1/internal/domain/logic/mail_registration"
	mailRepo "github.com/4nd3r5on/ctf1/internal/repository/mail_verification"
	"github.com/4nd3r5on/ctf1/internal/repository/users"
	"github.com/4nd3r5on/ctf1/pkg/chi_utils"
	chi "github.com/go-chi/chi/v5"
)

type ApiOpts struct {
	ApiPrefix  string
	SMTPConfig mail_registration.SMTPConfig
	UR         users.UsersRepository
	MVR        mailRepo.MailVerificationRepository
}

func NewAPI(ctx context.Context, opts ApiOpts) *chi.Mux {
	r := chi.NewMux()

	r.Use(chi_utils.EnableCors)
	r.Route(opts.ApiPrefix, func(r chi.Router) {
		user_api.NewUserAPI(ctx, r, user_api.UserApiOpts{
			SMTPConfig: opts.SMTPConfig,

			UR:  opts.UR,
			MVR: opts.MVR,
		})
	})

	return r
}
