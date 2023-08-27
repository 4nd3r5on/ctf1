package user_api

import (
	"context"
	"net/http"

	"github.com/4nd3r5on/ctf1/internal/domain/logic/mail_registration"
	mailRepo "github.com/4nd3r5on/ctf1/internal/repository/mail_verification"
	userRepo "github.com/4nd3r5on/ctf1/internal/repository/users"
	"github.com/go-chi/chi/v5"
)

type UserApiOpts struct {
	SMTPConfig mail_registration.SMTPConfig
	UR         userRepo.UsersRepository
	MVR        mailRepo.MailVerificationRepository
}

func NewUserAPI(ctx context.Context, r chi.Router, opts UserApiOpts) {
	r.Route("/user", func(r chi.Router) {
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		})
	})
}
