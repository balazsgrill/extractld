package extractld

import "time"

type Contact interface {
	Name() string
	Email() string
}

type Mail interface {
	Source() string
	Sender() Contact
	Recipients() []Contact
	Topic() string
	Body() string
	IsFlagged() bool
	SentTime() time.Time
}

type MailProvider interface {
	MailByDate(start time.Time, end time.Time) ([]Mail, error)
}
