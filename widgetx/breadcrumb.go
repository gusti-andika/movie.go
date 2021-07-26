package widgetx

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Breadcrumb struct {
	Container *fyne.Container
	home      *HyperlinkX
	genre     *widget.Label
}

func NewBreadcrumb(title string, action func()) *Breadcrumb {
	home := NewHyperlink(title, action)
	genre := widget.NewLabel("")
	c := container.NewMax(container.NewHBox(home, genre))

	return &Breadcrumb{
		Container: c,
		home:      home,
		genre:     genre,
	}

}

func (b *Breadcrumb) SetGenre(text string) {
	b.genre.SetText(" > " + text)
}

func (b *Breadcrumb) ClearGenre() {
	b.genre.SetText("")
}
