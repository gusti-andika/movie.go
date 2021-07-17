package data

import (
	"fmt"
	"testing"

	"github.com/gusti-andika/movie.go/genre"
)

func TestTreeDataInit(t *testing.T) {
	var d Node
	g := genre.Genre{Id: 1, Name: "Horror"}
	d = g

	fmt.Print(d)

	genres := make([]genre.Genre, 2)
	genres[0] = genre.Genre{Id: 1, Name: "Horror"}
	genres[1] = genre.Genre{Id: 2, Name: "Action"}

	treeData := TreeData{}
	treeData.Init(genres)
}
