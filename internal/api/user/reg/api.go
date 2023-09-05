package reg_api

import (
	"context"
	"log/slog"
	"net/http"

	cfg "github.com/4nd3r5on/ctf1/internal/config"
	mailRepo "github.com/4nd3r5on/ctf1/internal/repository/mail_verification"
	userRepo "github.com/4nd3r5on/ctf1/internal/repository/users"
	"github.com/go-chi/chi/v5"
)

type RegistrationApiOpts struct {
	SMTPConfig cfg.SMTPConfig
	URepo      userRepo.UsersRepository
	MRepo      mailRepo.MailVerificationRepository
	Logger     *slog.Logger

	LogoURL string
	// This URL should be the path where this code is called,
	// routes before this code, in format http(s)://example.com/api/v3/user
	ThisURL string
}

func NewRegistrationAPI(ctx context.Context, r chi.Router, opts RegistrationApiOpts) {
	mailCallbackURL := opts.ThisURL + "/reg/mail/callback"

	r.Route("/reg", func(r chi.Router) {
		r.Route("/mail", func(r chi.Router) {
			r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case http.MethodPost:
					postMailReg(ctx, w, r, opts.Logger, opts.SMTPConfig,
						opts.LogoURL, mailCallbackURL, opts.MRepo)
				default:
					w.WriteHeader(http.StatusMethodNotAllowed)
				}
			})
			r.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case http.MethodPost:
					postMailCallback(ctx, w, r, opts.Logger, opts.MRepo, opts.URepo)
				default:
					w.WriteHeader(http.StatusMethodNotAllowed)
				}
			})
		})
	})
}
