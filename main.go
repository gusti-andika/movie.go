package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/gusti-andika/movie.go/config"
	"github.com/gusti-andika/movie.go/data"
	"github.com/gusti-andika/movie.go/genre"
)

var (
	topWindow   fyne.Window
	contentView *fyne.Container
)

func main() {

	a := app.New()
	w := a.NewWindow("Movies")
	topWindow = w

	nav := data.TreeData{}

	genres, err := genre.List()
	if err != nil {
		panic(err)
	}

	nav.Init(genres)

	navView := makeNav(&nav)

	contentView := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(90, 90)))
	_container := container.NewMax(contentView)

	split := container.NewHSplit(navView, _container)
	w.SetContent(split)
	w.SetMaster()

	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}

func makeNav(nav *data.TreeData) fyne.CanvasObject {

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return (*nav).Uids[uid]
		},

		IsBranch: func(uid string) bool {
			children, ok := (*nav).Uids[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := (*nav).Nodes[uid]
			if !ok {
				fyne.LogError("Missing genre: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.NodeName())
		},
		OnSelected: func(uid string) {
			if t, ok := (*nav).Nodes[uid]; ok {
				selectGenre(t.(*genre.Genre))
			}
		},
	}

	return container.NewBorder(nil, nil, nil, nil, tree)
}

func selectGenre(g *genre.Genre) {
	m, _ := g.Movies()
	cards := make([]fyne.CanvasObject, len(m))

	for _, mm := range m {
		card := widget.NewCard(mm.Title, "", widget.NewLabel(mm.Desc))
		imageStr, _ := config.ImageURL(mm.Poster)
		uri := storage.NewURI(imageStr)
		card.Image = canvas.NewImageFromURI(uri)
		cards = append(cards, card)
	}

	contentView.Objects = cards
	contentView.Refresh()
}
