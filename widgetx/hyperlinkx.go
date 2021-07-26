package widgetx

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type HyperlinkX struct {
	widget.Hyperlink
	action func()
}

func NewHyperlink(text string, action func()) *HyperlinkX {
	hl := widget.NewHyperlink(text, nil)
	return &HyperlinkX{
		Hyperlink: *hl,
		action:    action,
	}
}

func (hlx *HyperlinkX) Tapped(p *fyne.PointEvent) {
	hlx.action()
}
