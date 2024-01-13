package ms

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/balazsgrill/extractld"
	"github.com/balazsgrill/extractld/ms/data"
	"github.com/balazsgrill/extractld/schema"
	"github.com/balazsgrill/oauthenticator/client"
	"github.com/balazsgrill/sparqlupdate"
	"github.com/knakk/rdf"
	"jaytaylor.com/html2text"
)

const (
	outlookurlprefix = "https://outlook.live.com/owa/"
)

type outlookProcessor struct {
	ms *client.Oauth2Client
}

func (p *outlookProcessor) MailByDate(start time.Time, end time.Time) ([]extractld.Mail, error) {
	msgs, err := MessagesByReceivedDay(p.ms, start, end)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	result := make([]extractld.Mail, 0, len(msgs.Value))
	for _, m := range msgs.Value {
		msg, err := GetMessageByID(p.ms, m.ID)
		if err == nil {
			result = append(result, msg)
		} else {
			log.Println(err)
		}
	}
	return result, nil
}

func MessagesByReceivedDay(ms *client.Oauth2Client, start time.Time, end time.Time) (*data.MessageList, error) {
	result := &data.MessageList{}
	//geturl := fmt.Sprintf("me/messages?$select=ID,Subject,WebLink&$filter=(receivedDateTime ge %s) and (receivedDateTime le %s)", start.Format(time.RFC3339), end.Format(time.RFC3339))
	geturl := "me/messages"
	values := &url.Values{}
	values.Add("$select", "ID,Subject,WebLink")
	values.Add("$filter", fmt.Sprintf("(receivedDateTime ge %s) and (receivedDateTime le %s)", start.Format(time.RFC3339), end.Format(time.RFC3339)))
	return result, parseget(ms, geturl+"?"+values.Encode(), result)
}

func GetMessageByID(ms *client.Oauth2Client, id string) (*data.Message, error) {
	result := &data.Message{}
	return result, parseget(ms, "me/messages/"+url.QueryEscape(id), result)
}

func MessageListOfInterval(ms *client.Oauth2Client, graph *sparqlupdate.Graph, start time.Time, end time.Time) error {
	msgl, err := MessagesByReceivedDay(ms, start, end)
	if err != nil {
		return err
	}
	for _, msg := range msgl.Value {
		ListedMessageToGraph(msg, start, graph)
	}
	return nil
}

func MessageToGraph(m *data.Message, graph *sparqlupdate.Graph) {
	msg := schema.Define(m.WebLink)
	graph.AddTriple(msg, schema.RDF_Type, schema.SIOCT_MailMessage)
	graph.AddTriple(msg, schema.DCT_created, rdf.NewTypedLiteral(m.SentDateTime.String(), schema.XSD_DateTime))
	sender := graph.NewBlank()
	graph.AddTriple(msg, schema.DCT_creator, sender)
	graph.AddTriple(sender, schema.FOAF_mbox, schema.Define("mailto:"+m.Sender_.Address))

	if m.Body_ != nil {
		var content string
		if m.Body_.ContentType == "html" {
			var err error
			content, err = html2text.FromString(m.Body_.Content)
			if err != nil {
				log.Println(err)
			}
		} else {
			content = m.Body_.Content
		}
		graph.AddTriple(msg, schema.SIOC_content, rdf.NewTypedLiteral(content, schema.XSD_String))
	}
}

func ListedMessageToGraph(msg data.ListedMessage, start time.Time, graph *sparqlupdate.Graph) {
	dateItem := DateUri(start)
	msgi := schema.Define(msg.WebLink)
	graph.AddTriple(dateItem, schema.DCT_hasPart, msgi)
	graph.AddTriple(msgi, schema.DCT_subject, rdf.NewTypedLiteral(msg.Subject, schema.XSD_String))
}

func (p *outlookProcessor) Process(urls string) (*sparqlupdate.Graph, error) {
	if !strings.HasPrefix(urls, outlookurlprefix) {
		return nil, nil
	}
	u, err := url.Parse(urls)
	if err != nil {
		return nil, err
	}
	itemid := u.Query().Get("ItemID")
	if itemid == "" {
		return nil, nil
	}
	log.Println("Getting message " + itemid)
	itemid = strings.ReplaceAll(itemid, "/", "-")
	itemid = strings.ReplaceAll(itemid, " ", "_")
	m, err := GetMessageByID(p.ms, itemid)
	if err != nil {
		return nil, err
	}
	result := sparqlupdate.New()
	MessageToGraph(m, result)
	return result, nil
}
