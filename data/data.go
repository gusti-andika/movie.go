package data

import "github.com/gusti-andika/movie.go/genre"

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
		_, ok := d.Uids[""]
		if !ok {
			d.Uids[""] = make([]string, len(list))
		}

		d.Uids[""] = append(d.Uids[""], t.NodeId())
		d.Nodes[t.NodeId()] = t
	}
}
