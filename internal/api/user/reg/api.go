package reg_api

import (
	"context"
	"net/http"

	"github.com/4nd3r5on/ctf1/internal/domain/logic/mail_registration"
	mailRepo "github.com/4nd3r5on/ctf1/internal/repository/mail_verification"
	userRepo "github.com/4nd3r5on/ctf1/internal/repository/users"
	"github.com/go-chi/chi/v5"
)

type RegistrationApiOpts struct {
	SMTPConfig mail_registration.SMTPConfig
	UR         userRepo.UsersRepository
	MVR        mailRepo.MailVerificationRepository
}

func NewRegistrationAPI(ctx context.Context, r chi.Router, opts RegistrationApiOpts) {
	r.Route("/reg", func(r chi.Router) {

		r.Route("/mail", func(r chi.Router) {
			r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case http.MethodPost:

				default:
					w.WriteHeader(http.StatusMethodNotAllowed)
				}

			})
			r.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
				switch r.Method {
				case http.MethodPost:

				default:
					w.WriteHeader(http.StatusMethodNotAllowed)
				}
			})
		})
	})
}
