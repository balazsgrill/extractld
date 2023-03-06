package extractld

import (
	"fmt"

	"github.com/balazsgrill/sparqlupdate"
)

type UrlProcessor interface {
	Process(url string) (*sparqlupdate.Graph, error)
}

type UrlProcessorList struct {
	list []UrlProcessor
}

func (list *UrlProcessorList) Add(processor UrlProcessor) {
	list.list = append(list.list, processor)
}

func (list *UrlProcessorList) AddAll(processors []UrlProcessor) {
	list.list = append(list.list, processors...)
}

func mergeError(e1, e2 error) error {
	if e1 == nil {
		return e2
	}
	if e2 == nil {
		return e1
	}
	return fmt.Errorf("%v\n%v", e1, e2)
}

func mergeGraph(g1, g2 *sparqlupdate.Graph) *sparqlupdate.Graph {
	if g1 == nil {
		return g2
	}
	if g2 == nil {
		return g1
	}
	g1.Merge(g2)
	return g1
}

func (list *UrlProcessorList) Process(url string) (*sparqlupdate.Graph, error) {
	var result *sparqlupdate.Graph
	var err error

	for _, processor := range list.list {

		r, e := processor.Process(url)
		err = mergeError(err, e)
		result = mergeGraph(result, r)
	}

	return result, err
}
