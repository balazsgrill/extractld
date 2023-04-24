package data

type CalendarView struct {
	Value []CalendarViewEntry `json:"value"`
}

type CalendarViewEntry struct {
	ID        string         `json:"id"`
	WebLink   string         `json:"webLink"`
	Subject   string         `json:"subject"`
	Body      *MessageBody   `json:"body"`
	Organizer *ContactInfo   `json:"organizer"`
	Attendees []*ContactInfo `json:"attendees"`
}
