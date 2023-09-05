package reg_api

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	cfg "github.com/4nd3r5on/ctf1/internal/config"
	"github.com/4nd3r5on/ctf1/internal/domain/logic/mail_registration"
	mailRepo "github.com/4nd3r5on/ctf1/internal/repository/mail_verification"
	userRepo "github.com/4nd3r5on/ctf1/internal/repository/users"
	"gitlab.com/4nd3rs0n/errorsx/domain_errors"
)

type postMailRegReqBody struct {
	EMailAddress string `json:"email"`
}

func postMailReg(ctx context.Context,
	w http.ResponseWriter, r *http.Request, logger *slog.Logger,
	smtpConfig cfg.SMTPConfig,
	logoURL string, callbackEndpoint string,
	mr mailRepo.MailVerificationRepository) {

	var body postMailRegReqBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		msg := "Failed to parse request body."
		logger.Debug(msg, slog.String("error", err.Error()))
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// TODO: Not hardcoded locale
	errx := mail_registration.NewEMailVerification(ctx, mr, smtpConfig,
		mail_registration.NewEMailVerificationOpts{
			EMail:            body.EMailAddress,
			Locale:           mail_registration.LocaleEN,
			LogoURL:          logoURL,
			CallbackEndpoint: callbackEndpoint,
		})
	if errx != nil {
		domain_errors.HttpHandleDomainErr(errx, logger, w, r)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Verification e-mail sent. Check your inbox"))
}

type postMailCallbackReqBody struct {
	Name           string `json:"name"`
	VerificationID string `json:"verificationID"`
	Username       string `json:"username"`
	Password       string `jsoon:"password"`
}

type postMailCallbackRespBody struct {
	UserID string `json:"userID"`
	// TODO: Put refresh token into a secure cookie
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}

func postMailCallback(ctx context.Context,
	w http.ResponseWriter, r *http.Request, logger *slog.Logger,
	mr mailRepo.MailVerificationRepository, ur userRepo.UsersRepository) {

	var body postMailCallbackReqBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		msg := "Failed to parse request body."
		logger.Debug(msg, slog.String("error", err.Error()))
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	UserID, rt, at, errx := mail_registration.EmailCallback(ctx, mail_registration.EmailCallbackOpts{
		Name:           body.Name,
		VerificationID: body.VerificationID,
		Username:       body.Username,
		Password:       body.Password,
	}, mr, ur)
	if errx != nil {
		domain_errors.HttpHandleDomainErr(errx, logger, w, r)
		return
	}

	resp := postMailCallbackRespBody{
		UserID:       UserID.String(),
		RefreshToken: rt,
		AccessToken:  at,
	}
	respJSON, err := json.Marshal(resp)
	if err != nil {
		msg := "Failed to marshal responce body"
		logger.Error(msg,
			slog.String("error", err.Error()),
			slog.String("user_id", UserID.String()))
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	w.Write(respJSON)
	w.WriteHeader(http.StatusOK)
}
