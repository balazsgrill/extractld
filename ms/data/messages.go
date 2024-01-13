package data

import (
	"time"

	"github.com/balazsgrill/extractld"
)

type Messages struct {
	Value []Message `json:"value"`
}

type MessageList struct {
	Value    []ListedMessage `json:"value"`
	NextLink string          `json:"@odata.nextlink"`
}

type ListedMessage struct {
	ID      string `json:"id"`
	Subject string `json:"subject"`
	WebLink string `json:"webLink"`
}

type Message struct {
	ID            string         `json:"id"`
	Subject       string         `json:"subject"`
	WebLink       string         `json:"webLink"`
	SentDateTime  time.Time      `json:"sentDateTime"`
	Body_         *MessageBody   `json:"body"`
	Sender_       *ContactInfo   `json:"sender"`
	From          *ContactInfo   `json:"from"`
	ToRecipients  []*ContactInfo `json:"toRecipients"`
	CcRecipients  []*ContactInfo `json:"ccRecipients"`
	BccRecipients []*ContactInfo `json:"bccRecipients"`
	ReplyTo       []*ContactInfo `json:"replyTo"`
	Flag          *MessageFlag   `json:"flag"`
}

type MessageBody struct {
	ContentType string `json:"contentType"`
	Content     string `json:"content"`
}

type MessageFlag struct {
	FlagStatus string `json:"flagStatus"`
}

type ContactInfo struct {
	EmailAddress `json:"emailAddress"`
}

type EmailAddress struct {
	Name_   string `json:"name"`
	Address string `json:"address"`
}

var _ extractld.Contact = &ContactInfo{}

func (d *EmailAddress) Name() string {
	return d.Name_
}

func (d *EmailAddress) Email() string {
	return d.Address
}

var _ extractld.Mail = &Message{}

func (d *Message) SentTime() time.Time {
	return d.SentDateTime
}

func (d *Message) Source() string {
	return d.WebLink
}

func (d *Message) Sender() extractld.Contact {
	return d.Sender_
}

func (d *Message) Recipients() []extractld.Contact {
	result := make([]extractld.Contact, len(d.ToRecipients))
	for i, r := range d.ToRecipients {
		result[i] = r
	}
	return result
}

func (d *Message) Topic() string {
	return d.Subject
}

func (d *Message) Body() string {
	if d.Body_ != nil {
		return d.Body_.Content
	}
	return ""
}

func (d *Message) IsFlagged() bool {
	if d.Flag != nil {
		return d.Flag.FlagStatus == "flagged"
	}
	return false
}
