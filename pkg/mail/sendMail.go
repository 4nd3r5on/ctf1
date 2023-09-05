package mail

import "net/smtp"

const GmailSMTP = "smtp.google.com:587"

type SendMailOptions struct {
	From         string
	To           string
	SMTPServer   string
	SMTPPassword string
	EmailSubject string
	EmailBody    string
	HTML         bool
}

func SendEmail(opts SendMailOptions) error {
	auth := smtp.PlainAuth("", opts.From, opts.SMTPPassword, opts.SMTPServer)

	from := "From: " + opts.From + "\n"
	to := "To: " + opts.To + "\n"
	subject := "MIME-version: 1.0;\n Subject: " + opts.EmailSubject + "\n"
	var mime string = "charset=\"UTF-8\";"
	if opts.HTML {
		mime = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	}
	body := opts.EmailBody
	msg := []byte(from + to + subject + mime + body)

	err := smtp.SendMail(opts.SMTPServer, auth, opts.From,
		[]string{opts.To}, msg)
	if err != nil {
		return err
	}
	return nil
}
