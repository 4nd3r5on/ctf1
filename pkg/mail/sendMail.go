package mail

import "net/smtp"

const GmailSMTP = "smtp.google.com:587"

type SendMailOptions struct {
	From         string
	To           string
	SMTPServer   string // smtp.google.com:587
	SMTPPassword string
	EmailSubject string
	EmailBody    string
}

func SendEmail(opts SendMailOptions) error {
	auth := smtp.PlainAuth("", opts.From, opts.SMTPPassword, opts.SMTPServer)

	msg := []byte("To: " + opts.To + "\r\n" +
		"Subject: " + opts.EmailSubject + "\r\n" +
		"\r\n" +
		opts.EmailBody + "\r\n")

	err := smtp.SendMail(opts.SMTPServer, auth, opts.From,
		[]string{opts.To}, msg)
	if err != nil {
		return err
	}
	return nil
}
