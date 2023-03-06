package schema

import "github.com/knakk/rdf"

const (
	RDF  = "http://www.w3.org/1999/02/22-rdf-syntax-ns#"
	RDFS = "http://www.w3.org/2000/01/rdf-schema#"
	DOAP = "http://usefulinc.com/ns/doap#"
	FOAF = "http://xmlns.com/foaf/0.1/"
	FLOW = "http://www.w3.org/2005/01/wf/flow#"
	DC   = "http://purl.org/dc/elements/1.1/"
	DCT  = "http://purl.org/dc/terms/"
)

func Define(iri string) rdf.IRI {
	term, _ := rdf.NewIRI(iri)
	return term
}

var (
	XSD_String   = Define("http://www.w3.org/2001/XMLSchema#string")
	XSD_DateTime = Define("http://www.w3.org/2001/XMLSchema#dateTime")

	RDF_Type   = Define(RDF + "type")
	RDF_HTML   = Define(RDF + "HTML")
	RDFS_Label = Define(RDFS + "label")

	DOAP_Project     = Define(DOAP + "Project")
	DOAP_name        = Define(DOAP + "name")
	DOAP_homepage    = Define(DOAP + "homepage")
	DOAP_description = Define(DOAP + "description")
	DOAP_bugDatabase = Define(DOAP + "bug-database")

	FLOW_Task        = Define(FLOW + "Task")
	FLOW_Open        = Define(FLOW + "Open")
	FLOW_Closed      = Define(FLOW + "Closed")
	FLOW_description = Define(FLOW + "description")
	FLOW_dependent   = Define(FLOW + "dependent")

	FLOW_Tracker = Define(FLOW + "Tracker")
	FLOW_tracker = Define(FLOW + "tracker")

	DCT_hasPart  = Define(DCT + "hasPart")
	DCT_isPartOf = Define(DCT + "isPartOf")
	DC_title     = Define(DC + "title")

	FOAF_mbox = Define(FOAF + "mbox")
)
