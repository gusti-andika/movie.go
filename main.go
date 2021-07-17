package main

import (
	"image"
	"net/http"

	_ "image/jpeg"
	_ "image/png"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
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

	//contentView = fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(90, 90)))
	contentView = container.NewGridWithColumns(2)
	_container := container.NewMax(container.NewVScroll(contentView))

	split := container.NewHSplit(navView, _container)
	split.Offset = 0.2
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
				selectGenre(t.(genre.Genre))
			}
		},
	}

	return container.NewBorder(nil, nil, nil, nil, tree)
}

func selectGenre(g genre.Genre) {
	m, _ := g.Movies()
	cards := []fyne.CanvasObject{}

	for _, mm := range m {
		desc := widget.NewLabel(mm.Desc)
		desc.Wrapping = fyne.TextWrapWord
		card := widget.NewCard(mm.Title, "", desc)
		imageStr, _ := config.ImageURL(mm.Poster)
		if i, f := loadImage(imageStr); f {
			card.Image = canvas.NewImageFromImage(i)
			card.Image.FillMode = canvas.ImageFillOriginal

		}
		cards = append(cards, card)
	}

	contentView.Objects = cards
	contentView.Refresh()
}

func loadImage(url string) (image.Image, bool) {
	res, err := http.Get(url)
	if err != nil {
		return nil, false
	}

	defer res.Body.Close()
	img, _, err := image.Decode(res.Body)
	if err != nil {
		return nil, false
	}

	return img, true
}
