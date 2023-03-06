package main

import (
	"log"
	"net/http"

	"github.com/balazsgrill/extractld"
	"github.com/balazsgrill/extractld/bc3"
	cmd "github.com/balazsgrill/oauthenticator/server"
)

type mainExtractor struct {
	cmd.Main
	list      *extractld.UrlProcessorList
	processor extractld.UrlProcessor
}

func (m *mainExtractor) Init() {
	m.Main.Init()
	processors, err := bc3.CreateProcessors(m.Provider)
	if err == nil {
		m.list.AddAll(processors)
	} else {
		log.Println(err)
	}

	http.HandleFunc("/extract", m.extract)
}

func main() {
	list := &extractld.UrlProcessorList{}
	main := &mainExtractor{
		list:      list,
		processor: list,
	}
	main.InitFlags()
	main.ParseFlags()
	main.Init()
	main.Start()
}

func (m *mainExtractor) extract(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	resource := r.URL.Query().Get("resource")
	log.Printf("Retrieving data about %s\n", resource)
	defer r.Body.Close()

	graph, err := m.processor.Process(resource)
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
