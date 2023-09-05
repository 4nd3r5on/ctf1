package mail_registration

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/4nd3r5on/ctf1/internal/config"
	"github.com/4nd3r5on/ctf1/internal/repository"
	mailRepo "github.com/4nd3r5on/ctf1/internal/repository/mail_verification"
	userRepo "github.com/4nd3r5on/ctf1/internal/repository/users"
	"github.com/4nd3r5on/ctf1/pkg/mail"
	"github.com/google/uuid"
	"gitlab.com/4nd3rs0n/errorsx/domain_errors"
)

type NewEMailVerificationOpts struct {
	Locale           int
	EMail            string
	LogoURL          string
	CallbackEndpoint string
}

// Sends varification mail to specified address
func NewEMailVerification(
	ctx context.Context,
	mr mailRepo.MailVerificationRepository,
	smtpConfig config.SMTPConfig, opts NewEMailVerificationOpts) domain_errors.DomainErr {

	// TODO: E-mail validation

	// Generating the verification ID
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		return domain_errors.NewDomainErr(err.Error(), domain_errors.DomainErrInfo{
			Log:        true,
			HttpCode:   http.StatusInternalServerError,
			ApiMessage: "Failed to generate mail approvement ID",
		})
	}
	verificationID := hex.EncodeToString(buf)

	// Adding mail to a repository
	mr.AddMail(ctx, opts.EMail, verificationID)

	// Sending an E-mail
	err = mail.SendEmail(mail.SendMailOptions{
		From:         smtpConfig.From,
		To:           opts.EMail,
		SMTPServer:   smtpConfig.SMTPServer,
		SMTPPassword: smtpConfig.SMTPPassword,
		EmailSubject: "E-mail Verification",
		EmailBody:    NewVerificationEmailBody(opts.Locale, opts.LogoURL),
	})
	if err != nil {
		return domain_errors.NewDomainErr(err.Error(), domain_errors.DomainErrInfo{
			Log:        true,
			HttpCode:   500,
			ApiMessage: "Failed to send verification E-mail",
			Vars:       map[string]any{"receiver": opts.EMail},
		})
	}
	return nil
}

type EmailCallbackOpts struct {
	Name           string
	VerificationID string
	Username       string
	Password       string
}

// E-mail callback used to finish registration and create an account
func EmailCallback(
	ctx context.Context,
	opts EmailCallbackOpts,
	mr mailRepo.MailVerificationRepository,
	ur userRepo.UsersRepository) (uuid.UUID, string, string, domain_errors.DomainErr) {

	// TODO: Validation
	var noImpl string = "Not implemented yet, I'm sorry"

	mail, err := mr.GetMailByID(ctx, opts.VerificationID)
	if err != nil {
		if err == repository.ErrEntityNotFound {
			return uuid.Nil, "", "", domain_errors.NewDomainErr(err.Error(), domain_errors.DomainErrInfo{
				HttpCode:   http.StatusNotFound,
				ApiMessage: "Verification ID not found. Start registration from the beginning",
			})
		}
		return uuid.Nil, "", "", domain_errors.NewDomainErr(err.Error(), domain_errors.DomainErrInfo{
			Log:        true,
			HttpCode:   http.StatusInternalServerError,
			ApiMessage: "Failed to get mail by verification ID.",
			Vars:       map[string]any{"verification_id": opts.VerificationID},
		})
	}

	UserID, err := ur.CreateUser(ctx, userRepo.CreateUserOpts{
		Name:         opts.Name,
		EMail:        mail,
		UsernameBase: opts.Username,
		Password:     opts.Password,
	})
	if err != nil {
		if err == repository.ErrEntityExists {
			msg := "User already taken"
			return uuid.Nil, "", "", domain_errors.NewDomainErr(msg, domain_errors.DomainErrInfo{
				HttpCode:   http.StatusConflict,
				ApiMessage: msg,
			})
		}
		return uuid.Nil, "", "", domain_errors.NewDomainErr(err.Error(), domain_errors.DomainErrInfo{
			Log:        true,
			HttpCode:   http.StatusInternalServerError,
			ApiMessage: "Failed to create a new user.",
			Vars: map[string]any{
				"name": opts.Name, "username": opts.Username, "email": mail},
		})
	}
	return UserID, noImpl, noImpl, nil
}
