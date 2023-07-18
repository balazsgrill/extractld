package ms

import (
	"fmt"
	"time"

	"github.com/balazsgrill/extractld/ms/data"
	"github.com/balazsgrill/extractld/schema"
	"github.com/balazsgrill/oauthenticator/client"
	"github.com/balazsgrill/sparqlupdate"
	"github.com/knakk/rdf"
)

func CalendarView(ms *client.Oauth2Client, start time.Time, end time.Time) (*data.CalendarView, error) {
	result := &data.CalendarView{}
	// Format: 2023-01-25T10:03:53.528Z
	return result, parseget(ms, fmt.Sprintf("me/calendar/calendarview?startdatetime=%s&enddatetime=%s", start.Format(time.RFC3339), end.Format(time.RFC3339)), result)
}

func CalendarViewEntryToGraph(cal data.CalendarViewEntry, start time.Time, graph *sparqlupdate.Graph) {
	dateItem := DateUri(start)
	msgi := schema.Define(cal.WebLink)
	graph.AddTriple(dateItem, schema.DCT_hasPart, msgi)
	graph.AddTriple(msgi, schema.DCT_subject, rdf.NewTypedLiteral(cal.Subject, schema.XSD_String))
}

func CalendarViewOfInterval(ms *client.Oauth2Client, graph *sparqlupdate.Graph, start time.Time, end time.Time) error {
	calv, err := CalendarView(ms, start, end)
	if err != nil {
		return err
	}
	for _, cal := range calv.Value {
		CalendarViewEntryToGraph(cal, start, graph)
	}
	return nil
}
