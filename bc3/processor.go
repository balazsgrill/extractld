package bc3

import (
	"context"
	"strings"

	"github.com/balazsgrill/basecamp3"
	"github.com/balazsgrill/extractld"
	"github.com/balazsgrill/oauthenticator"
	"github.com/balazsgrill/sparqlupdate"
)

type BC3Processor struct {
	*basecamp3.Basecamp
	Ctx basecamp3.ContextWithTokenPersistence
}

var _ extractld.UrlProcessor = &BC3Processor{}

const (
	urlprefix  = "https://3.basecamp.com/"
	clientType = "https://3.basecampapi.com#Basecamp3Client"
)

func CreateProcessors[C oauthenticator.Config](provider oauthenticator.Provider[C]) ([]extractld.UrlProcessor, error) {
	configs, err := provider.ConfigsOfType(clientType)
	if err != nil {
		return nil, err
	}
	var result []extractld.UrlProcessor
	for _, c := range configs {
		bc := basecamp3.NewByOauth(c.Config())
		tp := &tokenPersistence{
			Context:          context.Background(),
			TokenPersistence: provider.Token(c),
		}
		bc.TokenPersitence = tp.get
		result = append(result, &BC3Processor{
			Basecamp: bc,
			Ctx:      tp,
		})
	}

	return result, nil
}

func (b *BC3Processor) Process(url string) (*sparqlupdate.Graph, error) {
	spath, found := strings.CutPrefix(url, urlprefix)
	if !found {
		return nil, nil
	}
	path := strings.Split(spath, "/")

	var level HierarchyLevel = &Root{
		Basecamp: b.Basecamp,
		Ctx:      b.Ctx,
	}
	for _, l := range path {
		l2 := level.Get(l)
		if l2 == nil {
			return nil, nil
		}
		level = l2
	}

	result := sparqlupdate.New()
	return result, level.Extract(result)
}
