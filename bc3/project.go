package bc3

import (
	"github.com/balazsgrill/extractld/schema"
	"github.com/balazsgrill/sparqlupdate"
	"github.com/knakk/rdf"
)

func (p *project) Extract(g *sparqlupdate.Graph) error {
	project, err := p.Basecamp.Project(p.Ctx, p.accountId, p.projectid)
	if err != nil {
		return err
	}
	projectitem := schema.Define(project.AppURL)
	g.AddTriple(projectitem, schema.RDF_Type, schema.DOAP_Project)
	g.AddTriple(projectitem, schema.DOAP_name, rdf.NewTypedLiteral(project.Name, schema.XSD_String))
	g.AddTriple(projectitem, schema.DOAP_description, rdf.NewTypedLiteral(project.Description, schema.RDF_HTML))

	for _, dock := range project.Dock {
		item := schema.Define(dock.AppURL)
		g.AddTriple(projectitem, schema.DCT_hasPart, item)
		//g.AddTriple(item, schema.DCT_isPartOf, projectitem)
	}
	return nil
}
