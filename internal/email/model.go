package email

type Email struct {
	From string
	To   []string
	CC   []string

	Subject     string
	Body        string
	Attach      string
	Attachments []string
}
