package ms

import (
	"fmt"
	"time"

	"github.com/balazsgrill/extractld/schema"
	"github.com/balazsgrill/sparqlupdate"
	"github.com/knakk/rdf"
)

const (
	dateURLLayout = "date://2006/01/02"
	dateURLFormat = "date://%04d/%02d/%02d"
)

type dateProcessor struct {
	DateProcessor
}

type DateProcessor interface {
	ProcessDate(time.Time) (*sparqlupdate.Graph, error)
}

func DateUri(t time.Time) rdf.IRI {
	return schema.Define(fmt.Sprintf(dateURLFormat, t.Year(), t.Month(), t.Day()))
}

// date://YYYY/MM/DD
func (p *dateProcessor) Process(url string) (*sparqlupdate.Graph, error) {
	time, err := time.Parse(dateURLLayout, url)
	if err != nil {
		return nil, nil
	}
	return p.ProcessDate(time)
}
