package bc3

import (
	"strconv"

	"github.com/balazsgrill/basecamp3"
	"github.com/balazsgrill/sparqlupdate"
)

type HierarchyLevel interface {
	Get(string) HierarchyLevel
	Extract(*sparqlupdate.Graph) error
}

type Root struct {
	*basecamp3.Basecamp
	Ctx basecamp3.ContextWithTokenPersistence
}

func (r *Root) Extract(*sparqlupdate.Graph) error {
	return nil
}

func (r *Root) Get(s string) HierarchyLevel {
	accountId, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &account{
		Root:      r,
		accountId: accountId,
	}
}

type account struct {
	*Root
	accountId int
}

func (a *account) Get(s string) HierarchyLevel {
	switch s {
	case "projects":
		return &projects{
			account: a,
		}
	case "buckets":
		return &buckets{
			account: a,
		}
	}
	return nil
}

type projects struct {
	*account
}

func (p *projects) Get(s string) HierarchyLevel {
	projectId, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &project{
		projects:  p,
		projectid: projectId,
	}
}

type buckets struct {
	*account
}

func (p *buckets) Get(s string) HierarchyLevel {
	projectId, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &bucket{
		buckets:   p,
		projectid: projectId,
	}
}

type project struct {
	*projects
	projectid int
}

type bucket struct {
	*buckets
	projectid int
}

type todosets struct {
	*bucket
}

type todolists struct {
	*bucket
}

type todos struct {
	*bucket
}

func (b *bucket) Get(s string) HierarchyLevel {
	switch s {
	case "todosets":
		return &todosets{
			bucket: b,
		}
	case "todolists":
		return &todolists{
			bucket: b,
		}
	case "todos":
		return &todos{
			bucket: b,
		}
	}
	return nil
}

type todoset struct {
	*todosets
	todosetid int
}

func (t *todosets) Get(s string) HierarchyLevel {
	todosetid, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &todoset{
		todosets:  t,
		todosetid: todosetid,
	}
}

type todoset_todolists struct {
	*todoset
}

func (t *todoset) Get(s string) HierarchyLevel {
	switch s {
	case "todolists":
		return &todoset_todolists{
			todoset: t,
		}
	}
	return nil
}

type todolist struct {
	*todolists
	todolistid int
}

func (t *todolists) Get(s string) HierarchyLevel {
	tid, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &todolist{
		todolists:  t,
		todolistid: tid,
	}
}

type todo struct {
	*todos
	todoid int
}

func (t *todos) Get(s string) HierarchyLevel {
	tid, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &todo{
		todos:  t,
		todoid: tid,
	}
}
