package main

import (
	"net/url"

	_ "image/jpeg"
	_ "image/png"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/gusti-andika/movie.go/background"
	"github.com/gusti-andika/movie.go/config"
	"github.com/gusti-andika/movie.go/data"
	"github.com/gusti-andika/movie.go/genre"
	"github.com/gusti-andika/movie.go/widgetx"
)

var (
	topWindow   fyne.Window
	contentView *fyne.Container
	breadcrumb  *widgetx.Breadcrumb
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

	navView := createNavigatorView(&nav)

	//contentView = container.NewGridWithColumns(2)
	_container := container.NewMax(createMainView())
	split := container.NewHSplit(navView, _container)
	split.Offset = 0.2
	w.SetContent(split)
	w.SetMaster()

	w.Resize(fyne.NewSize(640, 460))
	go background.RefreshImage()
	w.ShowAndRun()
}

func createMainView() fyne.CanvasObject {
	breadcrumb = widgetx.NewBreadcrumb("Movie Golang App", func() {
		contentView.Objects = []fyne.CanvasObject{welcomeScreen()}
		contentView.Refresh()
	})

	top := container.NewVBox(breadcrumb.Container, widget.NewSeparator())

	contentView = container.NewGridWithColumns(2)
	contentView.Objects = []fyne.CanvasObject{welcomeScreen()}
	contentView.Refresh()
	return container.NewBorder(top, nil, nil, nil, container.NewVScroll(contentView))
}

func createNavigatorView(nav *data.TreeData) fyne.CanvasObject {

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
				if _, ok := (*nav).Uids[uid]; !ok {
					selectGenre(t.(genre.Genre))
				}
			}
		},
	}

	top := container.NewVBox(widget.NewLabel("GENRES"), widget.NewSeparator())
	return container.NewBorder(top, nil, nil, nil, tree)
}

func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}

func welcomeScreen() fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("Welcome to golang Movie Catalog", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container.NewHBox(
			widget.NewHyperlink("fyne.io", parseURL("https://fyne.io/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("themoviedb.org", parseURL("https://developers.themoviedb.org")),
		),
	))
}

func selectGenre(g genre.Genre) {
	m, _ := g.Movies()
	cards := []fyne.CanvasObject{}

	contentView.Objects = []fyne.CanvasObject{widget.NewProgressBarInfinite()}
	contentView.Refresh()

	for _, mm := range m {
		desc := widget.NewLabel(mm.Desc)
		desc.Wrapping = fyne.TextWrapWord
		card := widget.NewCard(mm.Title, "", desc)
		imageStr, _ := config.ImageURL(mm.Poster)
		background.LoadImage(imageStr, card)
		cards = append(cards, card)
	}

	breadcrumb.SetGenre(g.Name)
	contentView.Objects = cards
	contentView.Refresh()
}
