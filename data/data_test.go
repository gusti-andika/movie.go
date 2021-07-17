package data

import (
	"testing"

	"github.com/gusti-andika/movie.go/genre"
	"github.com/stretchr/testify/assert"
)

func TestTreeDataInit(t *testing.T) {

	genres := make([]genre.Genre, 3)
	genres[0] = genre.Genre{Id: 1, Name: "Horror"}
	genres[1] = genre.Genre{Id: 2, Name: "Action"}
	genres[2] = genre.Genre{Id: 3, Name: "Adventure"}

	treeData := TreeData{}
	treeData.Init(genres)

	assert.Equal(t, 3, len(treeData.Uids))

	expectedNodes := []string{"A", "H"}
	assert.Equal(t, expectedNodes, treeData.Uids[""])
	assert.Equal(t, genre.Genre{-1, "A"}, treeData.Nodes["A"])

	expectedNodes = []string{"Action", "Adventure"}
	assert.Equal(t, expectedNodes, treeData.Uids["A"])
	assert.Equal(t, genre.Genre{1, "Horror"}, treeData.Nodes["Horror"])

	expectedNodes = []string{"Horror"}
	assert.Equal(t, expectedNodes, treeData.Uids["H"])
}
