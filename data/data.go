package data

import (
	"sort"

	"github.com/gusti-andika/movie.go/genre"
)

type TreeData struct {
	Uids  map[string][]string
	Nodes map[string]Node
}

type Node interface {
	NodeId() string
	NodeName() string
}

func (d *TreeData) Init(list []genre.Genre) {
	d.Uids = make(map[string][]string)
	d.Nodes = make(map[string]Node, len(list))
	for _, t := range list {
		name := t.Name
		d.Nodes[name] = t

		first := string(name[0])
		if !hasElem(d.Uids[""], first) {
			d.Uids[""] = append(d.Uids[""], first)
			d.Nodes[first] = genre.Genre{Id: -1, Name: first}
		}

		d.Uids[first] = append(d.Uids[first], name)

	}

	if len(d.Uids[""]) > 0 {
		sort.Strings(d.Uids[""])
	}
}

func hasElem(slice []string, e string) bool {
	for _, v := range slice {
		if v == e {
			return true
		}
	}

	return false
}
