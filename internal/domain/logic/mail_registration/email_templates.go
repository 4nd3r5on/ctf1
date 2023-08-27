package mail_registration

const (
	LocaleEN = iota
	LocaleRS = iota
)

// TODO
func NewVerificationEmailBody(locale int) string {
	switch locale {
	case LocaleRS:
		return "" +
			"" +
			"" +
			"" +
			""
	default:
		return "" +
			"" +
			"" +
			"" +
			""
	}
}
