package app

import (
	"log"
	"net/http"

	"github.com/balazsgrill/extractld"
	"github.com/balazsgrill/extractld/bc3"
	"github.com/balazsgrill/extractld/ms"
	cmd "github.com/balazsgrill/oauthenticator/server"
)

type ExtractorApp struct {
	cmd.MainApp
	List      *extractld.UrlProcessorList
	Processor extractld.UrlProcessor
}

func (m *ExtractorApp) Init() {
	m.MainApp.Init()
	m.List = &extractld.UrlProcessorList{}
	processors, err := bc3.CreateProcessors(m.Provider)
	if err == nil {
		m.List.AddAll(processors)
	} else {
		log.Println(err)
	}
	processors, err = ms.CreateProcessors(m.Provider)
	if err == nil {
		m.List.AddAll(processors)
	} else {
		log.Println(err)
	}

	m.MainApp.HttpServeMux().HandleFunc("/extract", m.extract)
}

func (m *ExtractorApp) MailProviders() []extractld.MailProvider {
	return extractld.GetProcessors[extractld.MailProvider](m.List)
}

func (m *ExtractorApp) extract(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	resource := r.URL.Query().Get("resource")
	log.Printf("Retrieving data about %s\n", resource)
	defer r.Body.Close()

	graph, err := m.Processor.Process(resource)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	if graph != nil {
		w.WriteHeader(200)
		w.Write([]byte(graph.UpdateQuery(nil)))
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Unknown resource"))
	}
}
