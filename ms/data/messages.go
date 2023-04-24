package data

import "time"

type Messages struct {
	Value []Message `json:"value"`
}

type MessageList struct {
	Value []ListedMessage `json:"value"`
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
	Body          *MessageBody   `json:"body"`
	Sender        *ContactInfo   `json:"sender"`
	From          *ContactInfo   `json:"from"`
	ToRecipients  []*ContactInfo `json:"toRecipients"`
	CcRecipients  []*ContactInfo `json:"ccRecipients"`
	BccRecipients []*ContactInfo `json:"bccRecipients"`
	ReplyTo       []*ContactInfo `json:"replyTo"`
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
	Name    string `json:"name"`
	Address string `json:"address"`
}
