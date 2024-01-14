package ms

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/balazsgrill/extractld"
	"github.com/balazsgrill/oauthenticator"
	"github.com/balazsgrill/oauthenticator/client"
	"github.com/balazsgrill/sparqlupdate"
)

const (
	clientType = "https://graph.microsoft.com#MSGraphClient"
	RootURL    = "https://graph.microsoft.com/v1.0/"
)

type msDateProcessor struct {
	ms *client.Oauth2Client
}

func CreateProcessors(provider oauthenticator.Provider) ([]extractld.UrlProcessor, error) {
	configs, err := provider.ConfigsOfType(clientType)
	if err != nil {
		return nil, err
	}
	var result []extractld.UrlProcessor
	for _, c := range configs {
		msgraphclient := client.New(c.Config(), c.Token())
		result = append(result, &dateProcessor{
			DateProcessor: &msDateProcessor{
				ms: msgraphclient,
			},
		}, &outlookProcessor{
			ms: msgraphclient,
		})
	}

	return result, nil
}

func get(ms *client.Oauth2Client, request string) (resp *http.Response, err error) {
	return ms.Get(RootURL + request)
}

func parseget(ms *client.Oauth2Client, request string, value interface{}) (err error) {
	log.Println(request)
	resp, err := get(ms, request)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Println(string(data))
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(data))
	}
	log.Println(string(data))
	return json.Unmarshal(data, value)
}

func concaterr(e1, e2 error) error {
	if e1 == nil {
		return e2
	}
	if e2 == nil {
		return e1
	}
	return fmt.Errorf("%v ; %v", e1, e2)
}

func (p *msDateProcessor) ProcessDate(start time.Time) (*sparqlupdate.Graph, error) {
	end := start.Add(24 * time.Hour)
	result := sparqlupdate.New()
	err := MessageListOfInterval(p.ms, result, start, end)
	err = concaterr(err, CalendarViewOfInterval(p.ms, result, start, end))
	return result, err
}
