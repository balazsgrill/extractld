package bc3

import (
	"fmt"
	"log"

	"github.com/balazsgrill/basecamp3"
	"github.com/balazsgrill/extractld/schema"
	"github.com/balazsgrill/sparqlupdate"
	"github.com/knakk/rdf"
)

func (ts *todoset) Extract(g *sparqlupdate.Graph) error {
	data, err := ts.Basecamp.TodoSet(ts.Ctx, ts.accountId, ts.projectid, int64(ts.todosetid))
	if err != nil {
		return err
	}
	todosetitem := schema.Define(data.AppURL)
	g.AddTriple(todosetitem, schema.RDF_Type, schema.FLOW_Tracker)
	g.AddTriple(todosetitem, schema.RDFS_Label, rdf.NewTypedLiteral(data.Name, schema.XSD_String))
	g.AddTriple(todosetitem, schema.DC_title, rdf.NewTypedLiteral(data.Title, schema.XSD_String))
	//g.AddTriple(todosetitem, schema.DCT_hasPart, schema.Define(data.AppTodoslistsURL))
	// infer parent URL
	g.AddTriple(todosetitem, schema.DCT_isPartOf, schema.Define(fmt.Sprintf("%s%d/projects/%d", urlprefix, ts.accountId, ts.projectid)))

	todos, err := ts.Basecamp.TodoSet_Lists(ts.Ctx, ts.accountId, ts.projectid, int64(ts.todosetid))
	if err != nil {
		return err
	}
	todolistAsParts(g, todosetitem, todos)

	return nil
}

func (ts *todolist) Extract(g *sparqlupdate.Graph) error {
	todo, err := ts.Basecamp.TodoList(ts.Ctx, ts.accountId, ts.projectid, ts.todolistid)
	if err != nil {
		return err
	}
	ts.extractTodo(g, todo)
	return nil
}

func (ts *todo) Extract(g *sparqlupdate.Graph) error {
	todo, err := ts.Basecamp.Todo(ts.Ctx, ts.accountId, ts.projectid, ts.todoid)
	if err != nil {
		return err
	}
	ts.extractTodo(g, todo)
	return nil
}

func todolistAsParts(g *sparqlupdate.Graph, parent rdf.IRI, todos []basecamp3.Todo) {
	for _, t := range todos {
		g.AddTriple(parent, schema.DCT_hasPart, schema.Define(t.AppURL))
	}
}

func (r *bucket) extractTodo(g *sparqlupdate.Graph, todo basecamp3.Todo) rdf.IRI {
	todoitem := schema.Define(todo.AppURL)

	g.AddTriple(todoitem, schema.DCT_isPartOf, schema.Define(todo.Parent.AppURL))
	g.AddTriple(todoitem, schema.RDF_Type, schema.FLOW_Task)

	// TODO uri to todoset?
	//g.AddTriple(todoitem, schema.FLOW_tracker, tracker)
	g.AddTriple(todoitem, schema.RDFS_Label, rdf.NewTypedLiteral(todo.Name, schema.XSD_String))
	g.AddTriple(todoitem, schema.DC_title, rdf.NewTypedLiteral(todo.Title, schema.XSD_String))
	g.AddTriple(todoitem, schema.FLOW_description, rdf.NewTypedLiteral(todo.Description, schema.RDF_HTML))
	if todo.Completed {
		g.AddTriple(todoitem, schema.RDF_Type, schema.FLOW_Closed)
	}

	//if todo.Type == "Todo" {
	// TODO Leaf todo
	//}
	if todo.Type == "Todolist" {
		groups, err := r.Basecamp.TodoList_Groups(r.Ctx, r.accountId, r.projectid, int(todo.ID))

		if err != nil {
			log.Println(err)
		}
		todolistAsParts(g, todoitem, groups)

		todos, err := r.Basecamp.Todos(r.Ctx, r.accountId, r.projectid, int(todo.ID))
		if err != nil {
			log.Println(err)
		}
		todolistAsParts(g, todoitem, todos)
	}

	return todoitem
}
